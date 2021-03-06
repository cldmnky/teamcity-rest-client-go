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

// ReplaceAgentRequirementsReader is a Reader for the ReplaceAgentRequirements structure.
type ReplaceAgentRequirementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceAgentRequirementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceAgentRequirementsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceAgentRequirementsOK creates a ReplaceAgentRequirementsOK with default headers values
func NewReplaceAgentRequirementsOK() *ReplaceAgentRequirementsOK {
	return &ReplaceAgentRequirementsOK{}
}

/*ReplaceAgentRequirementsOK handles this case with default header values.

successful operation
*/
type ReplaceAgentRequirementsOK struct {
	Payload *models.AgentRequirements
}

func (o *ReplaceAgentRequirementsOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/agent-requirements][%d] replaceAgentRequirementsOK  %+v", 200, o.Payload)
}

func (o *ReplaceAgentRequirementsOK) GetPayload() *models.AgentRequirements {
	return o.Payload
}

func (o *ReplaceAgentRequirementsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentRequirements)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
