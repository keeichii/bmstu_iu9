#include <stdio.h>
#include <stdlib.h>
#include <ctype.h>
#include <string.h>

static const char *specList = "+-*/()";

typedef struct AVL {
    char *key;
    int val, height;
    struct AVL *left, *right;
} AVL;

static int avlHeight(AVL *r) {
    return r ? r->height : 0;
}

static int avlBalance(AVL *r) {
    return avlHeight(r->left) - avlHeight(r->right);
}

static void avlFixHeight(AVL *r) {
    int hl = avlHeight(r->left), hr = avlHeight(r->right);
    r->height = (hl > hr ? hl : hr) + 1;
}

static AVL* avlRotateRight(AVL* y) {
    AVL* x = y->left;
    y->left = x->right;
    x->right = y;
    avlFixHeight(y);
    avlFixHeight(x);
    return x;
}

static AVL* avlRotateLeft(AVL* x) {
    AVL* y = x->right;
    x->right = y->left;
    y->left = x;
    avlFixHeight(x);
    avlFixHeight(y);
    return y;
}

static AVL* avlBalanceNode(AVL* r) {
    avlFixHeight(r);
    int bal = avlBalance(r);
    if (bal > 1) {
        if (avlBalance(r->left) < 0) r->left = avlRotateLeft(r->left);
        return avlRotateRight(r);
    } else if (bal < -1) {
        if (avlBalance(r->right) > 0) r->right = avlRotateRight(r->right);
        return avlRotateLeft(r);
    }
    return r;
}

static AVL* avlInsert(AVL* r, const char* key, int val) {
    if (!r) {
        AVL* node = malloc(sizeof(AVL));
        node->key = malloc(strlen(key) + 1);
        strcpy(node->key, key);
        node->val = val;
        node->height = 1;
        node->left = node->right = NULL;
        return node;
    }
    if (strcmp(key, r->key) < 0) {
        r->left = avlInsert(r->left, key, val);
    } else if (strcmp(key, r->key) > 0) {
        r->right = avlInsert(r->right, key, val);
    }
    return avlBalanceNode(r);
}

static int avlLookup(AVL* r, const char* key) {
    while (r) {
        if (strcmp(key, r->key) < 0) {
            r = r->left;
        } else if (strcmp(key, r->key) > 0) {
            r = r->right;
        } else {
            return r->val;
        }
    }
    return -1;
}

static void avlFree(AVL* r) {
    if (!r) return;
    avlFree(r->left);
    avlFree(r->right);
    free(r->key);
    free(r);
}

int main() {
    int n;
    scanf("%d", &n);
    getchar();
    char *line = malloc((n + 2) * sizeof(char));
    fread(line, 1, n, stdin);
    line[n] = '\0';
    AVL *root = NULL;
    int identCount = 0;
    int i = 0;
    while (i < n) {
        while (i < n && isspace((unsigned char)line[i])) i++;
        if (i >= n) break;
        if (isdigit((unsigned char)line[i])) {
            long long val = 0;
            while (i < n && isdigit((unsigned char)line[i])) {
                val = val * 10 + (line[i] - '0');
                i++;
            }
            printf("const %lld\n", val);
            continue;
        }
        if (strchr(specList, line[i])) {
            printf("spec %d\n", (int)(strchr(specList, line[i]) - specList));
            i++;
            continue;
        }
        if (isalpha((unsigned char)line[i])) {
            int start = i;
            i++;
            while (i < n && (isalnum((unsigned char)line[i]))) i++;
            char ttt = line[i];
            line[i] = '\0';
            int found = avlLookup(root, line + start);
            if (found < 0) {
                root = avlInsert(root, line + start, identCount);
                printf("ident %d\n", identCount);
                identCount++;
            } else {
                printf("ident %d\n", found);
            }
            line[i] = ttt;
        }
    }
    free(line);
    avlFree(root);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
