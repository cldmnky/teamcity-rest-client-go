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

// NewServeBuildTypeParams creates a new ServeBuildTypeParams object
// with the default values initialized.
func NewServeBuildTypeParams() *ServeBuildTypeParams {
	var ()
	return &ServeBuildTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildTypeParamsWithTimeout creates a new ServeBuildTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildTypeParamsWithTimeout(timeout time.Duration) *ServeBuildTypeParams {
	var ()
	return &ServeBuildTypeParams{

		timeout: timeout,
	}
}

// NewServeBuildTypeParamsWithContext creates a new ServeBuildTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildTypeParamsWithContext(ctx context.Context) *ServeBuildTypeParams {
	var ()
	return &ServeBuildTypeParams{

		Context: ctx,
	}
}

// NewServeBuildTypeParamsWithHTTPClient creates a new ServeBuildTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildTypeParamsWithHTTPClient(client *http.Client) *ServeBuildTypeParams {
	var ()
	return &ServeBuildTypeParams{
		HTTPClient: client,
	}
}

/*ServeBuildTypeParams contains all the parameters to send to the API endpoint
for the serve build type operation typically these are written to a http.Request
*/
type ServeBuildTypeParams struct {

	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build type params
func (o *ServeBuildTypeParams) WithTimeout(timeout time.Duration) *ServeBuildTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build type params
func (o *ServeBuildTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build type params
func (o *ServeBuildTypeParams) WithContext(ctx context.Context) *ServeBuildTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build type params
func (o *ServeBuildTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build type params
func (o *ServeBuildTypeParams) WithHTTPClient(client *http.Client) *ServeBuildTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build type params
func (o *ServeBuildTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the serve build type params
func (o *ServeBuildTypeParams) WithBtLocator(btLocator string) *ServeBuildTypeParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the serve build type params
func (o *ServeBuildTypeParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the serve build type params
func (o *ServeBuildTypeParams) WithFields(fields *string) *ServeBuildTypeParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve build type params
func (o *ServeBuildTypeParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the serve build type params
func (o *ServeBuildTypeParams) WithProjectLocator(projectLocator string) *ServeBuildTypeParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the serve build type params
func (o *ServeBuildTypeParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

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
