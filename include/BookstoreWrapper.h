#include <stdlib.h>
#include <stdbool.h>

typedef struct CBook CBook;

void* initBookStore();
void freeBookStore(void*);

bool hasBook(void* bookStore, struct CBook* book);

void addBook(void* bookStore, struct CBook* book);

struct CBook* parse(char* name, int64_t);
