// Code generated by go-swagger; DO NOT EDIT.

package group_settings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new group settings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for group settings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GroupSettingsList(params *GroupSettingsListParams, authInfo runtime.ClientAuthInfoWriter) (*GroupSettingsListOK, error)

	GroupSettingsRead(params *GroupSettingsReadParams, authInfo runtime.ClientAuthInfoWriter) (*GroupSettingsReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GroupSettingsList lists group settings entities

  List GroupSettings entities

--- GroupSettings ---

Settings that determine functionality associated with group/pair/mob programming.
*/
func (a *Client) GroupSettingsList(params *GroupSettingsListParams, authInfo runtime.ClientAuthInfoWriter) (*GroupSettingsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGroupSettingsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "group_settings_list",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/group_settings/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GroupSettingsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GroupSettingsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for group_settings_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GroupSettingsRead gets a specific group settings entity by id value

  Get a specific GroupSettings entity by id value
*/
func (a *Client) GroupSettingsRead(params *GroupSettingsReadParams, authInfo runtime.ClientAuthInfoWriter) (*GroupSettingsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGroupSettingsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "group_settings_read",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/group_settings/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GroupSettingsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GroupSettingsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for group_settings_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
