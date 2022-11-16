#include <string>
#include <vector>
#include "gen-src/gen-cpp/Types_types.h"

namespace BookStore
{

struct Author
{
    std::string name;
    int age;
};

struct Book
{
    std::string name;
    int price;
    Author author;
};

class BookStore
{
private:
    std::vector<Book> books;
    CXX::Orders orders;

public:
    BookStore() = default;
    BookStore(const std::vector<Book>& books_) : books(books_) {}

    bool hasBook(const Book& book);
    void addBook(const Book& book);
    void addOrder(const CXX::Order& order);

    CXX::Orders getOrders();
};

}
