package main

/*
#include <stdio.h>
#include <stdint.h>

typedef struct Foo {
    int32_t a;
    int32_t b;
} Foo;

void pass_array(Foo **in) {
    int i;

    for(i = 0; i < 2; i++) {
        fprintf(stderr, "[%d, %d]", in[i]->a, in[i]->b);
    }
    fprintf(stderr, "\n");
}

void pass_struct(Foo *in) {
    fprintf(stderr, "[%d, %d]\n", in->a, in->b);
}

*/
import "C"

import (
	thriftTypes "cgo-thrift/gen-src/gen-go/types"
	bookstore "cgo-thrift/go-src/Bookstore"
	"fmt"
	"strings"
)

type Foo struct {
	A int32
	B int32
}

func main() {
	bs := bookstore.BookStoreCgo{}
	bs.Init()
	defer bs.Free()

	book := bookstore.Book{
		Name:  "Book1",
		Price: 112,
		Author: bookstore.Author{
			Name: "AWS",
			Age:  32,
		},
	}

	bs.AddBook(book)

	if bs.HasBook(book) {
		println("Add succed")
	}

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 100)
	order.BookName = strings.Repeat("H", 100)

	bs.AddOrder(order)

	orders := bs.GetOrdersByThrift()
	fmt.Println(orders.String())

	// foo := Foo{25, 26}
	// foos := []Foo{{25, 26}, {50, 51}}

	// // wrong result = [25, 0]
	// C.pass_struct((*C.struct_Foo)(unsafe.Pointer(&foo)))

	// // doesn't work at all, SIGSEGV
	// // C.pass_array((**_Ctype_Foo)(unsafe.Pointer(&foos[0])))

	// // wrong result = [25, 0], [50, 0]
	// out := make([]*C.struct_Foo, len(foos))
	// out[0] = (*C.struct_Foo)(unsafe.Pointer(&foos[0]))
	// out[1] = (*C.struct_Foo)(unsafe.Pointer(&foos[1]))
	// C.pass_array((**C.struct_Foo)(unsafe.Pointer(&out[0])))

	return
}
