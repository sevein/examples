// Code generated by goa v2.0.10, DO NOT EDIT.
//
// divider gRPC client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	"context"

	divider "goa.design/examples/error/gen/divider"
	dividerpb "goa.design/examples/error/gen/grpc/divider/pb"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildIntegerDivideFunc builds the remote method to invoke for "divider"
// service "integer_divide" endpoint.
func BuildIntegerDivideFunc(grpccli dividerpb.DividerClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.IntegerDivide(ctx, reqpb.(*dividerpb.IntegerDivideRequest), opts...)
		}
		return grpccli.IntegerDivide(ctx, &dividerpb.IntegerDivideRequest{}, opts...)
	}
}

// EncodeIntegerDivideRequest encodes requests sent to divider integer_divide
// endpoint.
func EncodeIntegerDivideRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*divider.IntOperands)
	if !ok {
		return nil, goagrpc.ErrInvalidType("divider", "integer_divide", "*divider.IntOperands", v)
	}
	return NewIntegerDivideRequest(payload), nil
}

// DecodeIntegerDivideResponse decodes responses from the divider
// integer_divide endpoint.
func DecodeIntegerDivideResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*dividerpb.IntegerDivideResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("divider", "integer_divide", "*dividerpb.IntegerDivideResponse", v)
	}
	res := NewIntegerDivideResult(message)
	return res, nil
}

// BuildDivideFunc builds the remote method to invoke for "divider" service
// "divide" endpoint.
func BuildDivideFunc(grpccli dividerpb.DividerClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Divide(ctx, reqpb.(*dividerpb.DivideRequest), opts...)
		}
		return grpccli.Divide(ctx, &dividerpb.DivideRequest{}, opts...)
	}
}

// EncodeDivideRequest encodes requests sent to divider divide endpoint.
func EncodeDivideRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*divider.FloatOperands)
	if !ok {
		return nil, goagrpc.ErrInvalidType("divider", "divide", "*divider.FloatOperands", v)
	}
	return NewDivideRequest(payload), nil
}

// DecodeDivideResponse decodes responses from the divider divide endpoint.
func DecodeDivideResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*dividerpb.DivideResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("divider", "divide", "*dividerpb.DivideResponse", v)
	}
	res := NewDivideResult(message)
	return res, nil
}
