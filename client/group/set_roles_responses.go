// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/models"
)

// SetRolesReader is a Reader for the SetRoles structure.
type SetRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSetRolesOK creates a SetRolesOK with default headers values
func NewSetRolesOK() *SetRolesOK {
	return &SetRolesOK{}
}

/*SetRolesOK handles this case with default header values.

successful operation
*/
type SetRolesOK struct {
	Payload *models.Roles
}

func (o *SetRolesOK) Error() string {
	return fmt.Sprintf("[PUT /app/rest/userGroups/{groupLocator}/roles][%d] setRolesOK  %+v", 200, o.Payload)
}

func (o *SetRolesOK) GetPayload() *models.Roles {
	return o.Payload
}

func (o *SetRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Roles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
