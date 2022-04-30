// Code generated by goa v3.7.3, DO NOT EDIT.
//
// chatter gRPC client
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package client

import (
	"context"

	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	chatterpb "goa.design/examples/streaming/gen/grpc/chatter/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goapb "goa.design/goa/v3/grpc/pb"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli chatterpb.ChatterClient
	opts    []grpc.CallOption
}

// EchoerClientStream implements the chatter.EchoerClientStream interface.
type EchoerClientStream struct {
	stream chatterpb.Chatter_EchoerClient
}

// ListenerClientStream implements the chatter.ListenerClientStream interface.
type ListenerClientStream struct {
	stream chatterpb.Chatter_ListenerClient
}

// SummaryClientStream implements the chatter.SummaryClientStream interface.
type SummaryClientStream struct {
	stream chatterpb.Chatter_SummaryClient
	view   string
}

// SubscribeClientStream implements the chatter.SubscribeClientStream interface.
type SubscribeClientStream struct {
	stream chatterpb.Chatter_SubscribeClient
}

// HistoryClientStream implements the chatter.HistoryClientStream interface.
type HistoryClientStream struct {
	stream chatterpb.Chatter_HistoryClient
	view   string
}

// NewClient instantiates gRPC client for all the chatter service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: chatterpb.NewChatterClient(cc),
		opts:    opts,
	}
}

// Login calls the "Login" function in chatterpb.ChatterClient interface.
func (c *Client) Login() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildLoginFunc(c.grpccli, c.opts...),
			EncodeLoginRequest,
			DecodeLoginResponse)
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

// Echoer calls the "Echoer" function in chatterpb.ChatterClient interface.
func (c *Client) Echoer() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildEchoerFunc(c.grpccli, c.opts...),
			EncodeEchoerRequest,
			DecodeEchoerResponse)
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

// Listener calls the "Listener" function in chatterpb.ChatterClient interface.
func (c *Client) Listener() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildListenerFunc(c.grpccli, c.opts...),
			EncodeListenerRequest,
			DecodeListenerResponse)
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

// Summary calls the "Summary" function in chatterpb.ChatterClient interface.
func (c *Client) Summary() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildSummaryFunc(c.grpccli, c.opts...),
			EncodeSummaryRequest,
			DecodeSummaryResponse)
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

// Subscribe calls the "Subscribe" function in chatterpb.ChatterClient
// interface.
func (c *Client) Subscribe() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildSubscribeFunc(c.grpccli, c.opts...),
			EncodeSubscribeRequest,
			DecodeSubscribeResponse)
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

// History calls the "History" function in chatterpb.ChatterClient interface.
func (c *Client) History() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildHistoryFunc(c.grpccli, c.opts...),
			EncodeHistoryRequest,
			DecodeHistoryResponse)
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

// Recv reads instances of "chatterpb.EchoerResponse" from the "echoer"
// endpoint gRPC stream.
func (s *EchoerClientStream) Recv() (string, error) {
	var res string
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	return NewEchoerResponse(v), nil
}

// Send streams instances of "chatterpb.EchoerStreamingRequest" to the "echoer"
// endpoint gRPC stream.
func (s *EchoerClientStream) Send(res string) error {
	v := NewProtoEchoerStreamingRequest(res)
	return s.stream.Send(v)
}

func (s *EchoerClientStream) Close() error {
	// Close the send direction of the stream
	return s.stream.CloseSend()
}

// Send streams instances of "chatterpb.ListenerStreamingRequest" to the
// "listener" endpoint gRPC stream.
func (s *ListenerClientStream) Send(res string) error {
	v := NewProtoListenerStreamingRequest(res)
	return s.stream.Send(v)
}

func (s *ListenerClientStream) Close() error {
	// synchronize and report any server error
	_, err := s.stream.CloseAndRecv()
	return err
}

// CloseAndRecv reads instances of "chatterpb.ChatSummaryCollection" from the
// "summary" endpoint gRPC stream.
func (s *SummaryClientStream) CloseAndRecv() (chatter.ChatSummaryCollection, error) {
	var res chatter.ChatSummaryCollection
	v, err := s.stream.CloseAndRecv()
	if err != nil {
		return res, err
	}
	proj := NewChatSummaryCollection(v)
	vres := chatterviews.ChatSummaryCollection{Projected: proj, View: "default"}
	if err := chatterviews.ValidateChatSummaryCollection(vres); err != nil {
		return nil, err
	}
	return chatter.NewChatSummaryCollection(vres), nil
}

// Send streams instances of "chatterpb.SummaryStreamingRequest" to the
// "summary" endpoint gRPC stream.
func (s *SummaryClientStream) Send(res string) error {
	v := NewProtoSummaryStreamingRequest(res)
	return s.stream.Send(v)
}

// Recv reads instances of "chatterpb.SubscribeResponse" from the "subscribe"
// endpoint gRPC stream.
func (s *SubscribeClientStream) Recv() (*chatter.Event, error) {
	var res *chatter.Event
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	if err = ValidateSubscribeResponse(v); err != nil {
		return res, err
	}
	return NewEvent(v), nil
}

// Recv reads instances of "chatterpb.HistoryResponse" from the "history"
// endpoint gRPC stream.
func (s *HistoryClientStream) Recv() (*chatter.ChatSummary, error) {
	var res *chatter.ChatSummary
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	proj := NewChatSummaryView(v)
	vres := &chatterviews.ChatSummary{Projected: proj, View: s.view}
	if err := chatterviews.ValidateChatSummary(vres); err != nil {
		return nil, err
	}
	return chatter.NewChatSummary(vres), nil
}

// SetView sets the view.
func (s *HistoryClientStream) SetView(view string) {
	s.view = view
}
