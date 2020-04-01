// Code generated by go-swagger; DO NOT EDIT.

package ticket_events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TicketEventsListReader is a Reader for the TicketEventsList structure.
type TicketEventsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketEventsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketEventsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTicketEventsListOK creates a TicketEventsListOK with default headers values
func NewTicketEventsListOK() *TicketEventsListOK {
	return &TicketEventsListOK{}
}

/*TicketEventsListOK handles this case with default header values.

TicketEventsListOK ticket events list o k
*/
type TicketEventsListOK struct {
	Payload interface{}
}

func (o *TicketEventsListOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/ticket_events/][%d] ticketEventsListOK  %+v", 200, o.Payload)
}

func (o *TicketEventsListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketEventsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
