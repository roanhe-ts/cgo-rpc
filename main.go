package main

import (
	thriftTypes "cgo-thrift/gen-src/gen-go/types"
	bookstore "cgo-thrift/go-src/Bookstore"
	"fmt"
	"strings"
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

	bs.AddBook(book)

	if bs.HasBook(book) {
		println("Add succed")
	}

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 100)
	order.BookName = strings.Repeat("H", 100)

	bs.AddOrder(order)

	orders := bs.GetOrders()
	fmt.Println(orders.String())

	return
}
