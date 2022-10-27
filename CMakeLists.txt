cmake_minimum_required(VERSION 3.20)
project(cgo-thrift)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)

set(GEN_SRC_DIR ${PROJECT_SOURCE_DIR}/gen_src)
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
        "Generating thrift cpp ..."
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
    ${CMAKE_SOURCE_DIR}/src/Bookstore.h    
    ${CMAKE_SOURCE_DIR}/src/Bookstore.cpp)

add_library(bookstore STATIC ${BOOKSTORE_SRC} ${GEN_SRC_DIR}/gen-cpp/Types_types.cpp ${GEN_SRC_DIR}/gen-cpp/Types_types.h)
target_link_libraries(bookstore PUBLIC thrift)

add_library(bookstore_c STATIC ${CMAKE_SOURCE_DIR}/src/BookstoreWrapper.cpp)
target_link_directories(bookstore_c PUBLIC "/usr/local/Cellar/thrift/0.16.0/lib")
target_link_libraries(bookstore_c PUBLIC bookstore thrift)


add_custom_target(bookstore_go
    go
    build
    -o ${PROJECT_BINARY_DIR}/bookstore_go
    ${PROJECT_SOURCE_DIR}/main.go
)
