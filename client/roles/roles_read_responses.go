// Code generated by go-swagger; DO NOT EDIT.

package roles

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RolesReadReader is a Reader for the RolesRead structure.
type RolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRolesReadOK creates a RolesReadOK with default headers values
func NewRolesReadOK() *RolesReadOK {
	return &RolesReadOK{}
}

/*RolesReadOK handles this case with default header values.

RolesReadOK roles read o k
*/
type RolesReadOK struct {
	Payload interface{}
}

func (o *RolesReadOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/roles/{id}/][%d] rolesReadOK  %+v", 200, o.Payload)
}

func (o *RolesReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}