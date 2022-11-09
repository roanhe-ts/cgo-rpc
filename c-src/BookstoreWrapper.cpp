#include "cpp-src/Bookstore.h"
#include "cpp-src/Serialization.h"
#include "thrift/protocol/TProtocol.h"
#include <cstddef>
#include <cstdlib>
#include <exception>
#include <memory>
#include <string>
#include <iostream>
#include <sys/types.h>

static ThriftSerializer global_serializer = ThriftSerializer(false, 1024);

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

bool hasBook(void* bookStore, void* buffer, uint32_t size)
{
    CXX::Book cxxbook = deserializeFromBinanry<CXX::Book>(buffer, size);    
    BookStore* bstore = static_cast<BookStore*>(bookStore);
    return bstore->hasBook(cxxbook);
}

void addBook(void* bookStore, void* buffer, uint32_t size)
{
    CXX::Book cxxbook = deserializeFromBinanry<CXX::Book>(buffer, size);
    BookStore* bstore = static_cast<BookStore*>(bookStore);
    bstore->addBook(cxxbook);
    return;
}

struct binary {
    void* buffer;
    uint32_t size;
};

binary* getOrders(void* bookStore)
{
    binary* res_binary = new binary();
    BookStore* bstore = static_cast<BookStore*>(bookStore);
    CXX::Orders orders = bstore->getOrders();
    global_serializer.serialize<CXX::Orders>(&orders, &res_binary->size, (uint8_t**)(&res_binary->buffer));
    return res_binary;
}

}
