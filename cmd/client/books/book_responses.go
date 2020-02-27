// Code generated by go-swagger; DO NOT EDIT.

package books

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/janPhil/bookstore/cmd/models"
)

// BookReader is a Reader for the Book structure.
type BookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewBookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBookOK creates a BookOK with default headers values
func NewBookOK() *BookOK {
	return &BookOK{}
}

/*BookOK handles this case with default header values.

Data structure representing a single book
*/
type BookOK struct {
	Payload *models.Book
}

func (o *BookOK) Error() string {
	return fmt.Sprintf("[GET /books/{id}][%d] bookOK  %+v", 200, o.Payload)
}

func (o *BookOK) GetPayload() *models.Book {
	return o.Payload
}

func (o *BookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Book)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookNotFound creates a BookNotFound with default headers values
func NewBookNotFound() *BookNotFound {
	return &BookNotFound{}
}

/*BookNotFound handles this case with default header values.

Generic error message returned as a string
*/
type BookNotFound struct {
	Payload *models.GenericError
}

func (o *BookNotFound) Error() string {
	return fmt.Sprintf("[GET /books/{id}][%d] bookNotFound  %+v", 404, o.Payload)
}

func (o *BookNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *BookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}