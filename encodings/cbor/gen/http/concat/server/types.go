// Code generated by goa v3.7.12, DO NOT EDIT.
//
// concat HTTP server types
//
// Command:
// $ goa gen goa.design/examples/encodings/cbor/design -o encodings/cbor

package server

import (
	concat "goa.design/examples/encodings/cbor/gen/concat"
)

// NewConcatPayload builds a concat service concat endpoint payload.
func NewConcatPayload(a string, b string) *concat.ConcatPayload {
	v := &concat.ConcatPayload{}
	v.A = a
	v.B = b

	return v
}
