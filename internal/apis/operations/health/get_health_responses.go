// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/music-api-golang/internal/models"
)

// GetHealthOKCode is the HTTP code returned for type GetHealthOK
const GetHealthOKCode int = 200

/*
GetHealthOK Success

swagger:response getHealthOK
*/
type GetHealthOK struct {

	/*
	  In: Body
	*/
	Payload *GetHealthOKBody `json:"body,omitempty"`
}

// NewGetHealthOK creates GetHealthOK with default headers values
func NewGetHealthOK() *GetHealthOK {

	return &GetHealthOK{}
}

// WithPayload adds the payload to the get health o k response
func (o *GetHealthOK) WithPayload(payload *GetHealthOKBody) *GetHealthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health o k response
func (o *GetHealthOK) SetPayload(payload *GetHealthOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHealthBadRequestCode is the HTTP code returned for type GetHealthBadRequest
const GetHealthBadRequestCode int = 400

/*
GetHealthBadRequest Unauthorized

swagger:response getHealthBadRequest
*/
type GetHealthBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHealthBadRequest creates GetHealthBadRequest with default headers values
func NewGetHealthBadRequest() *GetHealthBadRequest {

	return &GetHealthBadRequest{}
}

// WithPayload adds the payload to the get health bad request response
func (o *GetHealthBadRequest) WithPayload(payload *models.Error) *GetHealthBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health bad request response
func (o *GetHealthBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHealthUnauthorizedCode is the HTTP code returned for type GetHealthUnauthorized
const GetHealthUnauthorizedCode int = 401

/*
GetHealthUnauthorized Unauthorized

swagger:response getHealthUnauthorized
*/
type GetHealthUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHealthUnauthorized creates GetHealthUnauthorized with default headers values
func NewGetHealthUnauthorized() *GetHealthUnauthorized {

	return &GetHealthUnauthorized{}
}

// WithPayload adds the payload to the get health unauthorized response
func (o *GetHealthUnauthorized) WithPayload(payload *models.Error) *GetHealthUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health unauthorized response
func (o *GetHealthUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetHealthNotFoundCode is the HTTP code returned for type GetHealthNotFound
const GetHealthNotFoundCode int = 404

/*
GetHealthNotFound The specified resource was not found

swagger:response getHealthNotFound
*/
type GetHealthNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetHealthNotFound creates GetHealthNotFound with default headers values
func NewGetHealthNotFound() *GetHealthNotFound {

	return &GetHealthNotFound{}
}

// WithPayload adds the payload to the get health not found response
func (o *GetHealthNotFound) WithPayload(payload *models.Error) *GetHealthNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get health not found response
func (o *GetHealthNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetHealthNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}