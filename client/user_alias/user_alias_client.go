// Code generated by go-swagger; DO NOT EDIT.

package user_alias

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new user alias API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user alias API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	UserAliasList(params *UserAliasListParams, authInfo runtime.ClientAuthInfoWriter) (*UserAliasListOK, error)

	UserAliasRead(params *UserAliasReadParams, authInfo runtime.ClientAuthInfoWriter) (*UserAliasReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  UserAliasList lists user alias entities

  List UserAlias entities

--- UserAlias ---

A UserAlias is a person or process that creates artifacts like tickets, pull requests, commits, and comments.
UserAliases are automatically collected from all sources and have an apex User object.

--- General ---

All objects can be traversed to related objects. This can be achieved by using the double-underscore syntax.
For instance, query all UserAlias records with a specific Team name like so:

"?team\_\_name=Bobs%20Team".

You can easily traverse multiple levels of relationships by connecting related objects. For example:

"?foo\_\_bar\_\_baz\_\_name=Bob"

All available relationships are defined in the filter documentation with the note "filter-traversable object."
Additionally, you can limit the properties returned in an object by using the `fields` or `omit`
keywords as follows:

"?fields=foo,bar,bad" will return only these fields

or

"?omit=foo,bar,baz" will return all but the specified fields
*/
func (a *Client) UserAliasList(params *UserAliasListParams, authInfo runtime.ClientAuthInfoWriter) (*UserAliasListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserAliasListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "user_alias_list",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/user_alias/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserAliasListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserAliasListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user_alias_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UserAliasRead gets a specific user alias entity by id value

  Get a specific UserAlias entity by id value
*/
func (a *Client) UserAliasRead(params *UserAliasReadParams, authInfo runtime.ClientAuthInfoWriter) (*UserAliasReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserAliasReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "user_alias_read",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/user_alias/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UserAliasReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UserAliasReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for user_alias_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
