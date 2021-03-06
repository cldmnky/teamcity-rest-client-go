// Code generated by go-swagger; DO NOT EDIT.

package user

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

// NewRemoveGroupParams creates a new RemoveGroupParams object
// with the default values initialized.
func NewRemoveGroupParams() *RemoveGroupParams {
	var ()
	return &RemoveGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRemoveGroupParamsWithTimeout creates a new RemoveGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRemoveGroupParamsWithTimeout(timeout time.Duration) *RemoveGroupParams {
	var ()
	return &RemoveGroupParams{

		timeout: timeout,
	}
}

// NewRemoveGroupParamsWithContext creates a new RemoveGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewRemoveGroupParamsWithContext(ctx context.Context) *RemoveGroupParams {
	var ()
	return &RemoveGroupParams{

		Context: ctx,
	}
}

// NewRemoveGroupParamsWithHTTPClient creates a new RemoveGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRemoveGroupParamsWithHTTPClient(client *http.Client) *RemoveGroupParams {
	var ()
	return &RemoveGroupParams{
		HTTPClient: client,
	}
}

/*RemoveGroupParams contains all the parameters to send to the API endpoint
for the remove group operation typically these are written to a http.Request
*/
type RemoveGroupParams struct {

	/*Fields*/
	Fields *string
	/*GroupLocator*/
	GroupLocator string
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the remove group params
func (o *RemoveGroupParams) WithTimeout(timeout time.Duration) *RemoveGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the remove group params
func (o *RemoveGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the remove group params
func (o *RemoveGroupParams) WithContext(ctx context.Context) *RemoveGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the remove group params
func (o *RemoveGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the remove group params
func (o *RemoveGroupParams) WithHTTPClient(client *http.Client) *RemoveGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the remove group params
func (o *RemoveGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the remove group params
func (o *RemoveGroupParams) WithFields(fields *string) *RemoveGroupParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the remove group params
func (o *RemoveGroupParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithGroupLocator adds the groupLocator to the remove group params
func (o *RemoveGroupParams) WithGroupLocator(groupLocator string) *RemoveGroupParams {
	o.SetGroupLocator(groupLocator)
	return o
}

// SetGroupLocator adds the groupLocator to the remove group params
func (o *RemoveGroupParams) SetGroupLocator(groupLocator string) {
	o.GroupLocator = groupLocator
}

// WithUserLocator adds the userLocator to the remove group params
func (o *RemoveGroupParams) WithUserLocator(userLocator string) *RemoveGroupParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the remove group params
func (o *RemoveGroupParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *RemoveGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param groupLocator
	if err := r.SetPathParam("groupLocator", o.GroupLocator); err != nil {
		return err
	}

	// path param userLocator
	if err := r.SetPathParam("userLocator", o.UserLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
