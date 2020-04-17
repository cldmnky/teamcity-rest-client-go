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

// NewServeBuildWithProjectParams creates a new ServeBuildWithProjectParams object
// with the default values initialized.
func NewServeBuildWithProjectParams() *ServeBuildWithProjectParams {
	var ()
	return &ServeBuildWithProjectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildWithProjectParamsWithTimeout creates a new ServeBuildWithProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildWithProjectParamsWithTimeout(timeout time.Duration) *ServeBuildWithProjectParams {
	var ()
	return &ServeBuildWithProjectParams{

		timeout: timeout,
	}
}

// NewServeBuildWithProjectParamsWithContext creates a new ServeBuildWithProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildWithProjectParamsWithContext(ctx context.Context) *ServeBuildWithProjectParams {
	var ()
	return &ServeBuildWithProjectParams{

		Context: ctx,
	}
}

// NewServeBuildWithProjectParamsWithHTTPClient creates a new ServeBuildWithProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildWithProjectParamsWithHTTPClient(client *http.Client) *ServeBuildWithProjectParams {
	var ()
	return &ServeBuildWithProjectParams{
		HTTPClient: client,
	}
}

/*ServeBuildWithProjectParams contains all the parameters to send to the API endpoint
for the serve build with project operation typically these are written to a http.Request
*/
type ServeBuildWithProjectParams struct {

	/*BtLocator*/
	BtLocator string
	/*BuildLocator*/
	BuildLocator string
	/*Fields*/
	Fields *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build with project params
func (o *ServeBuildWithProjectParams) WithTimeout(timeout time.Duration) *ServeBuildWithProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build with project params
func (o *ServeBuildWithProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build with project params
func (o *ServeBuildWithProjectParams) WithContext(ctx context.Context) *ServeBuildWithProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build with project params
func (o *ServeBuildWithProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build with project params
func (o *ServeBuildWithProjectParams) WithHTTPClient(client *http.Client) *ServeBuildWithProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build with project params
func (o *ServeBuildWithProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the serve build with project params
func (o *ServeBuildWithProjectParams) WithBtLocator(btLocator string) *ServeBuildWithProjectParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the serve build with project params
func (o *ServeBuildWithProjectParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithBuildLocator adds the buildLocator to the serve build with project params
func (o *ServeBuildWithProjectParams) WithBuildLocator(buildLocator string) *ServeBuildWithProjectParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the serve build with project params
func (o *ServeBuildWithProjectParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the serve build with project params
func (o *ServeBuildWithProjectParams) WithFields(fields *string) *ServeBuildWithProjectParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve build with project params
func (o *ServeBuildWithProjectParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithProjectLocator adds the projectLocator to the serve build with project params
func (o *ServeBuildWithProjectParams) WithProjectLocator(projectLocator string) *ServeBuildWithProjectParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the serve build with project params
func (o *ServeBuildWithProjectParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildWithProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
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
