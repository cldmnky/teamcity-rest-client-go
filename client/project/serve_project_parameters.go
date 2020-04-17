// Code generated by go-swagger; DO NOT EDIT.

package project

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

// NewServeProjectParams creates a new ServeProjectParams object
// with the default values initialized.
func NewServeProjectParams() *ServeProjectParams {
	var ()
	return &ServeProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeProjectParamsWithTimeout creates a new ServeProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeProjectParamsWithTimeout(timeout time.Duration) *ServeProjectParams {
	var ()
	return &ServeProjectParams{

		timeout: timeout,
	}
}

// NewServeProjectParamsWithContext creates a new ServeProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeProjectParamsWithContext(ctx context.Context) *ServeProjectParams {
	var ()
	return &ServeProjectParams{

		Context: ctx,
	}
}

// NewServeProjectParamsWithHTTPClient creates a new ServeProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeProjectParamsWithHTTPClient(client *http.Client) *ServeProjectParams {
	var ()
	return &ServeProjectParams{
		HTTPClient: client,
	}
}

/*ServeProjectParams contains all the parameters to send to the API endpoint
for the serve project operation typically these are written to a http.Request
*/
type ServeProjectParams struct {

	/*Fields*/
	Fields *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve project params
func (o *ServeProjectParams) WithTimeout(timeout time.Duration) *ServeProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve project params
func (o *ServeProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve project params
func (o *ServeProjectParams) WithContext(ctx context.Context) *ServeProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve project params
func (o *ServeProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve project params
func (o *ServeProjectParams) WithHTTPClient(client *http.Client) *ServeProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve project params
func (o *ServeProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve project params
func (o *ServeProjectParams) WithFields(fields *string) *ServeProjectParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve project params
func (o *ServeProjectParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the serve project params
func (o *ServeProjectParams) WithProjectLocator(projectLocator string) *ServeProjectParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the serve project params
func (o *ServeProjectParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param projectLocator
	if err := r.SetPathParam("projectLocator", o.ProjectLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
