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
)

// NewGetSnapshotDepParams creates a new GetSnapshotDepParams object
// with the default values initialized.
func NewGetSnapshotDepParams() *GetSnapshotDepParams {
	var ()
	return &GetSnapshotDepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSnapshotDepParamsWithTimeout creates a new GetSnapshotDepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSnapshotDepParamsWithTimeout(timeout time.Duration) *GetSnapshotDepParams {
	var ()
	return &GetSnapshotDepParams{

		timeout: timeout,
	}
}

// NewGetSnapshotDepParamsWithContext creates a new GetSnapshotDepParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSnapshotDepParamsWithContext(ctx context.Context) *GetSnapshotDepParams {
	var ()
	return &GetSnapshotDepParams{

		Context: ctx,
	}
}

// NewGetSnapshotDepParamsWithHTTPClient creates a new GetSnapshotDepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSnapshotDepParamsWithHTTPClient(client *http.Client) *GetSnapshotDepParams {
	var ()
	return &GetSnapshotDepParams{
		HTTPClient: client,
	}
}

/*GetSnapshotDepParams contains all the parameters to send to the API endpoint
for the get snapshot dep operation typically these are written to a http.Request
*/
type GetSnapshotDepParams struct {

	/*BtLocator*/
	BtLocator string
	/*Fields*/
	Fields *string
	/*SnapshotDepLocator*/
	SnapshotDepLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get snapshot dep params
func (o *GetSnapshotDepParams) WithTimeout(timeout time.Duration) *GetSnapshotDepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get snapshot dep params
func (o *GetSnapshotDepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get snapshot dep params
func (o *GetSnapshotDepParams) WithContext(ctx context.Context) *GetSnapshotDepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get snapshot dep params
func (o *GetSnapshotDepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get snapshot dep params
func (o *GetSnapshotDepParams) WithHTTPClient(client *http.Client) *GetSnapshotDepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get snapshot dep params
func (o *GetSnapshotDepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the get snapshot dep params
func (o *GetSnapshotDepParams) WithBtLocator(btLocator string) *GetSnapshotDepParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the get snapshot dep params
func (o *GetSnapshotDepParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFields adds the fields to the get snapshot dep params
func (o *GetSnapshotDepParams) WithFields(fields *string) *GetSnapshotDepParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get snapshot dep params
func (o *GetSnapshotDepParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithSnapshotDepLocator adds the snapshotDepLocator to the get snapshot dep params
func (o *GetSnapshotDepParams) WithSnapshotDepLocator(snapshotDepLocator string) *GetSnapshotDepParams {
	o.SetSnapshotDepLocator(snapshotDepLocator)
	return o
}

// SetSnapshotDepLocator adds the snapshotDepLocator to the get snapshot dep params
func (o *GetSnapshotDepParams) SetSnapshotDepLocator(snapshotDepLocator string) {
	o.SnapshotDepLocator = snapshotDepLocator
}

// WriteToRequest writes these params to a swagger request
func (o *GetSnapshotDepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param snapshotDepLocator
	if err := r.SetPathParam("snapshotDepLocator", o.SnapshotDepLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
