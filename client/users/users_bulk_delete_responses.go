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

// UsersBulkDeleteReader is a Reader for the UsersBulkDelete structure.
type UsersBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersBulkDeleteCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUsersBulkDeleteCreated creates a UsersBulkDeleteCreated with default headers values
func NewUsersBulkDeleteCreated() *UsersBulkDeleteCreated {
	return &UsersBulkDeleteCreated{}
}

/*UsersBulkDeleteCreated handles this case with default header values.

UsersBulkDeleteCreated users bulk delete created
*/
type UsersBulkDeleteCreated struct {
	Payload interface{}
}

func (o *UsersBulkDeleteCreated) Error() string {
	return fmt.Sprintf("[POST /v3/customer/core/users/bulk_delete/][%d] usersBulkDeleteCreated  %+v", 201, o.Payload)
}

func (o *UsersBulkDeleteCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersBulkDeleteCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
