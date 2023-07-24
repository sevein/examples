// Code generated by goa v3.12.2, DO NOT EDIT.
//
// sommelier client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the sommelier service endpoint HTTP clients.
type Client struct {
	// Pick Doer is the HTTP client used to make requests to the pick endpoint.
	PickDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the sommelier service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		PickDoer:            doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Pick returns an endpoint that makes HTTP requests to the sommelier service
// pick server.
func (c *Client) Pick() goa.Endpoint {
	var (
		encodeRequest  = EncodePickRequest(c.encoder)
		decodeResponse = DecodePickResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v any) (any, error) {
		req, err := c.BuildPickRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PickDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("sommelier", "pick", err)
		}
		return decodeResponse(resp)
	}
}
