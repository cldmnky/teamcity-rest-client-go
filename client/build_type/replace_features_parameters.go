// Code generated by go-swagger; DO NOT EDIT.

package build_type

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

// NewReplaceFeaturesParams creates a new ReplaceFeaturesParams object
// with the default values initialized.
func NewReplaceFeaturesParams() *ReplaceFeaturesParams {
	var ()
	return &ReplaceFeaturesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceFeaturesParamsWithTimeout creates a new ReplaceFeaturesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceFeaturesParamsWithTimeout(timeout time.Duration) *ReplaceFeaturesParams {
	var ()
	return &ReplaceFeaturesParams{

		timeout: timeout,
	}
}

// NewReplaceFeaturesParamsWithContext creates a new ReplaceFeaturesParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceFeaturesParamsWithContext(ctx context.Context) *ReplaceFeaturesParams {
	var ()
	return &ReplaceFeaturesParams{

		Context: ctx,
	}
}

// NewReplaceFeaturesParamsWithHTTPClient creates a new ReplaceFeaturesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceFeaturesParamsWithHTTPClient(client *http.Client) *ReplaceFeaturesParams {
	var ()
	return &ReplaceFeaturesParams{
		HTTPClient: client,
	}
}

/*ReplaceFeaturesParams contains all the parameters to send to the API endpoint
for the replace features operation typically these are written to a http.Request
*/
type ReplaceFeaturesParams struct {

	/*Body*/
	Body *models.Features
	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace features params
func (o *ReplaceFeaturesParams) WithTimeout(timeout time.Duration) *ReplaceFeaturesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace features params
func (o *ReplaceFeaturesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace features params
func (o *ReplaceFeaturesParams) WithContext(ctx context.Context) *ReplaceFeaturesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace features params
func (o *ReplaceFeaturesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace features params
func (o *ReplaceFeaturesParams) WithHTTPClient(client *http.Client) *ReplaceFeaturesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace features params
func (o *ReplaceFeaturesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace features params
func (o *ReplaceFeaturesParams) WithBody(body *models.Features) *ReplaceFeaturesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace features params
func (o *ReplaceFeaturesParams) SetBody(body *models.Features) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the replace features params
func (o *ReplaceFeaturesParams) WithBtLocator(btLocator string) *ReplaceFeaturesParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the replace features params
func (o *ReplaceFeaturesParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the replace features params
func (o *ReplaceFeaturesParams) WithFields(fields *string) *ReplaceFeaturesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace features params
func (o *ReplaceFeaturesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceFeaturesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
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
