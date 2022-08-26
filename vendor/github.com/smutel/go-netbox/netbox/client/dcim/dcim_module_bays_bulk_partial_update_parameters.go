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

	"github.com/smutel/go-netbox/netbox/models"
)

// NewDcimModuleBaysBulkPartialUpdateParams creates a new DcimModuleBaysBulkPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModuleBaysBulkPartialUpdateParams() *DcimModuleBaysBulkPartialUpdateParams {
	return &DcimModuleBaysBulkPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModuleBaysBulkPartialUpdateParamsWithTimeout creates a new DcimModuleBaysBulkPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimModuleBaysBulkPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimModuleBaysBulkPartialUpdateParams {
	return &DcimModuleBaysBulkPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimModuleBaysBulkPartialUpdateParamsWithContext creates a new DcimModuleBaysBulkPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimModuleBaysBulkPartialUpdateParamsWithContext(ctx context.Context) *DcimModuleBaysBulkPartialUpdateParams {
	return &DcimModuleBaysBulkPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimModuleBaysBulkPartialUpdateParamsWithHTTPClient creates a new DcimModuleBaysBulkPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModuleBaysBulkPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimModuleBaysBulkPartialUpdateParams {
	return &DcimModuleBaysBulkPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimModuleBaysBulkPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim module bays bulk partial update operation.

   Typically these are written to a http.Request.
*/
type DcimModuleBaysBulkPartialUpdateParams struct {

	// Data.
	Data *models.WritableModuleBay

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim module bays bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBaysBulkPartialUpdateParams) WithDefaults() *DcimModuleBaysBulkPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim module bays bulk partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleBaysBulkPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim module bays bulk partial update params
func (o *DcimModuleBaysBulkPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimModuleBaysBulkPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim module bays bulk partial update params
func (o *DcimModuleBaysBulkPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim module bays bulk partial update params
func (o *DcimModuleBaysBulkPartialUpdateParams) WithContext(ctx context.Context) *DcimModuleBaysBulkPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim module bays bulk partial update params
func (o *DcimModuleBaysBulkPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim module bays bulk partial update params
func (o *DcimModuleBaysBulkPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimModuleBaysBulkPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim module bays bulk partial update params
func (o *DcimModuleBaysBulkPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim module bays bulk partial update params
func (o *DcimModuleBaysBulkPartialUpdateParams) WithData(data *models.WritableModuleBay) *DcimModuleBaysBulkPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim module bays bulk partial update params
func (o *DcimModuleBaysBulkPartialUpdateParams) SetData(data *models.WritableModuleBay) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModuleBaysBulkPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
