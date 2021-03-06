// Code generated by go-swagger; DO NOT EDIT.

package greeter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/wawa0210/go-grpc-journey/pkg/sdk/models"
)

// NewGreeterSayHelloParams creates a new GreeterSayHelloParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGreeterSayHelloParams() *GreeterSayHelloParams {
	return &GreeterSayHelloParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGreeterSayHelloParamsWithTimeout creates a new GreeterSayHelloParams object
// with the ability to set a timeout on a request.
func NewGreeterSayHelloParamsWithTimeout(timeout time.Duration) *GreeterSayHelloParams {
	return &GreeterSayHelloParams{
		timeout: timeout,
	}
}

// NewGreeterSayHelloParamsWithContext creates a new GreeterSayHelloParams object
// with the ability to set a context for a request.
func NewGreeterSayHelloParamsWithContext(ctx context.Context) *GreeterSayHelloParams {
	return &GreeterSayHelloParams{
		Context: ctx,
	}
}

// NewGreeterSayHelloParamsWithHTTPClient creates a new GreeterSayHelloParams object
// with the ability to set a custom HTTPClient for a request.
func NewGreeterSayHelloParamsWithHTTPClient(client *http.Client) *GreeterSayHelloParams {
	return &GreeterSayHelloParams{
		HTTPClient: client,
	}
}

/* GreeterSayHelloParams contains all the parameters to send to the API endpoint
   for the greeter say hello operation.

   Typically these are written to a http.Request.
*/
type GreeterSayHelloParams struct {

	// Body.
	Body *models.HelloworldHelloRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the greeter say hello params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GreeterSayHelloParams) WithDefaults() *GreeterSayHelloParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the greeter say hello params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GreeterSayHelloParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the greeter say hello params
func (o *GreeterSayHelloParams) WithTimeout(timeout time.Duration) *GreeterSayHelloParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the greeter say hello params
func (o *GreeterSayHelloParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the greeter say hello params
func (o *GreeterSayHelloParams) WithContext(ctx context.Context) *GreeterSayHelloParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the greeter say hello params
func (o *GreeterSayHelloParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the greeter say hello params
func (o *GreeterSayHelloParams) WithHTTPClient(client *http.Client) *GreeterSayHelloParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the greeter say hello params
func (o *GreeterSayHelloParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the greeter say hello params
func (o *GreeterSayHelloParams) WithBody(body *models.HelloworldHelloRequest) *GreeterSayHelloParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the greeter say hello params
func (o *GreeterSayHelloParams) SetBody(body *models.HelloworldHelloRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GreeterSayHelloParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
