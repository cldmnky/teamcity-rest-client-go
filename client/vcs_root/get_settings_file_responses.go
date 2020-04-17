// Code generated by go-swagger; DO NOT EDIT.

package vcs_root

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetSettingsFileReader is a Reader for the GetSettingsFile structure.
type GetSettingsFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSettingsFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSettingsFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSettingsFileOK creates a GetSettingsFileOK with default headers values
func NewGetSettingsFileOK() *GetSettingsFileOK {
	return &GetSettingsFileOK{}
}

/*GetSettingsFileOK handles this case with default header values.

successful operation
*/
type GetSettingsFileOK struct {
	Payload string
}

func (o *GetSettingsFileOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/vcs-roots/{vcsRootLocator}/settingsFile][%d] getSettingsFileOK  %+v", 200, o.Payload)
}

func (o *GetSettingsFileOK) GetPayload() string {
	return o.Payload
}

func (o *GetSettingsFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}