// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAllParametersReader is a Reader for the DeleteAllParameters structure.
type DeleteAllParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAllParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewDeleteAllParametersDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewDeleteAllParametersDefault creates a DeleteAllParametersDefault with default headers values
func NewDeleteAllParametersDefault(code int) *DeleteAllParametersDefault {
	return &DeleteAllParametersDefault{
		_statusCode: code,
	}
}

/*DeleteAllParametersDefault handles this case with default header values.

successful operation
*/
type DeleteAllParametersDefault struct {
	_statusCode int
}

// Code gets the status code for the delete all parameters default response
func (o *DeleteAllParametersDefault) Code() int {
	return o._statusCode
}

func (o *DeleteAllParametersDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/projects/{projectLocator}/projectFeatures/{featureLocator}/properties][%d] deleteAllParameters default ", o._statusCode)
}

func (o *DeleteAllParametersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
