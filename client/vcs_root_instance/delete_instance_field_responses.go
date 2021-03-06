// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteInstanceFieldReader is a Reader for the DeleteInstanceField structure.
type DeleteInstanceFieldReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstanceFieldReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteInstanceFieldDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteInstanceFieldDefault creates a DeleteInstanceFieldDefault with default headers values
func NewDeleteInstanceFieldDefault(code int) *DeleteInstanceFieldDefault {
	return &DeleteInstanceFieldDefault{
		_statusCode: code,
	}
}

/*DeleteInstanceFieldDefault handles this case with default header values.

successful operation
*/
type DeleteInstanceFieldDefault struct {
	_statusCode int
}

// Code gets the status code for the delete instance field default response
func (o *DeleteInstanceFieldDefault) Code() int {
	return o._statusCode
}

func (o *DeleteInstanceFieldDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/{field}][%d] deleteInstanceField default ", o._statusCode)
}

func (o *DeleteInstanceFieldDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
