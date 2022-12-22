set(TEST_SRC
    test_main.cpp)

add_executable(test-main ${TEST_SRC})
target_include_directories(test-main PRIVATE ${CMAKE_CURRENT_SOURCE_DIR})
target_link_libraries(test-main PRIVATE bookstorecli_cpp pthread gmock)
