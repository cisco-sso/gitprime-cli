// Code generated by go-swagger; DO NOT EDIT.

package user_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UserGroupsListReader is a Reader for the UserGroupsList structure.
type UserGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUserGroupsListOK creates a UserGroupsListOK with default headers values
func NewUserGroupsListOK() *UserGroupsListOK {
	return &UserGroupsListOK{}
}

/*UserGroupsListOK handles this case with default header values.

UserGroupsListOK user groups list o k
*/
type UserGroupsListOK struct {
	Payload interface{}
}

func (o *UserGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/user_groups/][%d] userGroupsListOK  %+v", 200, o.Payload)
}

func (o *UserGroupsListOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UserGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
