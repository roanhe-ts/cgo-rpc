package BookStore

/*
#cgo CFLAGS: -I/Users/hzq/VSCodeProject/cgo-thrift/include -Wvisibility
#cgo LDFLAGS: -L/Users/hzq/VSCodeProject/cgo-thrift/build -lbookstore_c -lbookstore -lstdc++
#include "BookStoreWrapper.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Book struct {
	Name  string
	Price int64
}

type BookStore struct {
	BookStoreCPtr unsafe.Pointer
}

// ParseBook Convert a Go struct Book to a C Book struct pointer
func ParseBook(book Book) *C.struct_CBook {
	cs := C.CString(book.Name)
	defer C.free(unsafe.Pointer(cs))
	var res = C.parse(cs, C.longlong(book.Price))
	return (*C.struct_CBook)(res)
}

func HasBook(bookStoreCPtr unsafe.Pointer, book Book) bool {
	cbook := ParseBook(book)
	return bool(C.hasBook(bookStoreCPtr, cbook))
}

func AddBook(bookStoreCPtr unsafe.Pointer, book Book) {
	cbook := ParseBook(book)
	C.addBook(bookStoreCPtr, cbook)
}

func InitBookStore() unsafe.Pointer {
	return C.initBookStore()
}
