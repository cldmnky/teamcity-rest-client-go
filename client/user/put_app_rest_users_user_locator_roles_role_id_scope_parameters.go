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

// NewPutAppRestUsersUserLocatorRolesRoleIDScopeParams creates a new PutAppRestUsersUserLocatorRolesRoleIDScopeParams object
// with the default values initialized.
func NewPutAppRestUsersUserLocatorRolesRoleIDScopeParams() *PutAppRestUsersUserLocatorRolesRoleIDScopeParams {
	var ()
	return &PutAppRestUsersUserLocatorRolesRoleIDScopeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAppRestUsersUserLocatorRolesRoleIDScopeParamsWithTimeout creates a new PutAppRestUsersUserLocatorRolesRoleIDScopeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAppRestUsersUserLocatorRolesRoleIDScopeParamsWithTimeout(timeout time.Duration) *PutAppRestUsersUserLocatorRolesRoleIDScopeParams {
	var ()
	return &PutAppRestUsersUserLocatorRolesRoleIDScopeParams{

		timeout: timeout,
	}
}

// NewPutAppRestUsersUserLocatorRolesRoleIDScopeParamsWithContext creates a new PutAppRestUsersUserLocatorRolesRoleIDScopeParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAppRestUsersUserLocatorRolesRoleIDScopeParamsWithContext(ctx context.Context) *PutAppRestUsersUserLocatorRolesRoleIDScopeParams {
	var ()
	return &PutAppRestUsersUserLocatorRolesRoleIDScopeParams{

		Context: ctx,
	}
}

// NewPutAppRestUsersUserLocatorRolesRoleIDScopeParamsWithHTTPClient creates a new PutAppRestUsersUserLocatorRolesRoleIDScopeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAppRestUsersUserLocatorRolesRoleIDScopeParamsWithHTTPClient(client *http.Client) *PutAppRestUsersUserLocatorRolesRoleIDScopeParams {
	var ()
	return &PutAppRestUsersUserLocatorRolesRoleIDScopeParams{
		HTTPClient: client,
	}
}

/*PutAppRestUsersUserLocatorRolesRoleIDScopeParams contains all the parameters to send to the API endpoint
for the put app rest users user locator roles role ID scope operation typically these are written to a http.Request
*/
type PutAppRestUsersUserLocatorRolesRoleIDScopeParams struct {

	/*RoleID*/
	RoleID string
	/*Scope*/
	Scope string
	/*UserLocator*/
	UserLocator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) WithTimeout(timeout time.Duration) *PutAppRestUsersUserLocatorRolesRoleIDScopeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) WithContext(ctx context.Context) *PutAppRestUsersUserLocatorRolesRoleIDScopeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) WithHTTPClient(client *http.Client) *PutAppRestUsersUserLocatorRolesRoleIDScopeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleID adds the roleID to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) WithRoleID(roleID string) *PutAppRestUsersUserLocatorRolesRoleIDScopeParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WithScope adds the scope to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) WithScope(scope string) *PutAppRestUsersUserLocatorRolesRoleIDScopeParams {
	o.SetScope(scope)
	return o
}

// SetScope adds the scope to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) SetScope(scope string) {
	o.Scope = scope
}

// WithUserLocator adds the userLocator to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) WithUserLocator(userLocator string) *PutAppRestUsersUserLocatorRolesRoleIDScopeParams {
	o.SetUserLocator(userLocator)
	return o
}

// SetUserLocator adds the userLocator to the put app rest users user locator roles role ID scope params
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) SetUserLocator(userLocator string) {
	o.UserLocator = userLocator
}

// WriteToRequest writes these params to a swagger request
func (o *PutAppRestUsersUserLocatorRolesRoleIDScopeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param roleId
	if err := r.SetPathParam("roleId", o.RoleID); err != nil {
		return err
	}

	// path param scope
	if err := r.SetPathParam("scope", o.Scope); err != nil {
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
