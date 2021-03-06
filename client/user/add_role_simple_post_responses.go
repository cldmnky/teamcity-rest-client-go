// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddRoleSimplePostReader is a Reader for the AddRoleSimplePost structure.
type AddRoleSimplePostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRoleSimplePostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewAddRoleSimplePostDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewAddRoleSimplePostDefault creates a AddRoleSimplePostDefault with default headers values
func NewAddRoleSimplePostDefault(code int) *AddRoleSimplePostDefault {
	return &AddRoleSimplePostDefault{
		_statusCode: code,
	}
}

/*AddRoleSimplePostDefault handles this case with default header values.

successful operation
*/
type AddRoleSimplePostDefault struct {
	_statusCode int
}

// Code gets the status code for the add role simple post default response
func (o *AddRoleSimplePostDefault) Code() int {
	return o._statusCode
}

func (o *AddRoleSimplePostDefault) Error() string {
	return fmt.Sprintf("[POST /app/rest/users/{userLocator}/roles/{roleId}/{scope}][%d] addRoleSimplePost default ", o._statusCode)
}

func (o *AddRoleSimplePostDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
