package bookstore

/*
#cgo CFLAGS: -I${SRCDIR}/../../include
#cgo LDFLAGS: -L${SRCDIR}/../../build -lbookstore_c -lbookstore -lthrift -lstdc++
#include "BookstoreWrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"bytes"
	thriftTypes "cgo-thrift/gen-src/gen-go/types"
	"fmt"
	"unsafe"

	"github.com/apache/thrift/lib/go/thrift"
)

type BookStoreCgo struct {
	BookStoreCPtr unsafe.Pointer
}

type Author struct {
	Name string
	Age  int32
}

type Book struct {
	Name   string
	Price  int32
	Author Author
}

func (bs *BookStoreCgo) HasBook(book Book) bool {

	return bool(
		C.hasBook(
			bs.BookStoreCPtr,
			C.CBook{
				name:  C.CString(book.Name),
				price: C.int32_t(book.Price),
				author: C.CAuthor{
					name: C.CString(book.Author.Name),
					age:  C.int32_t(book.Author.Age),
				},
			}),
	)
}

func (bs *BookStoreCgo) AddBook(book Book) {
	C.addBook(
		bs.BookStoreCPtr,
		C.CBook{
			name:  C.CString(book.Name),
			price: C.int32_t(book.Price),
			author: C.CAuthor{
				name: C.CString(book.Author.Name),
				age:  C.int32_t(book.Author.Age),
			},
		})
}

// Copy C binary and translate it to a Go binary
func TranslateCBinary2GoBinary(c *C.struct_Binary) *bytes.Buffer {
	if c == nil {
		return nil
	}

	var res bytes.Buffer

	for i := 0; i < int(c.size); i++ {
		begin := uintptr(c.buffer)
		curr := (*byte)(unsafe.Pointer(begin + uintptr(i)))
		res.WriteByte(*(curr))
	}

	return &res
}

func (bs *BookStoreCgo) GetOrders() (orders thriftTypes.Orders) {
	cbinary := C.getOrders(bs.BookStoreCPtr)
	gobinary := TranslateCBinary2GoBinary(cbinary)
	// Must free c binary after copy.
	C.free(unsafe.Pointer(cbinary))
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	mem_buffer.Buffer = gobinary
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)
	orders.Read(protocol)
	return orders
}

func (bs *BookStoreCgo) AddOrder(order thriftTypes.Order) {
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)

	printAndPanicError(order.Write(protocol))
	ptr := unsafe.Pointer(&mem_buffer.Bytes()[0])
	size := mem_buffer.Len()

	C.addOrder(bs.BookStoreCPtr, ptr, C.uint(size))
}

func InitBookStore() unsafe.Pointer {
	return C.initBookStore()
}

func printAndPanicError(err error) {
	if err != nil {
		fmt.Println("Error ", err)
		panic(err)
	}
}
