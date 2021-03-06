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

// NewDeleteSnapshotDepParams creates a new DeleteSnapshotDepParams object
// with the default values initialized.
func NewDeleteSnapshotDepParams() *DeleteSnapshotDepParams {
	var ()
	return &DeleteSnapshotDepParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSnapshotDepParamsWithTimeout creates a new DeleteSnapshotDepParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSnapshotDepParamsWithTimeout(timeout time.Duration) *DeleteSnapshotDepParams {
	var ()
	return &DeleteSnapshotDepParams{

		timeout: timeout,
	}
}

// NewDeleteSnapshotDepParamsWithContext creates a new DeleteSnapshotDepParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSnapshotDepParamsWithContext(ctx context.Context) *DeleteSnapshotDepParams {
	var ()
	return &DeleteSnapshotDepParams{

		Context: ctx,
	}
}

// NewDeleteSnapshotDepParamsWithHTTPClient creates a new DeleteSnapshotDepParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSnapshotDepParamsWithHTTPClient(client *http.Client) *DeleteSnapshotDepParams {
	var ()
	return &DeleteSnapshotDepParams{
		HTTPClient: client,
	}
}

/*DeleteSnapshotDepParams contains all the parameters to send to the API endpoint
for the delete snapshot dep operation typically these are written to a http.Request
*/
type DeleteSnapshotDepParams struct {

	/*BtLocator*/
	BtLocator string
	/*SnapshotDepLocator*/
	SnapshotDepLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete snapshot dep params
func (o *DeleteSnapshotDepParams) WithTimeout(timeout time.Duration) *DeleteSnapshotDepParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete snapshot dep params
func (o *DeleteSnapshotDepParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete snapshot dep params
func (o *DeleteSnapshotDepParams) WithContext(ctx context.Context) *DeleteSnapshotDepParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete snapshot dep params
func (o *DeleteSnapshotDepParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete snapshot dep params
func (o *DeleteSnapshotDepParams) WithHTTPClient(client *http.Client) *DeleteSnapshotDepParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete snapshot dep params
func (o *DeleteSnapshotDepParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBtLocator adds the btLocator to the delete snapshot dep params
func (o *DeleteSnapshotDepParams) WithBtLocator(btLocator string) *DeleteSnapshotDepParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the delete snapshot dep params
func (o *DeleteSnapshotDepParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithSnapshotDepLocator adds the snapshotDepLocator to the delete snapshot dep params
func (o *DeleteSnapshotDepParams) WithSnapshotDepLocator(snapshotDepLocator string) *DeleteSnapshotDepParams {
	o.SetSnapshotDepLocator(snapshotDepLocator)
	return o
}

// SetSnapshotDepLocator adds the snapshotDepLocator to the delete snapshot dep params
func (o *DeleteSnapshotDepParams) SetSnapshotDepLocator(snapshotDepLocator string) {
	o.SnapshotDepLocator = snapshotDepLocator
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSnapshotDepParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
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
