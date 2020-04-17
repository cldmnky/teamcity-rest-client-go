// Code generated by go-swagger; DO NOT EDIT.

package federation

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

// NewAddServerParams creates a new AddServerParams object
// with the default values initialized.
func NewAddServerParams() *AddServerParams {
	var ()
	return &AddServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddServerParamsWithTimeout creates a new AddServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddServerParamsWithTimeout(timeout time.Duration) *AddServerParams {
	var ()
	return &AddServerParams{

		timeout: timeout,
	}
}

// NewAddServerParamsWithContext creates a new AddServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddServerParamsWithContext(ctx context.Context) *AddServerParams {
	var ()
	return &AddServerParams{

		Context: ctx,
	}
}

// NewAddServerParamsWithHTTPClient creates a new AddServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddServerParamsWithHTTPClient(client *http.Client) *AddServerParams {
	var ()
	return &AddServerParams{
		HTTPClient: client,
	}
}

/*AddServerParams contains all the parameters to send to the API endpoint
for the add server operation typically these are written to a http.Request
*/
type AddServerParams struct {

	/*Body*/
	Body *models.Servers

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add server params
func (o *AddServerParams) WithTimeout(timeout time.Duration) *AddServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add server params
func (o *AddServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add server params
func (o *AddServerParams) WithContext(ctx context.Context) *AddServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add server params
func (o *AddServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add server params
func (o *AddServerParams) WithHTTPClient(client *http.Client) *AddServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add server params
func (o *AddServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add server params
func (o *AddServerParams) WithBody(body *models.Servers) *AddServerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add server params
func (o *AddServerParams) SetBody(body *models.Servers) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
