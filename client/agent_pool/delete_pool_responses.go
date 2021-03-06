// Code generated by go-swagger; DO NOT EDIT.

package agent_pool

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePoolReader is a Reader for the DeletePool structure.
type DeletePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeletePoolDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeletePoolDefault creates a DeletePoolDefault with default headers values
func NewDeletePoolDefault(code int) *DeletePoolDefault {
	return &DeletePoolDefault{
		_statusCode: code,
	}
}

/*DeletePoolDefault handles this case with default header values.

successful operation
*/
type DeletePoolDefault struct {
	_statusCode int
}

// Code gets the status code for the delete pool default response
func (o *DeletePoolDefault) Code() int {
	return o._statusCode
}

func (o *DeletePoolDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/agentPools/{agentPoolLocator}][%d] deletePool default ", o._statusCode)
}

func (o *DeletePoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
