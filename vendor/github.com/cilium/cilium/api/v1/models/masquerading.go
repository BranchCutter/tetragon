// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Masquerading Status of masquerading
//
// +k8s:deepcopy-gen=true
//
// swagger:model Masquerading
type Masquerading struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// enabled protocols
	EnabledProtocols *MasqueradingEnabledProtocols `json:"enabledProtocols,omitempty"`

	// Is BPF ip-masq-agent enabled
	IPMasqAgent bool `json:"ip-masq-agent,omitempty"`

	// mode
	// Enum: ["BPF","iptables"]
	Mode string `json:"mode,omitempty"`

	// This field is obsolete, please use snat-exclusion-cidr-v4 or snat-exclusion-cidr-v6.
	SnatExclusionCidr string `json:"snat-exclusion-cidr,omitempty"`

	// SnatExclusionCIDRv4 exempts SNAT from being performed on any packet sent to
	// an IPv4 address that belongs to this CIDR.
	SnatExclusionCidrV4 string `json:"snat-exclusion-cidr-v4,omitempty"`

	// SnatExclusionCIDRv6 exempts SNAT from being performed on any packet sent to
	// an IPv6 address that belongs to this CIDR.
	// For IPv6 we only do masquerading in iptables mode.
	SnatExclusionCidrV6 string `json:"snat-exclusion-cidr-v6,omitempty"`
}

// Validate validates this masquerading
func (m *Masquerading) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnabledProtocols(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Masquerading) validateEnabledProtocols(formats strfmt.Registry) error {
	if swag.IsZero(m.EnabledProtocols) { // not required
		return nil
	}

	if m.EnabledProtocols != nil {
		if err := m.EnabledProtocols.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enabledProtocols")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enabledProtocols")
			}
			return err
		}
	}

	return nil
}

var masqueradingTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BPF","iptables"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		masqueradingTypeModePropEnum = append(masqueradingTypeModePropEnum, v)
	}
}

const (

	// MasqueradingModeBPF captures enum value "BPF"
	MasqueradingModeBPF string = "BPF"

	// MasqueradingModeIptables captures enum value "iptables"
	MasqueradingModeIptables string = "iptables"
)

// prop value enum
func (m *Masquerading) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, masqueradingTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Masquerading) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this masquerading based on the context it is used
func (m *Masquerading) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnabledProtocols(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Masquerading) contextValidateEnabledProtocols(ctx context.Context, formats strfmt.Registry) error {

	if m.EnabledProtocols != nil {

		if swag.IsZero(m.EnabledProtocols) { // not required
			return nil
		}

		if err := m.EnabledProtocols.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enabledProtocols")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enabledProtocols")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Masquerading) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Masquerading) UnmarshalBinary(b []byte) error {
	var res Masquerading
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MasqueradingEnabledProtocols Is masquerading enabled
//
// swagger:model MasqueradingEnabledProtocols
type MasqueradingEnabledProtocols struct {

	// Is masquerading enabled for IPv4 traffic
	IPV4 bool `json:"ipv4,omitempty"`

	// Is masquerading enabled for IPv6 traffic
	IPV6 bool `json:"ipv6,omitempty"`
}

// Validate validates this masquerading enabled protocols
func (m *MasqueradingEnabledProtocols) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this masquerading enabled protocols based on context it is used
func (m *MasqueradingEnabledProtocols) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MasqueradingEnabledProtocols) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MasqueradingEnabledProtocols) UnmarshalBinary(b []byte) error {
	var res MasqueradingEnabledProtocols
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
