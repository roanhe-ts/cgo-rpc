package main

import (
	thriftTypes "cgo-thrift/gen_src/gen-go/types"
	bookstore "cgo-thrift/src/bookstore"
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
