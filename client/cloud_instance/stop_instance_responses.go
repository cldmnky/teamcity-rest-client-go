// Code generated by go-swagger; DO NOT EDIT.

package cloud_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// StopInstanceReader is a Reader for the StopInstance structure.
type StopInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewStopInstanceDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewStopInstanceDefault creates a StopInstanceDefault with default headers values
func NewStopInstanceDefault(code int) *StopInstanceDefault {
	return &StopInstanceDefault{
		_statusCode: code,
	}
}

/*StopInstanceDefault handles this case with default header values.

successful operation
*/
type StopInstanceDefault struct {
	_statusCode int
}

// Code gets the status code for the stop instance default response
func (o *StopInstanceDefault) Code() int {
	return o._statusCode
}

func (o *StopInstanceDefault) Error() string {
	return fmt.Sprintf("[DELETE /app/rest/cloud/instances/{instanceLocator}][%d] stopInstance default ", o._statusCode)
}

func (o *StopInstanceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
