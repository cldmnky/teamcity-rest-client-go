// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// ServeUserPropertiesReader is a Reader for the ServeUserProperties structure.
type ServeUserPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeUserPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServeUserPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeUserPropertiesOK creates a ServeUserPropertiesOK with default headers values
func NewServeUserPropertiesOK() *ServeUserPropertiesOK {
	return &ServeUserPropertiesOK{}
}

/*ServeUserPropertiesOK handles this case with default header values.

successful operation
*/
type ServeUserPropertiesOK struct {
	Payload *models.Properties
}

func (o *ServeUserPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/users/{userLocator}/properties][%d] serveUserPropertiesOK  %+v", 200, o.Payload)
}

func (o *ServeUserPropertiesOK) GetPayload() *models.Properties {
	return o.Payload
}

func (o *ServeUserPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Properties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
