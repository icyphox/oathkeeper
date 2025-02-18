// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Rule Rule swaggerRule is a single rule that will get checked on every HTTP request.
//
// swagger:model rule
type Rule struct {

	// Authenticators is a list of authentication handlers that will try and authenticate the provided credentials.
	// Authenticators are checked iteratively from index 0 to n and if the first authenticator to return a positive
	// result will be the one used.
	//
	// If you want the rule to first check a specific authenticator  before "falling back" to others, have that authenticator
	// as the first item in the array.
	Authenticators []*RuleHandler `json:"authenticators"`

	// authorizer
	Authorizer *RuleHandler `json:"authorizer,omitempty"`

	// Description is a human readable description of this rule.
	Description string `json:"description,omitempty"`

	// ID is the unique id of the rule. It can be at most 190 characters long, but the layout of the ID is up to you.
	// You will need this ID later on to update or delete the rule.
	ID string `json:"id,omitempty"`

	// match
	Match *RuleMatch `json:"match,omitempty"`

	// Mutators is a list of mutation handlers that transform the HTTP request. A common use case is generating a new set
	// of credentials (e.g. JWT) which then will be forwarded to the upstream server.
	//
	// Mutations are performed iteratively from index 0 to n and should all succeed in order for the HTTP request to be forwarded.
	Mutators []*RuleHandler `json:"mutators"`

	// upstream
	Upstream *Upstream `json:"upstream,omitempty"`
}

// Validate validates this rule
func (m *Rule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthorizer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMutators(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpstream(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rule) validateAuthenticators(formats strfmt.Registry) error {

	if swag.IsZero(m.Authenticators) { // not required
		return nil
	}

	for i := 0; i < len(m.Authenticators); i++ {
		if swag.IsZero(m.Authenticators[i]) { // not required
			continue
		}

		if m.Authenticators[i] != nil {
			if err := m.Authenticators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("authenticators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Rule) validateAuthorizer(formats strfmt.Registry) error {

	if swag.IsZero(m.Authorizer) { // not required
		return nil
	}

	if m.Authorizer != nil {
		if err := m.Authorizer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authorizer")
			}
			return err
		}
	}

	return nil
}

func (m *Rule) validateMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Match) { // not required
		return nil
	}

	if m.Match != nil {
		if err := m.Match.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("match")
			}
			return err
		}
	}

	return nil
}

func (m *Rule) validateMutators(formats strfmt.Registry) error {

	if swag.IsZero(m.Mutators) { // not required
		return nil
	}

	for i := 0; i < len(m.Mutators); i++ {
		if swag.IsZero(m.Mutators[i]) { // not required
			continue
		}

		if m.Mutators[i] != nil {
			if err := m.Mutators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mutators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Rule) validateUpstream(formats strfmt.Registry) error {

	if swag.IsZero(m.Upstream) { // not required
		return nil
	}

	if m.Upstream != nil {
		if err := m.Upstream.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upstream")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Rule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rule) UnmarshalBinary(b []byte) error {
	var res Rule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
