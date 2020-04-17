// Code generated by go-swagger; DO NOT EDIT.

package investigation

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

// NewReplaceInstanceParams creates a new ReplaceInstanceParams object
// with the default values initialized.
func NewReplaceInstanceParams() *ReplaceInstanceParams {
	var ()
	return &ReplaceInstanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceInstanceParamsWithTimeout creates a new ReplaceInstanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceInstanceParamsWithTimeout(timeout time.Duration) *ReplaceInstanceParams {
	var ()
	return &ReplaceInstanceParams{

		timeout: timeout,
	}
}

// NewReplaceInstanceParamsWithContext creates a new ReplaceInstanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceInstanceParamsWithContext(ctx context.Context) *ReplaceInstanceParams {
	var ()
	return &ReplaceInstanceParams{

		Context: ctx,
	}
}

// NewReplaceInstanceParamsWithHTTPClient creates a new ReplaceInstanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceInstanceParamsWithHTTPClient(client *http.Client) *ReplaceInstanceParams {
	var ()
	return &ReplaceInstanceParams{
		HTTPClient: client,
	}
}

/*ReplaceInstanceParams contains all the parameters to send to the API endpoint
for the replace instance operation typically these are written to a http.Request
*/
type ReplaceInstanceParams struct {

	/*Body*/
	Body *models.Investigation
	/*Fields*/
	Fields *string
	/*InvestigationLocator*/
	InvestigationLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace instance params
func (o *ReplaceInstanceParams) WithTimeout(timeout time.Duration) *ReplaceInstanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace instance params
func (o *ReplaceInstanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace instance params
func (o *ReplaceInstanceParams) WithContext(ctx context.Context) *ReplaceInstanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace instance params
func (o *ReplaceInstanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace instance params
func (o *ReplaceInstanceParams) WithHTTPClient(client *http.Client) *ReplaceInstanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace instance params
func (o *ReplaceInstanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace instance params
func (o *ReplaceInstanceParams) WithBody(body *models.Investigation) *ReplaceInstanceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace instance params
func (o *ReplaceInstanceParams) SetBody(body *models.Investigation) {
	o.Body = body
}

// WithFields adds the fields to the replace instance params
func (o *ReplaceInstanceParams) WithFields(fields *string) *ReplaceInstanceParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the replace instance params
func (o *ReplaceInstanceParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithInvestigationLocator adds the investigationLocator to the replace instance params
func (o *ReplaceInstanceParams) WithInvestigationLocator(investigationLocator string) *ReplaceInstanceParams {
	o.SetInvestigationLocator(investigationLocator)
	return o
}

// SetInvestigationLocator adds the investigationLocator to the replace instance params
func (o *ReplaceInstanceParams) SetInvestigationLocator(investigationLocator string) {
	o.InvestigationLocator = investigationLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceInstanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param investigationLocator
	if err := r.SetPathParam("investigationLocator", o.InvestigationLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
