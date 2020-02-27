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

// DeleteBookReader is a Reader for the DeleteBook structure.
type DeleteBookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDeleteBookCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteBookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewDeleteBookNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteBookCreated creates a DeleteBookCreated with default headers values
func NewDeleteBookCreated() *DeleteBookCreated {
	return &DeleteBookCreated{}
}

/*DeleteBookCreated handles this case with default header values.

No content is returned by this API endpoint
*/
type DeleteBookCreated struct {
}

func (o *DeleteBookCreated) Error() string {
	return fmt.Sprintf("[DELETE /books/{id}][%d] deleteBookCreated ", 201)
}

func (o *DeleteBookCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBookNotFound creates a DeleteBookNotFound with default headers values
func NewDeleteBookNotFound() *DeleteBookNotFound {
	return &DeleteBookNotFound{}
}

/*DeleteBookNotFound handles this case with default header values.

Generic error message returned as a string
*/
type DeleteBookNotFound struct {
	Payload *models.GenericError
}

func (o *DeleteBookNotFound) Error() string {
	return fmt.Sprintf("[DELETE /books/{id}][%d] deleteBookNotFound  %+v", 404, o.Payload)
}

func (o *DeleteBookNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteBookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBookNotImplemented creates a DeleteBookNotImplemented with default headers values
func NewDeleteBookNotImplemented() *DeleteBookNotImplemented {
	return &DeleteBookNotImplemented{}
}

/*DeleteBookNotImplemented handles this case with default header values.

Generic error message returned as a string
*/
type DeleteBookNotImplemented struct {
	Payload *models.GenericError
}

func (o *DeleteBookNotImplemented) Error() string {
	return fmt.Sprintf("[DELETE /books/{id}][%d] deleteBookNotImplemented  %+v", 501, o.Payload)
}

func (o *DeleteBookNotImplemented) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteBookNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}