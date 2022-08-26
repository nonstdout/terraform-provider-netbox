// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

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

// RearPortTemplate rear port template
//
// swagger:model RearPortTemplate
type RearPortTemplate struct {

	// Color
	// Max Length: 6
	// Pattern: ^[0-9a-f]{6}$
	Color string `json:"color,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device type
	DeviceType *NestedDeviceType `json:"device_type,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// module type
	ModuleType *NestedModuleType `json:"module_type,omitempty"`

	// Name
	//
	//
	// {module} is accepted as a substitution for the module bay position when attached to a module type.
	//
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Positions
	// Maximum: 1024
	// Minimum: 1
	Positions int64 `json:"positions,omitempty"`

	// type
	// Required: true
	Type *RearPortTemplateType `json:"type"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this rear port template
func (m *RearPortTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateColor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModuleType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePositions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RearPortTemplate) validateColor(formats strfmt.Registry) error {
	if swag.IsZero(m.Color) { // not required
		return nil
	}

	if err := validate.MaxLength("color", "body", m.Color, 6); err != nil {
		return err
	}

	if err := validate.Pattern("color", "body", m.Color, `^[0-9a-f]{6}$`); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validateDeviceType(formats strfmt.Registry) error {
	if swag.IsZero(m.DeviceType) { // not required
		return nil
	}

	if m.DeviceType != nil {
		if err := m.DeviceType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPortTemplate) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validateModuleType(formats strfmt.Registry) error {
	if swag.IsZero(m.ModuleType) { // not required
		return nil
	}

	if m.ModuleType != nil {
		if err := m.ModuleType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module_type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPortTemplate) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validatePositions(formats strfmt.Registry) error {
	if swag.IsZero(m.Positions) { // not required
		return nil
	}

	if err := validate.MinimumInt("positions", "body", m.Positions, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("positions", "body", m.Positions, 1024, false); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPortTemplate) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this rear port template based on the context it is used
func (m *RearPortTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeviceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModuleType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RearPortTemplate) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) contextValidateDeviceType(ctx context.Context, formats strfmt.Registry) error {

	if m.DeviceType != nil {
		if err := m.DeviceType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device_type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPortTemplate) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *RearPortTemplate) contextValidateModuleType(ctx context.Context, formats strfmt.Registry) error {

	if m.ModuleType != nil {
		if err := m.ModuleType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module_type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPortTemplate) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *RearPortTemplate) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RearPortTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RearPortTemplate) UnmarshalBinary(b []byte) error {
	var res RearPortTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RearPortTemplateType Type
//
// swagger:model RearPortTemplateType
type RearPortTemplateType struct {

	// label
	// Required: true
	// Enum: [8P8C 8P6C 8P4C 8P2C 6P6C 6P4C 6P2C 4P4C 4P2C GG45 TERA 4P TERA 2P TERA 1P 110 Punch BNC F Connector N Connector MRJ21 FC LC LC/PC LC/UPC LC/APC LSH LSH/PC LSH/UPC LSH/APC MPO MTRJ SC SC/PC SC/UPC SC/APC ST CS SN SMA 905 SMA 906 URM-P2 URM-P4 URM-P8 Splice Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [8p8c 8p6c 8p4c 8p2c 6p6c 6p4c 6p2c 4p4c 4p2c gg45 tera-4p tera-2p tera-1p 110-punch bnc f n mrj21 fc lc lc-pc lc-upc lc-apc lsh lsh-pc lsh-upc lsh-apc mpo mtrj sc sc-pc sc-upc sc-apc st cs sn sma-905 sma-906 urm-p2 urm-p4 urm-p8 splice other]
	Value *string `json:"value"`
}

// Validate validates this rear port template type
func (m *RearPortTemplateType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var rearPortTemplateTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8P8C","8P6C","8P4C","8P2C","6P6C","6P4C","6P2C","4P4C","4P2C","GG45","TERA 4P","TERA 2P","TERA 1P","110 Punch","BNC","F Connector","N Connector","MRJ21","FC","LC","LC/PC","LC/UPC","LC/APC","LSH","LSH/PC","LSH/UPC","LSH/APC","MPO","MTRJ","SC","SC/PC","SC/UPC","SC/APC","ST","CS","SN","SMA 905","SMA 906","URM-P2","URM-P4","URM-P8","Splice","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rearPortTemplateTypeTypeLabelPropEnum = append(rearPortTemplateTypeTypeLabelPropEnum, v)
	}
}

const (

	// RearPortTemplateTypeLabelNr8P8C captures enum value "8P8C"
	RearPortTemplateTypeLabelNr8P8C string = "8P8C"

	// RearPortTemplateTypeLabelNr8P6C captures enum value "8P6C"
	RearPortTemplateTypeLabelNr8P6C string = "8P6C"

	// RearPortTemplateTypeLabelNr8P4C captures enum value "8P4C"
	RearPortTemplateTypeLabelNr8P4C string = "8P4C"

	// RearPortTemplateTypeLabelNr8P2C captures enum value "8P2C"
	RearPortTemplateTypeLabelNr8P2C string = "8P2C"

	// RearPortTemplateTypeLabelNr6P6C captures enum value "6P6C"
	RearPortTemplateTypeLabelNr6P6C string = "6P6C"

	// RearPortTemplateTypeLabelNr6P4C captures enum value "6P4C"
	RearPortTemplateTypeLabelNr6P4C string = "6P4C"

	// RearPortTemplateTypeLabelNr6P2C captures enum value "6P2C"
	RearPortTemplateTypeLabelNr6P2C string = "6P2C"

	// RearPortTemplateTypeLabelNr4P4C captures enum value "4P4C"
	RearPortTemplateTypeLabelNr4P4C string = "4P4C"

	// RearPortTemplateTypeLabelNr4P2C captures enum value "4P2C"
	RearPortTemplateTypeLabelNr4P2C string = "4P2C"

	// RearPortTemplateTypeLabelGG45 captures enum value "GG45"
	RearPortTemplateTypeLabelGG45 string = "GG45"

	// RearPortTemplateTypeLabelTERA4P captures enum value "TERA 4P"
	RearPortTemplateTypeLabelTERA4P string = "TERA 4P"

	// RearPortTemplateTypeLabelTERA2P captures enum value "TERA 2P"
	RearPortTemplateTypeLabelTERA2P string = "TERA 2P"

	// RearPortTemplateTypeLabelTERA1P captures enum value "TERA 1P"
	RearPortTemplateTypeLabelTERA1P string = "TERA 1P"

	// RearPortTemplateTypeLabelNr110Punch captures enum value "110 Punch"
	RearPortTemplateTypeLabelNr110Punch string = "110 Punch"

	// RearPortTemplateTypeLabelBNC captures enum value "BNC"
	RearPortTemplateTypeLabelBNC string = "BNC"

	// RearPortTemplateTypeLabelFConnector captures enum value "F Connector"
	RearPortTemplateTypeLabelFConnector string = "F Connector"

	// RearPortTemplateTypeLabelNConnector captures enum value "N Connector"
	RearPortTemplateTypeLabelNConnector string = "N Connector"

	// RearPortTemplateTypeLabelMRJ21 captures enum value "MRJ21"
	RearPortTemplateTypeLabelMRJ21 string = "MRJ21"

	// RearPortTemplateTypeLabelFC captures enum value "FC"
	RearPortTemplateTypeLabelFC string = "FC"

	// RearPortTemplateTypeLabelLC captures enum value "LC"
	RearPortTemplateTypeLabelLC string = "LC"

	// RearPortTemplateTypeLabelLCPC captures enum value "LC/PC"
	RearPortTemplateTypeLabelLCPC string = "LC/PC"

	// RearPortTemplateTypeLabelLCUPC captures enum value "LC/UPC"
	RearPortTemplateTypeLabelLCUPC string = "LC/UPC"

	// RearPortTemplateTypeLabelLCAPC captures enum value "LC/APC"
	RearPortTemplateTypeLabelLCAPC string = "LC/APC"

	// RearPortTemplateTypeLabelLSH captures enum value "LSH"
	RearPortTemplateTypeLabelLSH string = "LSH"

	// RearPortTemplateTypeLabelLSHPC captures enum value "LSH/PC"
	RearPortTemplateTypeLabelLSHPC string = "LSH/PC"

	// RearPortTemplateTypeLabelLSHUPC captures enum value "LSH/UPC"
	RearPortTemplateTypeLabelLSHUPC string = "LSH/UPC"

	// RearPortTemplateTypeLabelLSHAPC captures enum value "LSH/APC"
	RearPortTemplateTypeLabelLSHAPC string = "LSH/APC"

	// RearPortTemplateTypeLabelMPO captures enum value "MPO"
	RearPortTemplateTypeLabelMPO string = "MPO"

	// RearPortTemplateTypeLabelMTRJ captures enum value "MTRJ"
	RearPortTemplateTypeLabelMTRJ string = "MTRJ"

	// RearPortTemplateTypeLabelSC captures enum value "SC"
	RearPortTemplateTypeLabelSC string = "SC"

	// RearPortTemplateTypeLabelSCPC captures enum value "SC/PC"
	RearPortTemplateTypeLabelSCPC string = "SC/PC"

	// RearPortTemplateTypeLabelSCUPC captures enum value "SC/UPC"
	RearPortTemplateTypeLabelSCUPC string = "SC/UPC"

	// RearPortTemplateTypeLabelSCAPC captures enum value "SC/APC"
	RearPortTemplateTypeLabelSCAPC string = "SC/APC"

	// RearPortTemplateTypeLabelST captures enum value "ST"
	RearPortTemplateTypeLabelST string = "ST"

	// RearPortTemplateTypeLabelCS captures enum value "CS"
	RearPortTemplateTypeLabelCS string = "CS"

	// RearPortTemplateTypeLabelSN captures enum value "SN"
	RearPortTemplateTypeLabelSN string = "SN"

	// RearPortTemplateTypeLabelSMA905 captures enum value "SMA 905"
	RearPortTemplateTypeLabelSMA905 string = "SMA 905"

	// RearPortTemplateTypeLabelSMA906 captures enum value "SMA 906"
	RearPortTemplateTypeLabelSMA906 string = "SMA 906"

	// RearPortTemplateTypeLabelURMDashP2 captures enum value "URM-P2"
	RearPortTemplateTypeLabelURMDashP2 string = "URM-P2"

	// RearPortTemplateTypeLabelURMDashP4 captures enum value "URM-P4"
	RearPortTemplateTypeLabelURMDashP4 string = "URM-P4"

	// RearPortTemplateTypeLabelURMDashP8 captures enum value "URM-P8"
	RearPortTemplateTypeLabelURMDashP8 string = "URM-P8"

	// RearPortTemplateTypeLabelSplice captures enum value "Splice"
	RearPortTemplateTypeLabelSplice string = "Splice"

	// RearPortTemplateTypeLabelOther captures enum value "Other"
	RearPortTemplateTypeLabelOther string = "Other"
)

// prop value enum
func (m *RearPortTemplateType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rearPortTemplateTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RearPortTemplateType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var rearPortTemplateTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["8p8c","8p6c","8p4c","8p2c","6p6c","6p4c","6p2c","4p4c","4p2c","gg45","tera-4p","tera-2p","tera-1p","110-punch","bnc","f","n","mrj21","fc","lc","lc-pc","lc-upc","lc-apc","lsh","lsh-pc","lsh-upc","lsh-apc","mpo","mtrj","sc","sc-pc","sc-upc","sc-apc","st","cs","sn","sma-905","sma-906","urm-p2","urm-p4","urm-p8","splice","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rearPortTemplateTypeTypeValuePropEnum = append(rearPortTemplateTypeTypeValuePropEnum, v)
	}
}

const (

	// RearPortTemplateTypeValueNr8p8c captures enum value "8p8c"
	RearPortTemplateTypeValueNr8p8c string = "8p8c"

	// RearPortTemplateTypeValueNr8p6c captures enum value "8p6c"
	RearPortTemplateTypeValueNr8p6c string = "8p6c"

	// RearPortTemplateTypeValueNr8p4c captures enum value "8p4c"
	RearPortTemplateTypeValueNr8p4c string = "8p4c"

	// RearPortTemplateTypeValueNr8p2c captures enum value "8p2c"
	RearPortTemplateTypeValueNr8p2c string = "8p2c"

	// RearPortTemplateTypeValueNr6p6c captures enum value "6p6c"
	RearPortTemplateTypeValueNr6p6c string = "6p6c"

	// RearPortTemplateTypeValueNr6p4c captures enum value "6p4c"
	RearPortTemplateTypeValueNr6p4c string = "6p4c"

	// RearPortTemplateTypeValueNr6p2c captures enum value "6p2c"
	RearPortTemplateTypeValueNr6p2c string = "6p2c"

	// RearPortTemplateTypeValueNr4p4c captures enum value "4p4c"
	RearPortTemplateTypeValueNr4p4c string = "4p4c"

	// RearPortTemplateTypeValueNr4p2c captures enum value "4p2c"
	RearPortTemplateTypeValueNr4p2c string = "4p2c"

	// RearPortTemplateTypeValueGg45 captures enum value "gg45"
	RearPortTemplateTypeValueGg45 string = "gg45"

	// RearPortTemplateTypeValueTeraDash4p captures enum value "tera-4p"
	RearPortTemplateTypeValueTeraDash4p string = "tera-4p"

	// RearPortTemplateTypeValueTeraDash2p captures enum value "tera-2p"
	RearPortTemplateTypeValueTeraDash2p string = "tera-2p"

	// RearPortTemplateTypeValueTeraDash1p captures enum value "tera-1p"
	RearPortTemplateTypeValueTeraDash1p string = "tera-1p"

	// RearPortTemplateTypeValueNr110DashPunch captures enum value "110-punch"
	RearPortTemplateTypeValueNr110DashPunch string = "110-punch"

	// RearPortTemplateTypeValueBnc captures enum value "bnc"
	RearPortTemplateTypeValueBnc string = "bnc"

	// RearPortTemplateTypeValueF captures enum value "f"
	RearPortTemplateTypeValueF string = "f"

	// RearPortTemplateTypeValueN captures enum value "n"
	RearPortTemplateTypeValueN string = "n"

	// RearPortTemplateTypeValueMrj21 captures enum value "mrj21"
	RearPortTemplateTypeValueMrj21 string = "mrj21"

	// RearPortTemplateTypeValueFc captures enum value "fc"
	RearPortTemplateTypeValueFc string = "fc"

	// RearPortTemplateTypeValueLc captures enum value "lc"
	RearPortTemplateTypeValueLc string = "lc"

	// RearPortTemplateTypeValueLcDashPc captures enum value "lc-pc"
	RearPortTemplateTypeValueLcDashPc string = "lc-pc"

	// RearPortTemplateTypeValueLcDashUpc captures enum value "lc-upc"
	RearPortTemplateTypeValueLcDashUpc string = "lc-upc"

	// RearPortTemplateTypeValueLcDashApc captures enum value "lc-apc"
	RearPortTemplateTypeValueLcDashApc string = "lc-apc"

	// RearPortTemplateTypeValueLsh captures enum value "lsh"
	RearPortTemplateTypeValueLsh string = "lsh"

	// RearPortTemplateTypeValueLshDashPc captures enum value "lsh-pc"
	RearPortTemplateTypeValueLshDashPc string = "lsh-pc"

	// RearPortTemplateTypeValueLshDashUpc captures enum value "lsh-upc"
	RearPortTemplateTypeValueLshDashUpc string = "lsh-upc"

	// RearPortTemplateTypeValueLshDashApc captures enum value "lsh-apc"
	RearPortTemplateTypeValueLshDashApc string = "lsh-apc"

	// RearPortTemplateTypeValueMpo captures enum value "mpo"
	RearPortTemplateTypeValueMpo string = "mpo"

	// RearPortTemplateTypeValueMtrj captures enum value "mtrj"
	RearPortTemplateTypeValueMtrj string = "mtrj"

	// RearPortTemplateTypeValueSc captures enum value "sc"
	RearPortTemplateTypeValueSc string = "sc"

	// RearPortTemplateTypeValueScDashPc captures enum value "sc-pc"
	RearPortTemplateTypeValueScDashPc string = "sc-pc"

	// RearPortTemplateTypeValueScDashUpc captures enum value "sc-upc"
	RearPortTemplateTypeValueScDashUpc string = "sc-upc"

	// RearPortTemplateTypeValueScDashApc captures enum value "sc-apc"
	RearPortTemplateTypeValueScDashApc string = "sc-apc"

	// RearPortTemplateTypeValueSt captures enum value "st"
	RearPortTemplateTypeValueSt string = "st"

	// RearPortTemplateTypeValueCs captures enum value "cs"
	RearPortTemplateTypeValueCs string = "cs"

	// RearPortTemplateTypeValueSn captures enum value "sn"
	RearPortTemplateTypeValueSn string = "sn"

	// RearPortTemplateTypeValueSmaDash905 captures enum value "sma-905"
	RearPortTemplateTypeValueSmaDash905 string = "sma-905"

	// RearPortTemplateTypeValueSmaDash906 captures enum value "sma-906"
	RearPortTemplateTypeValueSmaDash906 string = "sma-906"

	// RearPortTemplateTypeValueUrmDashP2 captures enum value "urm-p2"
	RearPortTemplateTypeValueUrmDashP2 string = "urm-p2"

	// RearPortTemplateTypeValueUrmDashP4 captures enum value "urm-p4"
	RearPortTemplateTypeValueUrmDashP4 string = "urm-p4"

	// RearPortTemplateTypeValueUrmDashP8 captures enum value "urm-p8"
	RearPortTemplateTypeValueUrmDashP8 string = "urm-p8"

	// RearPortTemplateTypeValueSplice captures enum value "splice"
	RearPortTemplateTypeValueSplice string = "splice"

	// RearPortTemplateTypeValueOther captures enum value "other"
	RearPortTemplateTypeValueOther string = "other"
)

// prop value enum
func (m *RearPortTemplateType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, rearPortTemplateTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RearPortTemplateType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rear port template type based on context it is used
func (m *RearPortTemplateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RearPortTemplateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RearPortTemplateType) UnmarshalBinary(b []byte) error {
	var res RearPortTemplateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
