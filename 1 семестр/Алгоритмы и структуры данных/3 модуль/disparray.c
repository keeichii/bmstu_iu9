#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Node {
    long long key, val;
    struct Node* next;
} Node;

static Node** table;
static long long M;

static void doAssign(long long i, long long v) {
    long long h = i % M;
    if (h < 0) h += M;
    Node* p = table[h];
    Node* prev = NULL;
    while (p) {
        if (p->key == i) {
            if (v == 0) {
                if (prev) prev->next = p->next;
                else table[h] = p->next;
                free(p);
            } else {
                p->val = v;
            }
            return;
        }
        prev = p;
        p = p->next;
    }
    if (v != 0) {
        Node* nn = malloc(sizeof(Node));
        nn->key = i;
        nn->val = v;
        nn->next = table[h];
        table[h] = nn;
    }
}

static long long doAt(long long i) {
    long long h = i % M;
    if (h < 0) h += M;
    Node* p = table[h];
    while (p) {
        if (p->key == i) return p->val;
        p = p->next;
    }
    return 0;
}

int main() {
    scanf("%lld", &M);
    table = calloc(M, sizeof(Node*));
    while (1) {
        char cmd[8];
        if (scanf("%s", cmd) != 1) break;
        if (strcmp(cmd, "END") == 0) break;
        if (strcmp(cmd, "ASSIGN") == 0) {
            long long i, v;
            scanf("%lld%lld", &i, &v);
            doAssign(i, v);
        } else if (strcmp(cmd, "AT") == 0) {
            long long i;
            scanf("%lld", &i);
            printf("%lld\n", doAt(i));
        }
    }
    for (long long i = 0; i < M; i++) {
        Node* p = table[i];
        while (p) {
            Node* nxt = p->next;
            free(p);
            p = nxt;
        }
    }
    free(table);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
