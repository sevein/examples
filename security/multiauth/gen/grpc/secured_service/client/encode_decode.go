// Code generated by goa v3.7.12, DO NOT EDIT.
//
// secured_service gRPC client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design -o security/multiauth

package client

import (
	"context"

	secured_servicepb "goa.design/examples/security/multiauth/gen/grpc/secured_service/pb"
	securedservice "goa.design/examples/security/multiauth/gen/secured_service"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildSigninFunc builds the remote method to invoke for "secured_service"
// service "signin" endpoint.
func BuildSigninFunc(grpccli secured_servicepb.SecuredServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Signin(ctx, reqpb.(*secured_servicepb.SigninRequest), opts...)
		}
		return grpccli.Signin(ctx, &secured_servicepb.SigninRequest{}, opts...)
	}
}

// EncodeSigninRequest encodes requests sent to secured_service signin endpoint.
func EncodeSigninRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*securedservice.SigninPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "signin", "*securedservice.SigninPayload", v)
	}
	(*md).Append("username", payload.Username)
	(*md).Append("password", payload.Password)
	return NewProtoSigninRequest(), nil
}

// DecodeSigninResponse decodes responses from the secured_service signin
// endpoint.
func DecodeSigninResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*secured_servicepb.SigninResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "signin", "*secured_servicepb.SigninResponse", v)
	}
	res := NewSigninResult(message)
	return res, nil
}

// BuildSecureFunc builds the remote method to invoke for "secured_service"
// service "secure" endpoint.
func BuildSecureFunc(grpccli secured_servicepb.SecuredServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Secure(ctx, reqpb.(*secured_servicepb.SecureRequest), opts...)
		}
		return grpccli.Secure(ctx, &secured_servicepb.SecureRequest{}, opts...)
	}
}

// EncodeSecureRequest encodes requests sent to secured_service secure endpoint.
func EncodeSecureRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*securedservice.SecurePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "secure", "*securedservice.SecurePayload", v)
	}
	(*md).Append("authorization", payload.Token)
	return NewProtoSecureRequest(payload), nil
}

// DecodeSecureResponse decodes responses from the secured_service secure
// endpoint.
func DecodeSecureResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*secured_servicepb.SecureResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "secure", "*secured_servicepb.SecureResponse", v)
	}
	res := NewSecureResult(message)
	return res, nil
}

// BuildDoublySecureFunc builds the remote method to invoke for
// "secured_service" service "doubly_secure" endpoint.
func BuildDoublySecureFunc(grpccli secured_servicepb.SecuredServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.DoublySecure(ctx, reqpb.(*secured_servicepb.DoublySecureRequest), opts...)
		}
		return grpccli.DoublySecure(ctx, &secured_servicepb.DoublySecureRequest{}, opts...)
	}
}

// EncodeDoublySecureRequest encodes requests sent to secured_service
// doubly_secure endpoint.
func EncodeDoublySecureRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*securedservice.DoublySecurePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "doubly_secure", "*securedservice.DoublySecurePayload", v)
	}
	(*md).Append("authorization", payload.Token)
	return NewProtoDoublySecureRequest(payload), nil
}

// DecodeDoublySecureResponse decodes responses from the secured_service
// doubly_secure endpoint.
func DecodeDoublySecureResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*secured_servicepb.DoublySecureResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "doubly_secure", "*secured_servicepb.DoublySecureResponse", v)
	}
	res := NewDoublySecureResult(message)
	return res, nil
}

// BuildAlsoDoublySecureFunc builds the remote method to invoke for
// "secured_service" service "also_doubly_secure" endpoint.
func BuildAlsoDoublySecureFunc(grpccli secured_servicepb.SecuredServiceClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb interface{}, opts ...grpc.CallOption) (interface{}, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.AlsoDoublySecure(ctx, reqpb.(*secured_servicepb.AlsoDoublySecureRequest), opts...)
		}
		return grpccli.AlsoDoublySecure(ctx, &secured_servicepb.AlsoDoublySecureRequest{}, opts...)
	}
}

// EncodeAlsoDoublySecureRequest encodes requests sent to secured_service
// also_doubly_secure endpoint.
func EncodeAlsoDoublySecureRequest(ctx context.Context, v interface{}, md *metadata.MD) (interface{}, error) {
	payload, ok := v.(*securedservice.AlsoDoublySecurePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "also_doubly_secure", "*securedservice.AlsoDoublySecurePayload", v)
	}
	if payload.OauthToken != nil {
		(*md).Append("oauth", *payload.OauthToken)
	}
	if payload.Token != nil {
		(*md).Append("authorization", *payload.Token)
	}
	return NewProtoAlsoDoublySecureRequest(payload), nil
}

// DecodeAlsoDoublySecureResponse decodes responses from the secured_service
// also_doubly_secure endpoint.
func DecodeAlsoDoublySecureResponse(ctx context.Context, v interface{}, hdr, trlr metadata.MD) (interface{}, error) {
	message, ok := v.(*secured_servicepb.AlsoDoublySecureResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "also_doubly_secure", "*secured_servicepb.AlsoDoublySecureResponse", v)
	}
	res := NewAlsoDoublySecureResult(message)
	return res, nil
}
