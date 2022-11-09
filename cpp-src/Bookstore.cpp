#include "Bookstore.h"
#include <map>
#include <string>

void BookStore::addBook(const CXX::Book &book_)
{
    books.push_back(book_);
}

bool BookStore::hasBook(const CXX::Book &book_)
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
    CXX::Orders orders;
    std::set<std::string> books {"Red son"};
    
    orders.entry["DDDD"] = books;
    return orders;
}
