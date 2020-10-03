// Code generated by goa v3.2.4, DO NOT EDIT.
//
// book service
//
// Command:
// $ goa gen book/design

package book

import (
	bookviews "book/gen/book/views"
	"context"

	goa "goa.design/goa/v3/pkg"
)

// The book service gives details of the book.
type Service interface {
	// Adds a new book to the book store.
	Create(context.Context, *Book) (res *Book, err error)
	// List all entries
	List(context.Context) (res []*Book, err error)
	// Updating the existing book
	Update(context.Context, *Book) (err error)
	// Remove book from storage
	Remove(context.Context, *RemovePayload) (err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "book"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [4]string{"create", "list", "update", "remove"}

// Book is the payload type of the book service create method.
type Book struct {
	// ID of the book
	ID uint32
	// Name of book
	Name string
	// Description of the book
	Description string
	// Price of the book
	Price uint32
}

// RemovePayload is the payload type of the book service remove method.
type RemovePayload struct {
	// ID of book to remove
	ID uint32
}

// MakeNotFound builds a goa.ServiceError from an error.
func MakeNotFound(err error) *goa.ServiceError {
	return &goa.ServiceError{
		Name:    "not-found",
		ID:      goa.NewErrorID(),
		Message: err.Error(),
	}
}

// NewBook initializes result type Book from viewed result type Book.
func NewBook(vres *bookviews.Book) *Book {
	return newBook(vres.Projected)
}

// NewViewedBook initializes viewed result type Book from result type Book
// using the given view.
func NewViewedBook(res *Book, view string) *bookviews.Book {
	p := newBookView(res)
	return &bookviews.Book{Projected: p, View: "default"}
}

// newBook converts projected type Book to service type Book.
func newBook(vres *bookviews.BookView) *Book {
	res := &Book{}
	if vres.ID != nil {
		res.ID = *vres.ID
	}
	if vres.Name != nil {
		res.Name = *vres.Name
	}
	if vres.Description != nil {
		res.Description = *vres.Description
	}
	if vres.Price != nil {
		res.Price = *vres.Price
	}
	return res
}

// newBookView projects result type Book to projected type BookView using the
// "default" view.
func newBookView(res *Book) *bookviews.BookView {
	vres := &bookviews.BookView{
		ID:          &res.ID,
		Name:        &res.Name,
		Description: &res.Description,
		Price:       &res.Price,
	}
	return vres
}
