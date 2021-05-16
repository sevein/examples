// Code generated by goa v3.4.0, DO NOT EDIT.
//
// text HTTP server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/encodings/text/design -o
// $(GOPATH)/src/goa.design/examples/encodings/text

package server

import (
	"context"
	"net/http"

	text "goa.design/examples/encodings/text/gen/text"
	goahttp "goa.design/goa/v3/http"
)

// EncodeConcatstringsResponse returns an encoder for responses returned by the
// text concatstrings endpoint.
func EncodeConcatstringsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(string)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "text/html")
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeConcatstringsRequest returns a decoder for requests sent to the text
// concatstrings endpoint.
func DecodeConcatstringsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			a string
			b string

			params = mux.Vars(r)
		)
		a = params["a"]
		b = params["b"]
		payload := NewConcatstringsPayload(a, b)

		return payload, nil
	}
}

// EncodeConcatbytesResponse returns an encoder for responses returned by the
// text concatbytes endpoint.
func EncodeConcatbytesResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.([]byte)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "text/html")
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeConcatbytesRequest returns a decoder for requests sent to the text
// concatbytes endpoint.
func DecodeConcatbytesRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			a string
			b string

			params = mux.Vars(r)
		)
		a = params["a"]
		b = params["b"]
		payload := NewConcatbytesPayload(a, b)

		return payload, nil
	}
}

// EncodeConcatstringfieldResponse returns an encoder for responses returned by
// the text concatstringfield endpoint.
func EncodeConcatstringfieldResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*text.MyConcatenation)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "text/html")
		enc := encoder(ctx, w)
		body := res.Stringfield
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeConcatstringfieldRequest returns a decoder for requests sent to the
// text concatstringfield endpoint.
func DecodeConcatstringfieldRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			a string
			b string

			params = mux.Vars(r)
		)
		a = params["a"]
		b = params["b"]
		payload := NewConcatstringfieldPayload(a, b)

		return payload, nil
	}
}

// EncodeConcatbytesfieldResponse returns an encoder for responses returned by
// the text concatbytesfield endpoint.
func EncodeConcatbytesfieldResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res, _ := v.(*text.MyConcatenation)
		ctx = context.WithValue(ctx, goahttp.ContentTypeKey, "text/html")
		enc := encoder(ctx, w)
		body := res.Bytesfield
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeConcatbytesfieldRequest returns a decoder for requests sent to the
// text concatbytesfield endpoint.
func DecodeConcatbytesfieldRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			a string
			b string

			params = mux.Vars(r)
		)
		a = params["a"]
		b = params["b"]
		payload := NewConcatbytesfieldPayload(a, b)

		return payload, nil
	}
}
