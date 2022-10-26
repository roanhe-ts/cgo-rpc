#include "Bookstore.h"

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
