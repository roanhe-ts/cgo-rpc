package bookstore

import (
	thriftTypes "cgo-thrift/gen-src/gen-go/types"
	"strings"
	"testing"
)

func BenchmarkCGoGetOrders2B(b *testing.B) {
	bookStore := BookStoreCgo{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 1)
	order.BookName = strings.Repeat("H", 1)

	bookStore.AddOrder(order)

	for i := 0; i < b.N; i++ {
		bookStore.GetOrders()
	}
}

func BenchmarkGolangGetOrders2B(b *testing.B) {
	bookStore := NewBookStoreGolang()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 1)
	order.BookName = strings.Repeat("H", 1)

	bookStore.AddOrder(order)

	for i := 0; i < b.N; i++ {
		bookStore.GetOrders()
	}
}

func BenchmarkCGoGetOrders200B(b *testing.B) {
	bookStore := BookStoreCgo{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 100)
	order.BookName = strings.Repeat("H", 100)

	bookStore.AddOrder(order)

	for i := 0; i < b.N; i++ {
		bookStore.GetOrders()
	}
}

func BenchmarkGolangGetOrders200B(b *testing.B) {
	bookStore := NewBookStoreGolang()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 100)
	order.BookName = strings.Repeat("H", 100)

	bookStore.AddOrder(order)

	for i := 0; i < b.N; i++ {
		bookStore.GetOrders()
	}
}

func BenchmarkCGoGetOrders1KB(b *testing.B) {
	bookStore := BookStoreCgo{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 500)
	order.BookName = strings.Repeat("H", 500)

	for i := 0; i < b.N; i++ {
		bookStore.AddOrder(order)
		bookStore.GetOrders()
	}
}

func BenchmarkGolangGetOrders1KB(b *testing.B) {
	bookStore := NewBookStoreGolang()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 500)
	order.BookName = strings.Repeat("H", 500)

	for i := 0; i < b.N; i++ {
		bookStore.AddOrder(order)
		bookStore.GetOrders()
	}
}

func BenchmarkCGoGetOrders10KB(b *testing.B) {
	bookStore := BookStoreCgo{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 5000)
	order.BookName = strings.Repeat("H", 5000)

	for i := 0; i < b.N; i++ {
		bookStore.AddOrder(order)
		bookStore.GetOrders()
	}
}

func BenchmarkGolangGetOrders10KB(b *testing.B) {
	bookStore := NewBookStoreGolang()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 5000)
	order.BookName = strings.Repeat("H", 5000)

	for i := 0; i < b.N; i++ {
		bookStore.AddOrder(order)
		bookStore.GetOrders()
	}
}

func BenchmarkCGoGetOrders100KB(b *testing.B) {
	bookStore := BookStoreCgo{}
	bookStore.BookStoreCPtr = InitBookStore()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 50000)
	order.BookName = strings.Repeat("H", 50000)

	for i := 0; i < b.N; i++ {
		bookStore.AddOrder(order)
		bookStore.GetOrders()
	}
}

func BenchmarkGolangGetOrders100KB(b *testing.B) {
	bookStore := NewBookStoreGolang()

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 50000)
	order.BookName = strings.Repeat("H", 50000)

	for i := 0; i < b.N; i++ {
		bookStore.AddOrder(order)
		bookStore.GetOrders()
	}
}
