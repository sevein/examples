// Code generated by goa v3.2.1, DO NOT EDIT.
//
// calc gRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/basic/design -o
// $(GOPATH)/src/goa.design/examples/basic

package server

import (
	"context"

	calc "goa.design/examples/basic/gen/calc"
	calcpb "goa.design/examples/basic/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeAddResponse encodes responses from the "calc" service "add" endpoint.
func EncodeAddResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(int)
	if !ok {
		return nil, goagrpc.ErrInvalidType("calc", "add", "int", v)
	}
	resp := NewAddResponse(result)
	return resp, nil
}

// DecodeAddRequest decodes requests sent to "calc" service "add" endpoint.
func DecodeAddRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *calcpb.AddRequest
		ok      bool
	)
	{
		if message, ok = v.(*calcpb.AddRequest); !ok {
			return nil, goagrpc.ErrInvalidType("calc", "add", "*calcpb.AddRequest", v)
		}
	}
	var payload *calc.AddPayload
	{
		payload = NewAddPayload(message)
	}
	return payload, nil
}
