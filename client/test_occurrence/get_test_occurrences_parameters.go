// Code generated by go-swagger; DO NOT EDIT.

package test_occurrence

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

// NewGetTestOccurrencesParams creates a new GetTestOccurrencesParams object
// with the default values initialized.
func NewGetTestOccurrencesParams() *GetTestOccurrencesParams {
	var ()
	return &GetTestOccurrencesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTestOccurrencesParamsWithTimeout creates a new GetTestOccurrencesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTestOccurrencesParamsWithTimeout(timeout time.Duration) *GetTestOccurrencesParams {
	var ()
	return &GetTestOccurrencesParams{

		timeout: timeout,
	}
}

// NewGetTestOccurrencesParamsWithContext creates a new GetTestOccurrencesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTestOccurrencesParamsWithContext(ctx context.Context) *GetTestOccurrencesParams {
	var ()
	return &GetTestOccurrencesParams{

		Context: ctx,
	}
}

// NewGetTestOccurrencesParamsWithHTTPClient creates a new GetTestOccurrencesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTestOccurrencesParamsWithHTTPClient(client *http.Client) *GetTestOccurrencesParams {
	var ()
	return &GetTestOccurrencesParams{
		HTTPClient: client,
	}
}

/*GetTestOccurrencesParams contains all the parameters to send to the API endpoint
for the get test occurrences operation typically these are written to a http.Request
*/
type GetTestOccurrencesParams struct {

	/*Fields*/
	Fields *string
	/*Locator*/
	Locator *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get test occurrences params
func (o *GetTestOccurrencesParams) WithTimeout(timeout time.Duration) *GetTestOccurrencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get test occurrences params
func (o *GetTestOccurrencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get test occurrences params
func (o *GetTestOccurrencesParams) WithContext(ctx context.Context) *GetTestOccurrencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get test occurrences params
func (o *GetTestOccurrencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get test occurrences params
func (o *GetTestOccurrencesParams) WithHTTPClient(client *http.Client) *GetTestOccurrencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get test occurrences params
func (o *GetTestOccurrencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get test occurrences params
func (o *GetTestOccurrencesParams) WithFields(fields *string) *GetTestOccurrencesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get test occurrences params
func (o *GetTestOccurrencesParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLocator adds the locator to the get test occurrences params
func (o *GetTestOccurrencesParams) WithLocator(locator *string) *GetTestOccurrencesParams {
	o.SetLocator(locator)
	return o
}

// SetLocator adds the locator to the get test occurrences params
func (o *GetTestOccurrencesParams) SetLocator(locator *string) {
	o.Locator = locator
}

// WriteToRequest writes these params to a swagger request
func (o *GetTestOccurrencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Locator != nil {

		// query param locator
		var qrLocator string
		if o.Locator != nil {
			qrLocator = *o.Locator
		}
		qLocator := qrLocator
		if qLocator != "" {
			if err := r.SetQueryParam("locator", qLocator); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
