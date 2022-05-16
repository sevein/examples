// Code generated by goa v3.7.5, DO NOT EDIT.
//
// secured_service gRPC server types
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design -o security/multiauth

package server

import (
	secured_servicepb "goa.design/examples/security/multiauth/gen/grpc/secured_service/pb"
	securedservice "goa.design/examples/security/multiauth/gen/secured_service"
)

// NewSigninPayload builds the payload of the "signin" endpoint of the
// "secured_service" service from the gRPC request type.
func NewSigninPayload(username string, password string) *securedservice.SigninPayload {
	v := &securedservice.SigninPayload{}
	v.Username = username
	v.Password = password
	return v
}

// NewProtoSigninResponse builds the gRPC response type from the result of the
// "signin" endpoint of the "secured_service" service.
func NewProtoSigninResponse(result *securedservice.Creds) *secured_servicepb.SigninResponse {
	message := &secured_servicepb.SigninResponse{
		Jwt:        result.JWT,
		ApiKey:     result.APIKey,
		OauthToken: result.OauthToken,
	}
	return message
}

// NewSecurePayload builds the payload of the "secure" endpoint of the
// "secured_service" service from the gRPC request type.
func NewSecurePayload(message *secured_servicepb.SecureRequest, token string) *securedservice.SecurePayload {
	v := &securedservice.SecurePayload{
		Fail: &message.Fail,
	}
	v.Token = token
	return v
}

// NewProtoSecureResponse builds the gRPC response type from the result of the
// "secure" endpoint of the "secured_service" service.
func NewProtoSecureResponse(result string) *secured_servicepb.SecureResponse {
	message := &secured_servicepb.SecureResponse{}
	message.Field = result
	return message
}

// NewDoublySecurePayload builds the payload of the "doubly_secure" endpoint of
// the "secured_service" service from the gRPC request type.
func NewDoublySecurePayload(message *secured_servicepb.DoublySecureRequest, token string) *securedservice.DoublySecurePayload {
	v := &securedservice.DoublySecurePayload{
		Key: message.Key,
	}
	v.Token = token
	return v
}

// NewProtoDoublySecureResponse builds the gRPC response type from the result
// of the "doubly_secure" endpoint of the "secured_service" service.
func NewProtoDoublySecureResponse(result string) *secured_servicepb.DoublySecureResponse {
	message := &secured_servicepb.DoublySecureResponse{}
	message.Field = result
	return message
}

// NewAlsoDoublySecurePayload builds the payload of the "also_doubly_secure"
// endpoint of the "secured_service" service from the gRPC request type.
func NewAlsoDoublySecurePayload(message *secured_servicepb.AlsoDoublySecureRequest, oauthToken *string, token *string) *securedservice.AlsoDoublySecurePayload {
	v := &securedservice.AlsoDoublySecurePayload{}
	if message.Username != "" {
		v.Username = &message.Username
	}
	if message.Password != "" {
		v.Password = &message.Password
	}
	if message.Key != "" {
		v.Key = &message.Key
	}
	v.OauthToken = oauthToken
	v.Token = token
	return v
}

// NewProtoAlsoDoublySecureResponse builds the gRPC response type from the
// result of the "also_doubly_secure" endpoint of the "secured_service" service.
func NewProtoAlsoDoublySecureResponse(result string) *secured_servicepb.AlsoDoublySecureResponse {
	message := &secured_servicepb.AlsoDoublySecureResponse{}
	message.Field = result
	return message
}
