// Code generated by goa v3.12.2, DO NOT EDIT.
//
// text service
//
// Command:
// $ goa gen goa.design/examples/encodings/text/design -o encodings/text

package text

import (
	"context"
)

// The text service performs operations on strings
type Service interface {
	// Concatstrings implements concatstrings.
	Concatstrings(context.Context, *ConcatstringsPayload) (res string, err error)
	// Concatbytes implements concatbytes.
	Concatbytes(context.Context, *ConcatbytesPayload) (res []byte, err error)
	// Concatstringfield implements concatstringfield.
	Concatstringfield(context.Context, *ConcatstringfieldPayload) (res *MyConcatenation, err error)
	// Concatbytesfield implements concatbytesfield.
	Concatbytesfield(context.Context, *ConcatbytesfieldPayload) (res *MyConcatenation, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "text"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"concatstrings", "concatbytes", "concatstringfield", "concatbytesfield"}

// ConcatbytesPayload is the payload type of the text service concatbytes
// method.
type ConcatbytesPayload struct {
	// Left operand
	A string
	// Right operand
	B string
}

// ConcatbytesfieldPayload is the payload type of the text service
// concatbytesfield method.
type ConcatbytesfieldPayload struct {
	// Left operand
	A string
	// Right operand
	B string
}

// ConcatstringfieldPayload is the payload type of the text service
// concatstringfield method.
type ConcatstringfieldPayload struct {
	// Left operand
	A string
	// Right operand
	B string
}

// ConcatstringsPayload is the payload type of the text service concatstrings
// method.
type ConcatstringsPayload struct {
	// Left operand
	A string
	// Right operand
	B string
}

// MyConcatenation is the result type of the text service concatstringfield
// method.
type MyConcatenation struct {
	Stringfield *string
	Bytesfield  []byte
}
