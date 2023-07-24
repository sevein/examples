// Code generated by goa v3.12.2, DO NOT EDIT.
//
// chatter gRPC client encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package client

import (
	"context"

	chatter "goa.design/examples/streaming/gen/chatter"
	chatterpb "goa.design/examples/streaming/gen/grpc/chatter/pb"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// BuildLoginFunc builds the remote method to invoke for "chatter" service
// "login" endpoint.
func BuildLoginFunc(grpccli chatterpb.ChatterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Login(ctx, reqpb.(*chatterpb.LoginRequest), opts...)
		}
		return grpccli.Login(ctx, &chatterpb.LoginRequest{}, opts...)
	}
}

// EncodeLoginRequest encodes requests sent to chatter login endpoint.
func EncodeLoginRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*chatter.LoginPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "login", "*chatter.LoginPayload", v)
	}
	(*md).Append("user", payload.User)
	(*md).Append("password", payload.Password)
	return NewProtoLoginRequest(), nil
}

// DecodeLoginResponse decodes responses from the chatter login endpoint.
func DecodeLoginResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	message, ok := v.(*chatterpb.LoginResponse)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "login", "*chatterpb.LoginResponse", v)
	}
	res := NewLoginResult(message)
	return res, nil
}

// BuildEchoerFunc builds the remote method to invoke for "chatter" service
// "echoer" endpoint.
func BuildEchoerFunc(grpccli chatterpb.ChatterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Echoer(ctx, opts...)
		}
		return grpccli.Echoer(ctx, opts...)
	}
}

// EncodeEchoerRequest encodes requests sent to chatter echoer endpoint.
func EncodeEchoerRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*chatter.EchoerPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "echoer", "*chatter.EchoerPayload", v)
	}
	(*md).Append("authorization", payload.Token)
	return nil, nil
}

// DecodeEchoerResponse decodes responses from the chatter echoer endpoint.
func DecodeEchoerResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	return &EchoerClientStream{
		stream: v.(chatterpb.Chatter_EchoerClient),
	}, nil
}

// BuildListenerFunc builds the remote method to invoke for "chatter" service
// "listener" endpoint.
func BuildListenerFunc(grpccli chatterpb.ChatterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Listener(ctx, opts...)
		}
		return grpccli.Listener(ctx, opts...)
	}
}

// EncodeListenerRequest encodes requests sent to chatter listener endpoint.
func EncodeListenerRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*chatter.ListenerPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "listener", "*chatter.ListenerPayload", v)
	}
	(*md).Append("authorization", payload.Token)
	return nil, nil
}

// DecodeListenerResponse decodes responses from the chatter listener endpoint.
func DecodeListenerResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	return &ListenerClientStream{
		stream: v.(chatterpb.Chatter_ListenerClient),
	}, nil
}

// BuildSummaryFunc builds the remote method to invoke for "chatter" service
// "summary" endpoint.
func BuildSummaryFunc(grpccli chatterpb.ChatterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Summary(ctx, opts...)
		}
		return grpccli.Summary(ctx, opts...)
	}
}

// EncodeSummaryRequest encodes requests sent to chatter summary endpoint.
func EncodeSummaryRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*chatter.SummaryPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "summary", "*chatter.SummaryPayload", v)
	}
	(*md).Append("authorization", payload.Token)
	return nil, nil
}

// DecodeSummaryResponse decodes responses from the chatter summary endpoint.
func DecodeSummaryResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	return &SummaryClientStream{
		stream: v.(chatterpb.Chatter_SummaryClient),
		view:   view,
	}, nil
}

// BuildSubscribeFunc builds the remote method to invoke for "chatter" service
// "subscribe" endpoint.
func BuildSubscribeFunc(grpccli chatterpb.ChatterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.Subscribe(ctx, reqpb.(*chatterpb.SubscribeRequest), opts...)
		}
		return grpccli.Subscribe(ctx, &chatterpb.SubscribeRequest{}, opts...)
	}
}

// EncodeSubscribeRequest encodes requests sent to chatter subscribe endpoint.
func EncodeSubscribeRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*chatter.SubscribePayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "subscribe", "*chatter.SubscribePayload", v)
	}
	(*md).Append("authorization", payload.Token)
	return NewProtoSubscribeRequest(), nil
}

// DecodeSubscribeResponse decodes responses from the chatter subscribe
// endpoint.
func DecodeSubscribeResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	return &SubscribeClientStream{
		stream: v.(chatterpb.Chatter_SubscribeClient),
	}, nil
}

// BuildHistoryFunc builds the remote method to invoke for "chatter" service
// "history" endpoint.
func BuildHistoryFunc(grpccli chatterpb.ChatterClient, cliopts ...grpc.CallOption) goagrpc.RemoteFunc {
	return func(ctx context.Context, reqpb any, opts ...grpc.CallOption) (any, error) {
		for _, opt := range cliopts {
			opts = append(opts, opt)
		}
		if reqpb != nil {
			return grpccli.History(ctx, reqpb.(*chatterpb.HistoryRequest), opts...)
		}
		return grpccli.History(ctx, &chatterpb.HistoryRequest{}, opts...)
	}
}

// EncodeHistoryRequest encodes requests sent to chatter history endpoint.
func EncodeHistoryRequest(ctx context.Context, v any, md *metadata.MD) (any, error) {
	payload, ok := v.(*chatter.HistoryPayload)
	if !ok {
		return nil, goagrpc.ErrInvalidType("chatter", "history", "*chatter.HistoryPayload", v)
	}
	if payload.View != nil {
		(*md).Append("view", *payload.View)
	}
	(*md).Append("authorization", payload.Token)
	return NewProtoHistoryRequest(), nil
}

// DecodeHistoryResponse decodes responses from the chatter history endpoint.
func DecodeHistoryResponse(ctx context.Context, v any, hdr, trlr metadata.MD) (any, error) {
	var view string
	{
		if vals := hdr.Get("goa-view"); len(vals) > 0 {
			view = vals[0]
		}
	}
	return &HistoryClientStream{
		stream: v.(chatterpb.Chatter_HistoryClient),
		view:   view,
	}, nil
}
