package handlers

import (
	"log"
	"net/http"

	"github.com/janPhil/bookstore/internal/values"
)

// swagger:route GET /books books AllBooks
// Return a list of books from the database
//
// responses:
// 200: booksResponse

// ListAll handles GET requests and returns all current books
func (h *BooksHandler) ListAll(wr http.ResponseWriter, r *http.Request) {
	h.logger.Println("[DEBUG] get all book")
	wr.Header().Add("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		books, err := h.storage.AllBooks()
		if err != nil {
			http.Error(wr, http.StatusText(500), 500)
		}
		err = books.ToJSON(wr)
		if err != nil {
			h.logger.Println("[ERROR] serializing Books")
		}
	}
}

// swagger:route GET /books/{id} books Book
// Returns a single book from the database
//
// responses:
// 200: bookResponse
// 404: errorResponse

// ListOne handles GET requests and returns a specific book
func (h *BooksHandler) ListOne(wr http.ResponseWriter, r *http.Request) {
	isbn := getBookId(r)
	h.logger.Println("[DEBUG] get record id", isbn)
	wr.Header().Add("Content-Type", "application/json")
	book, err := h.storage.Book(isbn)
	h.logger.Println("[DEBUG] book: ", book)
	if err != nil {
		http.Error(wr, http.StatusText(404), 404)
	}

	err = book.ToJSON(wr)
	if err != nil {
		h.logger.Println("[ERROR] serializing Books")
	}

}

// swagger:route POST /book book AddBook
// Create a new book
//
// responses:
//	200: bookResponse
//  501: errorResponse

func (h *BooksHandler) Create(wr http.ResponseWriter, r *http.Request) {
	h.logger.Println("[DEBUG] create book")
	wr.Header().Add("Content-Type", "application/json")
	book := &values.Book{}

	err := book.FromJSON(r.Body)
	if err != nil {
		h.logger.Println("[ERROR] deserializing Books")
	}

	err = h.storage.AddBook(book)
	if err != nil {
		http.Error(wr, http.StatusText(501), 501)
	}
}

// swagger:route DELETE /books/{id} books DeleteBook
// Delete a book by a given ID
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

func (h *BooksHandler) Delete(wr http.ResponseWriter, r *http.Request) {
	h.logger.Println("[DEBUG] create book")
	wr.Header().Add("Content-Type", "application/json")
	isbn := getBookId(r)
	err := h.storage.DeleteBook(isbn)
	if err != nil {
		http.Error(wr, http.StatusText(501), 501)
	}

}

type storageInterface interface {
	AllBooks() (values.Books, error)
	Book(id string) (values.Book, error)
	AddBook(b *values.Book) error
	DeleteBook(id string) error
}

type BooksHandler struct {
	logger  *log.Logger
	storage storageInterface
}

func NewBooksHandler(storage storageInterface, logger *log.Logger) *BooksHandler {
	return &BooksHandler{
		storage: storage,
		logger:  logger,
	}
}
