// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

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

// NewServeRootInstancesParams creates a new ServeRootInstancesParams object
// with the default values initialized.
func NewServeRootInstancesParams() *ServeRootInstancesParams {
	var ()
	return &ServeRootInstancesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeRootInstancesParamsWithTimeout creates a new ServeRootInstancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeRootInstancesParamsWithTimeout(timeout time.Duration) *ServeRootInstancesParams {
	var ()
	return &ServeRootInstancesParams{

		timeout: timeout,
	}
}

// NewServeRootInstancesParamsWithContext creates a new ServeRootInstancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeRootInstancesParamsWithContext(ctx context.Context) *ServeRootInstancesParams {
	var ()
	return &ServeRootInstancesParams{

		Context: ctx,
	}
}

// NewServeRootInstancesParamsWithHTTPClient creates a new ServeRootInstancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeRootInstancesParamsWithHTTPClient(client *http.Client) *ServeRootInstancesParams {
	var ()
	return &ServeRootInstancesParams{
		HTTPClient: client,
	}
}

/*ServeRootInstancesParams contains all the parameters to send to the API endpoint
for the serve root instances operation typically these are written to a http.Request
*/
type ServeRootInstancesParams struct {

	/*Fields*/
	Fields *string
	/*VcsRootLocator*/
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve root instances params
func (o *ServeRootInstancesParams) WithTimeout(timeout time.Duration) *ServeRootInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve root instances params
func (o *ServeRootInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve root instances params
func (o *ServeRootInstancesParams) WithContext(ctx context.Context) *ServeRootInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve root instances params
func (o *ServeRootInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve root instances params
func (o *ServeRootInstancesParams) WithHTTPClient(client *http.Client) *ServeRootInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve root instances params
func (o *ServeRootInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the serve root instances params
func (o *ServeRootInstancesParams) WithFields(fields *string) *ServeRootInstancesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve root instances params
func (o *ServeRootInstancesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithVcsRootLocator adds the vcsRootLocator to the serve root instances params
func (o *ServeRootInstancesParams) WithVcsRootLocator(vcsRootLocator string) *ServeRootInstancesParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the serve root instances params
func (o *ServeRootInstancesParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeRootInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param vcsRootLocator
	if err := r.SetPathParam("vcsRootLocator", o.VcsRootLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
