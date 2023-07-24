// Code generated by goa v3.12.2, DO NOT EDIT.
//
// default_service client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/security/hierarchy/design -o security/hierarchy

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the default_service service endpoint HTTP clients.
type Client struct {
	// Default Doer is the HTTP client used to make requests to the default
	// endpoint.
	DefaultDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the default_service service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		DefaultDoer:         doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Default returns an endpoint that makes HTTP requests to the default_service
// service default server.
func (c *Client) Default() goa.Endpoint {
	var (
		encodeRequest  = EncodeDefaultRequest(c.encoder)
		decodeResponse = DecodeDefaultResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildDefaultRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DefaultDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("default_service", "default", err)
		}
		return decodeResponse(resp)
	}
}
