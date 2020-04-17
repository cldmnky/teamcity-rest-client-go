// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetZippedReader is a Reader for the GetZipped structure.
type GetZippedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetZippedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetZippedDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetZippedDefault creates a GetZippedDefault with default headers values
func NewGetZippedDefault(code int) *GetZippedDefault {
	return &GetZippedDefault{
		_statusCode: code,
	}
}

/*GetZippedDefault handles this case with default header values.

successful operation
*/
type GetZippedDefault struct {
	_statusCode int
}

// Code gets the status code for the get zipped default response
func (o *GetZippedDefault) Code() int {
	return o._statusCode
}

func (o *GetZippedDefault) Error() string {
	return fmt.Sprintf("[GET /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/files/latest/archived{path}][%d] getZipped default ", o._statusCode)
}

func (o *GetZippedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}