// Code generated by goa v3.7.3, DO NOT EDIT.
//
// secured_service HTTP client types
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design -o security/multiauth

package client

import (
	securedservice "goa.design/examples/security/multiauth/gen/secured_service"
	goa "goa.design/goa/v3/pkg"
)

// SigninResponseBody is the type of the "secured_service" service "signin"
// endpoint HTTP response body.
type SigninResponseBody struct {
	// JWT token
	JWT *string `form:"jwt,omitempty" json:"jwt,omitempty" xml:"jwt,omitempty"`
	// API Key
	APIKey *string `form:"api_key,omitempty" json:"api_key,omitempty" xml:"api_key,omitempty"`
	// OAuth2 token
	OauthToken *string `form:"oauth_token,omitempty" json:"oauth_token,omitempty" xml:"oauth_token,omitempty"`
}

// NewSigninCredsOK builds a "secured_service" service "signin" endpoint result
// from a HTTP "OK" response.
func NewSigninCredsOK(body *SigninResponseBody) *securedservice.Creds {
	v := &securedservice.Creds{
		JWT:        *body.JWT,
		APIKey:     *body.APIKey,
		OauthToken: *body.OauthToken,
	}

	return v
}

// NewSigninUnauthorized builds a secured_service service signin endpoint
// unauthorized error.
func NewSigninUnauthorized(body string) securedservice.Unauthorized {
	v := securedservice.Unauthorized(body)

	return v
}

// NewSecureInvalidScopes builds a secured_service service secure endpoint
// invalid-scopes error.
func NewSecureInvalidScopes(body string) securedservice.InvalidScopes {
	v := securedservice.InvalidScopes(body)

	return v
}

// NewSecureUnauthorized builds a secured_service service secure endpoint
// unauthorized error.
func NewSecureUnauthorized(body string) securedservice.Unauthorized {
	v := securedservice.Unauthorized(body)

	return v
}

// NewDoublySecureInvalidScopes builds a secured_service service doubly_secure
// endpoint invalid-scopes error.
func NewDoublySecureInvalidScopes(body string) securedservice.InvalidScopes {
	v := securedservice.InvalidScopes(body)

	return v
}

// NewDoublySecureUnauthorized builds a secured_service service doubly_secure
// endpoint unauthorized error.
func NewDoublySecureUnauthorized(body string) securedservice.Unauthorized {
	v := securedservice.Unauthorized(body)

	return v
}

// NewAlsoDoublySecureInvalidScopes builds a secured_service service
// also_doubly_secure endpoint invalid-scopes error.
func NewAlsoDoublySecureInvalidScopes(body string) securedservice.InvalidScopes {
	v := securedservice.InvalidScopes(body)

	return v
}

// NewAlsoDoublySecureUnauthorized builds a secured_service service
// also_doubly_secure endpoint unauthorized error.
func NewAlsoDoublySecureUnauthorized(body string) securedservice.Unauthorized {
	v := securedservice.Unauthorized(body)

	return v
}

// ValidateSigninResponseBody runs the validations defined on SigninResponseBody
func ValidateSigninResponseBody(body *SigninResponseBody) (err error) {
	if body.JWT == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("jwt", "body"))
	}
	if body.APIKey == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("api_key", "body"))
	}
	if body.OauthToken == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("oauth_token", "body"))
	}
	return
}
