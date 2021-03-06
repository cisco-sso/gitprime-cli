// Code generated by go-swagger; DO NOT EDIT.

package ticket_projects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TicketProjectsListReader is a Reader for the TicketProjectsList structure.
type TicketProjectsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TicketProjectsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTicketProjectsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTicketProjectsListOK creates a TicketProjectsListOK with default headers values
func NewTicketProjectsListOK() *TicketProjectsListOK {
	return &TicketProjectsListOK{}
}

/*TicketProjectsListOK handles this case with default header values.

TicketProjectsListOK ticket projects list o k
*/
type TicketProjectsListOK struct {
	Payload interface{}
}

func (o *TicketProjectsListOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/ticket_projects/][%d] ticketProjectsListOK  %+v", 200, o.Payload)
}

func (o *TicketProjectsListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TicketProjectsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
