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

// CreatePoolReader is a Reader for the CreatePool structure.
type CreatePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreatePoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePoolOK creates a CreatePoolOK with default headers values
func NewCreatePoolOK() *CreatePoolOK {
	return &CreatePoolOK{}
}

/*CreatePoolOK handles this case with default header values.

successful operation
*/
type CreatePoolOK struct {
	Payload *models.AgentPool
}

func (o *CreatePoolOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/agentPools][%d] createPoolOK  %+v", 200, o.Payload)
}

func (o *CreatePoolOK) GetPayload() *models.AgentPool {
	return o.Payload
}

func (o *CreatePoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AgentPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
