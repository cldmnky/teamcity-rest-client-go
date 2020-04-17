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

// NewServeAggregatedBuildStatusParams creates a new ServeAggregatedBuildStatusParams object
// with the default values initialized.
func NewServeAggregatedBuildStatusParams() *ServeAggregatedBuildStatusParams {
	var ()
	return &ServeAggregatedBuildStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeAggregatedBuildStatusParamsWithTimeout creates a new ServeAggregatedBuildStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeAggregatedBuildStatusParamsWithTimeout(timeout time.Duration) *ServeAggregatedBuildStatusParams {
	var ()
	return &ServeAggregatedBuildStatusParams{

		timeout: timeout,
	}
}

// NewServeAggregatedBuildStatusParamsWithContext creates a new ServeAggregatedBuildStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeAggregatedBuildStatusParamsWithContext(ctx context.Context) *ServeAggregatedBuildStatusParams {
	var ()
	return &ServeAggregatedBuildStatusParams{

		Context: ctx,
	}
}

// NewServeAggregatedBuildStatusParamsWithHTTPClient creates a new ServeAggregatedBuildStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeAggregatedBuildStatusParamsWithHTTPClient(client *http.Client) *ServeAggregatedBuildStatusParams {
	var ()
	return &ServeAggregatedBuildStatusParams{
		HTTPClient: client,
	}
}

/*ServeAggregatedBuildStatusParams contains all the parameters to send to the API endpoint
for the serve aggregated build status operation typically these are written to a http.Request
*/
type ServeAggregatedBuildStatusParams struct {

	/*BuildLocator*/
	BuildLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve aggregated build status params
func (o *ServeAggregatedBuildStatusParams) WithTimeout(timeout time.Duration) *ServeAggregatedBuildStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve aggregated build status params
func (o *ServeAggregatedBuildStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve aggregated build status params
func (o *ServeAggregatedBuildStatusParams) WithContext(ctx context.Context) *ServeAggregatedBuildStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve aggregated build status params
func (o *ServeAggregatedBuildStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve aggregated build status params
func (o *ServeAggregatedBuildStatusParams) WithHTTPClient(client *http.Client) *ServeAggregatedBuildStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve aggregated build status params
func (o *ServeAggregatedBuildStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the serve aggregated build status params
func (o *ServeAggregatedBuildStatusParams) WithBuildLocator(buildLocator string) *ServeAggregatedBuildStatusParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the serve aggregated build status params
func (o *ServeAggregatedBuildStatusParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ServeAggregatedBuildStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param buildLocator
	if err := r.SetPathParam("buildLocator", o.BuildLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
