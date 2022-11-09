namespace cpp CXX
cpp_include "<unordered_set>"

struct Author {
    1: required string name
    2: required i32 age
}

struct Book {
    1: required string name
    2: required i32 price
    3: required Author author
}

struct Orders {
    # 1: required map<string, set cpp_type "std::set<Book, bool(*)(const Book&, const Book&)>" <Book>> entry
    1: required map<string, set<string>> entry
}