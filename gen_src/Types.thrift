namespace cpp CXX

struct Author {
    1: required string name
    2: required i32 age
}

struct Book {
    1: required string name
    2: required i32 price
    3: required Author author
}