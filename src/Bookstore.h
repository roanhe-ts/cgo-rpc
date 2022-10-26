#include <string>
#include <vector>

struct Book
{
    std::string name;
    int64_t price;

    Book(const std::string& name_, int64_t price_) : name(name_), price(price_){}
};

class BookStore
{
private:
    std::vector<Book> books;

public:
    BookStore() = default;
    BookStore(const std::vector<Book>& books_) : books(books_) {}

    bool hasBook(const Book& book);
    void addBook(const Book& book);
};