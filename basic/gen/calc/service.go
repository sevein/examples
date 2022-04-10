// Code generated by goa v3.7.2, DO NOT EDIT.
//
// calc service
//
// Command:
// $ goa gen goa.design/examples/basic/design -o basic

package calc

import (
	"context"
)

// The calc service performs operations on numbers
type Service interface {
	// Multiply implements multiply.
	Multiply(context.Context, *MultiplyPayload) (res int, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "calc"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [1]string{"multiply"}

// MultiplyPayload is the payload type of the calc service multiply method.
type MultiplyPayload struct {
	// Left operand
	A int
	// Right operand
	B int
}
