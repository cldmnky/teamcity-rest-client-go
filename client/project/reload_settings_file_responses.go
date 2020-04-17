// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// ReloadSettingsFileReader is a Reader for the ReloadSettingsFile structure.
type ReloadSettingsFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReloadSettingsFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReloadSettingsFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReloadSettingsFileOK creates a ReloadSettingsFileOK with default headers values
func NewReloadSettingsFileOK() *ReloadSettingsFileOK {
	return &ReloadSettingsFileOK{}
}

/*ReloadSettingsFileOK handles this case with default header values.

successful operation
*/
type ReloadSettingsFileOK struct {
	Payload *models.Project
}

func (o *ReloadSettingsFileOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/projects/{projectLocator}/latest][%d] reloadSettingsFileOK  %+v", 200, o.Payload)
}

func (o *ReloadSettingsFileOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *ReloadSettingsFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
