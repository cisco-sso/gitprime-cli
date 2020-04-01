// Code generated by go-swagger; DO NOT EDIT.

package ticket_comments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TicketCommentsReadReader is a Reader for the TicketCommentsRead structure.
type TicketCommentsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketCommentsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketCommentsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTicketCommentsReadOK creates a TicketCommentsReadOK with default headers values
func NewTicketCommentsReadOK() *TicketCommentsReadOK {
	return &TicketCommentsReadOK{}
}

/*TicketCommentsReadOK handles this case with default header values.

TicketCommentsReadOK ticket comments read o k
*/
type TicketCommentsReadOK struct {
	Payload interface{}
}

func (o *TicketCommentsReadOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/ticket_comments/{id}/][%d] ticketCommentsReadOK  %+v", 200, o.Payload)
}

func (o *TicketCommentsReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketCommentsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
