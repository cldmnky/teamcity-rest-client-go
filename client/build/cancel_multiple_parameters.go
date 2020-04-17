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

// NewCancelMultipleParams creates a new CancelMultipleParams object
// with the default values initialized.
func NewCancelMultipleParams() *CancelMultipleParams {
	var ()
	return &CancelMultipleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCancelMultipleParamsWithTimeout creates a new CancelMultipleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCancelMultipleParamsWithTimeout(timeout time.Duration) *CancelMultipleParams {
	var ()
	return &CancelMultipleParams{

		timeout: timeout,
	}
}

// NewCancelMultipleParamsWithContext creates a new CancelMultipleParams object
// with the default values initialized, and the ability to set a context for a request
func NewCancelMultipleParamsWithContext(ctx context.Context) *CancelMultipleParams {
	var ()
	return &CancelMultipleParams{

		Context: ctx,
	}
}

// NewCancelMultipleParamsWithHTTPClient creates a new CancelMultipleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCancelMultipleParamsWithHTTPClient(client *http.Client) *CancelMultipleParams {
	var ()
	return &CancelMultipleParams{
		HTTPClient: client,
	}
}

/*CancelMultipleParams contains all the parameters to send to the API endpoint
for the cancel multiple operation typically these are written to a http.Request
*/
type CancelMultipleParams struct {

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

// WithTimeout adds the timeout to the cancel multiple params
func (o *CancelMultipleParams) WithTimeout(timeout time.Duration) *CancelMultipleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel multiple params
func (o *CancelMultipleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel multiple params
func (o *CancelMultipleParams) WithContext(ctx context.Context) *CancelMultipleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel multiple params
func (o *CancelMultipleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel multiple params
func (o *CancelMultipleParams) WithHTTPClient(client *http.Client) *CancelMultipleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel multiple params
func (o *CancelMultipleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the cancel multiple params
func (o *CancelMultipleParams) WithBody(body *models.BuildCancelRequest) *CancelMultipleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the cancel multiple params
func (o *CancelMultipleParams) SetBody(body *models.BuildCancelRequest) {
	o.Body = body
}

// WithBuildLocator adds the buildLocator to the cancel multiple params
func (o *CancelMultipleParams) WithBuildLocator(buildLocator string) *CancelMultipleParams {
	o.SetBuildLocator(buildLocator)
	return o
}

// SetBuildLocator adds the buildLocator to the cancel multiple params
func (o *CancelMultipleParams) SetBuildLocator(buildLocator string) {
	o.BuildLocator = buildLocator
}

// WithFields adds the fields to the cancel multiple params
func (o *CancelMultipleParams) WithFields(fields *string) *CancelMultipleParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the cancel multiple params
func (o *CancelMultipleParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *CancelMultipleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
