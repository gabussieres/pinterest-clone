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

// User user
// swagger:model user
type User struct {

	// first name
	// Min Length: 1
	FirstName string `json:"first_name,omitempty"`

	// followers
	Followers int64 `json:"followers,omitempty"`

	// following
	Following int64 `json:"following,omitempty"`

	// id
	// Min Length: 1
	ID string `json:"id,omitempty"`

	// image url
	// Min Length: 1
	ImageURL string `json:"image_url,omitempty"`

	// last name
	// Min Length: 1
	LastName string `json:"last_name,omitempty"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateFirstName(formats strfmt.Registry) error {

	if swag.IsZero(m.FirstName) { // not required
		return nil
	}

	if err := validate.MinLength("first_name", "body", string(m.FirstName), 1); err != nil {
		return err
	}

	return nil
}

func (m *User) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.MinLength("id", "body", string(m.ID), 1); err != nil {
		return err
	}

	return nil
}

func (m *User) validateImageURL(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageURL) { // not required
		return nil
	}

	if err := validate.MinLength("image_url", "body", string(m.ImageURL), 1); err != nil {
		return err
	}

	return nil
}

func (m *User) validateLastName(formats strfmt.Registry) error {

	if swag.IsZero(m.LastName) { // not required
		return nil
	}

	if err := validate.MinLength("last_name", "body", string(m.LastName), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
