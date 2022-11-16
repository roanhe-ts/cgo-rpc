namespace cpp CXX
cpp_include "<unordered_set>"

struct Order {
    1: required string customer_name
    2: required string book_name
}

struct Orders {
    # 1: required map<string, set cpp_type "std::set<Book, bool(*)(const Book&, const Book&)>" <Book>> entry
    1: required map<string, set<string>> entry
}