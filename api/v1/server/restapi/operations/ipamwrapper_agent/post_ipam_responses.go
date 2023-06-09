// Code generated by go-swagger; DO NOT EDIT.

package ipamwrapper_agent

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/Inspur-Data/ipamwrapper/api/v1/models"
)

// PostIpamOKCode is the HTTP code returned for type PostIpamOK
const PostIpamOKCode int = 200

/*
PostIpamOK Success

swagger:response postIpamOK
*/
type PostIpamOK struct {

	/*
	  In: Body
	*/
	Payload *models.IpamAllocResponse `json:"body,omitempty"`
}

// NewPostIpamOK creates PostIpamOK with default headers values
func NewPostIpamOK() *PostIpamOK {

	return &PostIpamOK{}
}

// WithPayload adds the payload to the post ipam o k response
func (o *PostIpamOK) WithPayload(payload *models.IpamAllocResponse) *PostIpamOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post ipam o k response
func (o *PostIpamOK) SetPayload(payload *models.IpamAllocResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIpamOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostIpamFailureCode is the HTTP code returned for type PostIpamFailure
const PostIpamFailureCode int = 500

/*
PostIpamFailure Allocation failure

swagger:response postIpamFailure
*/
type PostIpamFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPostIpamFailure creates PostIpamFailure with default headers values
func NewPostIpamFailure() *PostIpamFailure {

	return &PostIpamFailure{}
}

// WithPayload adds the payload to the post ipam failure response
func (o *PostIpamFailure) WithPayload(payload models.Error) *PostIpamFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post ipam failure response
func (o *PostIpamFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostIpamFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
