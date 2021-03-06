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

// ServeBuildsReader is a Reader for the ServeBuilds structure.
type ServeBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServeBuildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeBuildsOK creates a ServeBuildsOK with default headers values
func NewServeBuildsOK() *ServeBuildsOK {
	return &ServeBuildsOK{}
}

/*ServeBuildsOK handles this case with default header values.

successful operation
*/
type ServeBuildsOK struct {
	Payload *models.Builds
}

func (o *ServeBuildsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}/buildTypes/{btLocator}/builds][%d] serveBuildsOK  %+v", 200, o.Payload)
}

func (o *ServeBuildsOK) GetPayload() *models.Builds {
	return o.Payload
}

func (o *ServeBuildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Builds)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
