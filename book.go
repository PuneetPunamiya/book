package bookapi

import (
	book "book/gen/book"
	"context"
	"log"
)

// book service example implementation.
// The example methods log the requests and return zero values.
type booksrvc struct {
	logger *log.Logger
}

// NewBook returns the book service implementation.
func NewBook(logger *log.Logger) book.Service {
	return &booksrvc{logger}
}

var bookStore = make([]*book.Book, 0)

// Adds a new book to the book store.
func (s *booksrvc) Create(ctx context.Context, p *book.Book) (res *book.Book, err error) {
	res = &book.Book{ID: p.ID, Name: p.Name, Description: p.Description, Price: p.Price}
	bookStore = append(bookStore, res)
	s.logger.Print("book.create")
	return res, nil
}

// List all entries
func (s *booksrvc) List(ctx context.Context) ([]*book.Book, error) {
	s.logger.Print("book.list")
	return bookStore, nil
}
