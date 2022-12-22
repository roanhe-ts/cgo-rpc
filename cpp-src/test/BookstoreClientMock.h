#include <cstdint>
#include <gmock/gmock-generated-function-mockers.h>
#include <gmock/gmock.h>
#ifndef BSTORE_MOCK_H
#define BSTORE_MOCK_H

#include "cpp-src/BookstoreClient.h"

using namespace apache::thrift;

class MockBoostoreClient : BookStoreClient {
public:
    MOCK_METHOD1(GetOrders, void(thrift::Orders&));
    MOCK_METHOD1(AddOrder, void(const thrift::Order));
    MOCK_METHOD1(AddBook, void(const thrift::Book));
    MOCK_METHOD1(HasBook, bool(const thrift::Book));
    MOCK_METHOD2(GetBookStoreName, void(void*, uint32_t));
};

#endif