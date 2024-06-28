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

	"github.com/nonstdout/go-netbox/v3/netbox/models"
)

// NewIpamL2vpnTerminationsBulkUpdateParams creates a new IpamL2vpnTerminationsBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamL2vpnTerminationsBulkUpdateParams() *IpamL2vpnTerminationsBulkUpdateParams {
	return &IpamL2vpnTerminationsBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamL2vpnTerminationsBulkUpdateParamsWithTimeout creates a new IpamL2vpnTerminationsBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamL2vpnTerminationsBulkUpdateParamsWithTimeout(timeout time.Duration) *IpamL2vpnTerminationsBulkUpdateParams {
	return &IpamL2vpnTerminationsBulkUpdateParams{
		timeout: timeout,
	}
}

// NewIpamL2vpnTerminationsBulkUpdateParamsWithContext creates a new IpamL2vpnTerminationsBulkUpdateParams object
// with the ability to set a context for a request.
func NewIpamL2vpnTerminationsBulkUpdateParamsWithContext(ctx context.Context) *IpamL2vpnTerminationsBulkUpdateParams {
	return &IpamL2vpnTerminationsBulkUpdateParams{
		Context: ctx,
	}
}

// NewIpamL2vpnTerminationsBulkUpdateParamsWithHTTPClient creates a new IpamL2vpnTerminationsBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamL2vpnTerminationsBulkUpdateParamsWithHTTPClient(client *http.Client) *IpamL2vpnTerminationsBulkUpdateParams {
	return &IpamL2vpnTerminationsBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamL2vpnTerminationsBulkUpdateParams contains all the parameters to send to the API endpoint

	for the ipam l2vpn terminations bulk update operation.

	Typically these are written to a http.Request.
*/
type IpamL2vpnTerminationsBulkUpdateParams struct {

	// Data.
	Data *models.WritableL2VPNTermination

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam l2vpn terminations bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnTerminationsBulkUpdateParams) WithDefaults() *IpamL2vpnTerminationsBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam l2vpn terminations bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnTerminationsBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam l2vpn terminations bulk update params
func (o *IpamL2vpnTerminationsBulkUpdateParams) WithTimeout(timeout time.Duration) *IpamL2vpnTerminationsBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam l2vpn terminations bulk update params
func (o *IpamL2vpnTerminationsBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam l2vpn terminations bulk update params
func (o *IpamL2vpnTerminationsBulkUpdateParams) WithContext(ctx context.Context) *IpamL2vpnTerminationsBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam l2vpn terminations bulk update params
func (o *IpamL2vpnTerminationsBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam l2vpn terminations bulk update params
func (o *IpamL2vpnTerminationsBulkUpdateParams) WithHTTPClient(client *http.Client) *IpamL2vpnTerminationsBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam l2vpn terminations bulk update params
func (o *IpamL2vpnTerminationsBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam l2vpn terminations bulk update params
func (o *IpamL2vpnTerminationsBulkUpdateParams) WithData(data *models.WritableL2VPNTermination) *IpamL2vpnTerminationsBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam l2vpn terminations bulk update params
func (o *IpamL2vpnTerminationsBulkUpdateParams) SetData(data *models.WritableL2VPNTermination) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamL2vpnTerminationsBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}