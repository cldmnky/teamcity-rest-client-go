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

// NewGetExampleNewProjectDescriptionParams creates a new GetExampleNewProjectDescriptionParams object
// with the default values initialized.
func NewGetExampleNewProjectDescriptionParams() *GetExampleNewProjectDescriptionParams {
	var ()
	return &GetExampleNewProjectDescriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetExampleNewProjectDescriptionParamsWithTimeout creates a new GetExampleNewProjectDescriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetExampleNewProjectDescriptionParamsWithTimeout(timeout time.Duration) *GetExampleNewProjectDescriptionParams {
	var ()
	return &GetExampleNewProjectDescriptionParams{

		timeout: timeout,
	}
}

// NewGetExampleNewProjectDescriptionParamsWithContext creates a new GetExampleNewProjectDescriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetExampleNewProjectDescriptionParamsWithContext(ctx context.Context) *GetExampleNewProjectDescriptionParams {
	var ()
	return &GetExampleNewProjectDescriptionParams{

		Context: ctx,
	}
}

// NewGetExampleNewProjectDescriptionParamsWithHTTPClient creates a new GetExampleNewProjectDescriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetExampleNewProjectDescriptionParamsWithHTTPClient(client *http.Client) *GetExampleNewProjectDescriptionParams {
	var ()
	return &GetExampleNewProjectDescriptionParams{
		HTTPClient: client,
	}
}

/*GetExampleNewProjectDescriptionParams contains all the parameters to send to the API endpoint
for the get example new project description operation typically these are written to a http.Request
*/
type GetExampleNewProjectDescriptionParams struct {

	/*ID*/
	ID *string
	/*ProjectLocator*/
	ProjectLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get example new project description params
func (o *GetExampleNewProjectDescriptionParams) WithTimeout(timeout time.Duration) *GetExampleNewProjectDescriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get example new project description params
func (o *GetExampleNewProjectDescriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get example new project description params
func (o *GetExampleNewProjectDescriptionParams) WithContext(ctx context.Context) *GetExampleNewProjectDescriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get example new project description params
func (o *GetExampleNewProjectDescriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get example new project description params
func (o *GetExampleNewProjectDescriptionParams) WithHTTPClient(client *http.Client) *GetExampleNewProjectDescriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get example new project description params
func (o *GetExampleNewProjectDescriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get example new project description params
func (o *GetExampleNewProjectDescriptionParams) WithID(id *string) *GetExampleNewProjectDescriptionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get example new project description params
func (o *GetExampleNewProjectDescriptionParams) SetID(id *string) {
	o.ID = id
}

// WithProjectLocator adds the projectLocator to the get example new project description params
func (o *GetExampleNewProjectDescriptionParams) WithProjectLocator(projectLocator string) *GetExampleNewProjectDescriptionParams {
	o.SetProjectLocator(projectLocator)
	return o
}

// SetProjectLocator adds the projectLocator to the get example new project description params
func (o *GetExampleNewProjectDescriptionParams) SetProjectLocator(projectLocator string) {
	o.ProjectLocator = projectLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetExampleNewProjectDescriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string
		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {
			if err := r.SetQueryParam("id", qID); err != nil {
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
