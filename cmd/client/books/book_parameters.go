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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewBookParams creates a new BookParams object
// with the default values initialized.
func NewBookParams() *BookParams {
	var ()
	return &BookParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBookParamsWithTimeout creates a new BookParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBookParamsWithTimeout(timeout time.Duration) *BookParams {
	var ()
	return &BookParams{

		timeout: timeout,
	}
}

// NewBookParamsWithContext creates a new BookParams object
// with the default values initialized, and the ability to set a context for a request
func NewBookParamsWithContext(ctx context.Context) *BookParams {
	var ()
	return &BookParams{

		Context: ctx,
	}
}

// NewBookParamsWithHTTPClient creates a new BookParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBookParamsWithHTTPClient(client *http.Client) *BookParams {
	var ()
	return &BookParams{
		HTTPClient: client,
	}
}

/*BookParams contains all the parameters to send to the API endpoint
for the book operation typically these are written to a http.Request
*/
type BookParams struct {

	/*ID
	  The id of the book for which the operation relates

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the book params
func (o *BookParams) WithTimeout(timeout time.Duration) *BookParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the book params
func (o *BookParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the book params
func (o *BookParams) WithContext(ctx context.Context) *BookParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the book params
func (o *BookParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the book params
func (o *BookParams) WithHTTPClient(client *http.Client) *BookParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the book params
func (o *BookParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the book params
func (o *BookParams) WithID(id int64) *BookParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the book params
func (o *BookParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *BookParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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