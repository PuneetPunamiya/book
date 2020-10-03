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

	Error("not-found", ErrorResult, "Book Not Found Error")

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

	//Method to update a specific book
	Method("update", func() {
		Description("Updating the existing book")

		Payload(Book)

		HTTP(func() {
			PATCH("/book/{id}")
			Response(StatusOK)
		})
	})

	//Method to remove a particular book
	Method("remove", func() {
		Description("Remove book from storage")
		Payload(func() {
			Attribute("id", UInt32, "ID of book to remove")
			Required("id")
		})

		HTTP(func() {
			DELETE("/book/{id}")

			Response(StatusOK)
			Response("not-found", StatusNotFound)
		})
	})
})
