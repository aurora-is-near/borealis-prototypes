//! Decoding of the [`proto`] types into the [`aurora_refiner_types`].
use crate::{access_key_permission_view, proto};
use aurora_refiner_types::near_block::{
    BlockView, ChunkHeaderView, ChunkView, ExecutionOutcomeWithOptionalReceipt, ExecutionOutcomeWithReceipt,
    IndexerBlockHeaderView, NEARBlock, Shard, TransactionWithOutcome,
};
use aurora_refiner_types::near_primitives::hash::CryptoHash;
use aurora_refiner_types::near_primitives::types::{AccountId, StoreKey, StoreValue};
use aurora_refiner_types::near_primitives::{self,views::validator_stake_view::ValidatorStakeView};
use aurora_refiner_types::near_primitives::views::{
    AccessKeyPermissionView, AccessKeyView, AccountView, ActionView, CostGasUsed, DataReceiverView,
    ExecutionMetadataView, ExecutionOutcomeView, ExecutionOutcomeWithIdView, ExecutionStatusView, ReceiptEnumView,
    ReceiptView, SignedTransactionView, StateChangeCauseView, StateChangeValueView, StateChangeWithCauseView,
    ValidatorStakeViewV1,
};
use itertools::{Either, Itertools};
use aurora_refiner_types::near_crypto::{ED25519PublicKey, PublicKey, Secp256K1PublicKey, Signature};
use aurora_refiner_types::near_primitives::account::{AccessKeyPermission, FunctionCallPermission};
use aurora_refiner_types::near_primitives::challenge::SlashedValidator;
use aurora_refiner_types::near_primitives::errors::{
    ActionError, ActionErrorKind, ActionsValidationError, CompilationError, FunctionCallError, HostError,
    InvalidAccessKeyError, InvalidTxError, MethodResolveError, PrepareError, ReceiptValidationError, TxExecutionError,
    WasmTrap,
};
use aurora_refiner_types::near_primitives::merkle::{Direction, MerklePathItem};
use aurora_refiner_types::near_primitives::transaction::{
    Action, AddKeyAction, CreateAccountAction, DeleteAccountAction, DeleteKeyAction, DeployContractAction,
    FunctionCallAction, StakeAction, TransferAction,
};

impl From<proto::Messages> for NEARBlock {
    fn from(value: proto::Messages) -> Self {
        let mut messages: (Vec<_>, Vec<_>) =
            value
                .into_inner()
                .into_iter()
                .partition_map(|v| match v.payload.unwrap() {
                    proto::message::Payload::NearBlockShard(shard) => Either::Left(Shard::from(shard)),
                    proto::message::Payload::NearBlockHeader(header) => Either::Right(BlockView::from(header)),
                    v => panic!("Unexpected variant {v:?}"),
                });

        Self {
            block: messages.1.pop().unwrap(),
            shards: messages.0,
        }
    }
}

impl From<proto::BlockShard> for NEARBlock {
    fn from(value: proto::BlockShard) -> Self {
        Self {
            block: BlockView::from(value.header.as_ref().unwrap().clone()),
            shards: vec![Shard::from(value)],
        }
    }
}

impl From<proto::BlockShard> for Shard {
    fn from(value: proto::BlockShard) -> Self {
        Self {
            shard_id: value.shard_id,
            chunk: value.chunk.map(Into::into),
            receipt_execution_outcomes: value.receipt_execution_outcomes.into_iter().map(Into::into).collect(),
            state_changes: value.state_changes.into_iter().map(Into::into).collect(),
        }
    }
}

impl From<proto::ExecutionOutcomeWithReceipt> for ExecutionOutcomeWithReceipt {
    fn from(value: proto::ExecutionOutcomeWithReceipt) -> Self {
        Self {
            execution_outcome: value.execution_outcome.unwrap().into(),
            receipt: value.receipt.unwrap().into(),
        }
    }
}

impl From<proto::ExecutionOutcomeView> for ExecutionOutcomeView {
    fn from(value: proto::ExecutionOutcomeView) -> Self {
        Self {
            logs: value.logs,
            receipt_ids: value
                .h256_receipt_ids
                .into_iter()
                .map(|v| CryptoHash(v.try_into().unwrap()))
                .collect(),
            gas_burnt: value.gas_burnt,
            tokens_burnt: u128::from_be_bytes(value.u128_tokens_burnt.try_into().unwrap()),
            executor_id: AccountId::try_from(value.executor_id).unwrap(),
            status: value.status.unwrap().into(),
            metadata: value.metadata.unwrap().into(),
        }
    }
}

impl From<proto::ExecutionStatusView> for ExecutionStatusView {
    fn from(value: proto::ExecutionStatusView) -> Self {
        match value.variant.unwrap() {
            proto::execution_status_view::Variant::Unknown(..) => ExecutionStatusView::Unknown,
            proto::execution_status_view::Variant::Failure(v) => ExecutionStatusView::Failure(v.error.unwrap().into()),
            proto::execution_status_view::Variant::SuccessValue(v) => ExecutionStatusView::SuccessValue(v.value),
            proto::execution_status_view::Variant::SuccessReceiptId(v) => {
                ExecutionStatusView::SuccessReceiptId(CryptoHash(v.h256_receipt_hash.try_into().unwrap()))
            }
        }
    }
}

impl From<proto::tx_execution_error::action_error::ActionErrorKind> for ActionErrorKind {
    fn from(value: proto::tx_execution_error::action_error::ActionErrorKind) -> Self {
        match value.variant.unwrap() {
            proto::tx_execution_error::action_error::action_error_kind::Variant::AccountAlreadyExists(v) => {
                Self::AccountAlreadyExists {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::AccountDoesNotExist(v) => {
                Self::AccountDoesNotExist {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::CreateAccountOnlyByRegistrar(v) => {
                Self::CreateAccountOnlyByRegistrar {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                    predecessor_id: AccountId::try_from(v.predecessor_id).unwrap(),
                    registrar_account_id: AccountId::try_from(v.registrar_account_id).unwrap(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::CreateAccountNotAllowed(v) => {
                Self::CreateAccountNotAllowed {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                    predecessor_id: AccountId::try_from(v.predecessor_id).unwrap(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::ActorNoPermission(v) => {
                Self::ActorNoPermission {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                    actor_id: AccountId::try_from(v.actor_id).unwrap(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::DeleteKeyDoesNotExist(v) => {
                Self::DeleteKeyDoesNotExist {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                    public_key: v.public_key.unwrap().into(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::AddKeyAlreadyExists(v) => {
                Self::AddKeyAlreadyExists {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                    public_key: v.public_key.unwrap().into(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::DeleteAccountStaking(v) => {
                Self::DeleteAccountStaking {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::LackBalanceForState(v) => {
                Self::LackBalanceForState {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                    amount: u128::from_be_bytes(v.u128_amount.try_into().unwrap()),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::TriesToUnstake(v) => {
                Self::TriesToUnstake {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::TriesToStake(v) => {
                Self::TriesToStake {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                    stake: u128::from_be_bytes(v.u128_stake.try_into().unwrap()),
                    balance: u128::from_be_bytes(v.u128_balance.try_into().unwrap()),
                    locked: u128::from_be_bytes(v.u128_locked.try_into().unwrap()),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::InsufficientStake(v) => {
                Self::InsufficientStake {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                    stake: u128::from_be_bytes(v.u128_stake.try_into().unwrap()),
                    minimum_stake: u128::from_be_bytes(v.u128_minimum_stake.try_into().unwrap()),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::FunctionCallError(v) => {
                Self::FunctionCallError(v.error.unwrap().into())
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::NewReceiptValidationError(v) => {
                Self::NewReceiptValidationError(v.error.unwrap().into())
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::OnlyImplicitAccountCreationAllowed(
                v,
            ) => Self::OnlyImplicitAccountCreationAllowed {
                account_id: AccountId::try_from(v.account_id).unwrap(),
            },
            proto::tx_execution_error::action_error::action_error_kind::Variant::DeleteAccountWithLargeState(v) => {
                Self::DeleteAccountWithLargeState {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::DelegateActionInvalidSignature(..) => {
                Self::DelegateActionInvalidSignature
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::DelegateActionSenderDoesNotMatchTxReceiver(error) => {
                Self::DelegateActionSenderDoesNotMatchTxReceiver {
                    sender_id: AccountId::try_from(error.sender_id).unwrap(),
                    receiver_id: AccountId::try_from(error.receiver_id).unwrap(),
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::DelegateActionExpired(..) => {
                Self::DelegateActionExpired
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::DelegateActionAccessKeyError(error) => {
                Self::DelegateActionAccessKeyError(
                    InvalidAccessKeyError::from(error.error.unwrap())
                )
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::DelegateActionInvalidNonce(error) => {
                Self::DelegateActionInvalidNonce {
                    delegate_nonce: error.delegate_nonce,
                    ak_nonce: error.ak_nonce,
                }
            }
            proto::tx_execution_error::action_error::action_error_kind::Variant::DelegateActionNonceTooLarge(error) => {
                Self::DelegateActionNonceTooLarge {
                    delegate_nonce: error.delegate_nonce,
                    upper_bound: error.upper_bound,
                }
            }
        }
    }
}

impl From<proto::ReceiptValidationError> for ReceiptValidationError {
    fn from(value: proto::ReceiptValidationError) -> Self {
        match value.variant.unwrap() {
            proto::receipt_validation_error::Variant::InvalidPredecessorId(v) => {
                ReceiptValidationError::InvalidPredecessorId {
                    account_id: v.account_id,
                }
            }
            proto::receipt_validation_error::Variant::InvalidReceiverId(v) => {
                ReceiptValidationError::InvalidReceiverId {
                    account_id: v.account_id,
                }
            }
            proto::receipt_validation_error::Variant::InvalidSignerId(v) => ReceiptValidationError::InvalidSignerId {
                account_id: v.account_id,
            },
            proto::receipt_validation_error::Variant::InvalidDataReceiverId(v) => {
                ReceiptValidationError::InvalidDataReceiverId {
                    account_id: v.account_id,
                }
            }
            proto::receipt_validation_error::Variant::ReturnedValueLengthExceeded(v) => {
                ReceiptValidationError::ReturnedValueLengthExceeded {
                    length: v.length,
                    limit: v.limit,
                }
            }
            proto::receipt_validation_error::Variant::NumberInputDataDependenciesExceeded(v) => {
                ReceiptValidationError::NumberInputDataDependenciesExceeded {
                    limit: v.limit,
                    number_of_input_data_dependencies: v.number_of_input_data_dependencies,
                }
            }
            proto::receipt_validation_error::Variant::ActionsValidation(v) => {
                ReceiptValidationError::ActionsValidation(v.error.unwrap().into())
            }
        }
    }
}

impl From<proto::ActionsValidationError> for ActionsValidationError {
    fn from(value: proto::ActionsValidationError) -> Self {
        match value.variant.unwrap() {
            proto::actions_validation_error::Variant::DeleteActionMustBeFinal(..) => {
                ActionsValidationError::DeleteActionMustBeFinal
            }
            proto::actions_validation_error::Variant::TotalPrepaidGasExceeded(v) => {
                ActionsValidationError::TotalPrepaidGasExceeded {
                    limit: v.limit,
                    total_prepaid_gas: v.total_prepaid_gas,
                }
            }
            proto::actions_validation_error::Variant::TotalNumberOfActionsExceeded(v) => {
                ActionsValidationError::TotalNumberOfActionsExceeded {
                    limit: v.limit,
                    total_number_of_actions: v.total_number_of_bytes,
                }
            }
            proto::actions_validation_error::Variant::AddKeyMethodNamesNumberOfBytesExceeded(v) => {
                ActionsValidationError::AddKeyMethodNamesNumberOfBytesExceeded {
                    limit: v.limit,
                    total_number_of_bytes: v.total_number_of_bytes,
                }
            }
            proto::actions_validation_error::Variant::AddKeyMethodNameLengthExceeded(v) => {
                ActionsValidationError::AddKeyMethodNameLengthExceeded {
                    limit: v.limit,
                    length: v.length,
                }
            }
            proto::actions_validation_error::Variant::IntegerOverflow(..) => ActionsValidationError::IntegerOverflow,
            proto::actions_validation_error::Variant::InvalidAccountId(v) => ActionsValidationError::InvalidAccountId {
                account_id: v.account_id,
            },
            proto::actions_validation_error::Variant::ContractSizeExceeded(v) => {
                ActionsValidationError::ContractSizeExceeded {
                    limit: v.limit,
                    size: v.size,
                }
            }
            proto::actions_validation_error::Variant::FunctionCallMethodNameLengthExceeded(v) => {
                ActionsValidationError::FunctionCallMethodNameLengthExceeded {
                    limit: v.limit,
                    length: v.length,
                }
            }
            proto::actions_validation_error::Variant::FunctionCallArgumentsLengthExceeded(v) => {
                ActionsValidationError::FunctionCallArgumentsLengthExceeded {
                    limit: v.limit,
                    length: v.length,
                }
            }
            proto::actions_validation_error::Variant::UnsuitableStakingKey(v) => {
                ActionsValidationError::UnsuitableStakingKey {
                    public_key: v.public_key.unwrap().into(),
                }
            }
            proto::actions_validation_error::Variant::FunctionCallZeroAttachedGas(..) => {
                ActionsValidationError::FunctionCallZeroAttachedGas
            }
            proto::actions_validation_error::Variant::DelegateActionMustBeOnlyOne(..) => {
                ActionsValidationError::DelegateActionMustBeOnlyOne
            }
            proto::actions_validation_error::Variant::UnsupportedProtocolFeature(v) => {
                ActionsValidationError::UnsupportedProtocolFeature {
                    protocol_feature: v.protocol_feature,
                    version: v.version,
                }
            }
        }
    }
}

impl From<proto::FunctionCallErrorSer> for FunctionCallError {
    fn from(value: proto::FunctionCallErrorSer) -> Self {
        match value.variant.unwrap() {
            proto::function_call_error_ser::Variant::CompilationError(v) => {
                Self::CompilationError(v.error.unwrap().into())
            }
            proto::function_call_error_ser::Variant::LinkError(v) => Self::LinkError { msg: v.msg },
            proto::function_call_error_ser::Variant::MethodResolveError(v) => {
                Self::MethodResolveError(proto::MethodResolveError::try_from(v.error).unwrap().into())
            }
            proto::function_call_error_ser::Variant::WasmTrap(v) => {
                Self::WasmTrap(proto::WasmTrap::try_from(v.error).unwrap().into())
            }
            proto::function_call_error_ser::Variant::WasmUnknownError(..) => Self::WasmUnknownError,
            proto::function_call_error_ser::Variant::HostError(v) => Self::HostError(v.error.unwrap().into()),
            proto::function_call_error_ser::Variant::ExecutionError(v) => Self::ExecutionError(v.message),
        }
    }
}

impl From<proto::CompilationError> for CompilationError {
    fn from(value: proto::CompilationError) -> Self {
        match value.variant.unwrap() {
            proto::compilation_error::Variant::CodeDoesNotExist(v) => CompilationError::CodeDoesNotExist {
                account_id: AccountId::try_from(v.account_id).unwrap(),
            },
            proto::compilation_error::Variant::PrepareError(error) => {
                CompilationError::PrepareError(proto::PrepareError::try_from(error.error).unwrap().into())
            }
            proto::compilation_error::Variant::WasmerCompileError(v) => {
                CompilationError::WasmerCompileError { msg: v.msg }
            }
        }
    }
}

impl From<proto::MethodResolveError> for MethodResolveError {
    fn from(value: proto::MethodResolveError) -> Self {
        match value {
            proto::MethodResolveError::MethodEmptyName => MethodResolveError::MethodEmptyName,
            proto::MethodResolveError::MethodNotFound => MethodResolveError::MethodNotFound,
            proto::MethodResolveError::MethodInvalidSignature => MethodResolveError::MethodInvalidSignature,
        }
    }
}

impl From<proto::HostError> for HostError {
    fn from(value: proto::HostError) -> Self {
        match value.variant.unwrap() {
            proto::host_error::Variant::BadUtf16(..) => HostError::BadUTF16,
            proto::host_error::Variant::BadUtf8(..) => HostError::BadUTF8,
            proto::host_error::Variant::GasExceeded(..) => HostError::GasExceeded,
            proto::host_error::Variant::GasLimitExceeded(..) => HostError::GasLimitExceeded,
            proto::host_error::Variant::BalanceExceeded(..) => HostError::BalanceExceeded,
            proto::host_error::Variant::EmptyMethodName(..) => HostError::EmptyMethodName,
            proto::host_error::Variant::GuestPanic(v) => HostError::GuestPanic { panic_msg: v.panic_msg },
            proto::host_error::Variant::IntegerOverflow(..) => HostError::IntegerOverflow,
            proto::host_error::Variant::InvalidPromiseIndex(v) => HostError::InvalidPromiseIndex {
                promise_idx: v.promise_idx,
            },
            proto::host_error::Variant::CannotAppendActionToJointPromise(..) => {
                HostError::CannotAppendActionToJointPromise
            }
            proto::host_error::Variant::CannotReturnJointPromise(..) => HostError::CannotReturnJointPromise,
            proto::host_error::Variant::InvalidPromiseResultIndex(v) => HostError::InvalidPromiseResultIndex {
                result_idx: v.result_idx,
            },
            proto::host_error::Variant::InvalidRegisterId(v) => HostError::InvalidRegisterId {
                register_id: v.register_id,
            },
            proto::host_error::Variant::IteratorWasInvalidated(v) => HostError::IteratorWasInvalidated {
                iterator_index: v.iterator_index,
            },
            proto::host_error::Variant::MemoryAccessViolation(..) => HostError::MemoryAccessViolation,
            proto::host_error::Variant::InvalidReceiptIndex(v) => HostError::InvalidReceiptIndex {
                receipt_index: v.receipt_index,
            },
            proto::host_error::Variant::InvalidIteratorIndex(v) => HostError::InvalidIteratorIndex {
                iterator_index: v.iterator_index,
            },
            proto::host_error::Variant::InvalidAccountId(..) => HostError::InvalidAccountId,
            proto::host_error::Variant::InvalidMethodName(..) => HostError::InvalidMethodName,
            proto::host_error::Variant::InvalidPublicKey(..) => HostError::InvalidPublicKey,
            proto::host_error::Variant::ProhibitedInView(v) => HostError::ProhibitedInView {
                method_name: v.method_name,
            },
            proto::host_error::Variant::NumberOfLogsExceeded(v) => HostError::NumberOfLogsExceeded { limit: v.limit },
            proto::host_error::Variant::KeyLengthExceeded(v) => HostError::KeyLengthExceeded {
                length: v.length,
                limit: v.limit,
            },
            proto::host_error::Variant::ValueLengthExceeded(v) => HostError::ValueLengthExceeded {
                length: v.length,
                limit: v.limit,
            },
            proto::host_error::Variant::TotalLogLengthExceeded(v) => HostError::TotalLogLengthExceeded {
                length: v.length,
                limit: v.limit,
            },
            proto::host_error::Variant::NumberPromisesExceeded(v) => HostError::NumberPromisesExceeded {
                number_of_promises: v.number_of_promises,
                limit: v.limit,
            },
            proto::host_error::Variant::NumberInputDataDependenciesExceeded(v) => {
                HostError::NumberInputDataDependenciesExceeded {
                    number_of_input_data_dependencies: v.number_of_input_data_dependencies,
                    limit: v.limit,
                }
            }
            proto::host_error::Variant::ReturnedValueLengthExceeded(v) => HostError::ReturnedValueLengthExceeded {
                length: v.length,
                limit: v.limit,
            },
            proto::host_error::Variant::ContractSizeExceeded(v) => HostError::ContractSizeExceeded {
                size: v.size,
                limit: v.limit,
            },
            proto::host_error::Variant::Deprecated(v) => HostError::Deprecated {
                method_name: v.method_name,
            },
            proto::host_error::Variant::EcrecoverError(v) => HostError::ECRecoverError { msg: v.msg },
            proto::host_error::Variant::AltBn128InvalidInput(v) => HostError::AltBn128InvalidInput { msg: v.msg },
            proto::host_error::Variant::Ed25519VerifyInvalidInput(v) => {
                HostError::Ed25519VerifyInvalidInput { msg: v.msg }
            }
        }
    }
}

impl From<proto::PrepareError> for PrepareError {
    fn from(value: proto::PrepareError) -> Self {
        match value {
            proto::PrepareError::Serialization => PrepareError::Serialization,
            proto::PrepareError::Deserialization => PrepareError::Deserialization,
            proto::PrepareError::InternalMemoryDeclared => PrepareError::InternalMemoryDeclared,
            proto::PrepareError::GasInstrumentation => PrepareError::GasInstrumentation,
            proto::PrepareError::StackHeightInstrumentation => PrepareError::StackHeightInstrumentation,
            proto::PrepareError::Instantiate => PrepareError::Instantiate,
            proto::PrepareError::Memory => PrepareError::Memory,
            proto::PrepareError::TooManyFunctions => PrepareError::TooManyFunctions,
            proto::PrepareError::TooManyLocals => PrepareError::TooManyLocals,
        }
    }
}

impl From<proto::WasmTrap> for WasmTrap {
    fn from(value: proto::WasmTrap) -> Self {
        match value {
            proto::WasmTrap::Unreachable => WasmTrap::Unreachable,
            proto::WasmTrap::IncorrectCallIndirectSignature => WasmTrap::IncorrectCallIndirectSignature,
            proto::WasmTrap::MemoryOutOfBounds => WasmTrap::MemoryOutOfBounds,
            proto::WasmTrap::CallIndirectOob => WasmTrap::CallIndirectOOB,
            proto::WasmTrap::IllegalArithmetic => WasmTrap::IllegalArithmetic,
            proto::WasmTrap::MisalignedAtomicAccess => WasmTrap::MisalignedAtomicAccess,
            proto::WasmTrap::IndirectCallToNull => WasmTrap::IndirectCallToNull,
            proto::WasmTrap::StackOverflow => WasmTrap::StackOverflow,
            proto::WasmTrap::GenericTrap => WasmTrap::GenericTrap,
        }
    }
}

impl From<proto::InvalidAccessKeyError> for InvalidAccessKeyError {
    fn from(value: proto::InvalidAccessKeyError) -> Self {
        match value.variant.unwrap() {
            proto::invalid_access_key_error::Variant::AccessKeyNotFound(v) => {
                InvalidAccessKeyError::AccessKeyNotFound {
                    public_key: v.public_key.unwrap().into(),
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                }
            }
            proto::invalid_access_key_error::Variant::ReceiverMismatch(v) => InvalidAccessKeyError::ReceiverMismatch {
                tx_receiver: AccountId::try_from(v.tx_receiver).unwrap(),
                ak_receiver: v.ak_receiver,
            },
            proto::invalid_access_key_error::Variant::MethodNameMismatch(v) => {
                InvalidAccessKeyError::MethodNameMismatch {
                    method_name: v.method_name,
                }
            }
            proto::invalid_access_key_error::Variant::RequiresFullAccess(..) => {
                InvalidAccessKeyError::RequiresFullAccess
            }
            proto::invalid_access_key_error::Variant::NotEnoughAllowance(v) => {
                InvalidAccessKeyError::NotEnoughAllowance {
                    public_key: v.public_key.unwrap().into(),
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                    allowance: u128::from_be_bytes(v.u128_allowance.try_into().unwrap()),
                    cost: u128::from_be_bytes(v.u128_cost.try_into().unwrap()),
                }
            }
            proto::invalid_access_key_error::Variant::DepositWithFunctionCall(..) => {
                InvalidAccessKeyError::DepositWithFunctionCall
            }
        }
    }
}

impl From<proto::tx_execution_error::invalid_tx_error::ActionsValidation> for ActionsValidationError {
    fn from(value: proto::tx_execution_error::invalid_tx_error::ActionsValidation) -> Self {
        match value.variant.unwrap() {
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::DeleteActionMustBeFinal(..) => ActionsValidationError::DeleteActionMustBeFinal,
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::TotalPrepaidGasExceeded(v) => ActionsValidationError::TotalPrepaidGasExceeded { total_prepaid_gas: v.total_prepaid_gas, limit: v.limit },
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::TotalNumberOfActionsExceeded(v) => ActionsValidationError::TotalNumberOfActionsExceeded { total_number_of_actions: v.total_number_of_actions, limit: v.limit },
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::AddKeyMethodNamesNumberOfBytesExceeded(v) => ActionsValidationError::AddKeyMethodNamesNumberOfBytesExceeded { total_number_of_bytes: v.total_number_of_bytes, limit: v.limit },
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::AddKeyMethodNameLengthExceeded(v) => ActionsValidationError::AddKeyMethodNameLengthExceeded { length: v.length, limit: v.limit },
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::IntegerOverflow(..) => ActionsValidationError::IntegerOverflow,
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::InvalidAccountId(v) => ActionsValidationError::InvalidAccountId { account_id: v.account_id },
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::ContractSizeExceeded(v) => ActionsValidationError::ContractSizeExceeded { size: v.size, limit: v.limit },
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::FunctionCallMethodNameLengthExceeded(v) => ActionsValidationError::FunctionCallMethodNameLengthExceeded { length: v.length, limit: v.limit },
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::FunctionCallArgumentsLengthExceeded(v) => ActionsValidationError::FunctionCallArgumentsLengthExceeded { length: v.length, limit: v.limit },
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::UnsuitableStakingKey(v) => ActionsValidationError::UnsuitableStakingKey { public_key: v.public_key.unwrap().into() },
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::FunctionCallZeroAttachedGas(..) => ActionsValidationError::FunctionCallZeroAttachedGas,
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::DelegateActionMustBeOnlyOne(..) => ActionsValidationError::DelegateActionMustBeOnlyOne,
            proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::UnsupportedProtocolFeature(v) => ActionsValidationError::UnsupportedProtocolFeature {
                protocol_feature: v.protocol_feature,
                version: v.version,
            },
        }
    }
}

impl From<proto::TxExecutionError> for TxExecutionError {
    fn from(value: proto::TxExecutionError) -> Self {
        match value.variant.unwrap() {
            proto::tx_execution_error::Variant::ActionError(v) => TxExecutionError::ActionError(ActionError {
                index: v.index,
                kind: v.kind.unwrap().into(),
            }),
            proto::tx_execution_error::Variant::InvalidTxError(error) => TxExecutionError::InvalidTxError(match error
                .variant
                .unwrap()
            {
                proto::tx_execution_error::invalid_tx_error::Variant::InvalidAccessKeyError(v) => {
                    InvalidTxError::InvalidAccessKeyError(v.into())
                }
                proto::tx_execution_error::invalid_tx_error::Variant::InvalidSignerId(v) => {
                    InvalidTxError::InvalidSignerId { signer_id: v.signer_id }
                }
                proto::tx_execution_error::invalid_tx_error::Variant::SignerDoesNotExist(v) => {
                    InvalidTxError::SignerDoesNotExist {
                        signer_id: AccountId::try_from(v.signer_id).unwrap(),
                    }
                }
                proto::tx_execution_error::invalid_tx_error::Variant::InvalidNonce(v) => InvalidTxError::InvalidNonce {
                    tx_nonce: v.tx_nonce,
                    ak_nonce: v.ak_nonce,
                },
                proto::tx_execution_error::invalid_tx_error::Variant::NonceTooLarge(v) => {
                    InvalidTxError::NonceTooLarge {
                        tx_nonce: v.tx_nonce,
                        upper_bound: v.upper_bound,
                    }
                }
                proto::tx_execution_error::invalid_tx_error::Variant::InvalidReceiverId(v) => {
                    InvalidTxError::InvalidReceiverId {
                        receiver_id: v.receiver_id,
                    }
                }
                proto::tx_execution_error::invalid_tx_error::Variant::InvalidSignature(..) => {
                    InvalidTxError::InvalidSignature
                }
                proto::tx_execution_error::invalid_tx_error::Variant::NotEnoughBalance(v) => {
                    InvalidTxError::NotEnoughBalance {
                        signer_id: AccountId::try_from(v.signer_id).unwrap(),
                        balance: u128::from_be_bytes(v.u128_balance.try_into().unwrap()),
                        cost: u128::from_be_bytes(v.u128_cost.try_into().unwrap()),
                    }
                }
                proto::tx_execution_error::invalid_tx_error::Variant::LackBalanceForState(v) => {
                    InvalidTxError::LackBalanceForState {
                        signer_id: AccountId::try_from(v.signer_id).unwrap(),
                        amount: u128::from_be_bytes(v.u128_amount.try_into().unwrap()),
                    }
                }
                proto::tx_execution_error::invalid_tx_error::Variant::CostOverflow(..) => InvalidTxError::CostOverflow,
                proto::tx_execution_error::invalid_tx_error::Variant::InvalidChain(..) => InvalidTxError::InvalidChain,
                proto::tx_execution_error::invalid_tx_error::Variant::Expired(..) => InvalidTxError::Expired,
                proto::tx_execution_error::invalid_tx_error::Variant::ActionsValidation(v) => {
                    InvalidTxError::ActionsValidation(v.into())
                }
                proto::tx_execution_error::invalid_tx_error::Variant::TransactionSizeExceeded(v) => {
                    InvalidTxError::TransactionSizeExceeded {
                        size: v.size,
                        limit: v.limit,
                    }
                }
            }),
        }
    }
}

impl From<proto::ExecutionMetadataView> for ExecutionMetadataView {
    fn from(value: proto::ExecutionMetadataView) -> Self {
        Self {
            version: value.version,
            gas_profile: value
                .gas_profile
                .map(|v| v.gas_profile.into_iter().map(Into::into).collect()),
        }
    }
}

impl From<proto::CostGasUsed> for CostGasUsed {
    fn from(value: proto::CostGasUsed) -> Self {
        CostGasUsed {
            cost_category: match proto::CostCategory::try_from(value.cost_category).unwrap() {
                proto::CostCategory::ActionCost => "ACTION_COST",
                proto::CostCategory::WasmHostCost => "WASM_HOST_COST",
            }
            .to_owned(),
            cost: match value.cost.unwrap().variant.unwrap() {
                proto::cost::Variant::ActionCost(v) => match proto::ActionCosts::try_from(v.value).unwrap() {
                    proto::ActionCosts::CreateAccount => "CREATE_ACCOUNT",
                    proto::ActionCosts::DeleteAccount => "DELETE_ACCOUNT",
                    proto::ActionCosts::DeployContract => "DEPLOY_CONTRACT",
                    proto::ActionCosts::FunctionCall => "FUNCTION_CALL",
                    proto::ActionCosts::Transfer => "TRANSFER",
                    proto::ActionCosts::Stake => "STAKE",
                    proto::ActionCosts::AddKey => "ADD_KEY",
                    proto::ActionCosts::DeleteKey => "DELETE_KEY",
                    proto::ActionCosts::ValueReturn => "VALUE_RETURN",
                    proto::ActionCosts::NewReceipt => "NEW_RECEIPT",
                    proto::ActionCosts::NewActionReceipt => "NEW_ACTION_RECEIPT",
                    proto::ActionCosts::NewDataReceiptBase => "NEW_DATA_RECEIPT_BASE",
                    proto::ActionCosts::NewDataReceiptByte => "NEW_DATA_RECEIPT_BYTE",
                    proto::ActionCosts::AddFullAccessKey => "ADD_FULL_ACCESS_KEY",
                    proto::ActionCosts::AddFunctionCallKeyBase => "ADD_FUNCTION_CALL_KEY_BASE",
                    proto::ActionCosts::AddFunctionCallKeyByte => "ADD_FUNCTION_CALL_KEY_BYTE",
                    proto::ActionCosts::DeployContractBase => "DEPLOY_CONTRACT_BASE",
                    proto::ActionCosts::DeployContractByte => "DEPLOY_CONTRACT_BYTE",
                    proto::ActionCosts::FunctionCallBase => "FUNCTION_CALL_BASE",
                    proto::ActionCosts::FunctionCallByte => "FUNCTION_CALL_BYTE",
                    proto::ActionCosts::Delegate => "DELEGATE",
                },
                proto::cost::Variant::ExtCost(v) => match proto::ExtCosts::try_from(v.value).unwrap() {
                    proto::ExtCosts::Base => "BASE",
                    proto::ExtCosts::ContractCompileBase => "CONTRACT_COMPILE_BASE",
                    proto::ExtCosts::ContractCompileBytes => "CONTRACT_COMPILE_BYTES",
                    proto::ExtCosts::ReadMemoryBase => "READ_MEMORY_BASE",
                    proto::ExtCosts::ReadMemoryByte => "READ_MEMORY_BYTE",
                    proto::ExtCosts::WriteMemoryBase => "WRITE_MEMORY_BASE",
                    proto::ExtCosts::WriteMemoryByte => "WRITE_MEMORY_BYTE",
                    proto::ExtCosts::ReadRegisterBase => "READ_REGISTER_BASE",
                    proto::ExtCosts::ReadRegisterByte => "READ_REGISTER_BYTE",
                    proto::ExtCosts::WriteRegisterBase => "WRITE_REGISTER_BASE",
                    proto::ExtCosts::WriteRegisterByte => "WRITE_REGISTER_BYTE",
                    proto::ExtCosts::Utf8DecodingBase => "UTF8_DECODING_BASE",
                    proto::ExtCosts::Utf8DecodingByte => "UTF8_DECODING_BYTE",
                    proto::ExtCosts::Utf16DecodingBase => "UTF16_DECODING_BASE",
                    proto::ExtCosts::Utf16DecodingByte => "UTF16_DECODING_BYTE",
                    proto::ExtCosts::Sha256Base => "SHA256_BASE",
                    proto::ExtCosts::Sha256Byte => "SHA256_BYTE",
                    proto::ExtCosts::Keccak256Base => "KECCAK256_BASE",
                    proto::ExtCosts::Keccak256Byte => "KECCAK256_BYTE",
                    proto::ExtCosts::Keccak512Base => "KECCAK512_BASE",
                    proto::ExtCosts::Keccak512Byte => "KECCAK512_BYTE",
                    proto::ExtCosts::Ripemd160Base => "RIPEMD160_BASE",
                    proto::ExtCosts::Ripemd160Block => "RIPEMD160_BLOCK",
                    proto::ExtCosts::EcrecoverBase => "ECRECOVER_BASE",
                    proto::ExtCosts::LogBase => "LOG_BASE",
                    proto::ExtCosts::LogByte => "LOG_BYTE",
                    proto::ExtCosts::StorageWriteBase => "STORAGE_WRITE_BASE",
                    proto::ExtCosts::StorageWriteKeyByte => "STORAGE_WRITE_KEY_BYTE",
                    proto::ExtCosts::StorageWriteValueByte => "STORAGE_WRITE_VALUE_BYTE",
                    proto::ExtCosts::StorageWriteEvictedByte => "STORAGE_WRITE_EVICTED_BYTE",
                    proto::ExtCosts::StorageReadBase => "STORAGE_READ_BASE",
                    proto::ExtCosts::StorageReadKeyByte => "STORAGE_READ_KEY_BYTE",
                    proto::ExtCosts::StorageReadValueByte => "STORAGE_READ_VALUE_BYTE",
                    proto::ExtCosts::StorageRemoveBase => "STORAGE_REMOVE_BASE",
                    proto::ExtCosts::StorageRemoveKeyByte => "STORAGE_REMOVE_KEY_BYTE",
                    proto::ExtCosts::StorageRemoveRetValueByte => "STORAGE_REMOVE_RET_VALUE_BYTE",
                    proto::ExtCosts::StorageHasKeyBase => "STORAGE_HAS_KEY_BASE",
                    proto::ExtCosts::StorageHasKeyByte => "STORAGE_HAS_KEY_BYTE",
                    proto::ExtCosts::StorageIterCreatePrefixBase => "STORAGE_ITER_CREATE_PREFIX_BASE",
                    proto::ExtCosts::StorageIterCreatePrefixByte => "STORAGE_ITER_CREATE_PREFIX_BYTE",
                    proto::ExtCosts::StorageIterCreateRangeBase => "STORAGE_ITER_CREATE_RANGE_BASE",
                    proto::ExtCosts::StorageIterCreateFromByte => "STORAGE_ITER_CREATE_FROM_BYTE",
                    proto::ExtCosts::StorageIterCreateToByte => "STORAGE_ITER_CREATE_TO_BYTE",
                    proto::ExtCosts::StorageIterNextBase => "STORAGE_ITER_NEXT_BASE",
                    proto::ExtCosts::StorageIterNextKeyByte => "STORAGE_ITER_NEXT_KEY_BYTE",
                    proto::ExtCosts::StorageIterNextValueByte => "STORAGE_ITER_NEXT_VALUE_BYTE",
                    proto::ExtCosts::TouchingTrieNode => "TOUCHING_TRIE_NODE",
                    proto::ExtCosts::ReadCachedTrieNode => "READ_CACHED_TRIE_NODE",
                    proto::ExtCosts::PromiseAndBase => "PROMISE_AND_BASE",
                    proto::ExtCosts::PromiseAndPerPromise => "PROMISE_AND_PER_PROMISE",
                    proto::ExtCosts::PromiseReturn => "PROMISE_RETURN",
                    proto::ExtCosts::ValidatorStakeBase => "VALIDATOR_STAKE_BASE",
                    proto::ExtCosts::ValidatorTotalStakeBase => "VALIDATOR_TOTAL_STAKE_BASE",
                    proto::ExtCosts::AltBn128G1MultiexpBase => "ALT_BN128_G1_MULTIEXP_BASE",
                    proto::ExtCosts::AltBn128G1MultiexpElement => "ALT_BN128_G1_MULTIEXP_ELEMENT",
                    proto::ExtCosts::AltBn128PairingCheckBase => "ALT_BN128_PAIRING_CHECK_BASE",
                    proto::ExtCosts::AltBn128PairingCheckElement => "ALT_BN128_PAIRING_CHECK_ELEMENT",
                    proto::ExtCosts::AltBn128G1SumBase => "ALT_BN128_G1_SUM_BASE",
                    proto::ExtCosts::AltBn128G1SumElement => "ALT_BN128_G1_SUM_ELEMENT",
                    proto::ExtCosts::Ed25519VerifyBase => "ED25519_VERIFY_BASE",
                    proto::ExtCosts::Ed25519VerifyByte => "ED25519_VERIFY_BYTE",
                    proto::ExtCosts::ContractLoadingBase => "CONTRACT_LOADING_BASE",
                    proto::ExtCosts::ContractLoadingBytes => "CONTRACT_LOADING_BYTES",
                },
                proto::cost::Variant::WasmInstruction(..) => "WASM_INSTRUCTION",
            }
            .to_owned(),
            gas_used: value.gas_used,
        }
    }
}

impl From<proto::ExecutionOutcomeWithIdView> for ExecutionOutcomeWithIdView {
    fn from(value: proto::ExecutionOutcomeWithIdView) -> Self {
        Self {
            proof: value.proof.into_iter().map(Into::into).collect(),
            block_hash: CryptoHash(value.h256_block_hash.try_into().unwrap()),
            id: CryptoHash(value.h256_id.try_into().unwrap()),
            outcome: value.outcome.unwrap().into(),
        }
    }
}

impl From<proto::MerklePathItem> for MerklePathItem {
    fn from(value: proto::MerklePathItem) -> Self {
        Self {
            hash: CryptoHash(value.h256_hash.try_into().unwrap()),
            direction: Direction::from(proto::Direction::try_from(value.direction).unwrap()),
        }
    }
}

impl From<proto::Direction> for Direction {
    fn from(value: proto::Direction) -> Self {
        match value {
            proto::Direction::Left => Direction::Left,
            proto::Direction::Right => Direction::Right,
        }
    }
}

impl From<proto::BlockHeaderView> for BlockView {
    fn from(value: proto::BlockHeaderView) -> Self {
        Self {
            author: AccountId::try_from(value.author).unwrap(),
            header: value.header.unwrap().into(),
        }
    }
}

impl From<proto::PartialBlockHeaderView> for BlockView {
    fn from(value: proto::PartialBlockHeaderView) -> Self {
        Self {
            author: AccountId::try_from(value.author).unwrap(),
            header: value.header.unwrap().into(),
        }
    }
}

impl From<proto::PartialIndexerBlockHeaderView> for IndexerBlockHeaderView {
    fn from(value: proto::PartialIndexerBlockHeaderView) -> Self {
        Self {
            height: value.height,
            prev_height: value.prev_height,
            epoch_id: CryptoHash(value.h256_epoch_id.try_into().unwrap()),
            next_epoch_id: CryptoHash(value.h256_next_epoch_id.try_into().unwrap()),
            hash: CryptoHash(value.h256_hash.try_into().unwrap()),
            prev_hash: CryptoHash(value.h256_prev_hash.try_into().unwrap()),
            prev_state_root: CryptoHash(value.h256_prev_state_root.try_into().unwrap()),
            chunk_receipts_root: CryptoHash(value.h256_chunk_receipts_root.try_into().unwrap()),
            chunk_headers_root: CryptoHash(value.h256_chunk_headers_root.try_into().unwrap()),
            chunk_tx_root: CryptoHash(value.h256_chunk_tx_root.try_into().unwrap()),
            outcome_root: CryptoHash(value.h256_outcome_root.try_into().unwrap()),
            chunks_included: value.chunks_included,
            challenges_root: CryptoHash(value.h256_challenges_root.try_into().unwrap()),
            timestamp: value.timestamp,
            timestamp_nanosec: value.timestamp_nanosec,
            random_value: CryptoHash(value.h256_random_value.try_into().unwrap()),
            validator_proposals: Vec::new(),
            chunk_mask: value.chunk_mask,
            gas_price: u128::from_be_bytes(value.u128_gas_price.try_into().unwrap()),
            block_ordinal: value.block_ordinal,
            total_supply: u128::from_be_bytes(value.u128_total_supply.try_into().unwrap()),
            challenges_result: Vec::new(),
            last_final_block: CryptoHash(value.h256_last_final_block.try_into().unwrap()),
            last_ds_final_block: CryptoHash(value.h256_last_ds_final_block.try_into().unwrap()),
            next_bp_hash: CryptoHash(value.h256_next_bp_hash.try_into().unwrap()),
            block_merkle_root: CryptoHash(value.h256_block_merkle_root.try_into().unwrap()),
            epoch_sync_data_hash: value
                .h256_epoch_sync_data_hash
                .map(|v| CryptoHash(v.try_into().unwrap())),
            approvals: Vec::new(),
            signature: value.signature.unwrap().into(),
            latest_protocol_version: value.latest_protocol_version,
        }
    }
}

impl From<proto::ChunkView> for ChunkView {
    fn from(value: proto::ChunkView) -> Self {
        Self {
            author: AccountId::try_from(value.author).unwrap(),
            header: value.header.unwrap().into(),
            transactions: value.transactions.into_iter().map(Into::into).collect(),
            receipts: value.receipts.into_iter().map(Into::into).collect(),
        }
    }
}

impl From<proto::ChunkHeaderView> for ChunkHeaderView {
    fn from(value: proto::ChunkHeaderView) -> Self {
        Self {
            chunk_hash: CryptoHash(value.h256_chunk_hash.try_into().unwrap()),
            prev_block_hash: CryptoHash(value.h256_prev_block_hash.try_into().unwrap()),
            outcome_root: CryptoHash(value.h256_outcome_root.try_into().unwrap()),
            prev_state_root: CryptoHash(value.h256_prev_state_root.try_into().unwrap()),
            encoded_merkle_root: CryptoHash(value.h256_encoded_merkle_root.try_into().unwrap()),
            encoded_length: value.encoded_length,
            height_created: value.height_created,
            height_included: value.height_included,
            shard_id: value.shard_id,
            gas_used: value.gas_used,
            gas_limit: value.gas_limit,
            validator_reward: u128::from_be_bytes(value.u128_validator_reward.try_into().unwrap()),
            balance_burnt: u128::from_be_bytes(value.u128_balance_burnt.try_into().unwrap()),
            outgoing_receipts_root: CryptoHash(value.h256_outgoing_receipts_root.try_into().unwrap()),
            tx_root: CryptoHash(value.h256_tx_root.try_into().unwrap()),
            validator_proposals: value.validator_proposals.into_iter().map(Into::into).collect(),
            signature: value.signature.unwrap().into(),
        }
    }
}

impl From<proto::TransactionWithOutcome> for TransactionWithOutcome {
    fn from(value: proto::TransactionWithOutcome) -> Self {
        Self {
            transaction: value.transaction.unwrap().into(),
            outcome: value.outcome.unwrap().into(),
        }
    }
}

impl From<proto::SignedTransactionView> for SignedTransactionView {
    fn from(value: proto::SignedTransactionView) -> Self {
        Self {
            signer_id: AccountId::try_from(value.signer_id).unwrap(),
            public_key: value.public_key.unwrap().into(),
            nonce: value.nonce,
            receiver_id: AccountId::try_from(value.receiver_id).unwrap(),
            actions: value.actions.into_iter().map(Into::into).collect(),
            signature: value.signature.unwrap().into(),
            hash: CryptoHash(value.h256_hash.try_into().unwrap()),
        }
    }
}

impl From<proto::Signature> for Signature {
    fn from(value: proto::Signature) -> Self {
        match value.variant.unwrap() {
            proto::signature::Variant::Ed25519(v) => {
                Signature::ED25519(ed25519_dalek::Signature::from_bytes(v.h512_value.as_slice()).unwrap())
            }
            proto::signature::Variant::Secp256k1(v) => {
                Signature::SECP256K1(From::<[u8; 65]>::from(v.h520_value.try_into().unwrap()))
            }
        }
    }
}

impl From<proto::ExecutionOutcomeWithOptionalReceipt> for ExecutionOutcomeWithOptionalReceipt {
    fn from(value: proto::ExecutionOutcomeWithOptionalReceipt) -> Self {
        Self {
            execution_outcome: value
                .execution_outcome
                .map(|v| ExecutionOutcomeWithIdView {
                    proof: v.proof.into_iter().map(Into::into).collect(),
                    block_hash: CryptoHash(v.h256_block_hash.try_into().unwrap()),
                    id: CryptoHash(v.h256_id.try_into().unwrap()),
                    outcome: v.outcome.unwrap().into(),
                })
                .unwrap(),
            receipt: value.receipt.map(|v| ReceiptView {
                predecessor_id: AccountId::try_from(v.predecessor_id).unwrap(),
                receiver_id: AccountId::try_from(v.receiver_id).unwrap(),
                receipt_id: CryptoHash(v.h256_receipt_id.try_into().unwrap()),
                receipt: v.receipt.unwrap().into(),
            }),
        }
    }
}

impl From<proto::IndexerBlockHeaderView> for IndexerBlockHeaderView {
    fn from(value: proto::IndexerBlockHeaderView) -> Self {
        Self {
            height: value.height,
            prev_height: value.prev_height,
            epoch_id: CryptoHash(value.h256_epoch_id.try_into().unwrap()),
            next_epoch_id: CryptoHash(value.h256_next_epoch_id.try_into().unwrap()),
            hash: CryptoHash(value.h256_hash.try_into().unwrap()),
            prev_hash: CryptoHash(value.h256_prev_hash.try_into().unwrap()),
            prev_state_root: CryptoHash(value.h256_prev_state_root.try_into().unwrap()),
            chunk_receipts_root: CryptoHash(value.h256_chunk_receipts_root.try_into().unwrap()),
            chunk_headers_root: CryptoHash(value.h256_chunk_headers_root.try_into().unwrap()),
            chunk_tx_root: CryptoHash(value.h256_chunk_tx_root.try_into().unwrap()),
            outcome_root: CryptoHash(value.h256_outcome_root.try_into().unwrap()),
            chunks_included: value.chunks_included,
            challenges_root: CryptoHash(value.h256_challenges_root.try_into().unwrap()),
            timestamp: value.timestamp,
            timestamp_nanosec: value.timestamp_nanosec,
            random_value: CryptoHash(value.h256_random_value.try_into().unwrap()),
            validator_proposals: value.validator_proposals.into_iter().map(Into::into).collect(),
            chunk_mask: value.chunk_mask,
            gas_price: u128::from_be_bytes(value.u128_gas_price.try_into().unwrap()),
            block_ordinal: value.block_ordinal,
            total_supply: u128::from_be_bytes(value.u128_total_supply.try_into().unwrap()),
            challenges_result: value.challenges_result.into_iter().map(Into::into).collect(),
            last_final_block: CryptoHash(value.h256_last_final_block.try_into().unwrap()),
            last_ds_final_block: CryptoHash(value.h256_last_ds_final_block.try_into().unwrap()),
            next_bp_hash: CryptoHash(value.h256_next_bp_hash.try_into().unwrap()),
            block_merkle_root: CryptoHash(value.h256_block_merkle_root.try_into().unwrap()),
            epoch_sync_data_hash: value
                .h256_epoch_sync_data_hash
                .map(|v| CryptoHash(v.try_into().unwrap())),
            approvals: value.approvals.into_iter().map(Into::into).collect(),
            signature: value.signature.unwrap().into(),
            latest_protocol_version: value.latest_protocol_version,
        }
    }
}

impl From<proto::ValidatorStakeView> for ValidatorStakeView {
    fn from(value: proto::ValidatorStakeView) -> Self {
        match value.variant.unwrap() {
            proto::validator_stake_view::Variant::V1(view) => ValidatorStakeView::V1(ValidatorStakeViewV1 {
                account_id: AccountId::try_from(view.account_id).unwrap(),
                public_key: view.public_key.unwrap().into(),
                stake: u128::from_be_bytes(view.u128_stake.try_into().unwrap()),
            }),
        }
    }
}

impl From<proto::SlashedValidator> for SlashedValidator {
    fn from(value: proto::SlashedValidator) -> Self {
        Self {
            account_id: AccountId::try_from(value.account_id).unwrap(),
            is_double_sign: value.is_double_sign,
        }
    }
}

impl From<proto::OptionalSignature> for Option<Signature> {
    fn from(value: proto::OptionalSignature) -> Self {
        value.value.map(Into::into)
    }
}

impl From<proto::ReceiptView> for ReceiptView {
    fn from(value: proto::ReceiptView) -> Self {
        Self {
            predecessor_id: AccountId::try_from(value.predecessor_id).unwrap(),
            receiver_id: AccountId::try_from(value.receiver_id).unwrap(),
            receipt_id: CryptoHash(value.h256_receipt_id.try_into().unwrap()),
            receipt: value.receipt.unwrap().into(),
        }
    }
}

impl From<proto::ReceiptEnumView> for ReceiptEnumView {
    fn from(value: proto::ReceiptEnumView) -> Self {
        match value.variant.unwrap() {
            proto::receipt_enum_view::Variant::Action(v) => ReceiptEnumView::Action {
                signer_id: AccountId::try_from(v.signer_id).unwrap(),
                signer_public_key: v.signer_public_key.unwrap().into(),
                gas_price: u128::from_be_bytes(v.u128_gas_price.try_into().unwrap()),
                output_data_receivers: v.output_data_receivers.into_iter().map(Into::into).collect(),
                input_data_ids: v
                    .h256_input_data_ids
                    .into_iter()
                    .map(|v| CryptoHash(v.try_into().unwrap()))
                    .collect(),
                actions: v.actions.into_iter().map(Into::into).collect(),
            },
            proto::receipt_enum_view::Variant::Data(v) => ReceiptEnumView::Data {
                data: v.data,
                data_id: CryptoHash(v.h256_data_id.try_into().unwrap()),
            },
        }
    }
}

impl From<proto::DataReceiverView> for DataReceiverView {
    fn from(value: proto::DataReceiverView) -> Self {
        Self {
            data_id: CryptoHash(value.h256_data_id.try_into().unwrap()),
            receiver_id: AccountId::try_from(value.receiver_id).unwrap(),
        }
    }
}

impl From<proto::action_view::delegate::delegate_action::NonDelegateAction>
    for near_primitives::delegate_action::NonDelegateAction
{
    fn from(value: proto::action_view::delegate::delegate_action::NonDelegateAction) -> Self {
        Self::try_from(match value.action.unwrap().variant.unwrap() {
            proto::action_view::Variant::CreateAccount(..) => Action::CreateAccount(CreateAccountAction {}),
            proto::action_view::Variant::DeployContract(v) => {
                Action::DeployContract(DeployContractAction { code: v.code })
            }
            proto::action_view::Variant::FunctionCall(v) => Action::FunctionCall(FunctionCallAction {
                method_name: v.method_name,
                args: v.args,
                gas: v.gas,
                deposit: u128::from_be_bytes(v.u128_deposit.try_into().unwrap()),
            }),
            proto::action_view::Variant::Transfer(v) => Action::Transfer(TransferAction {
                deposit: u128::from_be_bytes(v.u128_deposit.try_into().unwrap()),
            }),
            proto::action_view::Variant::Stake(v) => Action::Stake(StakeAction {
                stake: u128::from_be_bytes(v.u128_stake.try_into().unwrap()),
                public_key: v.public_key.unwrap().into(),
            }),
            proto::action_view::Variant::AddKey(v) => Action::AddKey(AddKeyAction {
                public_key: v.public_key.unwrap().into(),
                access_key: v.access_key.unwrap().into(),
            }),
            proto::action_view::Variant::DeleteKey(v) => Action::DeleteKey(DeleteKeyAction {
                public_key: v.public_key.unwrap().into(),
            }),
            proto::action_view::Variant::DeleteAccount(v) => Action::DeleteAccount(DeleteAccountAction {
                beneficiary_id: AccountId::try_from(v.beneficiary_id).unwrap(),
            }),
            proto::action_view::Variant::Delegate(..) => panic!("Non-delegate action cannot contain delegate action"),
        })
        .unwrap()
    }
}

impl From<proto::AccessKeyView> for near_primitives::account::AccessKey {
    fn from(value: proto::AccessKeyView) -> Self {
        Self {
            nonce: value.nonce,
            permission: value.permission.unwrap().into(),
        }
    }
}

impl From<proto::AccessKeyPermissionView> for AccessKeyPermission {
    fn from(value: proto::AccessKeyPermissionView) -> Self {
        match value.variant.unwrap() {
            access_key_permission_view::Variant::FunctionCall(v) => Self::FunctionCall(v.into()),
            access_key_permission_view::Variant::FullAccess(..) => Self::FullAccess,
        }
    }
}

impl From<proto::access_key_permission_view::FunctionCall> for FunctionCallPermission {
    fn from(value: proto::access_key_permission_view::FunctionCall) -> Self {
        Self {
            allowance: value.u128_allowance.map(|v| u128::from_be_bytes(v.try_into().unwrap())),
            receiver_id: value.receiver_id,
            method_names: value.method_names,
        }
    }
}

impl From<proto::action_view::delegate::DelegateAction> for near_primitives::delegate_action::DelegateAction {
    fn from(value: proto::action_view::delegate::DelegateAction) -> Self {
        Self {
            sender_id: AccountId::try_from(value.sender_id).unwrap(),
            receiver_id: AccountId::try_from(value.receiver_id).unwrap(),
            actions: value.actions.into_iter().map(Into::into).collect(),
            nonce: value.nonce,
            max_block_height: value.max_block_height,
            public_key: value.public_key.unwrap().into(),
        }
    }
}

impl From<proto::ActionView> for ActionView {
    fn from(value: proto::ActionView) -> Self {
        match value.variant.unwrap() {
            proto::action_view::Variant::CreateAccount(..) => ActionView::CreateAccount,
            proto::action_view::Variant::DeployContract(v) => ActionView::DeployContract { code: v.code },
            proto::action_view::Variant::FunctionCall(v) => ActionView::FunctionCall {
                method_name: v.method_name,
                args: v.args.into(),
                gas: v.gas,
                deposit: u128::from_be_bytes(v.u128_deposit.try_into().unwrap()),
            },
            proto::action_view::Variant::Transfer(v) => ActionView::Transfer {
                deposit: u128::from_be_bytes(v.u128_deposit.try_into().unwrap()),
            },
            proto::action_view::Variant::Stake(v) => ActionView::Stake {
                stake: u128::from_be_bytes(v.u128_stake.try_into().unwrap()),
                public_key: v.public_key.unwrap().into(),
            },
            proto::action_view::Variant::AddKey(v) => ActionView::AddKey {
                public_key: v.public_key.unwrap().into(),
                access_key: v.access_key.unwrap().into(),
            },
            proto::action_view::Variant::DeleteKey(v) => ActionView::DeleteKey {
                public_key: v.public_key.unwrap().into(),
            },
            proto::action_view::Variant::DeleteAccount(v) => ActionView::DeleteAccount {
                beneficiary_id: AccountId::try_from(v.beneficiary_id).unwrap(),
            },
            proto::action_view::Variant::Delegate(v) => ActionView::Delegate {
                delegate_action: v.delegate_action.unwrap().into(),
                signature: v.signature.unwrap().into(),
            },
        }
    }
}

impl From<proto::StateChangeWithCauseView> for StateChangeWithCauseView {
    fn from(value: proto::StateChangeWithCauseView) -> Self {
        Self {
            cause: value.cause.unwrap().into(),
            value: value.value.unwrap().into(),
        }
    }
}

impl From<proto::StateChangeCauseView> for StateChangeCauseView {
    fn from(value: proto::StateChangeCauseView) -> Self {
        match value.variant.unwrap() {
            proto::state_change_cause_view::Variant::NotWritableToDisk(..) => StateChangeCauseView::NotWritableToDisk,
            proto::state_change_cause_view::Variant::InitialState(..) => StateChangeCauseView::InitialState,
            proto::state_change_cause_view::Variant::TransactionProcessing(v) => {
                StateChangeCauseView::TransactionProcessing {
                    tx_hash: CryptoHash(v.h256_tx_hash.try_into().unwrap()),
                }
            }
            proto::state_change_cause_view::Variant::ActionReceiptProcessingStarted(v) => {
                StateChangeCauseView::ActionReceiptProcessingStarted {
                    receipt_hash: CryptoHash(v.h256_receipt_hash.try_into().unwrap()),
                }
            }
            proto::state_change_cause_view::Variant::ActionReceiptGasReward(v) => {
                StateChangeCauseView::ActionReceiptGasReward {
                    receipt_hash: CryptoHash(v.h256_receipt_hash.try_into().unwrap()),
                }
            }
            proto::state_change_cause_view::Variant::ReceiptProcessing(v) => StateChangeCauseView::ReceiptProcessing {
                receipt_hash: CryptoHash(v.h256_receipt_hash.try_into().unwrap()),
            },
            proto::state_change_cause_view::Variant::PostponedReceipt(v) => StateChangeCauseView::PostponedReceipt {
                receipt_hash: CryptoHash(v.h256_receipt_hash.try_into().unwrap()),
            },
            proto::state_change_cause_view::Variant::UpdatedDelayedReceipts(..) => {
                StateChangeCauseView::UpdatedDelayedReceipts
            }
            proto::state_change_cause_view::Variant::ValidatorAccountsUpdate(..) => {
                StateChangeCauseView::ValidatorAccountsUpdate
            }
            proto::state_change_cause_view::Variant::Migration(..) => StateChangeCauseView::Migration,
            proto::state_change_cause_view::Variant::Resharding(..) => StateChangeCauseView::Resharding,
        }
    }
}

impl From<proto::StateChangeValueView> for StateChangeValueView {
    fn from(value: proto::StateChangeValueView) -> Self {
        match value.variant.unwrap() {
            proto::state_change_value_view::Variant::AccountUpdate(v) => StateChangeValueView::AccountUpdate {
                account_id: AccountId::try_from(v.account_id).unwrap(),
                account: v.account.unwrap().into(),
            },
            proto::state_change_value_view::Variant::AccountDeletion(v) => StateChangeValueView::AccountDeletion {
                account_id: AccountId::try_from(v.account_id).unwrap(),
            },
            proto::state_change_value_view::Variant::AccessKeyUpdate(v) => StateChangeValueView::AccessKeyUpdate {
                account_id: AccountId::try_from(v.account_id).unwrap(),
                public_key: v.public_key.unwrap().into(),
                access_key: v.access_key.unwrap().into(),
            },
            proto::state_change_value_view::Variant::AccessKeyDeletion(v) => StateChangeValueView::AccessKeyDeletion {
                account_id: AccountId::try_from(v.account_id).unwrap(),
                public_key: v.public_key.unwrap().into(),
            },
            proto::state_change_value_view::Variant::DataUpdate(v) => StateChangeValueView::DataUpdate {
                account_id: AccountId::try_from(v.account_id).unwrap(),
                key: StoreKey::from(v.key),
                value: StoreValue::from(v.value),
            },
            proto::state_change_value_view::Variant::DataDeletion(v) => StateChangeValueView::DataDeletion {
                account_id: AccountId::try_from(v.account_id).unwrap(),
                key: StoreKey::from(v.key),
            },
            proto::state_change_value_view::Variant::ContractCodeUpdate(v) => {
                StateChangeValueView::ContractCodeUpdate {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                    code: v.code,
                }
            }
            proto::state_change_value_view::Variant::ContractCodeDeletion(v) => {
                StateChangeValueView::ContractCodeDeletion {
                    account_id: AccountId::try_from(v.account_id).unwrap(),
                }
            }
        }
    }
}

impl From<proto::PublicKey> for PublicKey {
    fn from(value: proto::PublicKey) -> Self {
        match value.variant.unwrap() {
            proto::public_key::Variant::Secp256k1(v) => {
                PublicKey::SECP256K1(Secp256K1PublicKey::try_from(v.h512_value.as_slice()).unwrap())
            }
            proto::public_key::Variant::Ed25519(v) => {
                PublicKey::ED25519(ED25519PublicKey::try_from(v.h256_value.as_slice()).unwrap())
            }
        }
    }
}

impl From<proto::AccessKeyView> for AccessKeyView {
    fn from(value: proto::AccessKeyView) -> Self {
        Self {
            nonce: value.nonce,
            permission: value.permission.unwrap().into(),
        }
    }
}

impl From<proto::AccessKeyPermissionView> for AccessKeyPermissionView {
    fn from(value: proto::AccessKeyPermissionView) -> Self {
        match value.variant.unwrap() {
            proto::access_key_permission_view::Variant::FunctionCall(v) => AccessKeyPermissionView::FunctionCall {
                allowance: v.u128_allowance.map(|v| u128::from_be_bytes(v.try_into().unwrap())),
                receiver_id: v.receiver_id,
                method_names: v.method_names,
            },
            proto::access_key_permission_view::Variant::FullAccess(..) => AccessKeyPermissionView::FullAccess,
        }
    }
}

impl From<proto::AccountView> for AccountView {
    fn from(value: proto::AccountView) -> Self {
        Self {
            amount: u128::from_be_bytes(value.u128_amount.try_into().unwrap()),
            locked: u128::from_be_bytes(value.u128_locked.try_into().unwrap()),
            code_hash: CryptoHash(value.h256_code_hash.try_into().unwrap()),
            storage_usage: value.storage_usage,
            storage_paid_at: value.storage_paid_at,
        }
    }
}
