// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IpamAllocArgs Alloc IP request args
//
// swagger:model IpamAllocArgs
type IpamAllocArgs struct {

	// container ID
	// Required: true
	ContainerID *string `json:"containerID"`

	// if name
	// Required: true
	IfName *string `json:"ifName"`

	// ipam
	Ipam string `json:"ipam,omitempty"`

	// net namespace
	// Required: true
	NetNamespace *string `json:"netNamespace"`

	// pod name
	// Required: true
	PodName *string `json:"podName"`

	// pod namespace
	// Required: true
	PodNamespace *string `json:"podNamespace"`
}

// Validate validates this ipam alloc args
func (m *IpamAllocArgs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContainerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIfName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePodNamespace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IpamAllocArgs) validateContainerID(formats strfmt.Registry) error {

	if err := validate.Required("containerID", "body", m.ContainerID); err != nil {
		return err
	}

	return nil
}

func (m *IpamAllocArgs) validateIfName(formats strfmt.Registry) error {

	if err := validate.Required("ifName", "body", m.IfName); err != nil {
		return err
	}

	return nil
}

func (m *IpamAllocArgs) validateNetNamespace(formats strfmt.Registry) error {

	if err := validate.Required("netNamespace", "body", m.NetNamespace); err != nil {
		return err
	}

	return nil
}

func (m *IpamAllocArgs) validatePodName(formats strfmt.Registry) error {

	if err := validate.Required("podName", "body", m.PodName); err != nil {
		return err
	}

	return nil
}

func (m *IpamAllocArgs) validatePodNamespace(formats strfmt.Registry) error {

	if err := validate.Required("podNamespace", "body", m.PodNamespace); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ipam alloc args based on context it is used
func (m *IpamAllocArgs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IpamAllocArgs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IpamAllocArgs) UnmarshalBinary(b []byte) error {
	var res IpamAllocArgs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

func (m *IpamAllocArgs) String() string{
	return fmt.Sprintf("",*m.IfName,m.Ipam,*m.PodName,*m.PodNamespace,*m.NetNamespace,*m.ContainerID)
}