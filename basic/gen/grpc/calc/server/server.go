// Code generated by goa v3.7.1, DO NOT EDIT.
//
// calc gRPC server
//
// Command:
// $ goa gen goa.design/examples/basic/design -o basic

package server

import (
	"context"

	calc "goa.design/examples/basic/gen/calc"
	calcpb "goa.design/examples/basic/gen/grpc/calc/pb"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
)

// Server implements the calcpb.CalcServer interface.
type Server struct {
	MultiplyH goagrpc.UnaryHandler
	calcpb.UnimplementedCalcServer
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the expr.
type ErrorNamer interface {
	ErrorName() string
}

// New instantiates the server struct with the calc service endpoints.
func New(e *calc.Endpoints, uh goagrpc.UnaryHandler) *Server {
	return &Server{
		MultiplyH: NewMultiplyHandler(e.Multiply, uh),
	}
}

// NewMultiplyHandler creates a gRPC handler which serves the "calc" service
// "multiply" endpoint.
func NewMultiplyHandler(endpoint goa.Endpoint, h goagrpc.UnaryHandler) goagrpc.UnaryHandler {
	if h == nil {
		h = goagrpc.NewUnaryHandler(endpoint, DecodeMultiplyRequest, EncodeMultiplyResponse)
	}
	return h
}

// Multiply implements the "Multiply" method in calcpb.CalcServer interface.
func (s *Server) Multiply(ctx context.Context, message *calcpb.MultiplyRequest) (*calcpb.MultiplyResponse, error) {
	ctx = context.WithValue(ctx, goa.MethodKey, "multiply")
	ctx = context.WithValue(ctx, goa.ServiceKey, "calc")
	resp, err := s.MultiplyH.Handle(ctx, message)
	if err != nil {
		return nil, goagrpc.EncodeError(err)
	}
	return resp.(*calcpb.MultiplyResponse), nil
}
