// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+nonstdout@users.noreply.github.com>
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

package ipam

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

	"github.com/nonstdout/go-netbox/v3/netbox/models"
)

// NewIpamServicesUpdateParams creates a new IpamServicesUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServicesUpdateParams() *IpamServicesUpdateParams {
	return &IpamServicesUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesUpdateParamsWithTimeout creates a new IpamServicesUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamServicesUpdateParamsWithTimeout(timeout time.Duration) *IpamServicesUpdateParams {
	return &IpamServicesUpdateParams{
		timeout: timeout,
	}
}

// NewIpamServicesUpdateParamsWithContext creates a new IpamServicesUpdateParams object
// with the ability to set a context for a request.
func NewIpamServicesUpdateParamsWithContext(ctx context.Context) *IpamServicesUpdateParams {
	return &IpamServicesUpdateParams{
		Context: ctx,
	}
}

// NewIpamServicesUpdateParamsWithHTTPClient creates a new IpamServicesUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServicesUpdateParamsWithHTTPClient(client *http.Client) *IpamServicesUpdateParams {
	return &IpamServicesUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamServicesUpdateParams contains all the parameters to send to the API endpoint

	for the ipam services update operation.

	Typically these are written to a http.Request.
*/
type IpamServicesUpdateParams struct {

	// Data.
	Data *models.WritableService

	/* ID.

	   A unique integer value identifying this service.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam services update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesUpdateParams) WithDefaults() *IpamServicesUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam services update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam services update params
func (o *IpamServicesUpdateParams) WithTimeout(timeout time.Duration) *IpamServicesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services update params
func (o *IpamServicesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services update params
func (o *IpamServicesUpdateParams) WithContext(ctx context.Context) *IpamServicesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services update params
func (o *IpamServicesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services update params
func (o *IpamServicesUpdateParams) WithHTTPClient(client *http.Client) *IpamServicesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services update params
func (o *IpamServicesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam services update params
func (o *IpamServicesUpdateParams) WithData(data *models.WritableService) *IpamServicesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam services update params
func (o *IpamServicesUpdateParams) SetData(data *models.WritableService) {
	o.Data = data
}

// WithID adds the id to the ipam services update params
func (o *IpamServicesUpdateParams) WithID(id int64) *IpamServicesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam services update params
func (o *IpamServicesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}