// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// GetPermissionsReader is a Reader for the GetPermissions structure.
type GetPermissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPermissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPermissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPermissionsOK creates a GetPermissionsOK with default headers values
func NewGetPermissionsOK() *GetPermissionsOK {
	return &GetPermissionsOK{}
}

/*GetPermissionsOK handles this case with default header values.

successful operation
*/
type GetPermissionsOK struct {
	Payload *models.PermissionAssignments
}

func (o *GetPermissionsOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/users/{userLocator}/permissions][%d] getPermissionsOK  %+v", 200, o.Payload)
}

func (o *GetPermissionsOK) GetPayload() *models.PermissionAssignments {
	return o.Payload
}

func (o *GetPermissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PermissionAssignments)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
