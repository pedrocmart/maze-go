// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BaseResponse base response
//
// swagger:model BaseResponse
type BaseResponse struct {

	// message
	Message string `json:"message,omitempty"`

	// status
	// Required: true
	Status int64 `json:"status"`

	// success
	// Required: true
	Success bool `json:"success"`
}

// Validate validates this base response
func (m *BaseResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", int64(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *BaseResponse) validateSuccess(formats strfmt.Registry) error {

	if err := validate.Required("success", "body", bool(m.Success)); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this base response based on context it is used
func (m *BaseResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseResponse) UnmarshalBinary(b []byte) error {
	var res BaseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
