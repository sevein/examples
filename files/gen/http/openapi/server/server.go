// Code generated by goa v3.7.6, DO NOT EDIT.
//
// openapi HTTP server
//
// Command:
// $ goa gen goa.design/examples/files/design -o files

package server

import (
	"context"
	"net/http"

	openapi "goa.design/examples/files/gen/openapi"
	goahttp "goa.design/goa/v3/http"
)

// Server lists the openapi service endpoint HTTP handlers.
type Server struct {
	Mounts              []*MountPoint
	GenHTTPOpenapiJSON  http.Handler
	GenHTTPOpenapiYaml  http.Handler
	GenHTTPOpenapi3JSON http.Handler
	GenHTTPOpenapi3Yaml http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the openapi service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *openapi.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
	fileSystemGenHTTPOpenapiJSON http.FileSystem,
	fileSystemGenHTTPOpenapiYaml http.FileSystem,
	fileSystemGenHTTPOpenapi3JSON http.FileSystem,
	fileSystemGenHTTPOpenapi3Yaml http.FileSystem,
) *Server {
	if fileSystemGenHTTPOpenapiJSON == nil {
		fileSystemGenHTTPOpenapiJSON = http.Dir(".")
	}
	if fileSystemGenHTTPOpenapiYaml == nil {
		fileSystemGenHTTPOpenapiYaml = http.Dir(".")
	}
	if fileSystemGenHTTPOpenapi3JSON == nil {
		fileSystemGenHTTPOpenapi3JSON = http.Dir(".")
	}
	if fileSystemGenHTTPOpenapi3Yaml == nil {
		fileSystemGenHTTPOpenapi3Yaml = http.Dir(".")
	}
	return &Server{
		Mounts: []*MountPoint{
			{"gen/http/openapi.json", "GET", "/openapi.json"},
			{"gen/http/openapi.yaml", "GET", "/openapi.yaml"},
			{"gen/http/openapi3.json", "GET", "/openapi3.json"},
			{"gen/http/openapi3.yaml", "GET", "/openapi3.yaml"},
		},
		GenHTTPOpenapiJSON:  http.FileServer(fileSystemGenHTTPOpenapiJSON),
		GenHTTPOpenapiYaml:  http.FileServer(fileSystemGenHTTPOpenapiYaml),
		GenHTTPOpenapi3JSON: http.FileServer(fileSystemGenHTTPOpenapi3JSON),
		GenHTTPOpenapi3Yaml: http.FileServer(fileSystemGenHTTPOpenapi3Yaml),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "openapi" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
}

// Mount configures the mux to serve the openapi endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGenHTTPOpenapiJSON(mux, goahttp.Replace("", "/gen/http/openapi.json", h.GenHTTPOpenapiJSON))
	MountGenHTTPOpenapiYaml(mux, goahttp.Replace("", "/gen/http/openapi.yaml", h.GenHTTPOpenapiYaml))
	MountGenHTTPOpenapi3JSON(mux, goahttp.Replace("", "/gen/http/openapi3.json", h.GenHTTPOpenapi3JSON))
	MountGenHTTPOpenapi3Yaml(mux, goahttp.Replace("", "/gen/http/openapi3.yaml", h.GenHTTPOpenapi3Yaml))
}

// Mount configures the mux to serve the openapi endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountGenHTTPOpenapiJSON configures the mux to serve GET request made to
// "/openapi.json".
func MountGenHTTPOpenapiJSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.json", h.ServeHTTP)
}

// MountGenHTTPOpenapiYaml configures the mux to serve GET request made to
// "/openapi.yaml".
func MountGenHTTPOpenapiYaml(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi.yaml", h.ServeHTTP)
}

// MountGenHTTPOpenapi3JSON configures the mux to serve GET request made to
// "/openapi3.json".
func MountGenHTTPOpenapi3JSON(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi3.json", h.ServeHTTP)
}

// MountGenHTTPOpenapi3Yaml configures the mux to serve GET request made to
// "/openapi3.yaml".
func MountGenHTTPOpenapi3Yaml(mux goahttp.Muxer, h http.Handler) {
	mux.Handle("GET", "/openapi3.yaml", h.ServeHTTP)
}
