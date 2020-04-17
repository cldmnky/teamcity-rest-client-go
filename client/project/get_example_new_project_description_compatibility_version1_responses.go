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

// GetExampleNewProjectDescriptionCompatibilityVersion1Reader is a Reader for the GetExampleNewProjectDescriptionCompatibilityVersion1 structure.
type GetExampleNewProjectDescriptionCompatibilityVersion1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExampleNewProjectDescriptionCompatibilityVersion1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExampleNewProjectDescriptionCompatibilityVersion1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetExampleNewProjectDescriptionCompatibilityVersion1OK creates a GetExampleNewProjectDescriptionCompatibilityVersion1OK with default headers values
func NewGetExampleNewProjectDescriptionCompatibilityVersion1OK() *GetExampleNewProjectDescriptionCompatibilityVersion1OK {
	return &GetExampleNewProjectDescriptionCompatibilityVersion1OK{}
}

/*GetExampleNewProjectDescriptionCompatibilityVersion1OK handles this case with default header values.

successful operation
*/
type GetExampleNewProjectDescriptionCompatibilityVersion1OK struct {
	Payload *models.NewProjectDescription
}

func (o *GetExampleNewProjectDescriptionCompatibilityVersion1OK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}/newProjectDescription][%d] getExampleNewProjectDescriptionCompatibilityVersion1OK  %+v", 200, o.Payload)
}

func (o *GetExampleNewProjectDescriptionCompatibilityVersion1OK) GetPayload() *models.NewProjectDescription {
	return o.Payload
}

func (o *GetExampleNewProjectDescriptionCompatibilityVersion1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NewProjectDescription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}