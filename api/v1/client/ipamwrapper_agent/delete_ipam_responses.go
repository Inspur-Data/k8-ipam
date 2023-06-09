// Code generated by go-swagger; DO NOT EDIT.

package ipamwrapper_agent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Inspur-Data/ipamwrapper/api/v1/models"
)

// DeleteIpamReader is a Reader for the DeleteIpam structure.
type DeleteIpamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteIpamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteIpamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteIpamFailure()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteIpamOK creates a DeleteIpamOK with default headers values
func NewDeleteIpamOK() *DeleteIpamOK {
	return &DeleteIpamOK{}
}

/*
DeleteIpamOK describes a response with status code 200, with default header values.

Success
*/
type DeleteIpamOK struct {
}

// IsSuccess returns true when this delete ipam o k response has a 2xx status code
func (o *DeleteIpamOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete ipam o k response has a 3xx status code
func (o *DeleteIpamOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete ipam o k response has a 4xx status code
func (o *DeleteIpamOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete ipam o k response has a 5xx status code
func (o *DeleteIpamOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete ipam o k response a status code equal to that given
func (o *DeleteIpamOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete ipam o k response
func (o *DeleteIpamOK) Code() int {
	return 200
}

func (o *DeleteIpamOK) Error() string {
	return fmt.Sprintf("[DELETE /ipam][%d] deleteIpamOK ", 200)
}

func (o *DeleteIpamOK) String() string {
	return fmt.Sprintf("[DELETE /ipam][%d] deleteIpamOK ", 200)
}

func (o *DeleteIpamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteIpamFailure creates a DeleteIpamFailure with default headers values
func NewDeleteIpamFailure() *DeleteIpamFailure {
	return &DeleteIpamFailure{}
}

/*
DeleteIpamFailure describes a response with status code 500, with default header values.

Addresses release failure
*/
type DeleteIpamFailure struct {
	Payload models.Error
}

// IsSuccess returns true when this delete ipam failure response has a 2xx status code
func (o *DeleteIpamFailure) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete ipam failure response has a 3xx status code
func (o *DeleteIpamFailure) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete ipam failure response has a 4xx status code
func (o *DeleteIpamFailure) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete ipam failure response has a 5xx status code
func (o *DeleteIpamFailure) IsServerError() bool {
	return true
}

// IsCode returns true when this delete ipam failure response a status code equal to that given
func (o *DeleteIpamFailure) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete ipam failure response
func (o *DeleteIpamFailure) Code() int {
	return 500
}

func (o *DeleteIpamFailure) Error() string {
	return fmt.Sprintf("[DELETE /ipam][%d] deleteIpamFailure  %+v", 500, o.Payload)
}

func (o *DeleteIpamFailure) String() string {
	return fmt.Sprintf("[DELETE /ipam][%d] deleteIpamFailure  %+v", 500, o.Payload)
}

func (o *DeleteIpamFailure) GetPayload() models.Error {
	return o.Payload
}

func (o *DeleteIpamFailure) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
