#ifndef BOOKSTORESERVER
#define BOOKSTORESERVER

#include "./gen-src/gen-cpp/BookStoreService.h"
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/server/TSimpleServer.h>
#include <thrift/transport/TServerSocket.h>
#include <thrift/transport/TBufferTransports.h>


class BookstoreServer : virtual public BookStoreServiceIf {
public:
  BookstoreServer() {}

  void Start();

  void GetOrders( ::thrift::Orders& _return);

  void AddOrder(const  ::thrift::Order& order);

  void AddBook(const  ::thrift::Book& book);

  bool HasBook(const  ::thrift::Book& book);

private:
    
};

#endif