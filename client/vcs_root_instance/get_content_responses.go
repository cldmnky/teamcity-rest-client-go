// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetContentReader is a Reader for the GetContent structure.
type GetContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	result := NewGetContentDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result
}

// NewGetContentDefault creates a GetContentDefault with default headers values
func NewGetContentDefault(code int) *GetContentDefault {
	return &GetContentDefault{
		_statusCode: code,
	}
}

/*GetContentDefault handles this case with default header values.

successful operation
*/
type GetContentDefault struct {
	_statusCode int
}

// Code gets the status code for the get content default response
func (o *GetContentDefault) Code() int {
	return o._statusCode
}

func (o *GetContentDefault) Error() string {
	return fmt.Sprintf("[GET /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/files/latest/content{path}][%d] getContent default ", o._statusCode)
}

func (o *GetContentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
