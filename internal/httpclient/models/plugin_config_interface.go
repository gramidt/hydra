// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginConfigInterface PluginConfigInterface The interface between Docker and the plugin
//
// swagger:model PluginConfigInterface
type PluginConfigInterface struct {

	// Protocol to use for clients connecting to the plugin.
	ProtocolScheme string `json:"ProtocolScheme,omitempty"`

	// socket
	// Required: true
	Socket *string `json:"Socket"`

	// types
	// Required: true
	Types []*PluginInterfaceType `json:"Types"`
}

// Validate validates this plugin config interface
func (m *PluginConfigInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSocket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginConfigInterface) validateSocket(formats strfmt.Registry) error {

	if err := validate.Required("Socket", "body", m.Socket); err != nil {
		return err
	}

	return nil
}

func (m *PluginConfigInterface) validateTypes(formats strfmt.Registry) error {

	if err := validate.Required("Types", "body", m.Types); err != nil {
		return err
	}

	for i := 0; i < len(m.Types); i++ {
		if swag.IsZero(m.Types[i]) { // not required
			continue
		}

		if m.Types[i] != nil {
			if err := m.Types[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Types" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginConfigInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginConfigInterface) UnmarshalBinary(b []byte) error {
	var res PluginConfigInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
