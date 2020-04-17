// Code generated by go-swagger; DO NOT EDIT.

package build_type

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// AddArtifactDepReader is a Reader for the AddArtifactDep structure.
type AddArtifactDepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddArtifactDepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddArtifactDepOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddArtifactDepOK creates a AddArtifactDepOK with default headers values
func NewAddArtifactDepOK() *AddArtifactDepOK {
	return &AddArtifactDepOK{}
}

/*AddArtifactDepOK handles this case with default header values.

successful operation
*/
type AddArtifactDepOK struct {
	Payload *models.ArtifactDependency
}

func (o *AddArtifactDepOK) Error() string {
	return fmt.Sprintf("[POST /app/rest/buildTypes/{btLocator}/artifact-dependencies][%d] addArtifactDepOK  %+v", 200, o.Payload)
}

func (o *AddArtifactDepOK) GetPayload() *models.ArtifactDependency {
	return o.Payload
}

func (o *AddArtifactDepOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ArtifactDependency)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
