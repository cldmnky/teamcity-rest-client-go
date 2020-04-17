// Code generated by go-swagger; DO NOT EDIT.

package agent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAgentReader is a Reader for the DeleteAgent structure.
type DeleteAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteAgentDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteAgentDefault creates a DeleteAgentDefault with default headers values
func NewDeleteAgentDefault(code int) *DeleteAgentDefault {
	return &DeleteAgentDefault{
		_statusCode: code,
	}
}

/*DeleteAgentDefault handles this case with default header values.

successful operation
*/
type DeleteAgentDefault struct {
	_statusCode int
}

// Code gets the status code for the delete agent default response
func (o *DeleteAgentDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAgentDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/agents/{agentLocator}][%d] deleteAgent default ", o._statusCode)
}

func (o *DeleteAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
