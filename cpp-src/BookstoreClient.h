#include <string>
#include <vector>
#include "gen-src/gen-cpp/Types_types.h"

namespace BookStoreClient
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

class BookStoreClient
{
private:
    std::vector<Book> books;
    thrift::Orders orders;

public:
    BookStoreClient() = default;
    BookStoreClient(const std::vector<Book>& books_) : books(books_) {}

    bool hasBook(const Book& book);
    void addBook(const Book& book);
    void addOrder(const thrift::Order& order);

    thrift::Orders getOrders();
};

}
