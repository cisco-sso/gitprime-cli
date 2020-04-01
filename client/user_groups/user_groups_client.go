// Code generated by go-swagger; DO NOT EDIT.

package user_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user groups API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user groups API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	UserGroupsList(params *UserGroupsListParams, authInfo runtime.ClientAuthInfoWriter) (*UserGroupsListOK, error)

	UserGroupsRead(params *UserGroupsReadParams, authInfo runtime.ClientAuthInfoWriter) (*UserGroupsReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UserGroupsList lists apex group entities

  List ApexGroup entities

--- ApexGroup ---

An ApexGroup is a collection of ApexUsers that authored a commit together etc.

--- General ---

All objects can be traversed to related objects. This can be achieved by using the double-underscore syntax.
For instance, query all ApexGroup records within an org with name like so:

"?org\_\_name=GitPrime".

You can easily traverse multiple levels of relationships by connecting related objects. For example:

"?foo\_\_bar\_\_baz\_\_name=Bob"

All available relationships are defined in the filter documentation with the note "filter-traversable object."
Additionally, you can limit the properties returned in an object by using the `fields` or `omit`
keywords as follows:

"?fields=foo,bar,bad" will return only these fields

or

"?omit=foo,bar,baz" will return all but the specified fields
*/
func (a *Client) UserGroupsList(params *UserGroupsListParams, authInfo runtime.ClientAuthInfoWriter) (*UserGroupsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "user_groups_list",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/user_groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserGroupsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserGroupsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user_groups_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserGroupsRead gets a specific apex group entity by id value

  Get a specific ApexGroup entity by id value
*/
func (a *Client) UserGroupsRead(params *UserGroupsReadParams, authInfo runtime.ClientAuthInfoWriter) (*UserGroupsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserGroupsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "user_groups_read",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/user_groups/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserGroupsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserGroupsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user_groups_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}