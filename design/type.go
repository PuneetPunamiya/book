package design

import . "goa.design/goa/v3/dsl"

var Book = ResultType("application/vnd.book", "Book", func() {
	Description("Details of a book")

	Attribute("id", UInt32, "ID of the book", func() {
		Example("id", 1)
	})
	Attribute("name", String, "Name of book", func() {
		Example("name", "book1")
		MaxLength(100)
	})
	Attribute("description", String, "Description of the book", func() {
		Example("name", "Books are human's best friend")
		MaxLength(100)
	})
	Attribute("price", UInt32, "Price of the book", func() {
		Example("price", 100)
	})

	Required("id", "name", "description", "price")
})
