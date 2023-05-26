// Code generated by go-swagger; DO NOT EDIT.

package health_check

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetHealthyOKCode is the HTTP code returned for type GetHealthyOK
const GetHealthyOKCode int = 200

/*
GetHealthyOK Success

swagger:response getHealthyOK
*/
type GetHealthyOK struct {
}

// NewGetHealthyOK creates GetHealthyOK with default headers values
func NewGetHealthyOK() *GetHealthyOK {

	return &GetHealthyOK{}
}

// WriteResponse to the client
func (o *GetHealthyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// GetHealthyInternalServerErrorCode is the HTTP code returned for type GetHealthyInternalServerError
const GetHealthyInternalServerErrorCode int = 500

/*
GetHealthyInternalServerError Failed

swagger:response getHealthyInternalServerError
*/
type GetHealthyInternalServerError struct {
}

// NewGetHealthyInternalServerError creates GetHealthyInternalServerError with default headers values
func NewGetHealthyInternalServerError() *GetHealthyInternalServerError {

	return &GetHealthyInternalServerError{}
}

// WriteResponse to the client
func (o *GetHealthyInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
