// Code generated by goa v3.13.2, DO NOT EDIT.
//
// calc gRPC server
//
// Command:
// $ goa gen goa.design/examples/error/design -o error

package server

import (
	"context"
	"errors"

	calc "goa.design/examples/error/gen/calc"
	calcpb "goa.design/examples/error/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc/codes"
)

// Server implements the calcpb.CalcServer interface.
type Server struct {
	DivideH goagrpc.UnaryHandler
	calcpb.UnimplementedCalcServer
}

// New instantiates the server struct with the calc service endpoints.
func New(e *calc.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		DivideH: NewDivideHandler(e.Divide, uh),
	}
}

// NewDivideHandler creates a gRPC handler which serves the "calc" service
// "divide" endpoint.
func NewDivideHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeDivideRequest, EncodeDivideResponse)
	}
	return h
}

// Divide implements the "Divide" method in calcpb.CalcServer interface.
func (s *Server) Divide(ctx context.Context, message *calcpb.DivideRequest) (*calcpb.DivideResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "divide")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.DivideH.Handle(ctx, message)
	if err != nil {
		var en goa.GoaErrorNamer
		if errors.As(err, &en) {
			switch en.GoaErrorName() {
			case "div_by_zero":
				var er *calc.DivByZero
				errors.As(err, &er)
				return nil, goagrpc.NewStatusError(codes.InvalidArgument, err, NewDivideDivByZeroError(er))
			case "timeout":
				return nil, goagrpc.NewStatusError(codes.DeadlineExceeded, err, goagrpc.NewErrorResponse(err))
			}
		}
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.DivideResponse), nil
}
