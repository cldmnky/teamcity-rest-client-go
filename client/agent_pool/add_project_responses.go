// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// AddProjectReader is a Reader for the AddProject structure.
type AddProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddProjectOK creates a AddProjectOK with default headers values
func NewAddProjectOK() *AddProjectOK {
	return &AddProjectOK{}
}

/*AddProjectOK handles this case with default header values.

successful operation
*/
type AddProjectOK struct {
	Payload *models.Project
}

func (o *AddProjectOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/agentPools/{agentPoolLocator}/projects][%d] addProjectOK  %+v", 200, o.Payload)
}

func (o *AddProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *AddProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
