// Code generated by go-swagger; DO NOT EDIT.

package available_repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AvailableReposReadReader is a Reader for the AvailableReposRead structure.
type AvailableReposReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AvailableReposReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAvailableReposReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAvailableReposReadOK creates a AvailableReposReadOK with default headers values
func NewAvailableReposReadOK() *AvailableReposReadOK {
	return &AvailableReposReadOK{}
}

/*AvailableReposReadOK handles this case with default header values.

AvailableReposReadOK available repos read o k
*/
type AvailableReposReadOK struct {
	Payload interface{}
}

func (o *AvailableReposReadOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/available_repos/{id}/][%d] availableReposReadOK  %+v", 200, o.Payload)
}

func (o *AvailableReposReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *AvailableReposReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
