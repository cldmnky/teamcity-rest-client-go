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

// ServeBuildTypeReader is a Reader for the ServeBuildType structure.
type ServeBuildTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServeBuildTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServeBuildTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewServeBuildTypeOK creates a ServeBuildTypeOK with default headers values
func NewServeBuildTypeOK() *ServeBuildTypeOK {
	return &ServeBuildTypeOK{}
}

/*ServeBuildTypeOK handles this case with default header values.

successful operation
*/
type ServeBuildTypeOK struct {
	Payload *models.BuildType
}

func (o *ServeBuildTypeOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}/buildTypes/{btLocator}][%d] serveBuildTypeOK  %+v", 200, o.Payload)
}

func (o *ServeBuildTypeOK) GetPayload() *models.BuildType {
	return o.Payload
}

func (o *ServeBuildTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
