package main

import (
	thriftTypes "cgo-thrift/gen_src/gen-go/types"
	bookStore "cgo-thrift/go_src/BookStore"
)

func main() {
	bs := bookStore.BookStore{}
	bs.BookStoreCPtr = bookStore.InitBookStore()

	book := thriftTypes.Book{
		Author: &thriftTypes.Author{
			Name: "XYZ",
			Age:  45,
		},
		Name:  "Book1",
		Price: 112,
	}

	bookStore.AddBook(bs.BookStoreCPtr, book)

	if bookStore.HasBook(bs.BookStoreCPtr, book) {
		println("Add succed")
	}

	return
}
