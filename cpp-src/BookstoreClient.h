#pragma once

#include <cstdint>
#include <cstring>
#include <memory>
#include <string>
#include <string_view>
#include <sys/socket.h>
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/protocol/TProtocol.h>
#include <thrift/transport/TBufferTransports.h>
#include <thrift/transport/TSocket.h>
#include <utility>
#include <vector>
#include "gen-src/gen-cpp/Types_types.h"
#include "gen-src/gen-cpp/BookStoreService.h"
#include <iostream>

using namespace apache::thrift;
using namespace apache::thrift::transport;
using namespace apache::thrift::protocol;

class BookStoreClient
{
public:
    BookStoreClient() : socket(new TSocket("localhost", 9090)), transport(new TBufferedTransport(socket)), protocol(new TBinaryProtocol(transport)) {}

    void GetOrders(thrift::Orders& _ret)
    {
        BookStoreServiceClient client(protocol);
        transport->open();
        client.GetOrders(_ret);
        transport->close();
    }

    void AddOrder(const thrift::Order order)
    {
        BookStoreServiceClient client(protocol);
        transport->open();
        client.AddOrder(order);
        transport->close();
    }

    void AddBook(const thrift::Book book)
    {
        BookStoreServiceClient client(protocol);
        transport->open();
        if (transport->isOpen())
        {
            std::cout << "Opne\n";
        }
        client.AddBook(book);
        transport->close();
    }

    bool HasBook(const thrift::Book book)
    {
        BookStoreServiceClient client(protocol);
        transport->open();
        bool res = client.HasBook(book);
        transport->close();
        return res;
    }

    // Memory is allocated by caller.
    void GetBookStoreName(void* buf, uint32_t size)
    {   
        BookStoreServiceClient client(protocol);
        transport->open();

        std::string _return((char*)buf);
        client.GetBookStoreName(_return, size);
        std::cout << "_return: " << _return <<" buf add: " << buf << " string add:" << static_cast<const void*>(_return.c_str()) << std::endl;

        mempcpy(buf, _return.c_str(), size);
        transport->close();

        return;
    }

private:
    std::shared_ptr<TTransport> socket;
    std::shared_ptr<TTransport> transport;
    std::shared_ptr<TProtocol> protocol;
};
