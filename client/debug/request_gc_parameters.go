// Code generated by go-swagger; DO NOT EDIT.

package debug

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
)

// NewRequestGcParams creates a new RequestGcParams object
// with the default values initialized.
func NewRequestGcParams() *RequestGcParams {

	return &RequestGcParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRequestGcParamsWithTimeout creates a new RequestGcParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRequestGcParamsWithTimeout(timeout time.Duration) *RequestGcParams {

	return &RequestGcParams{

		timeout: timeout,
	}
}

// NewRequestGcParamsWithContext creates a new RequestGcParams object
// with the default values initialized, and the ability to set a context for a request
func NewRequestGcParamsWithContext(ctx context.Context) *RequestGcParams {

	return &RequestGcParams{

		Context: ctx,
	}
}

// NewRequestGcParamsWithHTTPClient creates a new RequestGcParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRequestGcParamsWithHTTPClient(client *http.Client) *RequestGcParams {

	return &RequestGcParams{
		HTTPClient: client,
	}
}

/*RequestGcParams contains all the parameters to send to the API endpoint
for the request gc operation typically these are written to a http.Request
*/
type RequestGcParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the request gc params
func (o *RequestGcParams) WithTimeout(timeout time.Duration) *RequestGcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the request gc params
func (o *RequestGcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the request gc params
func (o *RequestGcParams) WithContext(ctx context.Context) *RequestGcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the request gc params
func (o *RequestGcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the request gc params
func (o *RequestGcParams) WithHTTPClient(client *http.Client) *RequestGcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the request gc params
func (o *RequestGcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *RequestGcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}