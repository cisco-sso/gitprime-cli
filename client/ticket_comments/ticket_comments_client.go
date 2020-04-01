// Code generated by go-swagger; DO NOT EDIT.

package ticket_comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ticket comments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ticket comments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TicketCommentsList(params *TicketCommentsListParams, authInfo runtime.ClientAuthInfoWriter) (*TicketCommentsListOK, error)

	TicketCommentsRead(params *TicketCommentsReadParams, authInfo runtime.ClientAuthInfoWriter) (*TicketCommentsReadOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TicketCommentsList lists ticket comment entities

  List TicketComment entities

--- Ticket Comment ---

A Ticket Comment is a comment made by an author related to a specific ticket.

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
func (a *Client) TicketCommentsList(params *TicketCommentsListParams, authInfo runtime.ClientAuthInfoWriter) (*TicketCommentsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTicketCommentsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ticket_comments_list",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/ticket_comments/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TicketCommentsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TicketCommentsListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ticket_comments_list: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TicketCommentsRead gets a specific ticket comment entity by id value

  Get a specific TicketComment entity by id value
*/
func (a *Client) TicketCommentsRead(params *TicketCommentsReadParams, authInfo runtime.ClientAuthInfoWriter) (*TicketCommentsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTicketCommentsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ticket_comments_read",
		Method:             "GET",
		PathPattern:        "/v3/customer/core/ticket_comments/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TicketCommentsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TicketCommentsReadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ticket_comments_read: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
