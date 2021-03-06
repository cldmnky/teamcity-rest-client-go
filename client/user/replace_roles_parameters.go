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

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// NewReplaceRolesParams creates a new ReplaceRolesParams object
// with the default values initialized.
func NewReplaceRolesParams() *ReplaceRolesParams {
	var ()
	return &ReplaceRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceRolesParamsWithTimeout creates a new ReplaceRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceRolesParamsWithTimeout(timeout time.Duration) *ReplaceRolesParams {
	var ()
	return &ReplaceRolesParams{

		timeout: timeout,
	}
}

// NewReplaceRolesParamsWithContext creates a new ReplaceRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceRolesParamsWithContext(ctx context.Context) *ReplaceRolesParams {
	var ()
	return &ReplaceRolesParams{

		Context: ctx,
	}
}

// NewReplaceRolesParamsWithHTTPClient creates a new ReplaceRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceRolesParamsWithHTTPClient(client *http.Client) *ReplaceRolesParams {
	var ()
	return &ReplaceRolesParams{
		HTTPClient: client,
	}
}

/*ReplaceRolesParams contains all the parameters to send to the API endpoint
for the replace roles operation typically these are written to a http.Request
*/
type ReplaceRolesParams struct {

	/*Body*/
	Body *models.Roles
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace roles params
func (o *ReplaceRolesParams) WithTimeout(timeout time.Duration) *ReplaceRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace roles params
func (o *ReplaceRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace roles params
func (o *ReplaceRolesParams) WithContext(ctx context.Context) *ReplaceRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace roles params
func (o *ReplaceRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace roles params
func (o *ReplaceRolesParams) WithHTTPClient(client *http.Client) *ReplaceRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace roles params
func (o *ReplaceRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the replace roles params
func (o *ReplaceRolesParams) WithBody(body *models.Roles) *ReplaceRolesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the replace roles params
func (o *ReplaceRolesParams) SetBody(body *models.Roles) {
	o.Body = body
}

// WithUserLocator adds the userLocator to the replace roles params
func (o *ReplaceRolesParams) WithUserLocator(userLocator string) *ReplaceRolesParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the replace roles params
func (o *ReplaceRolesParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
