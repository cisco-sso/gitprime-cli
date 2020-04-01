// Code generated by go-swagger; DO NOT EDIT.

package commits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new commits API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for commits API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CommitsList(params *CommitsListParams, authInfo runtime.ClientAuthInfoWriter) (*CommitsListOK, error)

	CommitsRead(params *CommitsReadParams, authInfo runtime.ClientAuthInfoWriter) (*CommitsReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CommitsList lists commit entities

  List Commit entities

--- Commit ---

A Commit represents a set of code changes committed locally and then pushed to a repository. All Commits have
associated metrics that describe the quantity and nature of the code that was committed as well as it's
author.

** Deprecation notice! Pair programming has caused the deprecation of user_alias_id. You will now need to use
the "aliases" field, which is an array of User Alias ids.

** As of November 2019, it is required to utilize a date range selection (author_date or author_local_date) criteria
 when utilizing the .agg method of this API. The date range must be less than or equal to one year.

It is always recommended that an integrator utilize the author_date or author_local_date as a selection criteria to
improve the performance of the API.

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
func (a *Client) CommitsList(params *CommitsListParams, authInfo runtime.ClientAuthInfoWriter) (*CommitsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommitsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "commits_list",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/commits/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CommitsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CommitsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for commits_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CommitsRead gets a specific commit entity by id value

  Get a specific Commit entity by id value
*/
func (a *Client) CommitsRead(params *CommitsReadParams, authInfo runtime.ClientAuthInfoWriter) (*CommitsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCommitsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "commits_read",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/commits/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CommitsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CommitsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for commits_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}