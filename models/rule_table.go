// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RuleTable rule table
// swagger:model RuleTable
type RuleTable struct {

	// default
	Default interface{} `json:"default,omitempty"`

	// maximum
	Maximum interface{} `json:"maximum,omitempty"`

	// minimum
	Minimum interface{} `json:"minimum,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// required
	Required bool `json:"required,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this rule table
func (m *RuleTable) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RuleTable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RuleTable) UnmarshalBinary(b []byte) error {
	var res RuleTable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
