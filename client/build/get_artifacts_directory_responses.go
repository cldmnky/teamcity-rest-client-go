// Code generated by go-swagger; DO NOT EDIT.

package build

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetArtifactsDirectoryReader is a Reader for the GetArtifactsDirectory structure.
type GetArtifactsDirectoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArtifactsDirectoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArtifactsDirectoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetArtifactsDirectoryOK creates a GetArtifactsDirectoryOK with default headers values
func NewGetArtifactsDirectoryOK() *GetArtifactsDirectoryOK {
	return &GetArtifactsDirectoryOK{}
}

/*GetArtifactsDirectoryOK handles this case with default header values.

successful operation
*/
type GetArtifactsDirectoryOK struct {
	Payload string
}

func (o *GetArtifactsDirectoryOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/builds/{buildLocator}/artifactsDirectory][%d] getArtifactsDirectoryOK  %+v", 200, o.Payload)
}

func (o *GetArtifactsDirectoryOK) GetPayload() string {
	return o.Payload
}

func (o *GetArtifactsDirectoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
