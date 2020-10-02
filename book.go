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

// Adds a new book to the book store.
func (s *booksrvc) Create(ctx context.Context, p *book.Book) (res *book.Book, err error) {
	res = &book.Book{}
	s.logger.Print("book.create")
	return
}

// List all entries
func (s *booksrvc) List(ctx context.Context) (res []*book.Book, err error) {
	s.logger.Print("book.list")
	return
}
