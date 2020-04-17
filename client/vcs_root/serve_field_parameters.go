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

// NewServeFieldParams creates a new ServeFieldParams object
// with the default values initialized.
func NewServeFieldParams() *ServeFieldParams {
	var ()
	return &ServeFieldParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeFieldParamsWithTimeout creates a new ServeFieldParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeFieldParamsWithTimeout(timeout time.Duration) *ServeFieldParams {
	var ()
	return &ServeFieldParams{

		timeout: timeout,
	}
}

// NewServeFieldParamsWithContext creates a new ServeFieldParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeFieldParamsWithContext(ctx context.Context) *ServeFieldParams {
	var ()
	return &ServeFieldParams{

		Context: ctx,
	}
}

// NewServeFieldParamsWithHTTPClient creates a new ServeFieldParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeFieldParamsWithHTTPClient(client *http.Client) *ServeFieldParams {
	var ()
	return &ServeFieldParams{
		HTTPClient: client,
	}
}

/*ServeFieldParams contains all the parameters to send to the API endpoint
for the serve field operation typically these are written to a http.Request
*/
type ServeFieldParams struct {

	/*Field*/
	Field string
	/*VcsRootLocator*/
	VcsRootLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve field params
func (o *ServeFieldParams) WithTimeout(timeout time.Duration) *ServeFieldParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve field params
func (o *ServeFieldParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve field params
func (o *ServeFieldParams) WithContext(ctx context.Context) *ServeFieldParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve field params
func (o *ServeFieldParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve field params
func (o *ServeFieldParams) WithHTTPClient(client *http.Client) *ServeFieldParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve field params
func (o *ServeFieldParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithField adds the field to the serve field params
func (o *ServeFieldParams) WithField(field string) *ServeFieldParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the serve field params
func (o *ServeFieldParams) SetField(field string) {
	o.Field = field
}

// WithVcsRootLocator adds the vcsRootLocator to the serve field params
func (o *ServeFieldParams) WithVcsRootLocator(vcsRootLocator string) *ServeFieldParams {
	o.SetVcsRootLocator(vcsRootLocator)
	return o
}

// SetVcsRootLocator adds the vcsRootLocator to the serve field params
func (o *ServeFieldParams) SetVcsRootLocator(vcsRootLocator string) {
	o.VcsRootLocator = vcsRootLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeFieldParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
		return err
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
