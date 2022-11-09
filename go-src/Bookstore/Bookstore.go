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

func HasBook(bookStoreCPtr unsafe.Pointer, book thriftTypes.Book) bool {
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)

	printAndPanicError(book.Write(context.Background(), protocol))

	ptr := unsafe.Pointer(&mem_buffer.Bytes()[0])
	size := mem_buffer.Len()

	return bool(C.hasBook(bookStoreCPtr, ptr, C.uint(size)))
}

func AddBook(bookStoreCPtr unsafe.Pointer, book thriftTypes.Book) {
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)
	printAndPanicError(book.Write(defaultCtx, protocol))
	ptr := unsafe.Pointer(&mem_buffer.Bytes()[0])
	size := mem_buffer.Len()

	C.addBook(bookStoreCPtr, ptr, C.uint(size))
}

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

func GetOrders(bookStoreCPtr unsafe.Pointer) (orders thriftTypes.Orders) {
	cbinary := C.getOrders(bookStoreCPtr)

	gobinary := TranslateCBinary2GoBinary(cbinary)
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	mem_buffer.Buffer = gobinary
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)
	orders.Read(defaultCtx, protocol)
	return orders
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
