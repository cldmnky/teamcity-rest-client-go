// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// ServeProjectReader is a Reader for the ServeProject structure.
type ServeProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServeProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeProjectOK creates a ServeProjectOK with default headers values
func NewServeProjectOK() *ServeProjectOK {
	return &ServeProjectOK{}
}

/*ServeProjectOK handles this case with default header values.

successful operation
*/
type ServeProjectOK struct {
	Payload *models.Project
}

func (o *ServeProjectOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}][%d] serveProjectOK  %+v", 200, o.Payload)
}

func (o *ServeProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *ServeProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
