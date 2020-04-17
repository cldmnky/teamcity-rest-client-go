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

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// NewSetBuildTypesOrderParams creates a new SetBuildTypesOrderParams object
// with the default values initialized.
func NewSetBuildTypesOrderParams() *SetBuildTypesOrderParams {
	var ()
	return &SetBuildTypesOrderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetBuildTypesOrderParamsWithTimeout creates a new SetBuildTypesOrderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetBuildTypesOrderParamsWithTimeout(timeout time.Duration) *SetBuildTypesOrderParams {
	var ()
	return &SetBuildTypesOrderParams{

		timeout: timeout,
	}
}

// NewSetBuildTypesOrderParamsWithContext creates a new SetBuildTypesOrderParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetBuildTypesOrderParamsWithContext(ctx context.Context) *SetBuildTypesOrderParams {
	var ()
	return &SetBuildTypesOrderParams{

		Context: ctx,
	}
}

// NewSetBuildTypesOrderParamsWithHTTPClient creates a new SetBuildTypesOrderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetBuildTypesOrderParamsWithHTTPClient(client *http.Client) *SetBuildTypesOrderParams {
	var ()
	return &SetBuildTypesOrderParams{
		HTTPClient: client,
	}
}

/*SetBuildTypesOrderParams contains all the parameters to send to the API endpoint
for the set build types order operation typically these are written to a http.Request
*/
type SetBuildTypesOrderParams struct {

	/*Body*/
	Body *models.BuildTypes
	/*Field*/
	Field string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set build types order params
func (o *SetBuildTypesOrderParams) WithTimeout(timeout time.Duration) *SetBuildTypesOrderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set build types order params
func (o *SetBuildTypesOrderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set build types order params
func (o *SetBuildTypesOrderParams) WithContext(ctx context.Context) *SetBuildTypesOrderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set build types order params
func (o *SetBuildTypesOrderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set build types order params
func (o *SetBuildTypesOrderParams) WithHTTPClient(client *http.Client) *SetBuildTypesOrderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set build types order params
func (o *SetBuildTypesOrderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set build types order params
func (o *SetBuildTypesOrderParams) WithBody(body *models.BuildTypes) *SetBuildTypesOrderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set build types order params
func (o *SetBuildTypesOrderParams) SetBody(body *models.BuildTypes) {
	o.Body = body
}

// WithField adds the field to the set build types order params
func (o *SetBuildTypesOrderParams) WithField(field string) *SetBuildTypesOrderParams {
	o.SetField(field)
	return o
}

// SetField adds the field to the set build types order params
func (o *SetBuildTypesOrderParams) SetField(field string) {
	o.Field = field
}

// WithProjectLocator adds the projectLocator to the set build types order params
func (o *SetBuildTypesOrderParams) WithProjectLocator(projectLocator string) *SetBuildTypesOrderParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the set build types order params
func (o *SetBuildTypesOrderParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *SetBuildTypesOrderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param field
	if err := r.SetPathParam("field", o.Field); err != nil {
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
