#include "BookstoreServer.h"

using namespace ::apache::thrift;
using namespace ::apache::thrift::protocol;
using namespace ::apache::thrift::transport;
using namespace ::apache::thrift::server;

void BookstoreServer::Start()
{
    int port = 9090;

    ::apache::thrift::stdcxx::shared_ptr<BookstoreServer> handler(this);
    ::apache::thrift::stdcxx::shared_ptr<TProcessor> processor(new BookStoreServiceProcessor(handler));
    ::apache::thrift::stdcxx::shared_ptr<TServerTransport> serverTransport(new TServerSocket(port));
    ::apache::thrift::stdcxx::shared_ptr<TTransportFactory> transportFactory(new TBufferedTransportFactory());
    ::apache::thrift::stdcxx::shared_ptr<TProtocolFactory> protocolFactory(new TBinaryProtocolFactory());

    TSimpleServer server(processor, serverTransport, transportFactory, protocolFactory);
    server.serve();
  }

void BookstoreServer::GetOrders( ::thrift::Orders& _return) {
    // Your implementation goes here
    printf("GetOrders\n");
  }

  void AddOrder(const  ::thrift::Order& order) {
    // Your implementation goes here
    printf("AddOrder\n");
  }

  void AddBook(const  ::thrift::Book& book) {
    // Your implementation goes here
    printf("AddBook\n");
  }

  bool HasBook(const  ::thrift::Book& book) {
    // Your implementation goes here
    printf("HasBook\n");
  }

};