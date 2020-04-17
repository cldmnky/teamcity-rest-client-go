// Code generated by go-swagger; DO NOT EDIT.

package cloud_instance

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

// NewStartInstanceParams creates a new StartInstanceParams object
// with the default values initialized.
func NewStartInstanceParams() *StartInstanceParams {
	var ()
	return &StartInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStartInstanceParamsWithTimeout creates a new StartInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStartInstanceParamsWithTimeout(timeout time.Duration) *StartInstanceParams {
	var ()
	return &StartInstanceParams{

		timeout: timeout,
	}
}

// NewStartInstanceParamsWithContext creates a new StartInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewStartInstanceParamsWithContext(ctx context.Context) *StartInstanceParams {
	var ()
	return &StartInstanceParams{

		Context: ctx,
	}
}

// NewStartInstanceParamsWithHTTPClient creates a new StartInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStartInstanceParamsWithHTTPClient(client *http.Client) *StartInstanceParams {
	var ()
	return &StartInstanceParams{
		HTTPClient: client,
	}
}

/*StartInstanceParams contains all the parameters to send to the API endpoint
for the start instance operation typically these are written to a http.Request
*/
type StartInstanceParams struct {

	/*Body*/
	Body *models.CloudInstance
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the start instance params
func (o *StartInstanceParams) WithTimeout(timeout time.Duration) *StartInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the start instance params
func (o *StartInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the start instance params
func (o *StartInstanceParams) WithContext(ctx context.Context) *StartInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the start instance params
func (o *StartInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the start instance params
func (o *StartInstanceParams) WithHTTPClient(client *http.Client) *StartInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the start instance params
func (o *StartInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the start instance params
func (o *StartInstanceParams) WithBody(body *models.CloudInstance) *StartInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the start instance params
func (o *StartInstanceParams) SetBody(body *models.CloudInstance) {
	o.Body = body
}

// WithFields adds the fields to the start instance params
func (o *StartInstanceParams) WithFields(fields *string) *StartInstanceParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the start instance params
func (o *StartInstanceParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *StartInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
