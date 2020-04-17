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

// NewServeBuildRelatedIssuesParams creates a new ServeBuildRelatedIssuesParams object
// with the default values initialized.
func NewServeBuildRelatedIssuesParams() *ServeBuildRelatedIssuesParams {
	var ()
	return &ServeBuildRelatedIssuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewServeBuildRelatedIssuesParamsWithTimeout creates a new ServeBuildRelatedIssuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewServeBuildRelatedIssuesParamsWithTimeout(timeout time.Duration) *ServeBuildRelatedIssuesParams {
	var ()
	return &ServeBuildRelatedIssuesParams{

		timeout: timeout,
	}
}

// NewServeBuildRelatedIssuesParamsWithContext creates a new ServeBuildRelatedIssuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewServeBuildRelatedIssuesParamsWithContext(ctx context.Context) *ServeBuildRelatedIssuesParams {
	var ()
	return &ServeBuildRelatedIssuesParams{

		Context: ctx,
	}
}

// NewServeBuildRelatedIssuesParamsWithHTTPClient creates a new ServeBuildRelatedIssuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewServeBuildRelatedIssuesParamsWithHTTPClient(client *http.Client) *ServeBuildRelatedIssuesParams {
	var ()
	return &ServeBuildRelatedIssuesParams{
		HTTPClient: client,
	}
}

/*ServeBuildRelatedIssuesParams contains all the parameters to send to the API endpoint
for the serve build related issues operation typically these are written to a http.Request
*/
type ServeBuildRelatedIssuesParams struct {

	/*BuildLocator*/
	BuildLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the serve build related issues params
func (o *ServeBuildRelatedIssuesParams) WithTimeout(timeout time.Duration) *ServeBuildRelatedIssuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the serve build related issues params
func (o *ServeBuildRelatedIssuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the serve build related issues params
func (o *ServeBuildRelatedIssuesParams) WithContext(ctx context.Context) *ServeBuildRelatedIssuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the serve build related issues params
func (o *ServeBuildRelatedIssuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the serve build related issues params
func (o *ServeBuildRelatedIssuesParams) WithHTTPClient(client *http.Client) *ServeBuildRelatedIssuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the serve build related issues params
func (o *ServeBuildRelatedIssuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBuildLocator adds the buildLocator to the serve build related issues params
func (o *ServeBuildRelatedIssuesParams) WithBuildLocator(buildLocator string) *ServeBuildRelatedIssuesParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the serve build related issues params
func (o *ServeBuildRelatedIssuesParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the serve build related issues params
func (o *ServeBuildRelatedIssuesParams) WithFields(fields *string) *ServeBuildRelatedIssuesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the serve build related issues params
func (o *ServeBuildRelatedIssuesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ServeBuildRelatedIssuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
