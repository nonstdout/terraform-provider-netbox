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

package wireless

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

// NewWirelessWirelessLanGroupsDeleteParams creates a new WirelessWirelessLanGroupsDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWirelessWirelessLanGroupsDeleteParams() *WirelessWirelessLanGroupsDeleteParams {
	return &WirelessWirelessLanGroupsDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWirelessWirelessLanGroupsDeleteParamsWithTimeout creates a new WirelessWirelessLanGroupsDeleteParams object
// with the ability to set a timeout on a request.
func NewWirelessWirelessLanGroupsDeleteParamsWithTimeout(timeout time.Duration) *WirelessWirelessLanGroupsDeleteParams {
	return &WirelessWirelessLanGroupsDeleteParams{
		timeout: timeout,
	}
}

// NewWirelessWirelessLanGroupsDeleteParamsWithContext creates a new WirelessWirelessLanGroupsDeleteParams object
// with the ability to set a context for a request.
func NewWirelessWirelessLanGroupsDeleteParamsWithContext(ctx context.Context) *WirelessWirelessLanGroupsDeleteParams {
	return &WirelessWirelessLanGroupsDeleteParams{
		Context: ctx,
	}
}

// NewWirelessWirelessLanGroupsDeleteParamsWithHTTPClient creates a new WirelessWirelessLanGroupsDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewWirelessWirelessLanGroupsDeleteParamsWithHTTPClient(client *http.Client) *WirelessWirelessLanGroupsDeleteParams {
	return &WirelessWirelessLanGroupsDeleteParams{
		HTTPClient: client,
	}
}

/*
WirelessWirelessLanGroupsDeleteParams contains all the parameters to send to the API endpoint

	for the wireless wireless lan groups delete operation.

	Typically these are written to a http.Request.
*/
type WirelessWirelessLanGroupsDeleteParams struct {

	/* ID.

	   A unique integer value identifying this Wireless LAN Group.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the wireless wireless lan groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLanGroupsDeleteParams) WithDefaults() *WirelessWirelessLanGroupsDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the wireless wireless lan groups delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WirelessWirelessLanGroupsDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the wireless wireless lan groups delete params
func (o *WirelessWirelessLanGroupsDeleteParams) WithTimeout(timeout time.Duration) *WirelessWirelessLanGroupsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the wireless wireless lan groups delete params
func (o *WirelessWirelessLanGroupsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the wireless wireless lan groups delete params
func (o *WirelessWirelessLanGroupsDeleteParams) WithContext(ctx context.Context) *WirelessWirelessLanGroupsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the wireless wireless lan groups delete params
func (o *WirelessWirelessLanGroupsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the wireless wireless lan groups delete params
func (o *WirelessWirelessLanGroupsDeleteParams) WithHTTPClient(client *http.Client) *WirelessWirelessLanGroupsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the wireless wireless lan groups delete params
func (o *WirelessWirelessLanGroupsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the wireless wireless lan groups delete params
func (o *WirelessWirelessLanGroupsDeleteParams) WithID(id int64) *WirelessWirelessLanGroupsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the wireless wireless lan groups delete params
func (o *WirelessWirelessLanGroupsDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *WirelessWirelessLanGroupsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}