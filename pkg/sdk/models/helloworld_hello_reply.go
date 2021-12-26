// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HelloworldHelloReply The response message containing the greetings
//
// swagger:model helloworldHelloReply
type HelloworldHelloReply struct {

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this helloworld hello reply
func (m *HelloworldHelloReply) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this helloworld hello reply based on context it is used
func (m *HelloworldHelloReply) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HelloworldHelloReply) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HelloworldHelloReply) UnmarshalBinary(b []byte) error {
	var res HelloworldHelloReply
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}