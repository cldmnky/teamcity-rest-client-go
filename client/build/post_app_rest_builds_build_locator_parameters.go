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

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// NewPostAppRestBuildsBuildLocatorParams creates a new PostAppRestBuildsBuildLocatorParams object
// with the default values initialized.
func NewPostAppRestBuildsBuildLocatorParams() *PostAppRestBuildsBuildLocatorParams {
	var ()
	return &PostAppRestBuildsBuildLocatorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAppRestBuildsBuildLocatorParamsWithTimeout creates a new PostAppRestBuildsBuildLocatorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAppRestBuildsBuildLocatorParamsWithTimeout(timeout time.Duration) *PostAppRestBuildsBuildLocatorParams {
	var ()
	return &PostAppRestBuildsBuildLocatorParams{

		timeout: timeout,
	}
}

// NewPostAppRestBuildsBuildLocatorParamsWithContext creates a new PostAppRestBuildsBuildLocatorParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAppRestBuildsBuildLocatorParamsWithContext(ctx context.Context) *PostAppRestBuildsBuildLocatorParams {
	var ()
	return &PostAppRestBuildsBuildLocatorParams{

		Context: ctx,
	}
}

// NewPostAppRestBuildsBuildLocatorParamsWithHTTPClient creates a new PostAppRestBuildsBuildLocatorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAppRestBuildsBuildLocatorParamsWithHTTPClient(client *http.Client) *PostAppRestBuildsBuildLocatorParams {
	var ()
	return &PostAppRestBuildsBuildLocatorParams{
		HTTPClient: client,
	}
}

/*PostAppRestBuildsBuildLocatorParams contains all the parameters to send to the API endpoint
for the post app rest builds build locator operation typically these are written to a http.Request
*/
type PostAppRestBuildsBuildLocatorParams struct {

	/*Body*/
	Body *models.BuildCancelRequest
	/*BuildLocator*/
	BuildLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) WithTimeout(timeout time.Duration) *PostAppRestBuildsBuildLocatorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) WithContext(ctx context.Context) *PostAppRestBuildsBuildLocatorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) WithHTTPClient(client *http.Client) *PostAppRestBuildsBuildLocatorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) WithBody(body *models.BuildCancelRequest) *PostAppRestBuildsBuildLocatorParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) SetBody(body *models.BuildCancelRequest) {
	o.Body = body
}

// WithBuildLocator adds the buildLocator to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) WithBuildLocator(buildLocator string) *PostAppRestBuildsBuildLocatorParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) WithFields(fields *string) *PostAppRestBuildsBuildLocatorParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the post app rest builds build locator params
func (o *PostAppRestBuildsBuildLocatorParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *PostAppRestBuildsBuildLocatorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
