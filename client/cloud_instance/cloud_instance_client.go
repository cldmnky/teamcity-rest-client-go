// Code generated by go-swagger; DO NOT EDIT.

package cloud_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new cloud instance API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for cloud instance API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ServeImage(params *ServeImageParams, authInfo runtime.ClientAuthInfoWriter) (*ServeImageOK, error)

	ServeImages(params *ServeImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ServeImagesOK, error)

	ServeProfile(params *ServeProfileParams, authInfo runtime.ClientAuthInfoWriter) (*ServeProfileOK, error)

	ServeProfiles(params *ServeProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*ServeProfilesOK, error)

	StartInstance(params *StartInstanceParams, authInfo runtime.ClientAuthInfoWriter) error

	StopInstance(params *StopInstanceParams, authInfo runtime.ClientAuthInfoWriter) error

	SetTransport(transport runtime.ClientTransport)
}

/*
  ServeImage serve image API
*/
func (a *Client) ServeImage(params *ServeImageParams, authInfo runtime.ClientAuthInfoWriter) (*ServeImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveImage",
		Method:             "GET",
		PathPattern:        "/app/rest/cloud/images/{imageLocator}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServeImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serveImage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ServeImages serve images API
*/
func (a *Client) ServeImages(params *ServeImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ServeImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveImages",
		Method:             "GET",
		PathPattern:        "/app/rest/cloud/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServeImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serveImages: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ServeProfile serve profile API
*/
func (a *Client) ServeProfile(params *ServeProfileParams, authInfo runtime.ClientAuthInfoWriter) (*ServeProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeProfileParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveProfile",
		Method:             "GET",
		PathPattern:        "/app/rest/cloud/profiles/{profileLocator}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeProfileReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServeProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serveProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ServeProfiles serve profiles API
*/
func (a *Client) ServeProfiles(params *ServeProfilesParams, authInfo runtime.ClientAuthInfoWriter) (*ServeProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServeProfilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "serveProfiles",
		Method:             "GET",
		PathPattern:        "/app/rest/cloud/profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ServeProfilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServeProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serveProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  StartInstance start instance API
*/
func (a *Client) StartInstance(params *StartInstanceParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartInstanceParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startInstance",
		Method:             "POST",
		PathPattern:        "/app/rest/cloud/instances",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StartInstanceReader{formats: a.formats},
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
  StopInstance stop instance API
*/
func (a *Client) StopInstance(params *StopInstanceParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopInstanceParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopInstance",
		Method:             "DELETE",
		PathPattern:        "/app/rest/cloud/instances/{instanceLocator}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &StopInstanceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
