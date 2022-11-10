package bookstore

import (
	thriftTypes "cgo-thrift/gen-src/gen-go/types"
)

type BookStoreGolang struct {
	books  []thriftTypes.Book
	orders thriftTypes.Orders
}

func NewBookStoreGolang() *BookStoreGolang {
	res := new(BookStoreGolang)
	res.books = make([]thriftTypes.Book, 0)
	res.orders = thriftTypes.Orders{}
	res.orders.Entry = make(map[string][]string)
	return res
}

func (bs *BookStoreGolang) HasBook(book thriftTypes.Book) bool {
	for _, value := range bs.books {
		if value.Equals(&book) {
			return true
		}
	}
	return false
}

func (bs *BookStoreGolang) AddBook(book thriftTypes.Book) {
	bs.books = append(bs.books, book)
}

func (bs *BookStoreGolang) GetOrders() thriftTypes.Orders {
	return bs.orders
}

func (bs *BookStoreGolang) AddOrder(order thriftTypes.Order) {
	orders := bs.orders.GetEntry()
	books, exists := orders[order.CustomerName]
	if exists {
		books = append(books, order.BookName)
		return
	} else {
		orders[order.CustomerName] = []string{order.BookName}
	}
}
