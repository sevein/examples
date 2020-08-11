// Code generated by goa v3.2.2, DO NOT EDIT.
//
// divider gRPC client types
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	divider "goa.design/examples/error/gen/divider"
	dividerpb "goa.design/examples/error/gen/grpc/divider/pb"
)

// NewIntegerDivideRequest builds the gRPC request type from the payload of the
// "integer_divide" endpoint of the "divider" service.
func NewIntegerDivideRequest(payload *divider.IntOperands) *dividerpb.IntegerDivideRequest {
	message := &dividerpb.IntegerDivideRequest{
		A: int32(payload.A),
		B: int32(payload.B),
	}
	return message
}

// NewIntegerDivideResult builds the result type of the "integer_divide"
// endpoint of the "divider" service from the gRPC response type.
func NewIntegerDivideResult(message *dividerpb.IntegerDivideResponse) int {
	result := int(message.Field)
	return result
}

// NewDivideRequest builds the gRPC request type from the payload of the
// "divide" endpoint of the "divider" service.
func NewDivideRequest(payload *divider.FloatOperands) *dividerpb.DivideRequest {
	message := &dividerpb.DivideRequest{
		A: payload.A,
		B: payload.B,
	}
	return message
}

// NewDivideResult builds the result type of the "divide" endpoint of the
// "divider" service from the gRPC response type.
func NewDivideResult(message *dividerpb.DivideResponse) float64 {
	result := message.Field
	return result
}
