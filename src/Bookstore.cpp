#include "BookStore.h"

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
