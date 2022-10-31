package main

import (
	bookstore "bookstore"
	thriftTypes "cgo-thrift/gen_src/gen-go/types"
)

func main() {
	bs := bookstore.BookStore{}
	bs.BookStoreCPtr = bookstore.InitBookStore()

	book := thriftTypes.Book{
		Author: &thriftTypes.Author{
			Name: "XYZ",
			Age:  45,
		},
		Name:  "Book1",
		Price: 112,
	}

	bookstore.AddBook(bs.BookStoreCPtr, book)

	if bookstore.HasBook(bs.BookStoreCPtr, book) {
		println("Add succed")
	}

	return
}
