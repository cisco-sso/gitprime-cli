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

// IntegrationsConnectionStatusReader is a Reader for the IntegrationsConnectionStatus structure.
type IntegrationsConnectionStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationsConnectionStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntegrationsConnectionStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIntegrationsConnectionStatusOK creates a IntegrationsConnectionStatusOK with default headers values
func NewIntegrationsConnectionStatusOK() *IntegrationsConnectionStatusOK {
	return &IntegrationsConnectionStatusOK{}
}

/*IntegrationsConnectionStatusOK handles this case with default header values.

IntegrationsConnectionStatusOK integrations connection status o k
*/
type IntegrationsConnectionStatusOK struct {
	Payload interface{}
}

func (o *IntegrationsConnectionStatusOK) Error() string {
	return fmt.Sprintf("[GET /v3/customer/core/integrations/{id}/connection_status/][%d] integrationsConnectionStatusOK  %+v", 200, o.Payload)
}

func (o *IntegrationsConnectionStatusOK) GetPayload() interface{} {
	return o.Payload
}

func (o *IntegrationsConnectionStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
