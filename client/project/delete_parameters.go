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

// NewDeleteParams creates a new DeleteParams object
// with the default values initialized.
func NewDeleteParams() *DeleteParams {
	var ()
	return &DeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteParamsWithTimeout creates a new DeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteParamsWithTimeout(timeout time.Duration) *DeleteParams {
	var ()
	return &DeleteParams{

		timeout: timeout,
	}
}

// NewDeleteParamsWithContext creates a new DeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteParamsWithContext(ctx context.Context) *DeleteParams {
	var ()
	return &DeleteParams{

		Context: ctx,
	}
}

// NewDeleteParamsWithHTTPClient creates a new DeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteParamsWithHTTPClient(client *http.Client) *DeleteParams {
	var ()
	return &DeleteParams{
		HTTPClient: client,
	}
}

/*DeleteParams contains all the parameters to send to the API endpoint
for the delete operation typically these are written to a http.Request
*/
type DeleteParams struct {

	/*FeatureLocator*/
	FeatureLocator string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete params
func (o *DeleteParams) WithTimeout(timeout time.Duration) *DeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete params
func (o *DeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete params
func (o *DeleteParams) WithContext(ctx context.Context) *DeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete params
func (o *DeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete params
func (o *DeleteParams) WithHTTPClient(client *http.Client) *DeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete params
func (o *DeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFeatureLocator adds the featureLocator to the delete params
func (o *DeleteParams) WithFeatureLocator(featureLocator string) *DeleteParams {
	o.SetFeatureLocator(featureLocator)
	return o
}

// SetFeatureLocator adds the featureLocator to the delete params
func (o *DeleteParams) SetFeatureLocator(featureLocator string) {
	o.FeatureLocator = featureLocator
}

// WithProjectLocator adds the projectLocator to the delete params
func (o *DeleteParams) WithProjectLocator(projectLocator string) *DeleteParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the delete params
func (o *DeleteParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param featureLocator
	if err := r.SetPathParam("featureLocator", o.FeatureLocator); err != nil {
		return err
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
