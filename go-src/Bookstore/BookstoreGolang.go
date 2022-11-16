package bookstore

import (
	thriftTypes "cgo-thrift/gen-src/gen-go/types"
)

type BookStoreGolang struct {
	books  []Book
	orders thriftTypes.Orders
}

func NewBookStoreGolang() *BookStoreGolang {
	res := new(BookStoreGolang)
	res.books = make([]Book, 0)
	res.orders = thriftTypes.Orders{}
	res.orders.Entry = make(map[string][]string)
	return res
}

func (bs *BookStoreGolang) HasBook(book Book) bool {
	for _, value := range bs.books {
		deep_equal := value.Author.Age == book.Author.Age && value.Author.Name == book.Author.Name &&
			value.Name == book.Name && value.Price == book.Price
		if deep_equal {
			return true
		}
	}
	return false
}

func (bs *BookStoreGolang) AddBook(book Book) {
	bs.books = append(bs.books, book)
}

func (bs *BookStoreGolang) GetOrders() thriftTypes.Orders {
	return bs.orders
}

func (bs *BookStoreGolang) AddOrder(order thriftTypes.Order) {
	orders := bs.orders.GetEntry()
	books, exists := orders[order.CustomerName]
	if exists {
		_ = append(books, order.BookName)
		return
	} else {
		orders[order.CustomerName] = []string{order.BookName}
	}
}
