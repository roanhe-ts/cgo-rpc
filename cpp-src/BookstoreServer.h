#ifndef BOOKSTORESERVER
#define BOOKSTORESERVER

#include "./gen-src/gen-cpp/BookStoreService.h"
#include <memory>
#include <string>
#include <unordered_set>
#include <thrift/protocol/TBinaryProtocol.h>
#include <thrift/server/TSimpleServer.h>
#include <thrift/transport/TServerSocket.h>
#include <thrift/transport/TBufferTransports.h>

using namespace thrift;

class BookstoreServer : virtual public BookStoreServiceIf {
public:
  BookstoreServer() {}

  void Start();
  void Shutdown();

  void GetOrders(Orders& _return);

  void AddOrder(const Order& order);

  void AddBook(const Book& book);

  bool HasBook(const Book& book);

private:
    std::vector<Book> books;
    std::unordered_set<std::string> book_sets;
    thrift::Orders orders;
    std::unique_ptr<apache::thrift::server::TSimpleServer> server;
};

#endif