#include "BookstoreServer.h"
#include "gen-src/gen-cpp/Types_types.h"
#include <memory>

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

    server = std::make_unique<apache::thrift::server::TSimpleServer>(processor, serverTransport, transportFactory, protocolFactory);
    server->serve();
}

void BookstoreServer::Shutdown()
{
    server->stop();
}

void BookstoreServer::GetOrders(Orders& _return)
{
    printf("GetOrders\n");   
    _return = orders;
    return;
}

void BookstoreServer::AddOrder(const Order& order)
{
    printf("AddOrder\n");
    orders.entry.emplace_back(order);
    return;
}

void BookstoreServer::AddBook(const Book& book)
{
    printf("AddBook\n");
    books.emplace_back(book);
    book_sets.insert(book.name);
    return;
}

bool BookstoreServer::HasBook(const Book& book)
{
    printf("HasBook\n");
    return book_sets.count(book.name);
}
