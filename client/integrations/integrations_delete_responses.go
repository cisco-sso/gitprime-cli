// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IntegrationsDeleteReader is a Reader for the IntegrationsDelete structure.
type IntegrationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIntegrationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIntegrationsDeleteNoContent creates a IntegrationsDeleteNoContent with default headers values
func NewIntegrationsDeleteNoContent() *IntegrationsDeleteNoContent {
	return &IntegrationsDeleteNoContent{}
}

/*IntegrationsDeleteNoContent handles this case with default header values.

IntegrationsDeleteNoContent integrations delete no content
*/
type IntegrationsDeleteNoContent struct {
	Payload interface{}
}

func (o *IntegrationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v3/customer/core/integrations/{id}/][%d] integrationsDeleteNoContent  %+v", 204, o.Payload)
}

func (o *IntegrationsDeleteNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *IntegrationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
