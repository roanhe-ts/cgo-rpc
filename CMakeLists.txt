cmake_minimum_required(VERSION 3.20)
project(cgo-thrift)

set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(GEN_SRC_DIR ${PROJECT_SOURCE_DIR}/gen-src)
set(THRIFT_HEADER /usr/local/include)

include_directories(${THRIFT_HEADER})

add_custom_command(
    OUTPUT 
        ${GEN_SRC_DIR}/gen-cpp/Types_types.cpp ${GEN_SRC_DIR}/gen-cpp/Types_types.h
    COMMAND 
        thrift -gen cpp ${GEN_SRC_DIR}/Types.thrift
    DEPENDS
        ${GEN_SRC_DIR}/Types.thrift
    WORKING_DIRECTORY
        ${GEN_SRC_DIR}
    COMMENT 
        "Generating thrift types cpp ..."
)

add_custom_command(
    OUTPUT 
        ${GEN_SRC_DIR}/gen-cpp/BookStoreService.cpp ${GEN_SRC_DIR}/gen-cpp/BookStoreService.h
    COMMAND 
        thrift -gen cpp ${GEN_SRC_DIR}/Bookstore.thrift
    DEPENDS
        ${GEN_SRC_DIR}/Bookstore.thrift
    WORKING_DIRECTORY
        ${GEN_SRC_DIR}
    COMMENT 
        "Generating thrift service cpp ..."
)

add_custom_command(
    OUTPUT 
        THRIFT_GO_GENS
    COMMAND 
        thrift -gen go ${GEN_SRC_DIR}/Types.thrift
    DEPENDS
        ${GEN_SRC_DIR}/Types.thrift
    WORKING_DIRECTORY
        ${GEN_SRC_DIR}
    COMMENT 
        "Generating thrift go ..."
)

set(BOOKSTORE_SRC
    ${CMAKE_SOURCE_DIR}/cpp-src/BookstoreServer.cpp
    ${CMAKE_SOURCE_DIR}/cpp-src/ServerMain.cpp)

add_executable(server
    ${BOOKSTORE_SRC} 
    ${GEN_SRC_DIR}/gen-cpp/Types_types.cpp 
    ${GEN_SRC_DIR}/gen-cpp/Types_types.h
    ${GEN_SRC_DIR}/gen-cpp/BookStoreService.cpp
    ${GEN_SRC_DIR}/gen-cpp/BookStoreService.h)

target_link_libraries(server PUBLIC thrift)
target_include_directories(server PRIVATE ${PROJECT_SOURCE_DIR})
target_compile_options(server PRIVATE "-std=c++14")   # std::make_unique

add_library(bookstorecli_c
    ${CMAKE_SOURCE_DIR}/c-src/WBookstoreClient.cpp
    ${GEN_SRC_DIR}/gen-cpp/Types_types.cpp 
    ${GEN_SRC_DIR}/gen-cpp/Types_types.h
    ${GEN_SRC_DIR}/gen-cpp/BookStoreService.cpp)
# target_link_directories(bookstore_c PUBLIC "/usr/local/Cellar/thrift/0.16.0/lib")
target_link_libraries(bookstorecli_c PUBLIC bookstore thrift)
target_include_directories(bookstorecli_c PRIVATE ${PROJECT_SOURCE_DIR})
target_compile_options(bookstorecli_c PRIVATE "-std=c++17")   # std::string_view

add_custom_target(go-cli
    go
    build
    -o ${PROJECT_BINARY_DIR}/go-cli
    ${PROJECT_SOURCE_DIR}/main.go

    DEPENDS
        THRIFT_GO_GENS
)
