#include "BookstoreClientMock.h"
#include "cpp-src/Serialization.h"
#include "gen-src/gen-cpp/Types_types.h"
#include "include/BookstoreWrapper.h"
#include <cstddef>
#include <cstdint>
#include <cstdlib>
#include <cstring>
#include <gmock/gmock-spec-builders.h>
#include <gmock/gmock.h>
#include <gtest/gtest.h>
#include <iostream>

using namespace testing;

int main(int argc, char **argv) {
  ::testing::InitGoogleMock(&argc, argv);

  ThriftSerializer local_serializer = ThriftSerializer(false, 1024);

  MockBoostoreClient mock_bsclient;
  thrift::Order order;
  order.book_name = "Red Sun";
  order.customer_name = "HZQ";

  Binary *binary_order = new Binary();
  local_serializer.serialize<thrift::Order>(
      &order, &binary_order->size, (uint8_t **)(&binary_order->buffer));

  EXPECT_CALL(mock_bsclient, AddOrder(_)).Times(1);
  addOrder(&mock_bsclient, binary_order->buffer, binary_order->size);

  free(binary_order->buffer);
  delete binary_order;

  thrift::Book book;
  book.author.age = 32;
  book.author.name = "ABCDEFG LIJKLMN OPQRST";
  book.name = "Red Sun";
  book.price = 58;

  CBook cbook;
  size_t strlen_ = strlen(book.name.c_str());
  cbook.name = (char *)malloc(strlen_ + 1);  // +1 for null
  std::cout << "book.name.size(): " << book.name.size() << " strlen: " <<  strlen_ << std::endl;
  strcpy(cbook.name, book.name.c_str());

  cbook.price = book.price;
  cbook.author.age = book.author.age;
  strlen_ = strlen(book.author.name.c_str());
  cbook.author.name = (char *)malloc(strlen_ + 1);
  std::cout << "author.name.size(): " << book.author.name.size() << " strlen: " << strlen_ << std::endl;
  strcpy(cbook.author.name, book.author.name.c_str());

  EXPECT_CALL(mock_bsclient, AddBook(book)).Times(1);
  addBook(&mock_bsclient, cbook);

  EXPECT_CALL(mock_bsclient, HasBook(book)).WillOnce(Return(true));
  EXPECT_TRUE(hasBook(&mock_bsclient, cbook));

    std::cout << "strlen author name " << strlen(cbook.author.name) << std::endl;
  for(size_t i = 0; i < strlen(cbook.author.name); ++ i) {
    std::cout << cbook.author.name[i];
  }
    
    std::cout << std::endl;


  free(cbook.name);
  free(cbook.author.name);
}