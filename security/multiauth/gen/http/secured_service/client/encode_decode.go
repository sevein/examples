// Code generated by goa v3.7.6, DO NOT EDIT.
//
// secured_service HTTP client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design -o security/multiauth

package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	securedservice "goa.design/examples/security/multiauth/gen/secured_service"
	goahttp "goa.design/goa/v3/http"
)

// BuildSigninRequest instantiates a HTTP request object with method and path
// set to call the "secured_service" service "signin" endpoint
func (c *Client) BuildSigninRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SigninSecuredServicePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("secured_service", "signin", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSigninRequest returns an encoder for requests sent to the
// secured_service signin server.
func EncodeSigninRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*securedservice.SigninPayload)
		if !ok {
			return goahttp.ErrInvalidType("secured_service", "signin", "*securedservice.SigninPayload", v)
		}
		req.SetBasicAuth(p.Username, p.Password)
		return nil
	}
}

// DecodeSigninResponse returns a decoder for responses returned by the
// secured_service signin endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeSigninResponse may return the following errors:
//	- "unauthorized" (type securedservice.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeSigninResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body SigninResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "signin", err)
			}
			err = ValidateSigninResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("secured_service", "signin", err)
			}
			res := NewSigninCredsOK(&body)
			return res, nil
		case http.StatusUnauthorized:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "signin", err)
			}
			return nil, NewSigninUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("secured_service", "signin", resp.StatusCode, string(body))
		}
	}
}

// BuildSecureRequest instantiates a HTTP request object with method and path
// set to call the "secured_service" service "secure" endpoint
func (c *Client) BuildSecureRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: SecureSecuredServicePath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("secured_service", "secure", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeSecureRequest returns an encoder for requests sent to the
// secured_service secure server.
func EncodeSecureRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*securedservice.SecurePayload)
		if !ok {
			return goahttp.ErrInvalidType("secured_service", "secure", "*securedservice.SecurePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		values := req.URL.Query()
		if p.Fail != nil {
			values.Add("fail", fmt.Sprintf("%v", *p.Fail))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeSecureResponse returns a decoder for responses returned by the
// secured_service secure endpoint. restoreBody controls whether the response
// body should be restored after having been read.
// DecodeSecureResponse may return the following errors:
//	- "invalid-scopes" (type securedservice.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type securedservice.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeSecureResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "secure", err)
			}
			return body, nil
		case http.StatusForbidden:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "secure", err)
			}
			return nil, NewSecureInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "secure", err)
			}
			return nil, NewSecureUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("secured_service", "secure", resp.StatusCode, string(body))
		}
	}
}

// BuildDoublySecureRequest instantiates a HTTP request object with method and
// path set to call the "secured_service" service "doubly_secure" endpoint
func (c *Client) BuildDoublySecureRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DoublySecureSecuredServicePath()}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("secured_service", "doubly_secure", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeDoublySecureRequest returns an encoder for requests sent to the
// secured_service doubly_secure server.
func EncodeDoublySecureRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*securedservice.DoublySecurePayload)
		if !ok {
			return goahttp.ErrInvalidType("secured_service", "doubly_secure", "*securedservice.DoublySecurePayload", v)
		}
		{
			head := p.Token
			if !strings.Contains(head, " ") {
				req.Header.Set("Authorization", "Bearer "+head)
			} else {
				req.Header.Set("Authorization", head)
			}
		}
		values := req.URL.Query()
		values.Add("k", p.Key)
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeDoublySecureResponse returns a decoder for responses returned by the
// secured_service doubly_secure endpoint. restoreBody controls whether the
// response body should be restored after having been read.
// DecodeDoublySecureResponse may return the following errors:
//	- "invalid-scopes" (type securedservice.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type securedservice.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeDoublySecureResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "doubly_secure", err)
			}
			return body, nil
		case http.StatusForbidden:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "doubly_secure", err)
			}
			return nil, NewDoublySecureInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "doubly_secure", err)
			}
			return nil, NewDoublySecureUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("secured_service", "doubly_secure", resp.StatusCode, string(body))
		}
	}
}

// BuildAlsoDoublySecureRequest instantiates a HTTP request object with method
// and path set to call the "secured_service" service "also_doubly_secure"
// endpoint
func (c *Client) BuildAlsoDoublySecureRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: AlsoDoublySecureSecuredServicePath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("secured_service", "also_doubly_secure", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeAlsoDoublySecureRequest returns an encoder for requests sent to the
// secured_service also_doubly_secure server.
func EncodeAlsoDoublySecureRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*securedservice.AlsoDoublySecurePayload)
		if !ok {
			return goahttp.ErrInvalidType("secured_service", "also_doubly_secure", "*securedservice.AlsoDoublySecurePayload", v)
		}
		if p.Token != nil {
			head := *p.Token
			req.Header.Set("X-Authorization", head)
		}
		values := req.URL.Query()
		if p.Key != nil {
			values.Add("k", *p.Key)
		}
		if p.OauthToken != nil {
			values.Add("oauth", *p.OauthToken)
		}
		req.URL.RawQuery = values.Encode()
		if p.Username != nil {
			if p.Password != nil {
				req.SetBasicAuth(*p.Username, *p.Password)
			}
		}
		return nil
	}
}

// DecodeAlsoDoublySecureResponse returns a decoder for responses returned by
// the secured_service also_doubly_secure endpoint. restoreBody controls
// whether the response body should be restored after having been read.
// DecodeAlsoDoublySecureResponse may return the following errors:
//	- "invalid-scopes" (type securedservice.InvalidScopes): http.StatusForbidden
//	- "unauthorized" (type securedservice.Unauthorized): http.StatusUnauthorized
//	- error: internal error
func DecodeAlsoDoublySecureResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
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
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "also_doubly_secure", err)
			}
			return body, nil
		case http.StatusForbidden:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "also_doubly_secure", err)
			}
			return nil, NewAlsoDoublySecureInvalidScopes(body)
		case http.StatusUnauthorized:
			var (
				body string
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("secured_service", "also_doubly_secure", err)
			}
			return nil, NewAlsoDoublySecureUnauthorized(body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("secured_service", "also_doubly_secure", resp.StatusCode, string(body))
		}
	}
}
