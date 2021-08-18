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

// Level Level
//
// swagger:model Level
type Level struct {

	// created at
	// Format: date-time
	CreatedAt *strfmt.DateTime `json:"createdAt,omitempty"`

	// Level ID
	// Example: 12
	ID int64 `json:"id,omitempty"`

	// maps
	// Example: [[5,2,4],[2,2,1],[1,2,5],[3,3,1],[2,2,0]]
	// Required: true
	Maps [][]int64 `json:"maps"`
}

// Validate validates this level
func (m *Level) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Level) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Level) validateMaps(formats strfmt.Registry) error {

	if err := validate.Required("maps", "body", m.Maps); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this level based on context it is used
func (m *Level) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Level) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Level) UnmarshalBinary(b []byte) error {
	var res Level
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
