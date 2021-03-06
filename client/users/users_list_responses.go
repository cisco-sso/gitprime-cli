// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UsersListReader is a Reader for the UsersList structure.
type UsersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersListOK creates a UsersListOK with default headers values
func NewUsersListOK() *UsersListOK {
	return &UsersListOK{}
}

/*UsersListOK handles this case with default header values.

UsersListOK users list o k
*/
type UsersListOK struct {
	Payload interface{}
}

func (o *UsersListOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/users/][%d] usersListOK  %+v", 200, o.Payload)
}

func (o *UsersListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
