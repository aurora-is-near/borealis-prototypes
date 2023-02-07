# Usage example

## Decoding exact message type
```go
import borealisproto "github.com/aurora-is-near/borealis-prototypes/go"
import "github.com/aurora-is-near/borealis-prototypes/go/payloads/rpc"

func decodeRpcRequest(data []byte) (*rpc.RpcMessage, error) {
	msg := borealisproto.Message{}
	if err := msg.UnmarshalVT(data); err != nil {
		return nil, err
	}
	rpcRequestPayload, ok := msg.Payload.(*borealisproto.Message_RpcRequest)
	if !ok {
		return nil, fmt.Errorf("not rpc request")
	}
	return rpcRequestPayload.RpcRequest, nil
}
```

## Decoding message of yet unknown type
```go
func decodeGeneric(data []byte) (any, error) {
	msg := borealisproto.Message{}
	if err := msg.UnmarshalVT(data); err != nil {
		return nil, err
	}
	switch t := msg.Payload.(type) {
	case *borealisproto.Message_RpcRequest:
		return t.RpcRequest, nil
	case *borealisproto.Message_RpcResponse:
		return t.RpcResponse, nil
	default:
		return nil, nil
	}
}
```

## Encoding message
```go
func encodeRpcResponse(rpcResponse *rpc.RpcMessage) ([]byte, error) {
	msg := &borealisproto.Message{
		Version: 0,
		Id:      "",
		Payload: &borealisproto.Message_RpcResponse{
			RpcResponse: rpcResponse,
		},
	}
	return msg.MarshalVT()
}
```
