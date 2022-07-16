// Code generated by goa v3.7.12, DO NOT EDIT.
//
// secured_service service
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design -o security/multiauth

package securedservice

import (
	"context"

	"goa.design/goa/v3/security"
)

// The secured service exposes endpoints that require valid authorization
// credentials.
type Service interface {
	// Creates a valid JWT
	Signin(context.Context, *SigninPayload) (res *Creds, err error)
	// This action is secured with the jwt scheme
	Secure(context.Context, *SecurePayload) (res string, err error)
	// This action is secured with the jwt scheme and also requires an API key
	// query string.
	DoublySecure(context.Context, *DoublySecurePayload) (res string, err error)
	// This action is secured with the jwt scheme and also requires an API key
	// header.
	AlsoDoublySecure(context.Context, *AlsoDoublySecurePayload) (res string, err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// BasicAuth implements the authorization logic for the Basic security scheme.
	BasicAuth(ctx context.Context, user, pass string, schema *security.BasicScheme) (context.Context, error)
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
	// APIKeyAuth implements the authorization logic for the APIKey security scheme.
	APIKeyAuth(ctx context.Context, key string, schema *security.APIKeyScheme) (context.Context, error)
	// OAuth2Auth implements the authorization logic for the OAuth2 security scheme.
	OAuth2Auth(ctx context.Context, token string, schema *security.OAuth2Scheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "secured_service"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"signin", "secure", "doubly_secure", "also_doubly_secure"}

// AlsoDoublySecurePayload is the payload type of the secured_service service
// also_doubly_secure method.
type AlsoDoublySecurePayload struct {
	// Username used to perform signin
	Username *string
	// Password used to perform signin
	Password *string
	// API key
	Key *string
	// JWT used for authentication
	Token      *string
	OauthToken *string
}

// Creds is the result type of the secured_service service signin method.
type Creds struct {
	// JWT token
	JWT string
	// API Key
	APIKey string
	// OAuth2 token
	OauthToken string
}

// DoublySecurePayload is the payload type of the secured_service service
// doubly_secure method.
type DoublySecurePayload struct {
	// API key
	Key string
	// JWT used for authentication
	Token string
}

// SecurePayload is the payload type of the secured_service service secure
// method.
type SecurePayload struct {
	// Whether to force auth failure even with a valid JWT
	Fail *bool
	// JWT used for authentication
	Token string
}

// Credentials used to authenticate to retrieve JWT token
type SigninPayload struct {
	// Username used to perform signin
	Username string
	// Password used to perform signin
	Password string
}

// Token scopes are invalid
type InvalidScopes string

// Credentials are invalid
type Unauthorized string

// Error returns an error description.
func (e InvalidScopes) Error() string {
	return "Token scopes are invalid"
}

// ErrorName returns "invalid-scopes".
func (e InvalidScopes) ErrorName() string {
	return "invalid-scopes"
}

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "Credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}
