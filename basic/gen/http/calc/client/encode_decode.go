// Code generated by goa v3.5.4, DO NOT EDIT.
//
// calc HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/basic/design -o
// $(GOPATH)/src/goa.design/examples/basic

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	calc "goa.design/examples/basic/gen/calc"
	goahttp "goa.design/goa/v3/http"
)

// BuildMultiplyRequest instantiates a HTTP request object with method and path
// set to call the "calc" service "multiply" endpoint
func (c *Client) BuildMultiplyRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*calc.MultiplyPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("calc", "multiply", "*calc.MultiplyPayload", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: MultiplyCalcPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("calc", "multiply", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeMultiplyResponse returns a decoder for responses returned by the calc
// multiply endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeMultiplyResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body int
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("calc", "multiply", err)
			}
			return body, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("calc", "multiply", resp.StatusCode, string(body))
		}
	}
}
