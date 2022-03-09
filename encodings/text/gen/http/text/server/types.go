// Code generated by goa v3.6.1, DO NOT EDIT.
//
// text HTTP server types
//
// Command:
// $ goa gen goa.design/examples/encodings/text/design -o
// $(GOPATH)/src/goa.design/examples/encodings/text

package server

import (
	text "goa.design/examples/encodings/text/gen/text"
)

// NewConcatstringsPayload builds a text service concatstrings endpoint payload.
func NewConcatstringsPayload(a string, b string) *text.ConcatstringsPayload {
	v := &text.ConcatstringsPayload{}
	v.A = a
	v.B = b

	return v
}

// NewConcatbytesPayload builds a text service concatbytes endpoint payload.
func NewConcatbytesPayload(a string, b string) *text.ConcatbytesPayload {
	v := &text.ConcatbytesPayload{}
	v.A = a
	v.B = b

	return v
}

// NewConcatstringfieldPayload builds a text service concatstringfield endpoint
// payload.
func NewConcatstringfieldPayload(a string, b string) *text.ConcatstringfieldPayload {
	v := &text.ConcatstringfieldPayload{}
	v.A = a
	v.B = b

	return v
}

// NewConcatbytesfieldPayload builds a text service concatbytesfield endpoint
// payload.
func NewConcatbytesfieldPayload(a string, b string) *text.ConcatbytesfieldPayload {
	v := &text.ConcatbytesfieldPayload{}
	v.A = a
	v.B = b

	return v
}
