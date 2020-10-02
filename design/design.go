package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("book", func() {
	Title("Book Store")
	Description("Service to perform CRUD operations using goa")
	Server("book", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("book", func() {
	Description("The book service gives details of the book.")

	//Method to add a new book
	Method("create", func() {
		Description("Adds a new book to the book store.")
		Payload(Book)
		Result(Book)
		HTTP(func() {
			POST("/")
			Response(StatusCreated)
		})
	})

	//Method to get all existing books
	Method("list", func() {
		Description("List all entries")
		Result(ArrayOf(Book))
		HTTP(func() {
			GET("/books")
			Response(StatusOK)
		})
	})
})