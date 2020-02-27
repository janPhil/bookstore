package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/janPhil/bookstore/internal/values"
	_ "github.com/lib/pq"
)

const (
	hostname     = "localhost"
	hostPort     = 5432
	username     = "postgres"
	password     = "password"
	databaseName = "bookstore"
)

var pgConString = fmt.Sprintf("port=%d host=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	hostPort, hostname, username, password, databaseName)

type BookStorage struct {
	db *sql.DB
}

func NewBookStorage() (*BookStorage, error) {
	db, err := newDatabaseConnection()
	if err != nil {
		log.Fatal("Failed to establish connection to database")
	}
	return &BookStorage{
		db: db,
	}, nil
}

func (s *BookStorage) AllBooks() (values.Books, error) {
	rows, err := s.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := values.Books{}
	for rows.Next() {
		bk := new(values.Book)
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}
		books = append(books, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (s *BookStorage) Book(id string) (values.Book, error) {
	statement := "SELECT * FROM books WHERE isbn=$1;"
	b := values.Book{}
	row := s.db.QueryRow(statement, id)
	switch err := row.Scan(&b.Isbn, &b.Title, &b.Author, &b.Price); err {
	case sql.ErrNoRows:
		return b, err
	default:
		fmt.Println(err)
	}
	return b, nil
}

func (s *BookStorage) AddBook(b *values.Book) error {
	statement := "INSERT INTO books (isbn, title, author, price) VALUES($1, $2, $3, $4)"
	_, err := s.db.Exec(statement, b.Isbn, b.Title, b.Author, b.Price)
	if err != nil {
		return err
	}
	return nil
}

func (s *BookStorage) DeleteBook(id string) error {
	statement := "DELETE FROM books WHERE isbn=$1"
	_, err := s.db.Exec(statement, id)
	if err != nil {
		return err
	}
	return nil
}

func newDatabaseConnection() (*sql.DB, error) {
	var err error
	db, err := sql.Open("postgres", pgConString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
