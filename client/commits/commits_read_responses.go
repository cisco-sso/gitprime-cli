// Code generated by go-swagger; DO NOT EDIT.

package commits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CommitsReadReader is a Reader for the CommitsRead structure.
type CommitsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CommitsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCommitsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCommitsReadOK creates a CommitsReadOK with default headers values
func NewCommitsReadOK() *CommitsReadOK {
	return &CommitsReadOK{}
}

/*CommitsReadOK handles this case with default header values.

CommitsReadOK commits read o k
*/
type CommitsReadOK struct {
	Payload interface{}
}

func (o *CommitsReadOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/commits/{id}/][%d] commitsReadOK  %+v", 200, o.Payload)
}

func (o *CommitsReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CommitsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}