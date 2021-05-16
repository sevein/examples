// Code generated by goa v3.4.0, DO NOT EDIT.
//
// calc HTTP client types
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	calc "goa.design/examples/error/gen/calc"
	goa "goa.design/goa/v3/pkg"
)

// DivideRequestBody is the type of the "calc" service "divide" endpoint HTTP
// request body.
type DivideRequestBody struct {
	Dividend int `form:"dividend" json:"dividend" xml:"dividend"`
	Divisor  int `form:"divisor" json:"divisor" xml:"divisor"`
}

// DivideResponseBody is the type of the "calc" service "divide" endpoint HTTP
// response body.
type DivideResponseBody struct {
	Quotient *int `form:"quotient,omitempty" json:"quotient,omitempty" xml:"quotient,omitempty"`
	Reminder *int `form:"reminder,omitempty" json:"reminder,omitempty" xml:"reminder,omitempty"`
}

// DivideDivByZeroResponseBody is the type of the "calc" service "divide"
// endpoint HTTP response body for the "div_by_zero" error.
type DivideDivByZeroResponseBody struct {
	// division by zero leads to infinity.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
}

// DivideTimeoutResponseBody is the type of the "calc" service "divide"
// endpoint HTTP response body for the "timeout" error.
type DivideTimeoutResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// NewDivideRequestBody builds the HTTP request body from the payload of the
// "divide" endpoint of the "calc" service.
func NewDivideRequestBody(p *calc.DividePayload) *DivideRequestBody {
	body := &DivideRequestBody{
		Dividend: p.Dividend,
		Divisor:  p.Divisor,
	}
	return body
}

// NewDivideResultOK builds a "calc" service "divide" endpoint result from a
// HTTP "OK" response.
func NewDivideResultOK(body *DivideResponseBody) *calc.DivideResult {
	v := &calc.DivideResult{
		Quotient: *body.Quotient,
		Reminder: *body.Reminder,
	}

	return v
}

// NewDivideDivByZero builds a calc service divide endpoint div_by_zero error.
func NewDivideDivByZero(body *DivideDivByZeroResponseBody) *calc.DivByZero {
	v := &calc.DivByZero{
		Message: *body.Message,
	}

	return v
}

// NewDivideTimeout builds a calc service divide endpoint timeout error.
func NewDivideTimeout(body *DivideTimeoutResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateDivideResponseBody runs the validations defined on DivideResponseBody
func ValidateDivideResponseBody(body *DivideResponseBody) (err error) {
	if body.Quotient == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("quotient", "body"))
	}
	if body.Reminder == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("reminder", "body"))
	}
	return
}

// ValidateDivideDivByZeroResponseBody runs the validations defined on
// divide_div_by_zero_response_body
func ValidateDivideDivByZeroResponseBody(body *DivideDivByZeroResponseBody) (err error) {
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	return
}

// ValidateDivideTimeoutResponseBody runs the validations defined on
// divide_timeout_response_body
func ValidateDivideTimeoutResponseBody(body *DivideTimeoutResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}
