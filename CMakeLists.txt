cmake_minimum_required(VERSION 3.20)
project(cgo-thrift)

set(CMAKE_CXX_STANDARD 11)

set(BOOKSTORE_SRC
    ${CMAKE_SOURCE_DIR}/src/Bookstore.h    
    ${CMAKE_SOURCE_DIR}/src/Bookstore.cpp)

add_library(bookstore STATIC ${BOOKSTORE_SRC})

add_library(bookstore_c STATIC ${CMAKE_SOURCE_DIR}/src/BookstoreWrapper.cpp)
target_link_libraries(bookstore_c PUBLIC bookstore)

add_executable(main main.c)
target_link_libraries(main bookstore_c)