// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewServePluginsParams creates a new ServePluginsParams object
// with the default values initialized.
func NewServePluginsParams() *ServePluginsParams {
	var ()
	return &ServePluginsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServePluginsParamsWithTimeout creates a new ServePluginsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServePluginsParamsWithTimeout(timeout time.Duration) *ServePluginsParams {
	var ()
	return &ServePluginsParams{

		timeout: timeout,
	}
}

// NewServePluginsParamsWithContext creates a new ServePluginsParams object
// with the default values initialized, and the ability to set a context for a request
func NewServePluginsParamsWithContext(ctx context.Context) *ServePluginsParams {
	var ()
	return &ServePluginsParams{

		Context: ctx,
	}
}

// NewServePluginsParamsWithHTTPClient creates a new ServePluginsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServePluginsParamsWithHTTPClient(client *http.Client) *ServePluginsParams {
	var ()
	return &ServePluginsParams{
		HTTPClient: client,
	}
}

/*ServePluginsParams contains all the parameters to send to the API endpoint
for the serve plugins operation typically these are written to a http.Request
*/
type ServePluginsParams struct {

	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve plugins params
func (o *ServePluginsParams) WithTimeout(timeout time.Duration) *ServePluginsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve plugins params
func (o *ServePluginsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve plugins params
func (o *ServePluginsParams) WithContext(ctx context.Context) *ServePluginsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve plugins params
func (o *ServePluginsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve plugins params
func (o *ServePluginsParams) WithHTTPClient(client *http.Client) *ServePluginsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve plugins params
func (o *ServePluginsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve plugins params
func (o *ServePluginsParams) WithFields(fields *string) *ServePluginsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve plugins params
func (o *ServePluginsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ServePluginsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
