#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAXLEVEL 32

static double frand() {
    return (double)rand() / RAND_MAX;
}

static int randomLevel() {
    int lvl = 1;
    while (frand() < 0.5 && lvl < MAXLEVEL) lvl++;
    return lvl;
}

typedef struct Node {
    int key;
    char *val;
    struct Node **next;
    int *span;
    int level;
} Node;

static Node *head;
static int maxLevel = 1;

static Node *createNode(int level, int k, const char *v) {
    Node *node = malloc(sizeof(Node));
    node->key = k;
    node->val = malloc(strlen(v) + 1);
    strcpy(node->val, v);
    node->level = level;
    node->next = calloc(level, sizeof(Node *));
    node->span = calloc(level, sizeof(int));
    return node;
}

static void init() {
    head = createNode(MAXLEVEL, -2147483647, "");
    for (int i = 0; i < MAXLEVEL; i++) {
        head->next[i] = NULL;
        head->span[i] = 0;
    }
    maxLevel = 1;
}

static void insertNode(int k, const char *v) {
    Node *update[MAXLEVEL];
    int rank[MAXLEVEL];
    Node *x = head;
    for (int i = maxLevel - 1; i >= 0; i--) {
        rank[i] = (i == maxLevel - 1 ? 0 : rank[i + 1]);
        while (x->next[i] && x->next[i]->key < k) {
            rank[i] += x->span[i];
            x = x->next[i];
        }
        update[i] = x;
    }
    int lvl = randomLevel();
    if (lvl > maxLevel) {
        for (int i = maxLevel; i < lvl; i++) {
            update[i] = head;
            head->span[i] = 0;
            rank[i] = 0;
        }
        maxLevel = lvl;
    }
    x = createNode(lvl, k, v);
    for (int i = 0; i < lvl; i++) {
        x->next[i] = update[i]->next[i];
        update[i]->next[i] = x;
        x->span[i] = update[i]->span[i] - (rank[0] - rank[i]);
        update[i]->span[i] = (rank[0] - rank[i]) + 1;
    }
    for (int i = lvl; i < maxLevel; i++) {
        update[i]->span[i]++;
    }
}

static Node *searchNode(int k) {
    Node *x = head;
    for (int i = maxLevel - 1; i >= 0; i--) {
        while (x->next[i] && x->next[i]->key < k) {
            x = x->next[i];
        }
    }
    x = x->next[0];
    if (x && x->key == k) return x;
    return NULL;
}

static void deleteNode(int k) {
    Node *update[MAXLEVEL];
    Node *x = head;
    for (int i = maxLevel - 1; i >= 0; i--) {
        while (x->next[i] && x->next[i]->key < k) {
            x = x->next[i];
        }
        update[i] = x;
    }
    x = x->next[0];
    if (!x || x->key != k) return;
    for (int i = 0; i < maxLevel; i++) {
        if (update[i]->next[i] == x) {
            update[i]->span[i] += x->span[i] - 1;
            update[i]->next[i] = x->next[i];
        } else {
            update[i]->span[i]--;
        }
    }
    while (maxLevel > 1 && !head->next[maxLevel - 1]) {
        maxLevel--;
    }
    free(x->val);
    free(x->next);
    free(x->span);
    free(x);
}

static int rankOfKey(int k) {
    int res = 0;
    Node *x = head;
    for (int i = maxLevel - 1; i >= 0; i--) {
        while (x->next[i] && x->next[i]->key < k) {
            res += x->span[i];
            x = x->next[i];
        }
    }
    x = x->next[0];
    if (!x || x->key != k) return -1;
    return res;
}

int main() {
    srand(1234567);
    init();
    while (1) {
        char cmd[16];
        if (scanf("%s", cmd) != 1) break;
        if (strcmp(cmd, "END") == 0) break;
        if (strcmp(cmd, "INSERT") == 0) {
            int k;
            char buf[32];
            scanf("%d %s", &k, buf);
            insertNode(k, buf);
        } else if (strcmp(cmd, "LOOKUP") == 0) {
            int k;
            scanf("%d", &k);
            Node *x = searchNode(k);
            if (!x) printf("?\n");
            else printf("%s\n", x->val);
        } else if (strcmp(cmd, "DELETE") == 0) {
            int k;
            scanf("%d", &k);
            deleteNode(k);
        } else if (strcmp(cmd, "RANK") == 0) {
            int k;
            scanf("%d", &k);
            printf("%d\n", rankOfKey(k));
        }
    }
    Node *cur = head->next[0];
    while (cur) {
        Node *ttt = cur->next[0];
        free(cur->val);
        free(cur->next);
        free(cur->span);
        free(cur);
        cur = ttt;
    }
    free(head->val);
    free(head->next);
    free(head->span);
    free(head);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
