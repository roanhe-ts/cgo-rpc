include "Types.thrift"

service BookStoreService {
    Types.Orders GetOrders()
    void AddOrder(1:Types.Order order)
    void AddBook(1:Types.Book book)
    bool HasBook(1:Types.Book book)
    binary GetBookStoreName(1: i32 size)
}