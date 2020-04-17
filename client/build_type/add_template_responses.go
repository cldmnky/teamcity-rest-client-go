// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// AddTemplateReader is a Reader for the AddTemplate structure.
type AddTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddTemplateOK creates a AddTemplateOK with default headers values
func NewAddTemplateOK() *AddTemplateOK {
	return &AddTemplateOK{}
}

/*AddTemplateOK handles this case with default header values.

successful operation
*/
type AddTemplateOK struct {
	Payload *models.BuildType
}

func (o *AddTemplateOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/buildTypes/{btLocator}/templates][%d] addTemplateOK  %+v", 200, o.Payload)
}

func (o *AddTemplateOK) GetPayload() *models.BuildType {
	return o.Payload
}

func (o *AddTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}