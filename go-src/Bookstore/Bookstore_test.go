package bookstore

import (
	thriftTypes "cgo-thrift/gen-src/gen-go/types"
	"strings"
	"testing"
)

func BenchmarkGetOrders2B(b *testing.B) {
	bookStore := BookStore{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 1)
	order.BookName = strings.Repeat("H", 1)

	bookStore.AddOrder(order)

	for i := 0; i < b.N; i++ {
		bookStore.GetOrders()
	}
}

func BenchmarkGetOrders200B(b *testing.B) {
	bookStore := BookStore{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 100)
	order.BookName = strings.Repeat("H", 100)

	bookStore.AddOrder(order)

	for i := 0; i < b.N; i++ {
		bookStore.GetOrders()
	}
}

func BenchmarkGetOrders400B(b *testing.B) {
	bookStore := BookStore{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 200)
	order.BookName = strings.Repeat("H", 200)

	bookStore.AddOrder(order)

	for i := 0; i < b.N; i++ {
		bookStore.GetOrders()
	}
}

func BenchmarkGetOrders1KB(b *testing.B) {
	bookStore := BookStore{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 500)
	order.BookName = strings.Repeat("H", 500)

	bookStore.AddOrder(order)

	for i := 0; i < b.N; i++ {
		bookStore.GetOrders()
	}
}

func BenchmarkGetOrders10KB(b *testing.B) {
	bookStore := BookStore{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 5000)
	order.BookName = strings.Repeat("H", 5000)

	bookStore.AddOrder(order)

	for i := 0; i < b.N; i++ {
		bookStore.GetOrders()
	}
}

func BenchmarkGetOrders100KB(b *testing.B) {
	bookStore := BookStore{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 50000)
	order.BookName = strings.Repeat("H", 50000)

	bookStore.AddOrder(order)

	for i := 0; i < b.N; i++ {
		bookStore.GetOrders()
	}
}
