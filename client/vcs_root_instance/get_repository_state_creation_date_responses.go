// Code generated by go-swagger; DO NOT EDIT.

package vcs_root_instance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetRepositoryStateCreationDateReader is a Reader for the GetRepositoryStateCreationDate structure.
type GetRepositoryStateCreationDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRepositoryStateCreationDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRepositoryStateCreationDateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepositoryStateCreationDateOK creates a GetRepositoryStateCreationDateOK with default headers values
func NewGetRepositoryStateCreationDateOK() *GetRepositoryStateCreationDateOK {
	return &GetRepositoryStateCreationDateOK{}
}

/*GetRepositoryStateCreationDateOK handles this case with default header values.

successful operation
*/
type GetRepositoryStateCreationDateOK struct {
	Payload string
}

func (o *GetRepositoryStateCreationDateOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/vcs-root-instances/{vcsRootInstanceLocator}/repositoryState/creationDate][%d] getRepositoryStateCreationDateOK  %+v", 200, o.Payload)
}

func (o *GetRepositoryStateCreationDateOK) GetPayload() string {
	return o.Payload
}

func (o *GetRepositoryStateCreationDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
