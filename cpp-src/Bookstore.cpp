#include "Bookstore.h"
#include <map>
#include <set>
#include <string>

namespace BookStore
{

void BookStore::addBook(const Book &book_)
{
    books.push_back(book_);
}

bool BookStore::hasBook(const Book &book_)
{
    for (const auto& book : books)
    {
        if (book.name == book_.name)
            return true;
    }

    return false;
}

CXX::Orders BookStore::getOrders()
{
    return orders;
}

void BookStore::addOrder(const CXX::Order &order)
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
