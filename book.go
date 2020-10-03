package bookapi

import (
	book "book/gen/book"
	"context"
	"fmt"
	"log"
)

// book service example implementation.
// The example methods log the requests and return zero values.
type booksrvc struct {
	logger *log.Logger
}

// Errors
var (
	notFoundError = book.MakeNotFound(fmt.Errorf("book not found"))
)

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

// Updating the existing book
func (s *booksrvc) Update(ctx context.Context, p *book.Book) (err error) {
	s.logger.Print("book.update")

	for i, book := range bookStore {
		if book.ID == p.ID {
			book.Name = p.Name
			book.Description = p.Description
			book.Price = p.Price
			bookStore = append(bookStore[:i], book)
		}
	}
	return
}

// Remove book from storage
func (s *booksrvc) Remove(ctx context.Context, p *book.RemovePayload) (err error) {
	s.logger.Print("book.remove")

	for i, book := range bookStore {
		if book.ID == p.ID {
			bookStore = append(bookStore[:i], bookStore[i+1:]...)
			s.logger.Printf("The event with ID %d has been deleted successfully", book.ID)
		} else {
			return notFoundError
		}
	}
	return
}
