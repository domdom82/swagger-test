// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InventoryItem inventory item
// swagger:model InventoryItem
type InventoryItem struct {

	// id
	// Required: true
	ID *strfmt.UUID `json:"id"`

	// manufacturer
	// Required: true
	Manufacturer *Manufacturer `json:"manufacturer"`

	// name
	// Required: true
	Name *string `json:"name"`

	// release date
	// Required: true
	ReleaseDate *int32 `json:"releaseDate"`
}

// Validate validates this inventory item
func (m *InventoryItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateManufacturer(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReleaseDate(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryItem) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.FormatOf("id", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InventoryItem) validateManufacturer(formats strfmt.Registry) error {

	if err := validate.Required("manufacturer", "body", m.Manufacturer); err != nil {
		return err
	}

	if m.Manufacturer != nil {

		if err := m.Manufacturer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("manufacturer")
			}
			return err
		}
	}

	return nil
}

func (m *InventoryItem) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *InventoryItem) validateReleaseDate(formats strfmt.Registry) error {

	if err := validate.Required("releaseDate", "body", m.ReleaseDate); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryItem) UnmarshalBinary(b []byte) error {
	var res InventoryItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
