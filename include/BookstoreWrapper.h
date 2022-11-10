#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>

void* initBookStore();
void freeBookStore(void*);

bool hasBook(void* bookStore, void* binanry_book, uint32_t size);

void addBook(void* bookStore, void* binanry_book, uint32_t size);

struct binary {
    void* buffer;
    uint32_t size;
};

struct binary* getOrders(void* bookStore);
void addOrder(void* bookStore, void* binary_order, uint32_t size);
