// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewAllBooksParams creates a new AllBooksParams object
// with the default values initialized.
func NewAllBooksParams() *AllBooksParams {

	return &AllBooksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAllBooksParamsWithTimeout creates a new AllBooksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAllBooksParamsWithTimeout(timeout time.Duration) *AllBooksParams {

	return &AllBooksParams{

		timeout: timeout,
	}
}

// NewAllBooksParamsWithContext creates a new AllBooksParams object
// with the default values initialized, and the ability to set a context for a request
func NewAllBooksParamsWithContext(ctx context.Context) *AllBooksParams {

	return &AllBooksParams{

		Context: ctx,
	}
}

// NewAllBooksParamsWithHTTPClient creates a new AllBooksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAllBooksParamsWithHTTPClient(client *http.Client) *AllBooksParams {

	return &AllBooksParams{
		HTTPClient: client,
	}
}

/*AllBooksParams contains all the parameters to send to the API endpoint
for the all books operation typically these are written to a http.Request
*/
type AllBooksParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the all books params
func (o *AllBooksParams) WithTimeout(timeout time.Duration) *AllBooksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the all books params
func (o *AllBooksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the all books params
func (o *AllBooksParams) WithContext(ctx context.Context) *AllBooksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the all books params
func (o *AllBooksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the all books params
func (o *AllBooksParams) WithHTTPClient(client *http.Client) *AllBooksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the all books params
func (o *AllBooksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AllBooksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}