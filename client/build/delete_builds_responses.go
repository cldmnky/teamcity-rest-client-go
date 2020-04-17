// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteBuildsReader is a Reader for the DeleteBuilds structure.
type DeleteBuildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBuildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteBuildsDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteBuildsDefault creates a DeleteBuildsDefault with default headers values
func NewDeleteBuildsDefault(code int) *DeleteBuildsDefault {
	return &DeleteBuildsDefault{
		_statusCode: code,
	}
}

/*DeleteBuildsDefault handles this case with default header values.

successful operation
*/
type DeleteBuildsDefault struct {
	_statusCode int
}

// Code gets the status code for the delete builds default response
func (o *DeleteBuildsDefault) Code() int {
	return o._statusCode
}

func (o *DeleteBuildsDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/builds][%d] deleteBuilds default ", o._statusCode)
}

func (o *DeleteBuildsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
