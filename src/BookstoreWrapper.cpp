#include "Bookstore.h"
#include <string>

extern "C"
{

struct CBook {
    char* name;
    int64_t price;

    CBook(char* name_, int64_t price_) : name(name_), price(price_) {}
};

void* initBookStore()
{
    BookStore* res = new BookStore();
    
    return res;
}

void freeBookStore(void* bookStore)
{
    BookStore* ptr = static_cast<BookStore*>(bookStore);
    
    delete ptr;
    
    return;
}

bool hasBook(void* bookStore, CBook* cbook)
{
    BookStore* bstore = static_cast<BookStore*>(bookStore);

    Book book(std::string(cbook->name), cbook->price);

    return bstore->hasBook(book);
}

void addBook(void* bookStore, CBook* cbook)
{
    BookStore* bstore = static_cast<BookStore*>(bookStore);

    Book book(std::string(cbook->name), cbook->price);

    bstore->addBook(book);
}

CBook* parse(char* name_, int64_t price_)
{
    return new CBook(name_, price_);
}

}
