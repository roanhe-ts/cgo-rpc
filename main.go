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

func main() {
	bs_client := bookstore.BookStoreCgoClient{}
	bs_client.Init()
	defer bs_client.Free()

	book := bookstore.Book{
		Name:  "Book1",
		Price: 112,
		Author: bookstore.Author{
			Name: "AWS",
			Age:  32,
		},
	}

	bs_client.AddBook(book)

	if bs_client.HasBook(book) {
		println("Add succed")
	}

	var order thriftTypes.Order
	order.CustomerName = strings.Repeat("Z", 10)
	order.BookName = strings.Repeat("H", 10)

	bs_client.AddOrder(order)

	orders := bs_client.GetOrdersByThrift()
	fmt.Println(orders.String())

	bs_name := make([]byte, 10)
	bs_client.GetBookStoreName(bs_name)
	fmt.Println(string(bs_name))

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
