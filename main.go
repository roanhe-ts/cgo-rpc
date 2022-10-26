package main

import (
	bookStore "cgo-thrift/go_src/BookStore"
)

func main() {
	bs := bookStore.BookStore{}
	bs.BookStoreCPtr = bookStore.InitBookStore()

	book := bookStore.Book{
		Name:  "Book1",
		Price: 112,
	}

	bookStore.AddBook(bs.BookStoreCPtr, book)

	if bookStore.HasBook(bs.BookStoreCPtr, book) {
		println("Add succed")
	}

	return
}
