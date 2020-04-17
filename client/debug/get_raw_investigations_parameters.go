// Code generated by go-swagger; DO NOT EDIT.

package debug

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

// NewGetRawInvestigationsParams creates a new GetRawInvestigationsParams object
// with the default values initialized.
func NewGetRawInvestigationsParams() *GetRawInvestigationsParams {
	var ()
	return &GetRawInvestigationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRawInvestigationsParamsWithTimeout creates a new GetRawInvestigationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRawInvestigationsParamsWithTimeout(timeout time.Duration) *GetRawInvestigationsParams {
	var ()
	return &GetRawInvestigationsParams{

		timeout: timeout,
	}
}

// NewGetRawInvestigationsParamsWithContext creates a new GetRawInvestigationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRawInvestigationsParamsWithContext(ctx context.Context) *GetRawInvestigationsParams {
	var ()
	return &GetRawInvestigationsParams{

		Context: ctx,
	}
}

// NewGetRawInvestigationsParamsWithHTTPClient creates a new GetRawInvestigationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRawInvestigationsParamsWithHTTPClient(client *http.Client) *GetRawInvestigationsParams {
	var ()
	return &GetRawInvestigationsParams{
		HTTPClient: client,
	}
}

/*GetRawInvestigationsParams contains all the parameters to send to the API endpoint
for the get raw investigations operation typically these are written to a http.Request
*/
type GetRawInvestigationsParams struct {

	/*Fields*/
	Fields *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get raw investigations params
func (o *GetRawInvestigationsParams) WithTimeout(timeout time.Duration) *GetRawInvestigationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get raw investigations params
func (o *GetRawInvestigationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get raw investigations params
func (o *GetRawInvestigationsParams) WithContext(ctx context.Context) *GetRawInvestigationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get raw investigations params
func (o *GetRawInvestigationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get raw investigations params
func (o *GetRawInvestigationsParams) WithHTTPClient(client *http.Client) *GetRawInvestigationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get raw investigations params
func (o *GetRawInvestigationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get raw investigations params
func (o *GetRawInvestigationsParams) WithFields(fields *string) *GetRawInvestigationsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get raw investigations params
func (o *GetRawInvestigationsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WriteToRequest writes these params to a swagger request
func (o *GetRawInvestigationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
