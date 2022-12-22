#include "cpp-src/BookstoreClient.h"
#include "cpp-src/Serialization.h"
#include "gen-src/gen-cpp/BookStoreService.h"
#include <cstddef>
#include <cstdint>
#include <cstdlib>
#include <memory>
#include <string>

static ThriftSerializer global_serializer = ThriftSerializer(false, 1024);

using namespace apache::thrift;
using namespace apache::thrift::protocol;
using namespace apache::thrift::transport;

#ifndef MAKE_TEST
extern "C"
{
#endif

typedef struct {
    char* name;
    int32_t age;
} CAuthor;

typedef struct {
    char* name;
    int32_t price;
    CAuthor author;
} CBook;

typedef struct Binary {
    void* buffer;
    uint32_t size;
} Binary;

static thrift::Book ToThriftBook(const CBook& book)
{
    thrift::Book thrift_book;
    thrift_book.name = std::string(book.name);
    thrift_book.price = book.price;
    thrift_book.author.name = std::string(book.author.name);
    thrift_book.author.age = book.author.age;

    return thrift_book;
}

void* initBSClient()
{
    return new BookStoreClient();
}

void freeBSClient(void* cpp_ptr)
{
    BookStoreClient* ptr = static_cast<BookStoreClient*>(cpp_ptr);
    delete ptr;
    return;
}

bool hasBook(void* cpp_ptr, CBook book)
{
    BookStoreClient* bsclient = static_cast<BookStoreClient*>(cpp_ptr);
    
    thrift::Book book_ = ToThriftBook(book);
    bool res = bsclient->HasBook(book_);
    return res;
}

void addBook(void* cpp_ptr, CBook book)
{
    BookStoreClient* bstore = static_cast<BookStoreClient*>(cpp_ptr);
    thrift::Book book_ = ToThriftBook(book);
    bstore->AddBook(book_);
    return;
}

void addOrder(void* cpp_ptr, void* buffer, uint32_t size)
{
    thrift::Order cxxorder = deserializeFromBinanry<thrift::Order>(buffer, size);
    BookStoreClient* bstore = static_cast<BookStoreClient*>(cpp_ptr);
    bstore->AddOrder(cxxorder);
}

Binary* getOrders(void* cpp_ptr)
{
    BookStoreClient* bsclient = static_cast<BookStoreClient*>(cpp_ptr);
    thrift::Orders orders;
    bsclient->GetOrders(orders);
    
    Binary* res_binary = new Binary();
    global_serializer.serialize<thrift::Orders>(&orders, &res_binary->size, (uint8_t**)(&res_binary->buffer));
    return res_binary;
}

void getBookStoreName(void* cpp_ptr, void* buf, uint32_t size)
{
    // Note: buf is allocated by golang.
    uint8_t* binary_buffer = static_cast<uint8_t*>(buf);
    std::shared_ptr<apache::thrift::transport::TMemoryBuffer> tmem_transport(
            new apache::thrift::transport::TMemoryBuffer(binary_buffer, size));
    std::shared_ptr<apache::thrift::protocol::TProtocol> proto = create_deserialize_protocol(tmem_transport, false);

    BookStoreClient* bsclient = static_cast<BookStoreClient*>(cpp_ptr);
    bsclient->GetBookStoreName(buf, size);
}

#ifndef MAKE_TEST
}
#endif
