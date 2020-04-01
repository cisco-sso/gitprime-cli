// Code generated by go-swagger; DO NOT EDIT.

package grants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GrantsReadReader is a Reader for the GrantsRead structure.
type GrantsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GrantsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGrantsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGrantsReadOK creates a GrantsReadOK with default headers values
func NewGrantsReadOK() *GrantsReadOK {
	return &GrantsReadOK{}
}

/*GrantsReadOK handles this case with default header values.

GrantsReadOK grants read o k
*/
type GrantsReadOK struct {
	Payload interface{}
}

func (o *GrantsReadOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/grants/{id}/][%d] grantsReadOK  %+v", 200, o.Payload)
}

func (o *GrantsReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GrantsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}