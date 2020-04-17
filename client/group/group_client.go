// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	AddRoleSimple(params *AddRoleSimpleParams, authInfo runtime.ClientAuthInfoWriter) (*AddRoleSimpleOK, error)

	DeleteGroup(params *DeleteGroupParams, authInfo runtime.ClientAuthInfoWriter) error

	GetParentGroups(params *GetParentGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetParentGroupsOK, error)

	GetProperties(params *GetPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPropertiesOK, error)

	ServeGroup(params *ServeGroupParams, authInfo runtime.ClientAuthInfoWriter) (*ServeGroupOK, error)

	ServeGroups(params *ServeGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*ServeGroupsOK, error)

	SetParentGroups(params *SetParentGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*SetParentGroupsOK, error)

	SetRoles(params *SetRolesParams, authInfo runtime.ClientAuthInfoWriter) (*SetRolesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddRoleSimple add role simple API
*/
func (a *Client) AddRoleSimple(params *AddRoleSimpleParams, authInfo runtime.ClientAuthInfoWriter) (*AddRoleSimpleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRoleSimpleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addRoleSimple",
		Method:             "POST",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/roles/{roleId}/{scope}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddRoleSimpleReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddRoleSimpleOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for addRoleSimple: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteGroup delete group API
*/
func (a *Client) DeleteGroup(params *DeleteGroupParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteGroupParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteGroup",
		Method:             "DELETE",
		PathPattern:        "/app/rest/userGroups/{groupLocator}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  GetParentGroups get parent groups API
*/
func (a *Client) GetParentGroups(params *GetParentGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*GetParentGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetParentGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getParentGroups",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/parent-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetParentGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetParentGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getParentGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetProperties get properties API
*/
func (a *Client) GetProperties(params *GetPropertiesParams, authInfo runtime.ClientAuthInfoWriter) (*GetPropertiesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPropertiesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getProperties",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/properties",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetPropertiesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPropertiesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getProperties: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ServeGroup serve group API
*/
func (a *Client) ServeGroup(params *ServeGroupParams, authInfo runtime.ClientAuthInfoWriter) (*ServeGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveGroup",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups/{groupLocator}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServeGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serveGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ServeGroups serve groups API
*/
func (a *Client) ServeGroups(params *ServeGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*ServeGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveGroups",
		Method:             "GET",
		PathPattern:        "/app/rest/userGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServeGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serveGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetParentGroups set parent groups API
*/
func (a *Client) SetParentGroups(params *SetParentGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*SetParentGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetParentGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setParentGroups",
		Method:             "PUT",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/parent-groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetParentGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetParentGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setParentGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetRoles set roles API
*/
func (a *Client) SetRoles(params *SetRolesParams, authInfo runtime.ClientAuthInfoWriter) (*SetRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetRolesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setRoles",
		Method:             "PUT",
		PathPattern:        "/app/rest/userGroups/{groupLocator}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SetRolesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetRolesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setRoles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
