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

// NewGetArtifactDependencyChangesParams creates a new GetArtifactDependencyChangesParams object
// with the default values initialized.
func NewGetArtifactDependencyChangesParams() *GetArtifactDependencyChangesParams {
	var ()
	return &GetArtifactDependencyChangesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetArtifactDependencyChangesParamsWithTimeout creates a new GetArtifactDependencyChangesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetArtifactDependencyChangesParamsWithTimeout(timeout time.Duration) *GetArtifactDependencyChangesParams {
	var ()
	return &GetArtifactDependencyChangesParams{

		timeout: timeout,
	}
}

// NewGetArtifactDependencyChangesParamsWithContext creates a new GetArtifactDependencyChangesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetArtifactDependencyChangesParamsWithContext(ctx context.Context) *GetArtifactDependencyChangesParams {
	var ()
	return &GetArtifactDependencyChangesParams{

		Context: ctx,
	}
}

// NewGetArtifactDependencyChangesParamsWithHTTPClient creates a new GetArtifactDependencyChangesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetArtifactDependencyChangesParamsWithHTTPClient(client *http.Client) *GetArtifactDependencyChangesParams {
	var ()
	return &GetArtifactDependencyChangesParams{
		HTTPClient: client,
	}
}

/*GetArtifactDependencyChangesParams contains all the parameters to send to the API endpoint
for the get artifact dependency changes operation typically these are written to a http.Request
*/
type GetArtifactDependencyChangesParams struct {

	/*BuildLocator*/
	BuildLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get artifact dependency changes params
func (o *GetArtifactDependencyChangesParams) WithTimeout(timeout time.Duration) *GetArtifactDependencyChangesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get artifact dependency changes params
func (o *GetArtifactDependencyChangesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get artifact dependency changes params
func (o *GetArtifactDependencyChangesParams) WithContext(ctx context.Context) *GetArtifactDependencyChangesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get artifact dependency changes params
func (o *GetArtifactDependencyChangesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get artifact dependency changes params
func (o *GetArtifactDependencyChangesParams) WithHTTPClient(client *http.Client) *GetArtifactDependencyChangesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get artifact dependency changes params
func (o *GetArtifactDependencyChangesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the get artifact dependency changes params
func (o *GetArtifactDependencyChangesParams) WithBuildLocator(buildLocator string) *GetArtifactDependencyChangesParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the get artifact dependency changes params
func (o *GetArtifactDependencyChangesParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the get artifact dependency changes params
func (o *GetArtifactDependencyChangesParams) WithFields(fields *string) *GetArtifactDependencyChangesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get artifact dependency changes params
func (o *GetArtifactDependencyChangesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetArtifactDependencyChangesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
