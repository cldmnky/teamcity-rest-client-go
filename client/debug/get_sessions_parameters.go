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
	"github.com/go-openapi/swag"
)

// NewGetSessionsParams creates a new GetSessionsParams object
// with the default values initialized.
func NewGetSessionsParams() *GetSessionsParams {
	var ()
	return &GetSessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSessionsParamsWithTimeout creates a new GetSessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSessionsParamsWithTimeout(timeout time.Duration) *GetSessionsParams {
	var ()
	return &GetSessionsParams{

		timeout: timeout,
	}
}

// NewGetSessionsParamsWithContext creates a new GetSessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSessionsParamsWithContext(ctx context.Context) *GetSessionsParams {
	var ()
	return &GetSessionsParams{

		Context: ctx,
	}
}

// NewGetSessionsParamsWithHTTPClient creates a new GetSessionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSessionsParamsWithHTTPClient(client *http.Client) *GetSessionsParams {
	var ()
	return &GetSessionsParams{
		HTTPClient: client,
	}
}

/*GetSessionsParams contains all the parameters to send to the API endpoint
for the get sessions operation typically these are written to a http.Request
*/
type GetSessionsParams struct {

	/*Fields*/
	Fields *string
	/*Manager*/
	Manager *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sessions params
func (o *GetSessionsParams) WithTimeout(timeout time.Duration) *GetSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sessions params
func (o *GetSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sessions params
func (o *GetSessionsParams) WithContext(ctx context.Context) *GetSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sessions params
func (o *GetSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sessions params
func (o *GetSessionsParams) WithHTTPClient(client *http.Client) *GetSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sessions params
func (o *GetSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get sessions params
func (o *GetSessionsParams) WithFields(fields *string) *GetSessionsParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get sessions params
func (o *GetSessionsParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithManager adds the manager to the get sessions params
func (o *GetSessionsParams) WithManager(manager *int64) *GetSessionsParams {
	o.SetManager(manager)
	return o
}

// SetManager adds the manager to the get sessions params
func (o *GetSessionsParams) SetManager(manager *int64) {
	o.Manager = manager
}

// WriteToRequest writes these params to a swagger request
func (o *GetSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Manager != nil {

		// query param manager
		var qrManager int64
		if o.Manager != nil {
			qrManager = *o.Manager
		}
		qManager := swag.FormatInt64(qrManager)
		if qManager != "" {
			if err := r.SetQueryParam("manager", qManager); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
