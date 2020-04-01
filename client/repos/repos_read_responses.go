// Code generated by go-swagger; DO NOT EDIT.

package repos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ReposReadReader is a Reader for the ReposRead structure.
type ReposReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReposReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReposReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReposReadOK creates a ReposReadOK with default headers values
func NewReposReadOK() *ReposReadOK {
	return &ReposReadOK{}
}

/*ReposReadOK handles this case with default header values.

ReposReadOK repos read o k
*/
type ReposReadOK struct {
	Payload interface{}
}

func (o *ReposReadOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/repos/{project}/][%d] reposReadOK  %+v", 200, o.Payload)
}

func (o *ReposReadOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ReposReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
