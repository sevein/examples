// Code generated by goa v3.7.12, DO NOT EDIT.
//
// chatter gRPC server types
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o streaming

package server

import (
	chatter "goa.design/examples/streaming/gen/chatter"
	chatterviews "goa.design/examples/streaming/gen/chatter/views"
	chatterpb "goa.design/examples/streaming/gen/grpc/chatter/pb"
)

// NewLoginPayload builds the payload of the "login" endpoint of the "chatter"
// service from the gRPC request type.
func NewLoginPayload(user string, password string) *chatter.LoginPayload {
	v := &chatter.LoginPayload{}
	v.User = user
	v.Password = password
	return v
}

// NewProtoLoginResponse builds the gRPC response type from the result of the
// "login" endpoint of the "chatter" service.
func NewProtoLoginResponse(result string) *chatterpb.LoginResponse {
	message := &chatterpb.LoginResponse{}
	message.Field = result
	return message
}

// NewEchoerPayload builds the payload of the "echoer" endpoint of the
// "chatter" service from the gRPC request type.
func NewEchoerPayload(token string) *chatter.EchoerPayload {
	v := &chatter.EchoerPayload{}
	v.Token = token
	return v
}

// NewProtoEchoerResponse builds the gRPC response type from the result of the
// "echoer" endpoint of the "chatter" service.
func NewProtoEchoerResponse(result string) *chatterpb.EchoerResponse {
	message := &chatterpb.EchoerResponse{}
	message.Field = result
	return message
}

func NewEchoerStreamingRequest(v *chatterpb.EchoerStreamingRequest) string {
	spayload := v.Field
	return spayload
}

// NewListenerPayload builds the payload of the "listener" endpoint of the
// "chatter" service from the gRPC request type.
func NewListenerPayload(token string) *chatter.ListenerPayload {
	v := &chatter.ListenerPayload{}
	v.Token = token
	return v
}

// NewProtoListenerResponse builds the gRPC response type from the result of
// the "listener" endpoint of the "chatter" service.
func NewProtoListenerResponse() *chatterpb.ListenerResponse {
	message := &chatterpb.ListenerResponse{}
	return message
}

func NewListenerStreamingRequest(v *chatterpb.ListenerStreamingRequest) string {
	spayload := v.Field
	return spayload
}

// NewSummaryPayload builds the payload of the "summary" endpoint of the
// "chatter" service from the gRPC request type.
func NewSummaryPayload(token string) *chatter.SummaryPayload {
	v := &chatter.SummaryPayload{}
	v.Token = token
	return v
}

// NewProtoChatSummaryCollection builds the gRPC response type from the result
// of the "summary" endpoint of the "chatter" service.
func NewProtoChatSummaryCollection(result chatterviews.ChatSummaryCollectionView) *chatterpb.ChatSummaryCollection {
	message := &chatterpb.ChatSummaryCollection{}
	message.Field = make([]*chatterpb.ChatSummary, len(result))
	for i, val := range result {
		message.Field[i] = &chatterpb.ChatSummary{}
		if val.Message != nil {
			message.Field[i].Message_ = *val.Message
		}
		if val.Length != nil {
			message.Field[i].Length = int32(*val.Length)
		}
		if val.SentAt != nil {
			message.Field[i].SentAt = *val.SentAt
		}
	}
	return message
}

func NewSummaryStreamingRequest(v *chatterpb.SummaryStreamingRequest) string {
	spayload := v.Field
	return spayload
}

// NewSubscribePayload builds the payload of the "subscribe" endpoint of the
// "chatter" service from the gRPC request type.
func NewSubscribePayload(token string) *chatter.SubscribePayload {
	v := &chatter.SubscribePayload{}
	v.Token = token
	return v
}

// NewProtoSubscribeResponse builds the gRPC response type from the result of
// the "subscribe" endpoint of the "chatter" service.
func NewProtoSubscribeResponse(result *chatter.Event) *chatterpb.SubscribeResponse {
	message := &chatterpb.SubscribeResponse{
		Message_: result.Message,
		Action:   result.Action,
		AddedAt:  result.AddedAt,
	}
	return message
}

// NewHistoryPayload builds the payload of the "history" endpoint of the
// "chatter" service from the gRPC request type.
func NewHistoryPayload(view *string, token string) *chatter.HistoryPayload {
	v := &chatter.HistoryPayload{}
	v.View = view
	v.Token = token
	return v
}

// NewProtoHistoryResponse builds the gRPC response type from the result of the
// "history" endpoint of the "chatter" service.
func NewProtoHistoryResponse(result *chatterviews.ChatSummaryView) *chatterpb.HistoryResponse {
	message := &chatterpb.HistoryResponse{}
	if result.Message != nil {
		message.Message_ = *result.Message
	}
	if result.Length != nil {
		message.Length = int32(*result.Length)
	}
	if result.SentAt != nil {
		message.SentAt = *result.SentAt
	}
	return message
}
