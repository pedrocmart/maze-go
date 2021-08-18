// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetLevelParams creates a new GetLevelParams object
//
// There are no default values defined in the spec.
func NewGetLevelParams() GetLevelParams {

	return GetLevelParams{}
}

// GetLevelParams contains all the bound params for the get level operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetLevel
type GetLevelParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: query
	*/
	LevelID int64
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetLevelParams() beforehand.
func (o *GetLevelParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qLevelID, qhkLevelID, _ := qs.GetOK("levelId")
	if err := o.bindLevelID(qLevelID, qhkLevelID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindLevelID binds and validates parameter LevelID from query.
func (o *GetLevelParams) bindLevelID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("levelId", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false

	if err := validate.RequiredString("levelId", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("levelId", "query", "int64", raw)
	}
	o.LevelID = value

	return nil
}