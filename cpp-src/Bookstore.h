#include <string>
#include <vector>
#include "gen-src/gen-cpp/Types_types.h"

class BookStore
{
private:
    std::vector<CXX::Book> books;
    CXX::Orders orders;

public:
    BookStore() = default;
    BookStore(const std::vector<CXX::Book>& books_) : books(books_) {}

    bool hasBook(const CXX::Book& book);
    void addBook(const CXX::Book& book);
    void addOrder(const CXX::Order& order);

    CXX::Orders getOrders();
};