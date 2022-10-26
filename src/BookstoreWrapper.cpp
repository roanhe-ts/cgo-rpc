#include "Bookstore.h"
#include "thrift/protocol/TProtocol.h"
#include <_types/_uint32_t.h>
#include <_types/_uint8_t.h>
#include <cstddef>
#include <cstdlib>
#include <exception>
#include <memory>
#include <string>
#include <thrift/TApplicationException.h>
#include <thrift/protocol/TCompactProtocol.h>
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/transport/TBufferTransports.h>
#include <iostream>

std::shared_ptr<apache::thrift::protocol::TProtocol> create_deserialize_protocol(
        std::shared_ptr<apache::thrift::transport::TMemoryBuffer> mem, 
        bool compact)
{
    if (compact) {
        apache::thrift::protocol::TCompactProtocolFactoryT<apache::thrift::transport::TMemoryBuffer>
                tproto_factory;
        return tproto_factory.getProtocol(mem);
    } else {
        apache::thrift::protocol::TBinaryProtocolFactoryT<apache::thrift::transport::TMemoryBuffer>
                tproto_factory;
        return tproto_factory.getProtocol(mem);
    }
}

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

bool hasBook(void* bookStore, void* go_book, uint32_t size)
{
    uint8_t* binary_book = static_cast<uint8_t*>(go_book);
    std::shared_ptr<apache::thrift::transport::TMemoryBuffer> tmem_transport(
            new apache::thrift::transport::TMemoryBuffer(binary_book, size));
    std::shared_ptr<apache::thrift::protocol::TProtocol> tproto = create_deserialize_protocol(tmem_transport, false);
    CXX::Book cxxbook;
    try {
        cxxbook.read(tproto.get());
    } catch (std::exception& e) {
        std::cout << "Couldn't deserialize thrift msg: " << e.what() << std::endl;
        exit(-1);
    }
    
    BookStore* bstore = static_cast<BookStore*>(bookStore);
    return bstore->hasBook(cxxbook);
}

void addBook(void* bookStore, void* go_book, uint32_t size)
{
    uint8_t* binary_book = static_cast<uint8_t*>(go_book);
    std::shared_ptr<apache::thrift::transport::TMemoryBuffer> tmem_transport(
            new apache::thrift::transport::TMemoryBuffer(binary_book, size));
    std::shared_ptr<apache::thrift::protocol::TProtocol> tproto = create_deserialize_protocol(tmem_transport, false);
    CXX::Book cxxbook;
    try {
        cxxbook.read(tproto.get());
    } catch (std::exception& e) {
        std::cout << "Couldn't deserialize thrift msg: " << e.what() << std::endl;
        exit(-1);
    }
    
    
    BookStore* bstore = static_cast<BookStore*>(bookStore);
    bstore->addBook(cxxbook);

    std::cout << "Added book: ";
    cxxbook.printTo(std::cout);

    return;
}

}
