// Code generated by goa v3.2.2, DO NOT EDIT.
//
// concat client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/encodings/cbor/design -o
// $(GOPATH)/src/goa.design/examples/encodings/cbor

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the concat service endpoint HTTP clients.
type Client struct {
	// Concat Doer is the HTTP client used to make requests to the concat endpoint.
	ConcatDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the concat service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		ConcatDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Concat returns an endpoint that makes HTTP requests to the concat service
// concat server.
func (c *Client) Concat() goa.Endpoint {
	var (
		decodeResponse = DecodeConcatResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildConcatRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ConcatDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("concat", "concat", err)
		}
		return decodeResponse(resp)
	}
}