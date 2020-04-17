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

// ReplaceSnapshotDepReader is a Reader for the ReplaceSnapshotDep structure.
type ReplaceSnapshotDepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceSnapshotDepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceSnapshotDepOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReplaceSnapshotDepOK creates a ReplaceSnapshotDepOK with default headers values
func NewReplaceSnapshotDepOK() *ReplaceSnapshotDepOK {
	return &ReplaceSnapshotDepOK{}
}

/*ReplaceSnapshotDepOK handles this case with default header values.

successful operation
*/
type ReplaceSnapshotDepOK struct {
	Payload *models.SnapshotDependency
}

func (o *ReplaceSnapshotDepOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/buildTypes/{btLocator}/snapshot-dependencies/{snapshotDepLocator}][%d] replaceSnapshotDepOK  %+v", 200, o.Payload)
}

func (o *ReplaceSnapshotDepOK) GetPayload() *models.SnapshotDependency {
	return o.Payload
}

func (o *ReplaceSnapshotDepOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotDependency)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}