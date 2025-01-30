#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Node {
    int xorv;
    long long count;
    struct Node* next;
} Node;

static Node** table;
static int CAP = 200003;

static unsigned int hsh(int x) {
    unsigned int y = (unsigned int)x;
    y ^= (y >> 16) * 0x45d9f3b;
    y ^= (y >> 16) * 0x45d9f3b;
    return y % CAP;
}

static long long getCount(int x) {
    unsigned int h = hsh(x);
    Node* p = table[h];
    while (p) {
        if (p->xorv == x) return p->count;
        p = p->next;
    }
    return 0;
}

static void incCount(int x, long long c) {
    unsigned int h = hsh(x);
    Node* p = table[h];
    while (p) {
        if (p->xorv == x) {
            p->count += c;
            return;
        }
        p = p->next;
    }
    Node* nn = malloc(sizeof(Node));
    nn->xorv = x;
    nn->count = c;
    nn->next = table[h];
    table[h] = nn;
}

int main() {
    int n;
    scanf("%d", &n);
    table = calloc(CAP, sizeof(Node*));
    long long ans = 0;
    incCount(0, 1);
    int cur = 0;
    for (int i = 0; i < n; i++) {
        int x;
        scanf("%d", &x);
        cur ^= x;
        long long cnt = getCount(cur);
        ans += cnt;
        incCount(cur, 1);
    }
    printf("%lld\n", ans);
    for (int i = 0; i < CAP; i++) {
        Node* p = table[i];
        while (p) {
            Node* nextNode = p->next;
            free(p);
            p = nextNode;
        }
    }
    free(table);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
