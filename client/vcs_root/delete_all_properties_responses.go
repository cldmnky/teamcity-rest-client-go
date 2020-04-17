// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAllPropertiesReader is a Reader for the DeleteAllProperties structure.
type DeleteAllPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAllPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteAllPropertiesDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteAllPropertiesDefault creates a DeleteAllPropertiesDefault with default headers values
func NewDeleteAllPropertiesDefault(code int) *DeleteAllPropertiesDefault {
	return &DeleteAllPropertiesDefault{
		_statusCode: code,
	}
}

/*DeleteAllPropertiesDefault handles this case with default header values.

successful operation
*/
type DeleteAllPropertiesDefault struct {
	_statusCode int
}

// Code gets the status code for the delete all properties default response
func (o *DeleteAllPropertiesDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAllPropertiesDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/vcs-roots/{vcsRootLocator}/properties][%d] deleteAllProperties default ", o._statusCode)
}

func (o *DeleteAllPropertiesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
