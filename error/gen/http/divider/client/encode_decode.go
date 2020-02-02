// Code generated by goa v2.0.10, DO NOT EDIT.
//
// divider HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"

	divider "goa.design/examples/error/gen/divider"
	goahttp "goa.design/goa/http"
)

// BuildIntegerDivideRequest instantiates a HTTP request object with method and
// path set to call the "divider" service "integer_divide" endpoint
func (c *Client) BuildIntegerDivideRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a int
		b int
	)
	{
		p, ok := v.(*divider.IntOperands)
		if !ok {
			return nil, goahttp.ErrInvalidType("divider", "integer_divide", "*divider.IntOperands", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: IntegerDivideDividerPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("divider", "integer_divide", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeIntegerDivideResponse returns a decoder for responses returned by the
// divider integer_divide endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeIntegerDivideResponse may return the following errors:
//	- "has_remainder" (type *goa.ServiceError): http.StatusExpectationFailed
//	- "div_by_zero" (type *goa.ServiceError): http.StatusBadRequest
//	- "timeout" (type *goa.ServiceError): http.StatusGatewayTimeout
//	- error: internal error
func DecodeIntegerDivideResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				return nil, goahttp.ErrDecodingError("divider", "integer_divide", err)
			}
			return body, nil
		case http.StatusExpectationFailed:
			var (
				body IntegerDivideHasRemainderResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("divider", "integer_divide", err)
			}
			err = ValidateIntegerDivideHasRemainderResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("divider", "integer_divide", err)
			}
			return nil, NewIntegerDivideHasRemainder(&body)
		case http.StatusBadRequest:
			var (
				body IntegerDivideDivByZeroResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("divider", "integer_divide", err)
			}
			err = ValidateIntegerDivideDivByZeroResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("divider", "integer_divide", err)
			}
			return nil, NewIntegerDivideDivByZero(&body)
		case http.StatusGatewayTimeout:
			var (
				body IntegerDivideTimeoutResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("divider", "integer_divide", err)
			}
			err = ValidateIntegerDivideTimeoutResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("divider", "integer_divide", err)
			}
			return nil, NewIntegerDivideTimeout(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("divider", "integer_divide", resp.StatusCode, string(body))
		}
	}
}

// BuildDivideRequest instantiates a HTTP request object with method and path
// set to call the "divider" service "divide" endpoint
func (c *Client) BuildDivideRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		a float64
		b float64
	)
	{
		p, ok := v.(*divider.FloatOperands)
		if !ok {
			return nil, goahttp.ErrInvalidType("divider", "divide", "*divider.FloatOperands", v)
		}
		a = p.A
		b = p.B
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DivideDividerPath(a, b)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("divider", "divide", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDivideResponse returns a decoder for responses returned by the divider
// divide endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeDivideResponse may return the following errors:
//	- "div_by_zero" (type *goa.ServiceError): http.StatusBadRequest
//	- "timeout" (type *goa.ServiceError): http.StatusGatewayTimeout
//	- error: internal error
func DecodeDivideResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body float64
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("divider", "divide", err)
			}
			return body, nil
		case http.StatusBadRequest:
			var (
				body DivideDivByZeroResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("divider", "divide", err)
			}
			err = ValidateDivideDivByZeroResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("divider", "divide", err)
			}
			return nil, NewDivideDivByZero(&body)
		case http.StatusGatewayTimeout:
			var (
				body DivideTimeoutResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("divider", "divide", err)
			}
			err = ValidateDivideTimeoutResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("divider", "divide", err)
			}
			return nil, NewDivideTimeout(&body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("divider", "divide", resp.StatusCode, string(body))
		}
	}
}
