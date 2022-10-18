// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/music-api-golang/internal/models"
)

// GetHealthHandlerFunc turns a function with the right signature into a get health handler
type GetHealthHandlerFunc func(GetHealthParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetHealthHandlerFunc) Handle(params GetHealthParams) middleware.Responder {
	return fn(params)
}

// GetHealthHandler interface for that can handle valid get health params
type GetHealthHandler interface {
	Handle(GetHealthParams) middleware.Responder
}

// NewGetHealth creates a new http.Handler for the get health operation
func NewGetHealth(ctx *middleware.Context, handler GetHealthHandler) *GetHealth {
	return &GetHealth{Context: ctx, Handler: handler}
}

/*
	GetHealth swagger:route GET /health Health getHealth

# Health check

Health check endpoint
*/
type GetHealth struct {
	Context *middleware.Context
	Handler GetHealthHandler
}

func (o *GetHealth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetHealthParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetHealthOKBody get health o k body
//
// swagger:model GetHealthOKBody
type GetHealthOKBody struct {

	// data
	Data *models.Health `json:"data,omitempty"`

	// success retrieve data
	// Example: app running well
	Message string `json:"message,omitempty"`
}

// Validate validates this get health o k body
func (o *GetHealthOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHealthOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	if o.Data != nil {
		if err := o.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHealthOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHealthOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get health o k body based on the context it is used
func (o *GetHealthOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetHealthOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	if o.Data != nil {
		if err := o.Data.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getHealthOK" + "." + "data")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getHealthOK" + "." + "data")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetHealthOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetHealthOKBody) UnmarshalBinary(b []byte) error {
	var res GetHealthOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}