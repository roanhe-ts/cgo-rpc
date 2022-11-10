package bookstore

/*
#cgo CFLAGS: -I${SRCDIR}/../../include -Wvisibility
#cgo LDFLAGS: -L${SRCDIR}/../../build -lbookstore_c -lbookstore -lthrift -lstdc++
#include "BookStoreWrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"bytes"
	thriftTypes "cgo-thrift/gen-src/gen-go/types"
	"context"
	"fmt"
	"unsafe"

	"github.com/apache/thrift/lib/go/thrift"
)

type BookStore struct {
	BookStoreCPtr unsafe.Pointer
}

var defaultCtx = context.Background()

func (bs *BookStore) HasBook(book thriftTypes.Book) bool {
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)

	printAndPanicError(book.Write(context.Background(), protocol))

	ptr := unsafe.Pointer(&mem_buffer.Bytes()[0])
	size := mem_buffer.Len()

	return bool(C.hasBook(bs.BookStoreCPtr, ptr, C.uint(size)))
}

func (bs *BookStore) AddBook(book thriftTypes.Book) {
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)
	printAndPanicError(book.Write(defaultCtx, protocol))
	ptr := unsafe.Pointer(&mem_buffer.Bytes()[0])
	size := mem_buffer.Len()

	C.addBook(bs.BookStoreCPtr, ptr, C.uint(size))
}

// Copy C binary and translate it to a Go binary
func TranslateCBinary2GoBinary(c *C.struct_binary) *bytes.Buffer {
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

func (bs *BookStore) GetOrders() (orders thriftTypes.Orders) {
	cbinary := C.getOrders(bs.BookStoreCPtr)
	gobinary := TranslateCBinary2GoBinary(cbinary)
	// Must free c binary after copy.
	C.free(unsafe.Pointer(cbinary))
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	mem_buffer.Buffer = gobinary
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)
	orders.Read(defaultCtx, protocol)
	return orders
}

func (bs *BookStore) AddOrder(order thriftTypes.Order) {
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)

	printAndPanicError(order.Write(defaultCtx, protocol))
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
