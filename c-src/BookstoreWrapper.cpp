#include "cpp-src/Bookstore.h"
#include "cpp-src/Serialization.h"
#include "include/BookstoreWrapper.h"
#include "gen-src/gen-cpp/Types_types.h"
#include "thrift/protocol/TProtocol.h"
#include <cstddef>
#include <cstdlib>
#include <exception>
#include <memory>
#include <string>
#include <iostream>
#include <sys/types.h>

static ThriftSerializer global_serializer = ThriftSerializer(false, 1024);

BookStore::Book* ToCppBook(CBook* book)
{
    BookStore::Book* bs_book = new BookStore::Book();

    bs_book->name = std::string(book->name);
    bs_book->price = book->price;
    bs_book->author.name = std::string(book->author.name);
    bs_book->author.age = book->author.age;

    return bs_book;
}

namespace BookStore
{

extern "C"
{

void* initBookStore()
{
    BookStore* res = new BookStore();
    
    return res;
}

void freeBookStore(void* bookStore)
{
    BookStore* ptr = static_cast<BookStore*>(bookStore);
    
    delete ptr;
    
    return;
}

bool hasBook(void* bookStore, CBook book)
{
    BookStore* bstore = static_cast<BookStore*>(bookStore);
    Book* book_ = ToCppBook(&book);
    bool res = bstore->hasBook(*book_);
    delete book_;
    return res;
}

void addBook(void* bookStore, CBook book)
{
    BookStore* bstore = static_cast<BookStore*>(bookStore);
    Book* book_ = ToCppBook(&book);
    bstore->addBook(*book_);
    delete book_;
    return;
}

void addOrder(void* bookStore, void* buffer, uint32_t size)
{
    CXX::Order cxxorder = deserializeFromBinanry<CXX::Order>(buffer, size);
    BookStore* bstore = static_cast<BookStore*>(bookStore);
    bstore->addOrder(cxxorder);
}

Binary* getOrders(void* bookStore)
{
    Binary* res_binary = new Binary();
    BookStore* bstore = static_cast<BookStore*>(bookStore);
    CXX::Orders orders = bstore->getOrders();
    global_serializer.serialize<CXX::Orders>(&orders, &res_binary->size, (uint8_t**)(&res_binary->buffer));
    return res_binary;
}


}

}
