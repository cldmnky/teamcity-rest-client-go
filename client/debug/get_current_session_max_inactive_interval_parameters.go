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

// NewGetCurrentSessionMaxInactiveIntervalParams creates a new GetCurrentSessionMaxInactiveIntervalParams object
// with the default values initialized.
func NewGetCurrentSessionMaxInactiveIntervalParams() *GetCurrentSessionMaxInactiveIntervalParams {

	return &GetCurrentSessionMaxInactiveIntervalParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrentSessionMaxInactiveIntervalParamsWithTimeout creates a new GetCurrentSessionMaxInactiveIntervalParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCurrentSessionMaxInactiveIntervalParamsWithTimeout(timeout time.Duration) *GetCurrentSessionMaxInactiveIntervalParams {

	return &GetCurrentSessionMaxInactiveIntervalParams{

		timeout: timeout,
	}
}

// NewGetCurrentSessionMaxInactiveIntervalParamsWithContext creates a new GetCurrentSessionMaxInactiveIntervalParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCurrentSessionMaxInactiveIntervalParamsWithContext(ctx context.Context) *GetCurrentSessionMaxInactiveIntervalParams {

	return &GetCurrentSessionMaxInactiveIntervalParams{

		Context: ctx,
	}
}

// NewGetCurrentSessionMaxInactiveIntervalParamsWithHTTPClient creates a new GetCurrentSessionMaxInactiveIntervalParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCurrentSessionMaxInactiveIntervalParamsWithHTTPClient(client *http.Client) *GetCurrentSessionMaxInactiveIntervalParams {

	return &GetCurrentSessionMaxInactiveIntervalParams{
		HTTPClient: client,
	}
}

/*GetCurrentSessionMaxInactiveIntervalParams contains all the parameters to send to the API endpoint
for the get current session max inactive interval operation typically these are written to a http.Request
*/
type GetCurrentSessionMaxInactiveIntervalParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get current session max inactive interval params
func (o *GetCurrentSessionMaxInactiveIntervalParams) WithTimeout(timeout time.Duration) *GetCurrentSessionMaxInactiveIntervalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get current session max inactive interval params
func (o *GetCurrentSessionMaxInactiveIntervalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get current session max inactive interval params
func (o *GetCurrentSessionMaxInactiveIntervalParams) WithContext(ctx context.Context) *GetCurrentSessionMaxInactiveIntervalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get current session max inactive interval params
func (o *GetCurrentSessionMaxInactiveIntervalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get current session max inactive interval params
func (o *GetCurrentSessionMaxInactiveIntervalParams) WithHTTPClient(client *http.Client) *GetCurrentSessionMaxInactiveIntervalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get current session max inactive interval params
func (o *GetCurrentSessionMaxInactiveIntervalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrentSessionMaxInactiveIntervalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
