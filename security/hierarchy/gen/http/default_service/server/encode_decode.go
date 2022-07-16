// Code generated by goa v3.7.11, DO NOT EDIT.
//
// default_service HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design -o security/hierarchy

package server

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeDefaultResponse returns an encoder for responses returned by the
// default_service default endpoint.
func EncodeDefaultResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeDefaultRequest returns a decoder for requests sent to the
// default_service default endpoint.
func DecodeDefaultRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		payload := NewDefaultPayload()
		user, pass, ok := r.BasicAuth()
		if !ok {
			return nil, goa.MissingFieldError("Authorization", "header")
		}
		payload.Username = user
		payload.Password = pass

		return payload, nil
	}
}
