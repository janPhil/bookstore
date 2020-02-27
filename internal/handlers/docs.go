// Package classification of Bookstore
//
// Documentation for Bookstore
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"github.com/janPhil/bookstore/internal/values"
)

// Data structure representing a single book
// swagger:response bookResponse
type bookResponseWrapper struct {
	// Newly created book
	// in: body
	Body values.Book
}

// List of books
// swagger:response booksResponse
type booksResponseWrapper struct {
	// All current books
	// in: body
	Body []values.Book
}

// Generic error message returned as a string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct {
}

// swagger:parameters DeleteBook Book
type bookIDParamsWrapper struct {
	// The id of the book for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
