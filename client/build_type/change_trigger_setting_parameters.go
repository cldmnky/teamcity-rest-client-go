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

// NewChangeTriggerSettingParams creates a new ChangeTriggerSettingParams object
// with the default values initialized.
func NewChangeTriggerSettingParams() *ChangeTriggerSettingParams {
	var ()
	return &ChangeTriggerSettingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChangeTriggerSettingParamsWithTimeout creates a new ChangeTriggerSettingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChangeTriggerSettingParamsWithTimeout(timeout time.Duration) *ChangeTriggerSettingParams {
	var ()
	return &ChangeTriggerSettingParams{

		timeout: timeout,
	}
}

// NewChangeTriggerSettingParamsWithContext creates a new ChangeTriggerSettingParams object
// with the default values initialized, and the ability to set a context for a request
func NewChangeTriggerSettingParamsWithContext(ctx context.Context) *ChangeTriggerSettingParams {
	var ()
	return &ChangeTriggerSettingParams{

		Context: ctx,
	}
}

// NewChangeTriggerSettingParamsWithHTTPClient creates a new ChangeTriggerSettingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChangeTriggerSettingParamsWithHTTPClient(client *http.Client) *ChangeTriggerSettingParams {
	var ()
	return &ChangeTriggerSettingParams{
		HTTPClient: client,
	}
}

/*ChangeTriggerSettingParams contains all the parameters to send to the API endpoint
for the change trigger setting operation typically these are written to a http.Request
*/
type ChangeTriggerSettingParams struct {

	/*Body*/
	Body string
	/*BtLocator*/
	BtLocator string
	/*FieldName*/
	FieldName string
	/*TriggerLocator*/
	TriggerLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the change trigger setting params
func (o *ChangeTriggerSettingParams) WithTimeout(timeout time.Duration) *ChangeTriggerSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change trigger setting params
func (o *ChangeTriggerSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change trigger setting params
func (o *ChangeTriggerSettingParams) WithContext(ctx context.Context) *ChangeTriggerSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change trigger setting params
func (o *ChangeTriggerSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change trigger setting params
func (o *ChangeTriggerSettingParams) WithHTTPClient(client *http.Client) *ChangeTriggerSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change trigger setting params
func (o *ChangeTriggerSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change trigger setting params
func (o *ChangeTriggerSettingParams) WithBody(body string) *ChangeTriggerSettingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change trigger setting params
func (o *ChangeTriggerSettingParams) SetBody(body string) {
	o.Body = body
}

// WithBtLocator adds the btLocator to the change trigger setting params
func (o *ChangeTriggerSettingParams) WithBtLocator(btLocator string) *ChangeTriggerSettingParams {
	o.SetBtLocator(btLocator)
	return o
}

// SetBtLocator adds the btLocator to the change trigger setting params
func (o *ChangeTriggerSettingParams) SetBtLocator(btLocator string) {
	o.BtLocator = btLocator
}

// WithFieldName adds the fieldName to the change trigger setting params
func (o *ChangeTriggerSettingParams) WithFieldName(fieldName string) *ChangeTriggerSettingParams {
	o.SetFieldName(fieldName)
	return o
}

// SetFieldName adds the fieldName to the change trigger setting params
func (o *ChangeTriggerSettingParams) SetFieldName(fieldName string) {
	o.FieldName = fieldName
}

// WithTriggerLocator adds the triggerLocator to the change trigger setting params
func (o *ChangeTriggerSettingParams) WithTriggerLocator(triggerLocator string) *ChangeTriggerSettingParams {
	o.SetTriggerLocator(triggerLocator)
	return o
}

// SetTriggerLocator adds the triggerLocator to the change trigger setting params
func (o *ChangeTriggerSettingParams) SetTriggerLocator(triggerLocator string) {
	o.TriggerLocator = triggerLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeTriggerSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param btLocator
	if err := r.SetPathParam("btLocator", o.BtLocator); err != nil {
		return err
	}

	// path param fieldName
	if err := r.SetPathParam("fieldName", o.FieldName); err != nil {
		return err
	}

	// path param triggerLocator
	if err := r.SetPathParam("triggerLocator", o.TriggerLocator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
