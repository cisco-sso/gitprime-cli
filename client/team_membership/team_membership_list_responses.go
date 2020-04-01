// Code generated by go-swagger; DO NOT EDIT.

package team_membership

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TeamMembershipListReader is a Reader for the TeamMembershipList structure.
type TeamMembershipListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TeamMembershipListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTeamMembershipListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTeamMembershipListOK creates a TeamMembershipListOK with default headers values
func NewTeamMembershipListOK() *TeamMembershipListOK {
	return &TeamMembershipListOK{}
}

/*TeamMembershipListOK handles this case with default header values.

TeamMembershipListOK team membership list o k
*/
type TeamMembershipListOK struct {
	Payload interface{}
}

func (o *TeamMembershipListOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/team_membership/][%d] teamMembershipListOK  %+v", 200, o.Payload)
}

func (o *TeamMembershipListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *TeamMembershipListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}