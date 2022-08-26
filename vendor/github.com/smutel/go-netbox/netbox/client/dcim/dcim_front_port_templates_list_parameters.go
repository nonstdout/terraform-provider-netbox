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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDcimFrontPortTemplatesListParams creates a new DcimFrontPortTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimFrontPortTemplatesListParams() *DcimFrontPortTemplatesListParams {
	return &DcimFrontPortTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimFrontPortTemplatesListParamsWithTimeout creates a new DcimFrontPortTemplatesListParams object
// with the ability to set a timeout on a request.
func NewDcimFrontPortTemplatesListParamsWithTimeout(timeout time.Duration) *DcimFrontPortTemplatesListParams {
	return &DcimFrontPortTemplatesListParams{
		timeout: timeout,
	}
}

// NewDcimFrontPortTemplatesListParamsWithContext creates a new DcimFrontPortTemplatesListParams object
// with the ability to set a context for a request.
func NewDcimFrontPortTemplatesListParamsWithContext(ctx context.Context) *DcimFrontPortTemplatesListParams {
	return &DcimFrontPortTemplatesListParams{
		Context: ctx,
	}
}

// NewDcimFrontPortTemplatesListParamsWithHTTPClient creates a new DcimFrontPortTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimFrontPortTemplatesListParamsWithHTTPClient(client *http.Client) *DcimFrontPortTemplatesListParams {
	return &DcimFrontPortTemplatesListParams{
		HTTPClient: client,
	}
}

/* DcimFrontPortTemplatesListParams contains all the parameters to send to the API endpoint
   for the dcim front port templates list operation.

   Typically these are written to a http.Request.
*/
type DcimFrontPortTemplatesListParams struct {

	// Color.
	Color *string

	// ColorIc.
	ColorIc *string

	// ColorIe.
	ColorIe *string

	// ColorIew.
	ColorIew *string

	// ColorIsw.
	ColorIsw *string

	// Colorn.
	Colorn *string

	// ColorNic.
	ColorNic *string

	// ColorNie.
	ColorNie *string

	// ColorNiew.
	ColorNiew *string

	// ColorNisw.
	ColorNisw *string

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// DevicetypeID.
	DevicetypeID *string

	// DevicetypeIDn.
	DevicetypeIDn *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// ModuletypeID.
	ModuletypeID *string

	// ModuletypeIDn.
	ModuletypeIDn *string

	// Name.
	Name *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	// Q.
	Q *string

	// Type.
	Type *string

	// Typen.
	Typen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim front port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesListParams) WithDefaults() *DcimFrontPortTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim front port templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimFrontPortTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithTimeout(timeout time.Duration) *DcimFrontPortTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithContext(ctx context.Context) *DcimFrontPortTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithHTTPClient(client *http.Client) *DcimFrontPortTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithColor adds the color to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithColor(color *string) *DcimFrontPortTemplatesListParams {
	o.SetColor(color)
	return o
}

// SetColor adds the color to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetColor(color *string) {
	o.Color = color
}

// WithColorIc adds the colorIc to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithColorIc(colorIc *string) *DcimFrontPortTemplatesListParams {
	o.SetColorIc(colorIc)
	return o
}

// SetColorIc adds the colorIc to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetColorIc(colorIc *string) {
	o.ColorIc = colorIc
}

// WithColorIe adds the colorIe to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithColorIe(colorIe *string) *DcimFrontPortTemplatesListParams {
	o.SetColorIe(colorIe)
	return o
}

// SetColorIe adds the colorIe to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetColorIe(colorIe *string) {
	o.ColorIe = colorIe
}

// WithColorIew adds the colorIew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithColorIew(colorIew *string) *DcimFrontPortTemplatesListParams {
	o.SetColorIew(colorIew)
	return o
}

// SetColorIew adds the colorIew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetColorIew(colorIew *string) {
	o.ColorIew = colorIew
}

// WithColorIsw adds the colorIsw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithColorIsw(colorIsw *string) *DcimFrontPortTemplatesListParams {
	o.SetColorIsw(colorIsw)
	return o
}

// SetColorIsw adds the colorIsw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetColorIsw(colorIsw *string) {
	o.ColorIsw = colorIsw
}

// WithColorn adds the colorn to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithColorn(colorn *string) *DcimFrontPortTemplatesListParams {
	o.SetColorn(colorn)
	return o
}

// SetColorn adds the colorN to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetColorn(colorn *string) {
	o.Colorn = colorn
}

// WithColorNic adds the colorNic to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithColorNic(colorNic *string) *DcimFrontPortTemplatesListParams {
	o.SetColorNic(colorNic)
	return o
}

// SetColorNic adds the colorNic to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetColorNic(colorNic *string) {
	o.ColorNic = colorNic
}

// WithColorNie adds the colorNie to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithColorNie(colorNie *string) *DcimFrontPortTemplatesListParams {
	o.SetColorNie(colorNie)
	return o
}

// SetColorNie adds the colorNie to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetColorNie(colorNie *string) {
	o.ColorNie = colorNie
}

// WithColorNiew adds the colorNiew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithColorNiew(colorNiew *string) *DcimFrontPortTemplatesListParams {
	o.SetColorNiew(colorNiew)
	return o
}

// SetColorNiew adds the colorNiew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetColorNiew(colorNiew *string) {
	o.ColorNiew = colorNiew
}

// WithColorNisw adds the colorNisw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithColorNisw(colorNisw *string) *DcimFrontPortTemplatesListParams {
	o.SetColorNisw(colorNisw)
	return o
}

// SetColorNisw adds the colorNisw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetColorNisw(colorNisw *string) {
	o.ColorNisw = colorNisw
}

// WithCreated adds the created to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithCreated(created *string) *DcimFrontPortTemplatesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithCreatedGte(createdGte *string) *DcimFrontPortTemplatesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithCreatedLte(createdLte *string) *DcimFrontPortTemplatesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDevicetypeID adds the devicetypeID to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithDevicetypeID(devicetypeID *string) *DcimFrontPortTemplatesListParams {
	o.SetDevicetypeID(devicetypeID)
	return o
}

// SetDevicetypeID adds the devicetypeId to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetDevicetypeID(devicetypeID *string) {
	o.DevicetypeID = devicetypeID
}

// WithDevicetypeIDn adds the devicetypeIDn to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithDevicetypeIDn(devicetypeIDn *string) *DcimFrontPortTemplatesListParams {
	o.SetDevicetypeIDn(devicetypeIDn)
	return o
}

// SetDevicetypeIDn adds the devicetypeIdN to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetDevicetypeIDn(devicetypeIDn *string) {
	o.DevicetypeIDn = devicetypeIDn
}

// WithID adds the id to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithID(id *string) *DcimFrontPortTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithIDGt(iDGt *string) *DcimFrontPortTemplatesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithIDGte(iDGte *string) *DcimFrontPortTemplatesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithIDLt(iDLt *string) *DcimFrontPortTemplatesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithIDLte(iDLte *string) *DcimFrontPortTemplatesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithIDn(iDn *string) *DcimFrontPortTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithLastUpdated(lastUpdated *string) *DcimFrontPortTemplatesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGte adds the lastUpdatedGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *DcimFrontPortTemplatesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLte adds the lastUpdatedLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *DcimFrontPortTemplatesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLimit adds the limit to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithLimit(limit *int64) *DcimFrontPortTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithModuletypeID adds the moduletypeID to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithModuletypeID(moduletypeID *string) *DcimFrontPortTemplatesListParams {
	o.SetModuletypeID(moduletypeID)
	return o
}

// SetModuletypeID adds the moduletypeId to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetModuletypeID(moduletypeID *string) {
	o.ModuletypeID = moduletypeID
}

// WithModuletypeIDn adds the moduletypeIDn to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithModuletypeIDn(moduletypeIDn *string) *DcimFrontPortTemplatesListParams {
	o.SetModuletypeIDn(moduletypeIDn)
	return o
}

// SetModuletypeIDn adds the moduletypeIdN to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetModuletypeIDn(moduletypeIDn *string) {
	o.ModuletypeIDn = moduletypeIDn
}

// WithName adds the name to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithName(name *string) *DcimFrontPortTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameIc adds the nameIc to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameIc(nameIc *string) *DcimFrontPortTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameIe(nameIe *string) *DcimFrontPortTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameIew(nameIew *string) *DcimFrontPortTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameIsw(nameIsw *string) *DcimFrontPortTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNamen(namen *string) *DcimFrontPortTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameNic(nameNic *string) *DcimFrontPortTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameNie(nameNie *string) *DcimFrontPortTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameNiew(nameNiew *string) *DcimFrontPortTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithNameNisw(nameNisw *string) *DcimFrontPortTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithOffset(offset *int64) *DcimFrontPortTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQ adds the q to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithQ(q *string) *DcimFrontPortTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WithType adds the typeVar to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithType(typeVar *string) *DcimFrontPortTemplatesListParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WithTypen adds the typen to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) WithTypen(typen *string) *DcimFrontPortTemplatesListParams {
	o.SetTypen(typen)
	return o
}

// SetTypen adds the typeN to the dcim front port templates list params
func (o *DcimFrontPortTemplatesListParams) SetTypen(typen *string) {
	o.Typen = typen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimFrontPortTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Color != nil {

		// query param color
		var qrColor string

		if o.Color != nil {
			qrColor = *o.Color
		}
		qColor := qrColor
		if qColor != "" {

			if err := r.SetQueryParam("color", qColor); err != nil {
				return err
			}
		}
	}

	if o.ColorIc != nil {

		// query param color__ic
		var qrColorIc string

		if o.ColorIc != nil {
			qrColorIc = *o.ColorIc
		}
		qColorIc := qrColorIc
		if qColorIc != "" {

			if err := r.SetQueryParam("color__ic", qColorIc); err != nil {
				return err
			}
		}
	}

	if o.ColorIe != nil {

		// query param color__ie
		var qrColorIe string

		if o.ColorIe != nil {
			qrColorIe = *o.ColorIe
		}
		qColorIe := qrColorIe
		if qColorIe != "" {

			if err := r.SetQueryParam("color__ie", qColorIe); err != nil {
				return err
			}
		}
	}

	if o.ColorIew != nil {

		// query param color__iew
		var qrColorIew string

		if o.ColorIew != nil {
			qrColorIew = *o.ColorIew
		}
		qColorIew := qrColorIew
		if qColorIew != "" {

			if err := r.SetQueryParam("color__iew", qColorIew); err != nil {
				return err
			}
		}
	}

	if o.ColorIsw != nil {

		// query param color__isw
		var qrColorIsw string

		if o.ColorIsw != nil {
			qrColorIsw = *o.ColorIsw
		}
		qColorIsw := qrColorIsw
		if qColorIsw != "" {

			if err := r.SetQueryParam("color__isw", qColorIsw); err != nil {
				return err
			}
		}
	}

	if o.Colorn != nil {

		// query param color__n
		var qrColorn string

		if o.Colorn != nil {
			qrColorn = *o.Colorn
		}
		qColorn := qrColorn
		if qColorn != "" {

			if err := r.SetQueryParam("color__n", qColorn); err != nil {
				return err
			}
		}
	}

	if o.ColorNic != nil {

		// query param color__nic
		var qrColorNic string

		if o.ColorNic != nil {
			qrColorNic = *o.ColorNic
		}
		qColorNic := qrColorNic
		if qColorNic != "" {

			if err := r.SetQueryParam("color__nic", qColorNic); err != nil {
				return err
			}
		}
	}

	if o.ColorNie != nil {

		// query param color__nie
		var qrColorNie string

		if o.ColorNie != nil {
			qrColorNie = *o.ColorNie
		}
		qColorNie := qrColorNie
		if qColorNie != "" {

			if err := r.SetQueryParam("color__nie", qColorNie); err != nil {
				return err
			}
		}
	}

	if o.ColorNiew != nil {

		// query param color__niew
		var qrColorNiew string

		if o.ColorNiew != nil {
			qrColorNiew = *o.ColorNiew
		}
		qColorNiew := qrColorNiew
		if qColorNiew != "" {

			if err := r.SetQueryParam("color__niew", qColorNiew); err != nil {
				return err
			}
		}
	}

	if o.ColorNisw != nil {

		// query param color__nisw
		var qrColorNisw string

		if o.ColorNisw != nil {
			qrColorNisw = *o.ColorNisw
		}
		qColorNisw := qrColorNisw
		if qColorNisw != "" {

			if err := r.SetQueryParam("color__nisw", qColorNisw); err != nil {
				return err
			}
		}
	}

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.DevicetypeID != nil {

		// query param devicetype_id
		var qrDevicetypeID string

		if o.DevicetypeID != nil {
			qrDevicetypeID = *o.DevicetypeID
		}
		qDevicetypeID := qrDevicetypeID
		if qDevicetypeID != "" {

			if err := r.SetQueryParam("devicetype_id", qDevicetypeID); err != nil {
				return err
			}
		}
	}

	if o.DevicetypeIDn != nil {

		// query param devicetype_id__n
		var qrDevicetypeIDn string

		if o.DevicetypeIDn != nil {
			qrDevicetypeIDn = *o.DevicetypeIDn
		}
		qDevicetypeIDn := qrDevicetypeIDn
		if qDevicetypeIDn != "" {

			if err := r.SetQueryParam("devicetype_id__n", qDevicetypeIDn); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.ModuletypeID != nil {

		// query param moduletype_id
		var qrModuletypeID string

		if o.ModuletypeID != nil {
			qrModuletypeID = *o.ModuletypeID
		}
		qModuletypeID := qrModuletypeID
		if qModuletypeID != "" {

			if err := r.SetQueryParam("moduletype_id", qModuletypeID); err != nil {
				return err
			}
		}
	}

	if o.ModuletypeIDn != nil {

		// query param moduletype_id__n
		var qrModuletypeIDn string

		if o.ModuletypeIDn != nil {
			qrModuletypeIDn = *o.ModuletypeIDn
		}
		qModuletypeIDn := qrModuletypeIDn
		if qModuletypeIDn != "" {

			if err := r.SetQueryParam("moduletype_id__n", qModuletypeIDn); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.Typen != nil {

		// query param type__n
		var qrTypen string

		if o.Typen != nil {
			qrTypen = *o.Typen
		}
		qTypen := qrTypen
		if qTypen != "" {

			if err := r.SetQueryParam("type__n", qTypen); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
