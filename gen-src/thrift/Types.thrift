# cpp_include "<unordered_set>"
namespace cpp thrift

struct Order {
    1: required string customer_name
    2: required string book_name
}

struct Orders {
    # 1: required map<string, set cpp_type "std::set<Book, bool(*)(const Book&, const Book&)>" <Book>> entry
    1: required list<Order> entry
}

struct Author {
    1: required string name
    2: required i32 age
}

struct Book {
    1: required string name
    2: required i32 price
    3: Author author
}
