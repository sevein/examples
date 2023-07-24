// Code generated by goa v3.12.2, DO NOT EDIT.
//
// secured_service gRPC client
//
// Command:
// $ goa gen goa.design/examples/security/multiauth/design -o security/multiauth

package client

import (
	"context"

	secured_servicepb "goa.design/examples/security/multiauth/gen/grpc/secured_service/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli secured_servicepb.SecuredServiceClient
	opts    []grpc.CallOption
}

// NewClient instantiates gRPC client for all the secured_service service
// servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: secured_servicepb.NewSecuredServiceClient(cc),
		opts:    opts,
	}
}

// Signin calls the "Signin" function in secured_servicepb.SecuredServiceClient
// interface.
func (c *Client) Signin() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildSigninFunc(c.grpccli, c.opts...),
			EncodeSigninRequest,
			DecodeSigninResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// Secure calls the "Secure" function in secured_servicepb.SecuredServiceClient
// interface.
func (c *Client) Secure() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildSecureFunc(c.grpccli, c.opts...),
			EncodeSecureRequest,
			DecodeSecureResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// DoublySecure calls the "DoublySecure" function in
// secured_servicepb.SecuredServiceClient interface.
func (c *Client) DoublySecure() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildDoublySecureFunc(c.grpccli, c.opts...),
			EncodeDoublySecureRequest,
			DecodeDoublySecureResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}

// AlsoDoublySecure calls the "AlsoDoublySecure" function in
// secured_servicepb.SecuredServiceClient interface.
func (c *Client) AlsoDoublySecure() goa.Endpoint {
	return func(ctx context.Context, v any) (any, error) {
		inv := goagrpc.NewInvoker(
			BuildAlsoDoublySecureFunc(c.grpccli, c.opts...),
			EncodeAlsoDoublySecureRequest,
			DecodeAlsoDoublySecureResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			resp := goagrpc.DecodeError(err)
			switch message := resp.(type) {
			case *goapb.ErrorResponse:
				return nil, goagrpc.NewServiceError(message)
			default:
				return nil, goa.Fault(err.Error())
			}
		}
		return res, nil
	}
}
