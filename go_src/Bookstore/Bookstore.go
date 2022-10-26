package BookStore

/*
#cgo CFLAGS: -I/Users/hzq/VSCodeProject/cgo-thrift/include -Wvisibility
#cgo LDFLAGS: -L/Users/hzq/VSCodeProject/cgo-thrift/build -lbookstore_c -lbookstore -lthrift -lstdc++
#include "BookStoreWrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	thriftTypes "cgo-thrift/gen_src/gen-go/types"
	"context"
	"unsafe"
	"fmt"

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

func InitBookStore() unsafe.Pointer {
	return C.initBookStore()
}

func printAndPanicError(err error) {
	if err != nil {
		fmt.Println("Error ", err)
		panic(err)
	}
}