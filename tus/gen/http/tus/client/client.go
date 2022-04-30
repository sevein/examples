// Code generated by goa v3.7.3, DO NOT EDIT.
//
// tus client HTTP transport
//
// Command:
// $ goa gen goa.design/examples/tus/design -o tus

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the tus service endpoint HTTP clients.
type Client struct {
	// Head Doer is the HTTP client used to make requests to the head endpoint.
	HeadDoer goahttp.Doer

	// Patch Doer is the HTTP client used to make requests to the patch endpoint.
	PatchDoer goahttp.Doer

	// Options Doer is the HTTP client used to make requests to the options
	// endpoint.
	OptionsDoer goahttp.Doer

	// Post Doer is the HTTP client used to make requests to the post endpoint.
	PostDoer goahttp.Doer

	// Delete Doer is the HTTP client used to make requests to the delete endpoint.
	DeleteDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the tus service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		HeadDoer:            doer,
		PatchDoer:           doer,
		OptionsDoer:         doer,
		PostDoer:            doer,
		DeleteDoer:          doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// Head returns an endpoint that makes HTTP requests to the tus service head
// server.
func (c *Client) Head() goa.Endpoint {
	var (
		encodeRequest  = EncodeHeadRequest(c.encoder)
		decodeResponse = DecodeHeadResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildHeadRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.HeadDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("tus", "head", err)
		}
		return decodeResponse(resp)
	}
}

// Patch returns an endpoint that makes HTTP requests to the tus service patch
// server.
func (c *Client) Patch() goa.Endpoint {
	var (
		encodeRequest  = EncodePatchRequest(c.encoder)
		decodeResponse = DecodePatchResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPatchRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PatchDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("tus", "patch", err)
		}
		return decodeResponse(resp)
	}
}

// Options returns an endpoint that makes HTTP requests to the tus service
// options server.
func (c *Client) Options() goa.Endpoint {
	var (
		decodeResponse = DecodeOptionsResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildOptionsRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.OptionsDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("tus", "options", err)
		}
		return decodeResponse(resp)
	}
}

// Post returns an endpoint that makes HTTP requests to the tus service post
// server.
func (c *Client) Post() goa.Endpoint {
	var (
		encodeRequest  = EncodePostRequest(c.encoder)
		decodeResponse = DecodePostResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPostRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PostDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("tus", "post", err)
		}
		return decodeResponse(resp)
	}
}

// Delete returns an endpoint that makes HTTP requests to the tus service
// delete server.
func (c *Client) Delete() goa.Endpoint {
	var (
		encodeRequest  = EncodeDeleteRequest(c.encoder)
		decodeResponse = DecodeDeleteResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("tus", "delete", err)
		}
		return decodeResponse(resp)
	}
}
