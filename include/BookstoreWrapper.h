#include <stdint.h>
#include <stdlib.h>
#include <stdbool.h>

void* initBookStore();
void freeBookStore(void*);

typedef struct {
    char* name;
    int32_t age;
} CAuthor;

typedef struct {
    char* name;
    int32_t price;
    CAuthor author;
} CBook;

bool hasBook(void* bookStore, CBook book);

void addBook(void* bookStore, CBook book);

typedef struct Binary {
    void* buffer;
    uint32_t size;
} Binary;

Binary* getOrders(void* bookStore);
void addOrder(void* bookStore, void* binary_order, uint32_t size);

