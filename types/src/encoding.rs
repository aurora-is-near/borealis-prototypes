use crate::proto;
use borealis_types::message::Message;
use borealis_types::payloads::events::{
    BlockView, ChunkHeaderView, ChunkView, ExecutionOutcomeWithOptionalReceipt, ExecutionOutcomeWithReceipt,
    IndexerBlockHeaderView, Shard, TransactionWithOutcome,
};
use borealis_types::payloads::NEARBlock;
use borealis_types::types::views::{
    ExecutionMetadataView, ExecutionOutcomeView, ExecutionOutcomeWithIdView, ExecutionStatusView, ReceiptView,
    StateChangeCauseView, StateChangeValueView, StateChangeWithCauseView,
};
use near_crypto::{PublicKey, Signature};
use near_primitives::challenge::SlashedValidator;
use near_primitives::errors::{
    ActionErrorKind, ActionsValidationError, InvalidAccessKeyError, InvalidTxError, ReceiptValidationError,
    TxExecutionError,
};
use near_primitives::merkle::{Direction, MerklePathItem};
use near_primitives::views::validator_stake_view::ValidatorStakeView;
use near_primitives::views::{
    AccessKeyPermissionView, AccessKeyView, AccountView, ActionView, CostGasUsed, DataReceiverView, ReceiptEnumView,
    SignedTransactionView,
};
use near_vm_errors::{CompilationError, FunctionCallErrorSer, HostError, MethodResolveError, PrepareError, WasmTrap};
use std::iter::once;

impl From<Message<NEARBlock>> for proto::Messages {
    fn from(value: Message<NEARBlock>) -> Self {
        let height = value.payload.block.header.height;

        value
            .payload
            .shards
            .into_iter()
            .map(proto::BlockShard::from)
            .map(|v| proto::Message {
                version: value.version as u32,
                id: format!("{}.{}", height, v.shard_id),
                payload: Some(proto::message::Payload::NearBlockShard(v)),
            })
            .chain(once(proto::Message {
                version: value.version as u32,
                id: height.to_string(),
                payload: Some(proto::message::Payload::NearBlockHeader(value.payload.block.into())),
            }))
            .collect::<Vec<proto::Message>>()
            .into()
    }
}

impl From<Shard> for proto::BlockShard {
    fn from(value: Shard) -> Self {
        Self {
            shard_id: value.shard_id,
            chunk: value.chunk.map(Into::into),
            receipt_execution_outcomes: value.receipt_execution_outcomes.into_iter().map(Into::into).collect(),
            state_changes: value.state_changes.into_iter().map(Into::into).collect(),
        }
    }
}

impl From<ExecutionOutcomeWithReceipt> for proto::ExecutionOutcomeWithReceipt {
    fn from(value: ExecutionOutcomeWithReceipt) -> Self {
        Self {
            execution_outcome: Some(value.execution_outcome.into()),
            receipt: Some(value.receipt.into()),
        }
    }
}

impl From<ExecutionOutcomeView> for proto::ExecutionOutcomeView {
    fn from(value: ExecutionOutcomeView) -> Self {
        Self {
            logs: value.logs,
            h256_receipt_ids: value.receipt_ids.into_iter().map(|v| v.0.to_vec()).collect(),
            gas_burnt: value.gas_burnt,
            u128_tokens_burnt: value.tokens_burnt.to_be_bytes().to_vec(),
            executor_id: value.executor_id.into(),
            status: Some(value.status.into()),
            metadata: Some(value.metadata.into()),
        }
    }
}

impl From<ExecutionStatusView> for proto::ExecutionStatusView {
    fn from(value: ExecutionStatusView) -> Self {
        Self {
            variant: Some(match value {
                ExecutionStatusView::Unknown => {
                    proto::execution_status_view::Variant::Unknown(proto::execution_status_view::Unknown {})
                }
                ExecutionStatusView::Failure(error) => {
                    proto::execution_status_view::Variant::Failure(proto::execution_status_view::Failure {
                        error: Some(error.into()),
                    })
                }
                ExecutionStatusView::SuccessValue(value) => {
                    proto::execution_status_view::Variant::SuccessValue(proto::execution_status_view::SuccessValue {
                        value,
                    })
                }
                ExecutionStatusView::SuccessReceiptId(receipt_hash) => {
                    proto::execution_status_view::Variant::SuccessReceiptId(
                        proto::execution_status_view::SuccessReceiptId {
                            h256_receipt_hash: receipt_hash.0.to_vec(),
                        },
                    )
                }
            }),
        }
    }
}

impl From<ActionErrorKind> for proto::tx_execution_error::action_error::ActionErrorKind {
    fn from(value: ActionErrorKind) -> Self {
        Self {
            variant: Some(match value {
                ActionErrorKind::AccountAlreadyExists { account_id } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::AccountAlreadyExists(
                        proto::tx_execution_error::action_error::action_error_kind::AccountAlreadyExists {
                            account_id: account_id.to_string(),
                        }
                    )
                }
                ActionErrorKind::AccountDoesNotExist { account_id } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::AccountDoesNotExist(
                        proto::tx_execution_error::action_error::action_error_kind::AccountDoesNotExist {
                            account_id: account_id.to_string(),
                        }
                    )
                }
                ActionErrorKind::CreateAccountOnlyByRegistrar { account_id,registrar_account_id, predecessor_id  } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::CreateAccountOnlyByRegistrar(
                        proto::tx_execution_error::action_error::action_error_kind::CreateAccountOnlyByRegistrar {
                            account_id: account_id.to_string(),
                            registrar_account_id: registrar_account_id.to_string(),
                            predecessor_id: predecessor_id.to_string(),
                        }
                    )
                }
                ActionErrorKind::CreateAccountNotAllowed { account_id, predecessor_id } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::CreateAccountNotAllowed(
                        proto::tx_execution_error::action_error::action_error_kind::CreateAccountNotAllowed {
                            account_id: account_id.to_string(),
                            predecessor_id: predecessor_id.to_string(),
                        }
                    )
                }
                ActionErrorKind::ActorNoPermission { account_id, actor_id } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::ActorNoPermission(
                        proto::tx_execution_error::action_error::action_error_kind::ActorNoPermission {
                            account_id: account_id.to_string(),
                            actor_id: actor_id.to_string(),
                        }
                    )
                }
                ActionErrorKind::DeleteKeyDoesNotExist { account_id, public_key } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::DeleteKeyDoesNotExist(
                        proto::tx_execution_error::action_error::action_error_kind::DeleteKeyDoesNotExist {
                            account_id: account_id.to_string(),
                            public_key: Some(public_key.into()),
                        }
                    )
                }
                ActionErrorKind::AddKeyAlreadyExists { account_id, public_key } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::AddKeyAlreadyExists(
                        proto::tx_execution_error::action_error::action_error_kind::AddKeyAlreadyExists {
                            account_id: account_id.to_string(),
                            public_key: Some(public_key.into()),
                        }
                    )
                }
                ActionErrorKind::DeleteAccountStaking { account_id } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::DeleteAccountStaking(
                        proto::tx_execution_error::action_error::action_error_kind::DeleteAccountStaking {
                            account_id: account_id.to_string(),
                        }
                    )
                }
                ActionErrorKind::LackBalanceForState { account_id, amount } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::LackBalanceForState(
                        proto::tx_execution_error::action_error::action_error_kind::LackBalanceForState {
                            account_id: account_id.to_string(),
                            u128_amount: amount.to_be_bytes().to_vec(),
                        }
                    )
                }
                ActionErrorKind::TriesToUnstake { account_id } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::TriesToUnstake(
                        proto::tx_execution_error::action_error::action_error_kind::TriesToUnstake {
                            account_id: account_id.to_string(),
                        }
                    )
                }
                ActionErrorKind::TriesToStake { account_id, stake, balance, locked } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::TriesToStake(
                        proto::tx_execution_error::action_error::action_error_kind::TriesToStake {
                            account_id: account_id.to_string(),
                            u128_stake: stake.to_be_bytes().to_vec(),
                            u128_balance: balance.to_be_bytes().to_vec(),
                            u128_locked: locked.to_be_bytes().to_vec(),
                        }
                    )
                }
                ActionErrorKind::InsufficientStake { account_id, stake, minimum_stake } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::InsufficientStake(
                        proto::tx_execution_error::action_error::action_error_kind::InsufficientStake {
                            account_id: account_id.to_string(),
                            u128_stake: stake.to_be_bytes().to_vec(),
                            u128_minimum_stake: minimum_stake.to_be_bytes().to_vec(),
                        }
                    )
                }
                ActionErrorKind::FunctionCallError(error) => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::FunctionCallError(
                        proto::tx_execution_error::action_error::action_error_kind::FunctionCallError {
                            error: Some(error.into()),
                        }
                    )
                }
                ActionErrorKind::NewReceiptValidationError(error) => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::NewReceiptValidationError(
                        proto::tx_execution_error::action_error::action_error_kind::NewReceiptValidationError {
                            error: Some(error.into()),
                        }
                    )
                }
                ActionErrorKind::OnlyImplicitAccountCreationAllowed { account_id } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::OnlyImplicitAccountCreationAllowed(
                        proto::tx_execution_error::action_error::action_error_kind::OnlyImplicitAccountCreationAllowed {
                            account_id: account_id.to_string(),
                        }
                    )
                }
                ActionErrorKind::DeleteAccountWithLargeState { account_id } => {
                    proto::tx_execution_error::action_error::action_error_kind::Variant::DeleteAccountWithLargeState(
                        proto::tx_execution_error::action_error::action_error_kind::DeleteAccountWithLargeState {
                            account_id: account_id.to_string(),
                        }
                    )
                }
            }),
        }
    }
}

impl From<ReceiptValidationError> for proto::ReceiptValidationError {
    fn from(value: ReceiptValidationError) -> Self {
        Self {
            variant: Some(match value {
                ReceiptValidationError::InvalidPredecessorId { account_id } => {
                    proto::receipt_validation_error::Variant::InvalidPredecessorId(
                        proto::receipt_validation_error::InvalidPredecessorId { account_id },
                    )
                }
                ReceiptValidationError::InvalidReceiverId { account_id } => {
                    proto::receipt_validation_error::Variant::InvalidReceiverId(
                        proto::receipt_validation_error::InvalidReceiverId { account_id },
                    )
                }
                ReceiptValidationError::InvalidSignerId { account_id } => {
                    proto::receipt_validation_error::Variant::InvalidSignerId(
                        proto::receipt_validation_error::InvalidSignerId { account_id },
                    )
                }
                ReceiptValidationError::InvalidDataReceiverId { account_id } => {
                    proto::receipt_validation_error::Variant::InvalidDataReceiverId(
                        proto::receipt_validation_error::InvalidDataReceiverId { account_id },
                    )
                }
                ReceiptValidationError::ReturnedValueLengthExceeded { length, limit } => {
                    proto::receipt_validation_error::Variant::ReturnedValueLengthExceeded(
                        proto::receipt_validation_error::ReturnedValueLengthExceeded { length, limit },
                    )
                }
                ReceiptValidationError::NumberInputDataDependenciesExceeded {
                    number_of_input_data_dependencies,
                    limit,
                } => proto::receipt_validation_error::Variant::NumberInputDataDependenciesExceeded(
                    proto::receipt_validation_error::NumberInputDataDependenciesExceeded {
                        number_of_input_data_dependencies,
                        limit,
                    },
                ),
                ReceiptValidationError::ActionsValidation(error) => {
                    proto::receipt_validation_error::Variant::ActionsValidation(
                        proto::receipt_validation_error::ActionsValidation {
                            error: Some(error.into()),
                        },
                    )
                }
            }),
        }
    }
}

impl From<ActionsValidationError> for proto::ActionsValidationError {
    fn from(value: ActionsValidationError) -> Self {
        Self {
            variant: Some(match value {
                ActionsValidationError::DeleteActionMustBeFinal => {
                    proto::actions_validation_error::Variant::DeleteActionMustBeFinal(
                        proto::actions_validation_error::DeleteActionMustBeFinal {},
                    )
                }
                ActionsValidationError::TotalPrepaidGasExceeded {
                    total_prepaid_gas,
                    limit,
                } => proto::actions_validation_error::Variant::TotalPrepaidGasExceeded(
                    proto::actions_validation_error::TotalPrepaidGasExceeded {
                        total_prepaid_gas,
                        limit,
                    },
                ),
                ActionsValidationError::TotalNumberOfActionsExceeded {
                    total_number_of_actions,
                    limit,
                } => proto::actions_validation_error::Variant::TotalNumberOfActionsExceeded(
                    proto::actions_validation_error::TotalNumberOfActionsExceeded {
                        total_number_of_bytes: total_number_of_actions,
                        limit,
                    },
                ),
                ActionsValidationError::AddKeyMethodNamesNumberOfBytesExceeded {
                    total_number_of_bytes,
                    limit,
                } => proto::actions_validation_error::Variant::AddKeyMethodNamesNumberOfBytesExceeded(
                    proto::actions_validation_error::AddKeyMethodNamesNumberOfBytesExceeded {
                        total_number_of_bytes,
                        limit,
                    },
                ),
                ActionsValidationError::AddKeyMethodNameLengthExceeded { length, limit } => {
                    proto::actions_validation_error::Variant::AddKeyMethodNameLengthExceeded(
                        proto::actions_validation_error::AddKeyMethodNameLengthExceeded { length, limit },
                    )
                }
                ActionsValidationError::IntegerOverflow => proto::actions_validation_error::Variant::IntegerOverflow(
                    proto::actions_validation_error::IntegerOverflow {},
                ),
                ActionsValidationError::InvalidAccountId { account_id } => {
                    proto::actions_validation_error::Variant::InvalidAccountId(
                        proto::actions_validation_error::InvalidAccountId { account_id },
                    )
                }
                ActionsValidationError::ContractSizeExceeded { size, limit } => {
                    proto::actions_validation_error::Variant::ContractSizeExceeded(
                        proto::actions_validation_error::ContractSizeExceeded { size, limit },
                    )
                }
                ActionsValidationError::FunctionCallMethodNameLengthExceeded { length, limit } => {
                    proto::actions_validation_error::Variant::FunctionCallMethodNameLengthExceeded(
                        proto::actions_validation_error::FunctionCallMethodNameLengthExceeded { length, limit },
                    )
                }
                ActionsValidationError::FunctionCallArgumentsLengthExceeded { length, limit } => {
                    proto::actions_validation_error::Variant::FunctionCallArgumentsLengthExceeded(
                        proto::actions_validation_error::FunctionCallArgumentsLengthExceeded { length, limit },
                    )
                }
                ActionsValidationError::UnsuitableStakingKey { public_key } => {
                    proto::actions_validation_error::Variant::UnsuitableStakingKey(
                        proto::actions_validation_error::UnsuitableStakingKey {
                            public_key: Some(public_key.into()),
                        },
                    )
                }
                ActionsValidationError::FunctionCallZeroAttachedGas => {
                    proto::actions_validation_error::Variant::FunctionCallZeroAttachedGas(
                        proto::actions_validation_error::FunctionCallZeroAttachedGas {},
                    )
                }
            }),
        }
    }
}

impl From<FunctionCallErrorSer> for proto::FunctionCallErrorSer {
    fn from(value: FunctionCallErrorSer) -> Self {
        Self {
            variant: Some(match value {
                FunctionCallErrorSer::CompilationError(error) => {
                    proto::function_call_error_ser::Variant::CompilationError(
                        proto::function_call_error_ser::CompilationError {
                            error: Some(error.into()),
                        },
                    )
                }
                FunctionCallErrorSer::LinkError { msg } => {
                    proto::function_call_error_ser::Variant::LinkError(proto::function_call_error_ser::LinkError {
                        msg,
                    })
                }
                FunctionCallErrorSer::MethodResolveError(error) => {
                    proto::function_call_error_ser::Variant::MethodResolveError(
                        proto::function_call_error_ser::MethodResolveError {
                            error: proto::MethodResolveError::from(error) as i32,
                        },
                    )
                }
                FunctionCallErrorSer::WasmTrap(error) => {
                    proto::function_call_error_ser::Variant::WasmTrap(proto::function_call_error_ser::WasmTrap {
                        error: proto::WasmTrap::from(error) as i32,
                    })
                }
                FunctionCallErrorSer::WasmUnknownError => proto::function_call_error_ser::Variant::WasmUnknownError(
                    proto::function_call_error_ser::WasmUnknownError {},
                ),
                FunctionCallErrorSer::HostError(error) => {
                    proto::function_call_error_ser::Variant::HostError(proto::function_call_error_ser::HostError {
                        error: Some(error.into()),
                    })
                }
                FunctionCallErrorSer::_EVMError => panic!("Deprecated error _EVMError"),
                FunctionCallErrorSer::ExecutionError(message) => {
                    proto::function_call_error_ser::Variant::ExecutionError(
                        proto::function_call_error_ser::ExecutionError { message },
                    )
                }
            }),
        }
    }
}

impl From<CompilationError> for proto::CompilationError {
    fn from(value: CompilationError) -> Self {
        Self {
            variant: Some(match value {
                CompilationError::CodeDoesNotExist { account_id } => {
                    proto::compilation_error::Variant::CodeDoesNotExist(proto::compilation_error::CodeDoesNotExist {
                        account_id: account_id.to_string(),
                    })
                }
                CompilationError::PrepareError(error) => {
                    proto::compilation_error::Variant::PrepareError(proto::compilation_error::PrepareError {
                        error: proto::PrepareError::from(error) as i32,
                    })
                }
                CompilationError::WasmerCompileError { msg } => proto::compilation_error::Variant::WasmerCompileError(
                    proto::compilation_error::WasmerCompileError { msg },
                ),
                CompilationError::UnsupportedCompiler { msg } => {
                    proto::compilation_error::Variant::UnsupportedCompiler(
                        proto::compilation_error::UnsupportedCompiler { msg },
                    )
                }
            }),
        }
    }
}

impl From<MethodResolveError> for proto::MethodResolveError {
    fn from(value: MethodResolveError) -> Self {
        match value {
            MethodResolveError::MethodEmptyName => proto::MethodResolveError::MethodEmptyName,
            MethodResolveError::MethodNotFound => proto::MethodResolveError::MethodNotFound,
            MethodResolveError::MethodInvalidSignature => proto::MethodResolveError::MethodInvalidSignature,
        }
    }
}

impl From<HostError> for proto::HostError {
    fn from(value: HostError) -> Self {
        Self {
            variant: Some(match value {
                HostError::BadUTF16 => proto::host_error::Variant::BadUtf16(proto::host_error::BadUtf16 {}),
                HostError::BadUTF8 => proto::host_error::Variant::BadUtf8(proto::host_error::BadUtf8 {}),
                HostError::GasExceeded => proto::host_error::Variant::GasExceeded(proto::host_error::GasExceeded {}),
                HostError::GasLimitExceeded => {
                    proto::host_error::Variant::GasLimitExceeded(proto::host_error::GasLimitExceeded {})
                }
                HostError::BalanceExceeded => {
                    proto::host_error::Variant::BalanceExceeded(proto::host_error::BalanceExceeded {})
                }
                HostError::EmptyMethodName => {
                    proto::host_error::Variant::EmptyMethodName(proto::host_error::EmptyMethodName {})
                }
                HostError::GuestPanic { panic_msg } => {
                    proto::host_error::Variant::GuestPanic(proto::host_error::GuestPanic { panic_msg })
                }
                HostError::IntegerOverflow => {
                    proto::host_error::Variant::IntegerOverflow(proto::host_error::IntegerOverflow {})
                }
                HostError::InvalidPromiseIndex { promise_idx } => {
                    proto::host_error::Variant::InvalidPromiseIndex(proto::host_error::InvalidPromiseIndex {
                        promise_idx,
                    })
                }
                HostError::CannotAppendActionToJointPromise => {
                    proto::host_error::Variant::CannotAppendActionToJointPromise(
                        proto::host_error::CannotAppendActionToJointPromise {},
                    )
                }
                HostError::CannotReturnJointPromise => {
                    proto::host_error::Variant::CannotReturnJointPromise(proto::host_error::CannotReturnJointPromise {})
                }
                HostError::InvalidPromiseResultIndex { result_idx } => {
                    proto::host_error::Variant::InvalidPromiseResultIndex(
                        proto::host_error::InvalidPromiseResultIndex { result_idx },
                    )
                }
                HostError::InvalidRegisterId { register_id } => {
                    proto::host_error::Variant::InvalidRegisterId(proto::host_error::InvalidRegisterId { register_id })
                }
                HostError::IteratorWasInvalidated { iterator_index } => {
                    proto::host_error::Variant::IteratorWasInvalidated(proto::host_error::IteratorWasInvalidated {
                        iterator_index,
                    })
                }
                HostError::MemoryAccessViolation => {
                    proto::host_error::Variant::MemoryAccessViolation(proto::host_error::MemoryAccessViolation {})
                }
                HostError::InvalidReceiptIndex { receipt_index } => {
                    proto::host_error::Variant::InvalidReceiptIndex(proto::host_error::InvalidReceiptIndex {
                        receipt_index,
                    })
                }
                HostError::InvalidIteratorIndex { iterator_index } => {
                    proto::host_error::Variant::InvalidIteratorIndex(proto::host_error::InvalidIteratorIndex {
                        iterator_index,
                    })
                }
                HostError::InvalidAccountId => {
                    proto::host_error::Variant::InvalidAccountId(proto::host_error::InvalidAccountId {})
                }
                HostError::InvalidMethodName => {
                    proto::host_error::Variant::InvalidMethodName(proto::host_error::InvalidMethodName {})
                }
                HostError::InvalidPublicKey => {
                    proto::host_error::Variant::InvalidPublicKey(proto::host_error::InvalidPublicKey {})
                }
                HostError::ProhibitedInView { method_name } => {
                    proto::host_error::Variant::ProhibitedInView(proto::host_error::ProhibitedInView { method_name })
                }
                HostError::NumberOfLogsExceeded { limit } => {
                    proto::host_error::Variant::NumberOfLogsExceeded(proto::host_error::NumberOfLogsExceeded { limit })
                }
                HostError::KeyLengthExceeded { length, limit } => {
                    proto::host_error::Variant::KeyLengthExceeded(proto::host_error::KeyLengthExceeded {
                        length,
                        limit,
                    })
                }
                HostError::ValueLengthExceeded { length, limit } => {
                    proto::host_error::Variant::ValueLengthExceeded(proto::host_error::ValueLengthExceeded {
                        length,
                        limit,
                    })
                }
                HostError::TotalLogLengthExceeded { length, limit } => {
                    proto::host_error::Variant::TotalLogLengthExceeded(proto::host_error::TotalLogLengthExceeded {
                        length,
                        limit,
                    })
                }
                HostError::NumberPromisesExceeded {
                    number_of_promises,
                    limit,
                } => proto::host_error::Variant::NumberPromisesExceeded(proto::host_error::NumberPromisesExceeded {
                    number_of_promises,
                    limit,
                }),
                HostError::NumberInputDataDependenciesExceeded {
                    number_of_input_data_dependencies,
                    limit,
                } => proto::host_error::Variant::NumberInputDataDependenciesExceeded(
                    proto::host_error::NumberInputDataDependenciesExceeded {
                        number_of_input_data_dependencies,
                        limit,
                    },
                ),
                HostError::ReturnedValueLengthExceeded { length, limit } => {
                    proto::host_error::Variant::ReturnedValueLengthExceeded(
                        proto::host_error::ReturnedValueLengthExceeded { length, limit },
                    )
                }
                HostError::ContractSizeExceeded { size, limit } => {
                    proto::host_error::Variant::ContractSizeExceeded(proto::host_error::ContractSizeExceeded {
                        size,
                        limit,
                    })
                }
                HostError::Deprecated { method_name } => {
                    proto::host_error::Variant::Deprecated(proto::host_error::Deprecated { method_name })
                }
                HostError::ECRecoverError { msg } => {
                    proto::host_error::Variant::EcrecoverError(proto::host_error::EcRecoverError { msg })
                }
                HostError::AltBn128InvalidInput { msg } => {
                    proto::host_error::Variant::AltBn128InvalidInput(proto::host_error::AltBn128InvalidInput { msg })
                }
            }),
        }
    }
}

impl From<PrepareError> for proto::PrepareError {
    fn from(value: PrepareError) -> Self {
        match value {
            PrepareError::Serialization => proto::PrepareError::Serialization,
            PrepareError::Deserialization => proto::PrepareError::Deserialization,
            PrepareError::InternalMemoryDeclared => proto::PrepareError::InternalMemoryDeclared,
            PrepareError::GasInstrumentation => proto::PrepareError::GasInstrumentation,
            PrepareError::StackHeightInstrumentation => proto::PrepareError::StackHeightInstrumentation,
            PrepareError::Instantiate => proto::PrepareError::Instantiate,
            PrepareError::Memory => proto::PrepareError::Memory,
            PrepareError::TooManyFunctions => proto::PrepareError::TooManyFunctions,
            PrepareError::TooManyLocals => proto::PrepareError::TooManyLocals,
        }
    }
}

impl From<WasmTrap> for proto::WasmTrap {
    fn from(value: WasmTrap) -> Self {
        match value {
            WasmTrap::Unreachable => proto::WasmTrap::Unreachable,
            WasmTrap::IncorrectCallIndirectSignature => proto::WasmTrap::IncorrectCallIndirectSignature,
            WasmTrap::MemoryOutOfBounds => proto::WasmTrap::MemoryOutOfBounds,
            WasmTrap::CallIndirectOOB => proto::WasmTrap::CallIndirectOob,
            WasmTrap::IllegalArithmetic => proto::WasmTrap::IllegalArithmetic,
            WasmTrap::MisalignedAtomicAccess => proto::WasmTrap::MisalignedAtomicAccess,
            WasmTrap::IndirectCallToNull => proto::WasmTrap::IndirectCallToNull,
            WasmTrap::StackOverflow => proto::WasmTrap::StackOverflow,
            WasmTrap::GenericTrap => proto::WasmTrap::GenericTrap,
        }
    }
}

impl From<InvalidAccessKeyError> for proto::tx_execution_error::invalid_tx_error::InvalidAccessKeyError {
    fn from(value: InvalidAccessKeyError) -> Self {
        Self {
            variant: Some(match value {
                InvalidAccessKeyError::AccessKeyNotFound { public_key, account_id } => {
                    proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::Variant::AccessKeyNotFound(
                        proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::AccessKeyNotFound {
                            account_id: account_id.to_string(),
                            public_key: Some(public_key.into()),
                        }
                    )
                }
                InvalidAccessKeyError::ReceiverMismatch { tx_receiver, ak_receiver } => {
                    proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::Variant::ReceiverMismatch(
                        proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::ReceiverMismatch {
                            tx_receiver: tx_receiver.to_string(),
                            ak_receiver,
                        }
                    )
                }
                InvalidAccessKeyError::MethodNameMismatch { method_name } => {
                    proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::Variant::MethodNameMismatch(
                        proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::MethodNameMismatch {
                            method_name,
                        }
                    )
                }
                InvalidAccessKeyError::RequiresFullAccess => {
                    proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::Variant::RequiresFullAccess(
                        proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::RequiresFullAccess {}
                    )
                }
                InvalidAccessKeyError::NotEnoughAllowance { account_id, public_key, allowance, cost } => {
                    proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::Variant::NotEnoughAllowance(
                        proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::NotEnoughAllowance {
                            account_id: account_id.to_string(),
                            public_key: Some(public_key.into()),
                            u128_allowance: allowance.to_be_bytes().to_vec(),
                            u128_cost: cost.to_be_bytes().to_vec(),
                        }
                    )
                }
                InvalidAccessKeyError::DepositWithFunctionCall => {
                    proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::Variant::DepositWithFunctionCall(
                        proto::tx_execution_error::invalid_tx_error::invalid_access_key_error::DepositWithFunctionCall {}
                    )
                }
            }),
        }
    }
}

impl From<ActionsValidationError> for proto::tx_execution_error::invalid_tx_error::ActionsValidation {
    fn from(value: ActionsValidationError) -> Self {
        Self {
            variant: Some(match value {
                ActionsValidationError::DeleteActionMustBeFinal => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::DeleteActionMustBeFinal(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::DeleteActionMustBeFinal {}
                    )
                }
                ActionsValidationError::TotalPrepaidGasExceeded { total_prepaid_gas, limit } => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::TotalPrepaidGasExceeded(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::TotalPrepaidGasExceeded { total_prepaid_gas, limit }
                    )
                }
                ActionsValidationError::TotalNumberOfActionsExceeded { total_number_of_actions, limit } => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::TotalNumberOfActionsExceeded(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::TotalNumberOfActionsExceeded { total_number_of_actions, limit }
                    )
                }
                ActionsValidationError::AddKeyMethodNamesNumberOfBytesExceeded { total_number_of_bytes, limit } => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::AddKeyMethodNamesNumberOfBytesExceeded(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::AddKeyMethodNamesNumberOfBytesExceeded { total_number_of_bytes, limit }
                    )
                }
                ActionsValidationError::AddKeyMethodNameLengthExceeded { length, limit } => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::AddKeyMethodNameLengthExceeded(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::AddKeyMethodNameLengthExceeded { length, limit }
                    )
                }
                ActionsValidationError::IntegerOverflow => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::IntegerOverflow(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::IntegerOverflow {}
                    )
                }
                ActionsValidationError::InvalidAccountId { account_id } => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::InvalidAccountId(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::InvalidAccountId { account_id }
                    )
                }
                ActionsValidationError::ContractSizeExceeded { size, limit } => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::ContractSizeExceeded(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::ContractSizeExceeded { size, limit }
                    )
                }
                ActionsValidationError::FunctionCallMethodNameLengthExceeded { length, limit } => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::FunctionCallMethodNameLengthExceeded(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::FunctionCallMethodNameLengthExceeded { length, limit }
                    )
                }
                ActionsValidationError::FunctionCallArgumentsLengthExceeded { length, limit } => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::FunctionCallArgumentsLengthExceeded(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::FunctionCallArgumentsLengthExceeded { length, limit }
                    )
                }
                ActionsValidationError::UnsuitableStakingKey { public_key } => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::UnsuitableStakingKey(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::UnsuitableStakingKey { public_key: Some(public_key.into()) }
                    )
                }
                ActionsValidationError::FunctionCallZeroAttachedGas => {
                    proto::tx_execution_error::invalid_tx_error::actions_validation::Variant::FunctionCallZeroAttachedGas(
                        proto::tx_execution_error::invalid_tx_error::actions_validation::FunctionCallZeroAttachedGas {}
                    )
                }
            }),
        }
    }
}

impl From<TxExecutionError> for proto::TxExecutionError {
    fn from(value: TxExecutionError) -> Self {
        Self {
            variant: Some(match value {
                TxExecutionError::ActionError(error) => {
                    proto::tx_execution_error::Variant::ActionError(proto::tx_execution_error::ActionError {
                        index: error.index,
                        kind: Some(proto::tx_execution_error::action_error::ActionErrorKind::from(
                            error.kind,
                        )),
                    })
                }
                TxExecutionError::InvalidTxError(error) => {
                    proto::tx_execution_error::Variant::InvalidTxError(proto::tx_execution_error::InvalidTxError {
                        variant: Some(match error {
                            InvalidTxError::InvalidAccessKeyError(error) => {
                                proto::tx_execution_error::invalid_tx_error::Variant::InvalidAccessKeyError(
                                    error.into(),
                                )
                            }
                            InvalidTxError::InvalidSignerId { signer_id } => {
                                proto::tx_execution_error::invalid_tx_error::Variant::InvalidSignerId(
                                    proto::tx_execution_error::invalid_tx_error::InvalidSignerId { signer_id },
                                )
                            }
                            InvalidTxError::SignerDoesNotExist { signer_id } => {
                                proto::tx_execution_error::invalid_tx_error::Variant::SignerDoesNotExist(
                                    proto::tx_execution_error::invalid_tx_error::SignerDoesNotExist {
                                        signer_id: signer_id.to_string(),
                                    },
                                )
                            }
                            InvalidTxError::InvalidNonce { ak_nonce, tx_nonce } => {
                                proto::tx_execution_error::invalid_tx_error::Variant::InvalidNonce(
                                    proto::tx_execution_error::invalid_tx_error::InvalidNonce { tx_nonce, ak_nonce },
                                )
                            }
                            InvalidTxError::NonceTooLarge { tx_nonce, upper_bound } => {
                                proto::tx_execution_error::invalid_tx_error::Variant::NonceTooLarge(
                                    proto::tx_execution_error::invalid_tx_error::NonceTooLarge {
                                        tx_nonce,
                                        upper_bound,
                                    },
                                )
                            }
                            InvalidTxError::InvalidReceiverId { receiver_id } => {
                                proto::tx_execution_error::invalid_tx_error::Variant::InvalidReceiverId(
                                    proto::tx_execution_error::invalid_tx_error::InvalidReceiverId { receiver_id },
                                )
                            }
                            InvalidTxError::InvalidSignature => {
                                proto::tx_execution_error::invalid_tx_error::Variant::InvalidSignature(
                                    proto::tx_execution_error::invalid_tx_error::InvalidSignature {},
                                )
                            }
                            InvalidTxError::NotEnoughBalance {
                                balance,
                                cost,
                                signer_id,
                            } => proto::tx_execution_error::invalid_tx_error::Variant::NotEnoughBalance(
                                proto::tx_execution_error::invalid_tx_error::NotEnoughBalance {
                                    signer_id: signer_id.to_string(),
                                    u128_balance: balance.to_be_bytes().to_vec(),
                                    u128_cost: cost.to_be_bytes().to_vec(),
                                },
                            ),
                            InvalidTxError::LackBalanceForState { amount, signer_id } => {
                                proto::tx_execution_error::invalid_tx_error::Variant::LackBalanceForState(
                                    proto::tx_execution_error::invalid_tx_error::LackBalanceForState {
                                        signer_id: signer_id.to_string(),
                                        u128_amount: amount.to_be_bytes().to_vec(),
                                    },
                                )
                            }
                            InvalidTxError::CostOverflow => {
                                proto::tx_execution_error::invalid_tx_error::Variant::CostOverflow(
                                    proto::tx_execution_error::invalid_tx_error::CostOverflow {},
                                )
                            }
                            InvalidTxError::InvalidChain => {
                                proto::tx_execution_error::invalid_tx_error::Variant::InvalidChain(
                                    proto::tx_execution_error::invalid_tx_error::InvalidChain {},
                                )
                            }
                            InvalidTxError::Expired => proto::tx_execution_error::invalid_tx_error::Variant::Expired(
                                proto::tx_execution_error::invalid_tx_error::Expired {},
                            ),
                            InvalidTxError::ActionsValidation(error) => {
                                proto::tx_execution_error::invalid_tx_error::Variant::ActionsValidation(error.into())
                            }
                            InvalidTxError::TransactionSizeExceeded { size, limit } => {
                                proto::tx_execution_error::invalid_tx_error::Variant::TransactionSizeExceeded(
                                    proto::tx_execution_error::invalid_tx_error::TransactionSizeExceeded {
                                        size,
                                        limit,
                                    },
                                )
                            }
                        }),
                    })
                }
            }),
        }
    }
}

impl From<ExecutionMetadataView> for proto::ExecutionMetadataView {
    fn from(value: ExecutionMetadataView) -> Self {
        Self {
            version: value.version,
            gas_profile: value
                .gas_profile
                .map(|v| v.into_iter().map(Into::into).collect())
                .map(|v| proto::execution_metadata_view::RepeatedCostGasUsed { gas_profile: v }),
        }
    }
}

impl From<CostGasUsed> for proto::CostGasUsed {
    fn from(value: CostGasUsed) -> Self {
        Self {
            cost_category: match value.cost_category.as_str() {
                "ACTION_COST" => proto::CostCategory::ActionCost,
                "WASM_HOST_COST" => proto::CostCategory::WasmHostCost,
                v => panic!("Unknown variant {v}"),
            } as i32,
            cost: Some(proto::Cost {
                variant: Some(match value.cost.as_str() {
                    "CREATE_ACCOUNT" => proto::cost::Variant::ActionCost(proto::cost::ActionCost {
                        value: proto::ActionCosts::CreateAccount as i32,
                    }),
                    "DELETE_ACCOUNT" => proto::cost::Variant::ActionCost(proto::cost::ActionCost {
                        value: proto::ActionCosts::DeleteAccount as i32,
                    }),
                    "DEPLOY_CONTRACT" => proto::cost::Variant::ActionCost(proto::cost::ActionCost {
                        value: proto::ActionCosts::DeployContract as i32,
                    }),
                    "FUNCTION_CALL" => proto::cost::Variant::ActionCost(proto::cost::ActionCost {
                        value: proto::ActionCosts::FunctionCall as i32,
                    }),
                    "TRANSFER" => proto::cost::Variant::ActionCost(proto::cost::ActionCost {
                        value: proto::ActionCosts::Transfer as i32,
                    }),
                    "STAKE" => proto::cost::Variant::ActionCost(proto::cost::ActionCost {
                        value: proto::ActionCosts::Stake as i32,
                    }),
                    "ADD_KEY" => proto::cost::Variant::ActionCost(proto::cost::ActionCost {
                        value: proto::ActionCosts::AddKey as i32,
                    }),
                    "DELETE_KEY" => proto::cost::Variant::ActionCost(proto::cost::ActionCost {
                        value: proto::ActionCosts::DeleteKey as i32,
                    }),
                    "VALUE_RETURN" => proto::cost::Variant::ActionCost(proto::cost::ActionCost {
                        value: proto::ActionCosts::ValueReturn as i32,
                    }),
                    "NEW_RECEIPT" => proto::cost::Variant::ActionCost(proto::cost::ActionCost {
                        value: proto::ActionCosts::NewReceipt as i32,
                    }),
                    "BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Base as i32,
                    }),
                    "CONTRACT_LOADING_BASE" | "CONTRACT_COMPILE_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::ContractLoadingBase as i32,
                    }),
                    "CONTRACT_LOADING_BYTES" | "CONTRACT_COMPILE_BYTES" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::ContractLoadingBytes as i32,
                    }),
                    "READ_MEMORY_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::ReadMemoryBase as i32,
                    }),
                    "READ_MEMORY_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::ReadMemoryByte as i32,
                    }),
                    "WRITE_MEMORY_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::WriteMemoryBase as i32,
                    }),
                    "WRITE_MEMORY_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::WriteMemoryByte as i32,
                    }),
                    "READ_REGISTER_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::ReadRegisterBase as i32,
                    }),
                    "READ_REGISTER_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::ReadRegisterByte as i32,
                    }),
                    "WRITE_REGISTER_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::WriteRegisterBase as i32,
                    }),
                    "WRITE_REGISTER_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::WriteRegisterByte as i32,
                    }),
                    "UTF8_DECODING_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Utf8DecodingBase as i32,
                    }),
                    "UTF8_DECODING_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Utf8DecodingByte as i32,
                    }),
                    "UTF16_DECODING_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Utf16DecodingBase as i32,
                    }),
                    "UTF16_DECODING_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Utf16DecodingByte as i32,
                    }),
                    "SHA256_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Sha256Base as i32,
                    }),
                    "SHA256_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Sha256Byte as i32,
                    }),
                    "KECCAK256_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Keccak256Base as i32,
                    }),
                    "KECCAK256_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Keccak256Byte as i32,
                    }),
                    "KECCAK512_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Keccak512Base as i32,
                    }),
                    "KECCAK512_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Keccak512Byte as i32,
                    }),
                    "RIPEMD160_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Ripemd160Base as i32,
                    }),
                    "RIPEMD160_BLOCK" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::Ripemd160Block as i32,
                    }),
                    "ECRECOVER_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::EcrecoverBase as i32,
                    }),
                    "LOG_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::LogBase as i32,
                    }),
                    "LOG_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::LogByte as i32,
                    }),
                    "STORAGE_WRITE_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageWriteBase as i32,
                    }),
                    "STORAGE_WRITE_KEY_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageWriteKeyByte as i32,
                    }),
                    "STORAGE_WRITE_VALUE_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageWriteValueByte as i32,
                    }),
                    "STORAGE_WRITE_EVICTED_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageWriteEvictedByte as i32,
                    }),
                    "STORAGE_READ_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageReadBase as i32,
                    }),
                    "STORAGE_READ_KEY_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageReadKeyByte as i32,
                    }),
                    "STORAGE_READ_VALUE_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageReadValueByte as i32,
                    }),
                    "STORAGE_REMOVE_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageRemoveBase as i32,
                    }),
                    "STORAGE_REMOVE_KEY_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageRemoveKeyByte as i32,
                    }),
                    "STORAGE_REMOVE_RET_VALUE_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageRemoveRetValueByte as i32,
                    }),
                    "STORAGE_HAS_KEY_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageHasKeyBase as i32,
                    }),
                    "STORAGE_HAS_KEY_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageHasKeyByte as i32,
                    }),
                    "STORAGE_ITER_CREATE_PREFIX_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageIterCreatePrefixBase as i32,
                    }),
                    "STORAGE_ITER_CREATE_PREFIX_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageIterCreatePrefixByte as i32,
                    }),
                    "STORAGE_ITER_CREATE_RANGE_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageIterCreateRangeBase as i32,
                    }),
                    "STORAGE_ITER_CREATE_FROM_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageIterCreateFromByte as i32,
                    }),
                    "STORAGE_ITER_CREATE_TO_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageIterCreateToByte as i32,
                    }),
                    "STORAGE_ITER_NEXT_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageIterNextBase as i32,
                    }),
                    "STORAGE_ITER_NEXT_KEY_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageIterNextKeyByte as i32,
                    }),
                    "STORAGE_ITER_NEXT_VALUE_BYTE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::StorageIterNextValueByte as i32,
                    }),
                    "TOUCHING_TRIE_NODE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::TouchingTrieNode as i32,
                    }),
                    "READ_CACHED_TRIE_NODE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::ReadCachedTrieNode as i32,
                    }),
                    "PROMISE_AND_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::PromiseAndBase as i32,
                    }),
                    "PROMISE_AND_PER_PROMISE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::PromiseAndPerPromise as i32,
                    }),
                    "PROMISE_RETURN" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::PromiseReturn as i32,
                    }),
                    "VALIDATOR_STAKE_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::ValidatorStakeBase as i32,
                    }),
                    "VALIDATOR_TOTAL_STAKE_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::ValidatorTotalStakeBase as i32,
                    }),
                    "ALT_BN128_G1_MULTIEXP_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::AltBn128G1MultiexpBase as i32,
                    }),
                    "ALT_BN128_G1_MULTIEXP_ELEMENT" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::AltBn128G1MultiexpElement as i32,
                    }),
                    "ALT_BN128_PAIRING_CHECK_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::AltBn128PairingCheckBase as i32,
                    }),
                    "ALT_BN128_PAIRING_CHECK_ELEMENT" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::AltBn128PairingCheckElement as i32,
                    }),
                    "ALT_BN128_G1_SUM_BASE" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::AltBn128G1SumBase as i32,
                    }),
                    "ALT_BN128_G1_SUM_ELEMENT" => proto::cost::Variant::ExtCost(proto::cost::ExtCost {
                        value: proto::ExtCosts::AltBn128G1SumElement as i32,
                    }),
                    "WASM_INSTRUCTION" => proto::cost::Variant::WasmInstruction(proto::cost::WasmInstruction {}),
                    v => panic!("Unknown variant {v}"),
                }),
            }),
            gas_used: value.gas_used,
        }
    }
}

impl From<ExecutionOutcomeWithIdView> for proto::ExecutionOutcomeWithIdView {
    fn from(value: ExecutionOutcomeWithIdView) -> Self {
        Self {
            proof: value.proof.into_iter().map(Into::into).collect(),
            h256_block_hash: value.block_hash.0.to_vec(),
            h256_id: value.id.0.to_vec(),
            outcome: Some(value.outcome.into()),
        }
    }
}

impl From<MerklePathItem> for proto::MerklePathItem {
    fn from(value: MerklePathItem) -> Self {
        Self {
            h256_hash: value.hash.0.to_vec(),
            direction: proto::Direction::from(value.direction) as i32,
        }
    }
}

impl From<Direction> for proto::Direction {
    fn from(value: Direction) -> Self {
        match value {
            Direction::Left => proto::Direction::Left,
            Direction::Right => proto::Direction::Right,
        }
    }
}

impl From<BlockView> for proto::BlockHeaderView {
    fn from(value: BlockView) -> Self {
        Self {
            author: value.author.into(),
            header: Some(value.header.into()),
        }
    }
}

impl From<ChunkView> for proto::ChunkView {
    fn from(value: ChunkView) -> Self {
        Self {
            author: value.author.into(),
            header: Some(value.header.into()),
            transactions: value.transactions.into_iter().map(Into::into).collect(),
            receipts: value.receipts.into_iter().map(Into::into).collect(),
        }
    }
}

impl From<ChunkHeaderView> for proto::ChunkHeaderView {
    fn from(value: ChunkHeaderView) -> Self {
        Self {
            h256_chunk_hash: value.chunk_hash.0.to_vec(),
            h256_prev_block_hash: value.prev_block_hash.0.to_vec(),
            h256_outcome_root: value.outcome_root.0.to_vec(),
            h256_prev_state_root: value.prev_state_root.0.to_vec(),
            h256_encoded_merkle_root: value.encoded_merkle_root.0.to_vec(),
            encoded_length: value.encoded_length,
            height_created: value.height_created,
            height_included: value.height_included,
            shard_id: value.shard_id,
            gas_used: value.gas_used,
            gas_limit: value.gas_limit,
            u128_validator_reward: value.validator_reward.to_be_bytes().to_vec(),
            u128_balance_burnt: value.balance_burnt.to_be_bytes().to_vec(),
            h256_outgoing_receipts_root: value.outgoing_receipts_root.0.to_vec(),
            h256_tx_root: value.tx_root.0.to_vec(),
            validator_proposals: value.validator_proposals.into_iter().map(Into::into).collect(),
            signature: Some(value.signature.into()),
        }
    }
}

impl From<TransactionWithOutcome> for proto::TransactionWithOutcome {
    fn from(value: TransactionWithOutcome) -> Self {
        Self {
            transaction: Some(proto::SignedTransactionView::from(value.transaction)),
            outcome: Some(proto::ExecutionOutcomeWithOptionalReceipt::from(value.outcome)),
        }
    }
}

impl From<SignedTransactionView> for proto::SignedTransactionView {
    fn from(value: SignedTransactionView) -> Self {
        Self {
            signer_id: value.signer_id.to_string(),
            public_key: Some(proto::PublicKey::from(value.public_key)),
            nonce: value.nonce,
            receiver_id: value.receiver_id.to_string(),
            actions: value.actions.into_iter().map(Into::into).collect(),
            signature: Some(proto::Signature::from(value.signature)),
            h256_hash: value.hash.0.to_vec(),
        }
    }
}

impl From<Signature> for proto::Signature {
    fn from(value: Signature) -> Self {
        Self {
            variant: Some(match &value {
                Signature::ED25519(signature) => proto::signature::Variant::Ed25519(proto::signature::Ed25519 {
                    h512_value: signature.as_ref().to_vec(),
                }),
                Signature::SECP256K1(_) => proto::signature::Variant::Secp256k1(proto::signature::Secp256k1 {
                    h520_value: bs58::decode(value.to_string().split(':').collect::<Vec<&str>>()[1])
                        .into_vec()
                        .unwrap(),
                }),
            }),
        }
    }
}

impl From<ExecutionOutcomeWithOptionalReceipt> for proto::ExecutionOutcomeWithOptionalReceipt {
    fn from(value: ExecutionOutcomeWithOptionalReceipt) -> Self {
        Self {
            execution_outcome: Some(proto::ExecutionOutcomeWithIdView {
                proof: value.execution_outcome.proof.into_iter().map(Into::into).collect(),
                h256_block_hash: value.execution_outcome.block_hash.0.to_vec(),
                h256_id: value.execution_outcome.id.0.to_vec(),
                outcome: Some(proto::ExecutionOutcomeView::from(value.execution_outcome.outcome)),
            }),
            receipt: value.receipt.map(|v| proto::ReceiptView {
                predecessor_id: v.predecessor_id.to_string(),
                receiver_id: v.receiver_id.to_string(),
                h256_receipt_id: v.receipt_id.0.to_vec(),
                receipt: Some(proto::ReceiptEnumView::from(v.receipt)),
            }),
        }
    }
}

impl From<IndexerBlockHeaderView> for proto::IndexerBlockHeaderView {
    fn from(value: IndexerBlockHeaderView) -> Self {
        Self {
            height: value.height,
            prev_height: value.prev_height,
            h256_epoch_id: value.epoch_id.0.to_vec(),
            h256_next_epoch_id: value.next_epoch_id.0.to_vec(),
            h256_hash: value.hash.0.to_vec(),
            h256_prev_hash: value.prev_hash.0.to_vec(),
            h256_prev_state_root: value.prev_state_root.0.to_vec(),
            h256_chunk_receipts_root: value.chunk_receipts_root.0.to_vec(),
            h256_chunk_headers_root: value.chunk_headers_root.0.to_vec(),
            h256_chunk_tx_root: value.chunk_tx_root.0.to_vec(),
            h256_outcome_root: value.outcome_root.0.to_vec(),
            chunks_included: value.chunks_included,
            h256_challenges_root: value.challenges_root.0.to_vec(),
            timestamp: value.timestamp,
            timestamp_nanosec: value.timestamp_nanosec,
            h256_random_value: value.random_value.0.to_vec(),
            validator_proposals: value.validator_proposals.into_iter().map(Into::into).collect(),
            chunk_mask: value.chunk_mask,
            u128_gas_price: value.gas_price.to_be_bytes().to_vec(),
            block_ordinal: value.block_ordinal,
            u128_total_supply: value.total_supply.to_be_bytes().to_vec(),
            challenges_result: value.challenges_result.into_iter().map(Into::into).collect(),
            h256_last_final_block: value.last_final_block.0.to_vec(),
            h256_last_ds_final_block: value.last_ds_final_block.0.to_vec(),
            h256_next_bp_hash: value.next_bp_hash.0.to_vec(),
            h256_block_merkle_root: value.block_merkle_root.0.to_vec(),
            h256_epoch_sync_data_hash: value.epoch_sync_data_hash.map(|v| v.0.to_vec()),
            approvals: value.approvals.into_iter().map(Into::into).collect(),
            signature: Some(value.signature.into()),
            latest_protocol_version: value.latest_protocol_version,
        }
    }
}

impl From<ValidatorStakeView> for proto::ValidatorStakeView {
    fn from(value: ValidatorStakeView) -> Self {
        Self {
            variant: Some(match value {
                ValidatorStakeView::V1(view) => {
                    proto::validator_stake_view::Variant::V1(proto::validator_stake_view::ValidatorStakeViewV1 {
                        account_id: view.account_id.to_string(),
                        public_key: Some(view.public_key.into()),
                        u128_stake: view.stake.to_be_bytes().to_vec(),
                    })
                }
            }),
        }
    }
}

impl From<SlashedValidator> for proto::SlashedValidator {
    fn from(value: SlashedValidator) -> Self {
        Self {
            account_id: value.account_id.to_string(),
            is_double_sign: value.is_double_sign,
        }
    }
}

impl From<Option<Signature>> for proto::OptionalSignature {
    fn from(value: Option<Signature>) -> Self {
        Self {
            value: value.map(Into::into),
        }
    }
}

impl From<ReceiptView> for proto::ReceiptView {
    fn from(value: ReceiptView) -> Self {
        Self {
            predecessor_id: value.predecessor_id.to_string(),
            receiver_id: value.receiver_id.to_string(),
            h256_receipt_id: value.receipt_id.0.to_vec(),
            receipt: Some(value.receipt.into()),
        }
    }
}

impl From<ReceiptEnumView> for proto::ReceiptEnumView {
    fn from(value: ReceiptEnumView) -> Self {
        Self {
            variant: Some(match value {
                ReceiptEnumView::Action {
                    actions,
                    gas_price,
                    signer_id,
                    input_data_ids,
                    output_data_receivers,
                    signer_public_key,
                } => proto::receipt_enum_view::Variant::Action(proto::receipt_enum_view::Action {
                    signer_id: signer_id.to_string(),
                    signer_public_key: Some(signer_public_key.into()),
                    u128_gas_price: gas_price.to_be_bytes().to_vec(),
                    output_data_receivers: output_data_receivers.into_iter().map(Into::into).collect(),
                    h256_input_data_ids: input_data_ids.into_iter().map(|v| v.0.to_vec()).collect(),
                    actions: actions.into_iter().map(Into::into).collect(),
                }),
                ReceiptEnumView::Data { data, data_id } => {
                    proto::receipt_enum_view::Variant::Data(proto::receipt_enum_view::Data {
                        data,
                        h256_data_id: data_id.0.to_vec(),
                    })
                }
            }),
        }
    }
}

impl From<DataReceiverView> for proto::DataReceiverView {
    fn from(value: DataReceiverView) -> Self {
        Self {
            h256_data_id: value.data_id.0.to_vec(),
            receiver_id: value.receiver_id.to_string(),
        }
    }
}

impl From<ActionView> for proto::ActionView {
    fn from(value: ActionView) -> Self {
        Self {
            variant: Some(match value {
                ActionView::CreateAccount => {
                    proto::action_view::Variant::CreateAccount(proto::action_view::CreateAccount {})
                }
                ActionView::DeployContract { code } => {
                    proto::action_view::Variant::DeployContract(proto::action_view::DeployContract { code })
                }
                ActionView::FunctionCall {
                    args,
                    gas,
                    deposit,
                    method_name,
                } => proto::action_view::Variant::FunctionCall(proto::action_view::FunctionCall {
                    method_name,
                    args,
                    gas,
                    u128_deposit: deposit.to_be_bytes().to_vec(),
                }),
                ActionView::Transfer { deposit } => {
                    proto::action_view::Variant::Transfer(proto::action_view::Transfer {
                        u128_deposit: deposit.to_be_bytes().to_vec(),
                    })
                }
                ActionView::Stake { stake, public_key } => {
                    proto::action_view::Variant::Stake(proto::action_view::Stake {
                        u128_stake: stake.to_be_bytes().to_vec(),
                        public_key: Some(public_key.into()),
                    })
                }
                ActionView::AddKey { access_key, public_key } => {
                    proto::action_view::Variant::AddKey(proto::action_view::AddKey {
                        public_key: Some(public_key.into()),
                        access_key: Some(access_key.into()),
                    })
                }
                ActionView::DeleteKey { public_key } => {
                    proto::action_view::Variant::DeleteKey(proto::action_view::DeleteKey {
                        public_key: Some(public_key.into()),
                    })
                }
                ActionView::DeleteAccount { beneficiary_id } => {
                    proto::action_view::Variant::DeleteAccount(proto::action_view::DeleteAccount {
                        beneficiary_id: beneficiary_id.to_string(),
                    })
                }
            }),
        }
    }
}

impl From<StateChangeWithCauseView> for proto::StateChangeWithCauseView {
    fn from(value: StateChangeWithCauseView) -> Self {
        Self {
            cause: Some(value.cause.into()),
            value: Some(value.value.into()),
        }
    }
}

impl From<StateChangeCauseView> for proto::StateChangeCauseView {
    fn from(value: StateChangeCauseView) -> Self {
        Self {
            variant: Some(match value {
                StateChangeCauseView::NotWritableToDisk => proto::state_change_cause_view::Variant::NotWritableToDisk(
                    proto::state_change_cause_view::NotWritableToDisk {},
                ),
                StateChangeCauseView::InitialState => proto::state_change_cause_view::Variant::InitialState(
                    proto::state_change_cause_view::InitialState {},
                ),
                StateChangeCauseView::TransactionProcessing { tx_hash } => {
                    proto::state_change_cause_view::Variant::TransactionProcessing(
                        proto::state_change_cause_view::TransactionProcessing {
                            h256_tx_hash: tx_hash.0.to_vec(),
                        },
                    )
                }
                StateChangeCauseView::ActionReceiptProcessingStarted { receipt_hash } => {
                    proto::state_change_cause_view::Variant::ActionReceiptProcessingStarted(
                        proto::state_change_cause_view::ActionReceiptProcessingStarted {
                            h256_receipt_hash: receipt_hash.0.to_vec(),
                        },
                    )
                }
                StateChangeCauseView::ActionReceiptGasReward { receipt_hash } => {
                    proto::state_change_cause_view::Variant::ActionReceiptGasReward(
                        proto::state_change_cause_view::ActionReceiptGasReward {
                            h256_receipt_hash: receipt_hash.0.to_vec(),
                        },
                    )
                }
                StateChangeCauseView::ReceiptProcessing { receipt_hash } => {
                    proto::state_change_cause_view::Variant::ReceiptProcessing(
                        proto::state_change_cause_view::ReceiptProcessing {
                            h256_receipt_hash: receipt_hash.0.to_vec(),
                        },
                    )
                }
                StateChangeCauseView::PostponedReceipt { receipt_hash } => {
                    proto::state_change_cause_view::Variant::PostponedReceipt(
                        proto::state_change_cause_view::PostponedReceipt {
                            h256_receipt_hash: receipt_hash.0.to_vec(),
                        },
                    )
                }
                StateChangeCauseView::UpdatedDelayedReceipts => {
                    proto::state_change_cause_view::Variant::UpdatedDelayedReceipts(
                        proto::state_change_cause_view::UpdatedDelayedReceipts {},
                    )
                }
                StateChangeCauseView::ValidatorAccountsUpdate => {
                    proto::state_change_cause_view::Variant::ValidatorAccountsUpdate(
                        proto::state_change_cause_view::ValidatorAccountsUpdate {},
                    )
                }
                StateChangeCauseView::Migration => {
                    proto::state_change_cause_view::Variant::Migration(proto::state_change_cause_view::Migration {})
                }
                StateChangeCauseView::Resharding => {
                    proto::state_change_cause_view::Variant::Resharding(proto::state_change_cause_view::Resharding {})
                }
            }),
        }
    }
}

impl From<StateChangeValueView> for proto::StateChangeValueView {
    fn from(value: StateChangeValueView) -> Self {
        Self {
            variant: Some(match value {
                StateChangeValueView::AccountUpdate { account_id, account } => {
                    proto::state_change_value_view::Variant::AccountUpdate(
                        proto::state_change_value_view::AccountUpdate {
                            account_id: account_id.into(),
                            account: Some(account.into()),
                        },
                    )
                }
                StateChangeValueView::AccountDeletion { account_id } => {
                    proto::state_change_value_view::Variant::AccountDeletion(
                        proto::state_change_value_view::AccountDeletion {
                            account_id: account_id.into(),
                        },
                    )
                }
                StateChangeValueView::AccessKeyUpdate {
                    account_id,
                    access_key,
                    public_key,
                } => proto::state_change_value_view::Variant::AccessKeyUpdate(
                    proto::state_change_value_view::AccessKeyUpdate {
                        account_id: account_id.into(),
                        access_key: Some(access_key.into()),
                        public_key: Some(public_key.into()),
                    },
                ),
                StateChangeValueView::AccessKeyDeletion { account_id, public_key } => {
                    proto::state_change_value_view::Variant::AccessKeyDeletion(
                        proto::state_change_value_view::AccessKeyDeletion {
                            account_id: account_id.into(),
                            public_key: Some(public_key.into()),
                        },
                    )
                }
                StateChangeValueView::DataUpdate { account_id, value, key } => {
                    proto::state_change_value_view::Variant::DataUpdate(proto::state_change_value_view::DataUpdate {
                        account_id: account_id.into(),
                        value: AsRef::<Vec<u8>>::as_ref(&value).to_vec(),
                        key: AsRef::<Vec<u8>>::as_ref(&key).to_vec(),
                    })
                }
                StateChangeValueView::DataDeletion { account_id, key } => {
                    proto::state_change_value_view::Variant::DataDeletion(
                        proto::state_change_value_view::DataDeletion {
                            account_id: account_id.into(),
                            key: AsRef::<Vec<u8>>::as_ref(&key).to_vec(),
                        },
                    )
                }
                StateChangeValueView::ContractCodeUpdate { account_id, code } => {
                    proto::state_change_value_view::Variant::ContractCodeUpdate(
                        proto::state_change_value_view::ContractCodeUpdate {
                            account_id: account_id.into(),
                            code,
                        },
                    )
                }
                StateChangeValueView::ContractCodeDeletion { account_id } => {
                    proto::state_change_value_view::Variant::ContractCodeDeletion(
                        proto::state_change_value_view::ContractCodeDeletion {
                            account_id: account_id.into(),
                        },
                    )
                }
            }),
        }
    }
}

impl From<PublicKey> for proto::PublicKey {
    fn from(value: PublicKey) -> Self {
        Self {
            variant: Some(match value {
                PublicKey::ED25519(key) => proto::public_key::Variant::Ed25519(proto::public_key::Ed25519 {
                    h256_value: key.0.to_vec(),
                }),
                PublicKey::SECP256K1(key) => proto::public_key::Variant::Secp256k1(proto::public_key::Secp256k1 {
                    h512_value: key.as_ref().to_vec(),
                }),
            }),
        }
    }
}

impl From<AccessKeyView> for proto::AccessKeyView {
    fn from(value: AccessKeyView) -> Self {
        Self {
            nonce: value.nonce,
            permission: Some(value.permission.into()),
        }
    }
}

impl From<AccessKeyPermissionView> for proto::AccessKeyPermissionView {
    fn from(value: AccessKeyPermissionView) -> Self {
        Self {
            variant: Some(match value {
                AccessKeyPermissionView::FunctionCall {
                    allowance,
                    method_names,
                    receiver_id,
                } => proto::access_key_permission_view::Variant::FunctionCall(
                    proto::access_key_permission_view::FunctionCall {
                        u128_allowance: allowance.map(|v| v.to_be_bytes().to_vec()),
                        receiver_id,
                        method_names,
                    },
                ),
                AccessKeyPermissionView::FullAccess => proto::access_key_permission_view::Variant::FullAccess(
                    proto::access_key_permission_view::FullAccess {},
                ),
            }),
        }
    }
}

impl From<AccountView> for proto::AccountView {
    fn from(value: AccountView) -> Self {
        Self {
            u128_amount: value.amount.to_be_bytes().to_vec(),
            u128_locked: value.locked.to_be_bytes().to_vec(),
            h256_code_hash: value.code_hash.0.to_vec(),
            storage_usage: value.storage_usage,
            storage_paid_at: value.storage_paid_at,
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use crate::proto::signature::Variant;
    use near_crypto::{KeyType, Secp256K1Signature, SecretKey};
    use sha2::Digest;

    #[test]
    fn test_secp256k1_signature_is_raw() {
        let data = sha2::Sha256::digest(b"123").to_vec();
        let sk = SecretKey::from_seed(KeyType::SECP256K1, "test");
        let expected_signature = sk.sign(&data);
        let proto_signature = proto::Signature::from(expected_signature.clone());

        match proto_signature.variant.unwrap() {
            Variant::Ed25519(..) => {
                panic!("Wrong variant Ed25519, expected Secp256k1");
            }
            Variant::Secp256k1(signature) => {
                let actual_signature =
                    Signature::SECP256K1(Secp256K1Signature::try_from(signature.h520_value.as_slice()).unwrap());

                assert_eq!(expected_signature, actual_signature);
            }
        }
    }

    #[test]
    fn test_ed25519_signature_is_raw() {
        let data = sha2::Sha256::digest(b"123").to_vec();
        let sk = SecretKey::from_seed(KeyType::ED25519, "test");
        let expected_signature = sk.sign(&data);
        let proto_signature = proto::Signature::from(expected_signature.clone());

        match proto_signature.variant.unwrap() {
            Variant::Ed25519(signature) => {
                let actual_signature =
                    Signature::ED25519(ed25519_dalek::Signature::try_from(signature.h512_value.as_slice()).unwrap());

                assert_eq!(expected_signature, actual_signature);
            }
            Variant::Secp256k1(..) => {
                panic!("Wrong variant Secp256k1, expected Ed25519");
            }
        }
    }
}
