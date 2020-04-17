// Code generated by go-swagger; DO NOT EDIT.

package build

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

// NewServeBuildStatusIconParams creates a new ServeBuildStatusIconParams object
// with the default values initialized.
func NewServeBuildStatusIconParams() *ServeBuildStatusIconParams {
	var ()
	return &ServeBuildStatusIconParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildStatusIconParamsWithTimeout creates a new ServeBuildStatusIconParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildStatusIconParamsWithTimeout(timeout time.Duration) *ServeBuildStatusIconParams {
	var ()
	return &ServeBuildStatusIconParams{

		timeout: timeout,
	}
}

// NewServeBuildStatusIconParamsWithContext creates a new ServeBuildStatusIconParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildStatusIconParamsWithContext(ctx context.Context) *ServeBuildStatusIconParams {
	var ()
	return &ServeBuildStatusIconParams{

		Context: ctx,
	}
}

// NewServeBuildStatusIconParamsWithHTTPClient creates a new ServeBuildStatusIconParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildStatusIconParamsWithHTTPClient(client *http.Client) *ServeBuildStatusIconParams {
	var ()
	return &ServeBuildStatusIconParams{
		HTTPClient: client,
	}
}

/*ServeBuildStatusIconParams contains all the parameters to send to the API endpoint
for the serve build status icon operation typically these are written to a http.Request
*/
type ServeBuildStatusIconParams struct {

	/*BuildLocator*/
	BuildLocator string
	/*Suffix*/
	Suffix string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build status icon params
func (o *ServeBuildStatusIconParams) WithTimeout(timeout time.Duration) *ServeBuildStatusIconParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build status icon params
func (o *ServeBuildStatusIconParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build status icon params
func (o *ServeBuildStatusIconParams) WithContext(ctx context.Context) *ServeBuildStatusIconParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build status icon params
func (o *ServeBuildStatusIconParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build status icon params
func (o *ServeBuildStatusIconParams) WithHTTPClient(client *http.Client) *ServeBuildStatusIconParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build status icon params
func (o *ServeBuildStatusIconParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the serve build status icon params
func (o *ServeBuildStatusIconParams) WithBuildLocator(buildLocator string) *ServeBuildStatusIconParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the serve build status icon params
func (o *ServeBuildStatusIconParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithSuffix adds the suffix to the serve build status icon params
func (o *ServeBuildStatusIconParams) WithSuffix(suffix string) *ServeBuildStatusIconParams {
	o.SetSuffix(suffix)
	return o
}

// SetSuffix adds the suffix to the serve build status icon params
func (o *ServeBuildStatusIconParams) SetSuffix(suffix string) {
	o.Suffix = suffix
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildStatusIconParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	// path param suffix
	if err := r.SetPathParam("suffix", o.Suffix); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
