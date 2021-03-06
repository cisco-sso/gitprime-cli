// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new teams API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for teams API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TeamsCreate(params *TeamsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsCreateCreated, error)

	TeamsDelete(params *TeamsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsDeleteNoContent, error)

	TeamsList(params *TeamsListParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsListOK, error)

	TeamsRead(params *TeamsReadParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsReadOK, error)

	TeamsUpdate(params *TeamsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TeamsCreate creates a new team

  Create a new team

The data that is entered is a JSON String.  For example:

{ "visibility": "show", "name": "A New Team", "parent": null, "description": "Some team from the API" }

If you'd like to create a child team, pass the team ID of a parent:

{ "visibility": "show", "name": "A Child Team", "parent": 100, "description": "A child team of Team ID 100" }
*/
func (a *Client) TeamsCreate(params *TeamsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_create",
		Method:             "POST",
		PathPattern:        "/v3/customer/core/teams/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TeamsCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for teams_create: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TeamsDelete deletes a specific team entity by id value

  Delete a specific Team entity by id value
*/
func (a *Client) TeamsDelete(params *TeamsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_delete",
		Method:             "DELETE",
		PathPattern:        "/v3/customer/core/teams/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TeamsDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for teams_delete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TeamsList lists team entities

  List Team entities

--- Team ---

A Team represents a user-defined group of authors.

--- General ---

All objects can be traversed to related objects. This can be achieved by using the double-underscore syntax.
For instance, query all Author records with a specific Team name like so:

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
func (a *Client) TeamsList(params *TeamsListParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_list",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/teams/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TeamsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for teams_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TeamsRead gets a specific team entity by id value

  Get a specific Team entity by id value
*/
func (a *Client) TeamsRead(params *TeamsReadParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_read",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/teams/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TeamsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for teams_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TeamsUpdate updates a specific team entity by id value

  Update a specific Team entity by id value
*/
func (a *Client) TeamsUpdate(params *TeamsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*TeamsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTeamsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "teams_update",
		Method:             "PUT",
		PathPattern:        "/v3/customer/core/teams/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TeamsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TeamsUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for teams_update: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
