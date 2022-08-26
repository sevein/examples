// Code generated by goa v3.8.4, DO NOT EDIT.
//
// sommelier gRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/examples/cellar/design -o cellar

package server

import (
	"context"

	sommelierpb "goa.design/examples/cellar/gen/grpc/sommelier/pb"
	sommelier "goa.design/examples/cellar/gen/sommelier"
	sommelierviews "goa.design/examples/cellar/gen/sommelier/views"
	goagrpc "goa.design/goa/v3/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodePickResponse encodes responses from the "sommelier" service "pick"
// endpoint.
func EncodePickResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	vres, ok := v.(sommelierviews.StoredBottleCollection)
	if !ok {
		return nil, goagrpc.ErrInvalidType("sommelier", "pick", "sommelierviews.StoredBottleCollection", v)
	}
	result := vres.Projected
	(*hdr).Append("goa-view", vres.View)
	resp := NewProtoStoredBottleCollection(result)
	return resp, nil
}

// DecodePickRequest decodes requests sent to "sommelier" service "pick"
// endpoint.
func DecodePickRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		message *sommelierpb.PickRequest
		ok      bool
	)
	{
		if message, ok = v.(*sommelierpb.PickRequest); !ok {
			return nil, goagrpc.ErrInvalidType("sommelier", "pick", "*sommelierpb.PickRequest", v)
		}
	}
	var payload *sommelier.Criteria
	{
		payload = NewPickPayload(message)
	}
	return payload, nil
}
