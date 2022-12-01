#include "BookstoreClient.h"
#include <map>
#include <set>
#include <string>

namespace BookStoreClient
{

void BookStoreClient::addBook(const Book &book_)
{
    books.push_back(book_);
}

bool BookStoreClient::hasBook(const Book &book_)
{
    for (const auto& book : books)
    {
        if (book.name == book_.name)
            return true;
    }

    return false;
}

thrift::Orders BookStoreClient::getOrders()
{
    return orders;
}

void BookStoreClient::addOrder(const thrift::Order &order)
{
    if (orders.entry.find(order.customer_name) != orders.entry.end())
    {
        orders.entry[order.customer_name].insert(order.book_name);
    }
    else
    {
        orders.entry.insert(
            std::pair<std::string, std::set<std::string>>(order.customer_name, std::set<std::string>{order.book_name}));
    }
}

}
