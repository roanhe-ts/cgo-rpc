package bookstore

/*
#cgo CFLAGS: -I${SRCDIR}/../../include
#cgo LDFLAGS: -L${SRCDIR}/../../build -lbookstorecli_c -lthrift -lstdc++
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

type BookStoreCgoClient struct {
	cpointer unsafe.Pointer
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

func (bs *BookStoreCgoClient) Init() {
	bs.cpointer = C.initBSClient()
}

func (bs *BookStoreCgoClient) Free() {
	C.freeBSClient(bs.cpointer)
	bs.cpointer = nil
}

func (bs *BookStoreCgoClient) HasBook(book Book) bool {
	book_name := C.CString(book.Name)
	author_name := C.CString(book.Author.Name)
	defer C.free(unsafe.Pointer(author_name))
	defer C.free(unsafe.Pointer(book_name))

	return bool(
		C.hasBook(
			bs.cpointer,
			C.CBook{
				name:  book_name,
				price: C.int32_t(book.Price),
				author: C.CAuthor{
					name: author_name,
					age:  C.int32_t(book.Author.Age),
				},
			}),
	)
}

func (bs *BookStoreCgoClient) AddBook(book Book) {
	book_name := C.CString(book.Name)
	author_name := C.CString(book.Author.Name)
	defer C.free(unsafe.Pointer(author_name))
	defer C.free(unsafe.Pointer(book_name))

	C.addBook(
		bs.cpointer,
		C.CBook{
			name:  book_name,
			price: C.int32_t(book.Price),
			author: C.CAuthor{
				name: author_name,
				age:  C.int32_t(book.Author.Age),
			},
		})
}

// Copy C binary and translate it to a Go binary
func TranslateCBinary2GoBinary(c *C.struct_Binary) *bytes.Buffer {
	if c == nil {
		return nil
	}

	return bytes.NewBuffer(unsafe.Slice((*byte)(c.buffer), int32(c.size)))
}

func (bs *BookStoreCgoClient) GetOrdersByThrift() (orders thriftTypes.Orders) {
	cbinary := C.getOrders(bs.cpointer)
	// Must free c binary after copy.
	defer C.free(unsafe.Pointer(cbinary))

	gobinary := TranslateCBinary2GoBinary(cbinary)
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	mem_buffer.Buffer = gobinary
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)
	orders.Read(protocol)
	return orders
}

func (bs *BookStoreCgoClient) AddOrder(order thriftTypes.Order) {
	mem_buffer := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(mem_buffer)

	printAndPanicError(order.Write(protocol))
	ptr := unsafe.Pointer(&mem_buffer.Bytes()[0])
	size := mem_buffer.Len()

	C.addOrder(bs.cpointer, ptr, C.uint(size))
}

func (bs *BookStoreCgoClient) GetBookStoreName(name []byte) {
	size := len(name)
	C.getBookStoreName(bs.cpointer, unsafe.Pointer(&name[0]), C.uint32_t(size))
}

func printAndPanicError(err error) {
	if err != nil {
		fmt.Println("Error ", err)
		panic(err)
	}
}
