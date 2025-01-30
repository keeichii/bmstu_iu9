#include <stdio.h>
#include <stdlib.h>

struct Elem {
    struct Elem *prev, *next;
    int v;
};

int main(){
    int n;
    scanf("%d", &n);
    if (n < 1) {
        return 0;
    }
    struct Elem *head = malloc(sizeof(struct Elem));
    scanf("%d", &head->v);
    head->prev = head;
    head->next = head;
    struct Elem *tail = head;
    for (int i = 1; i < n; i++) {
        struct Elem *node = malloc(sizeof(struct Elem));
        scanf("%d", &node->v);
        node->prev = tail;
        node->next = head;
        tail->next = node;
        head->prev = node;
        tail = node;
    }
    struct Elem *current = head->next;
    while (current != head) {
        int key = current->v;
        struct Elem *pos = current->prev;
        while (pos != tail && pos->v > key) {
            pos->next->v = pos->v;
            pos = pos->prev;
        }
        pos->next->v = key;
        current = current->next;
    }
    struct Elem *p = head;
    for (int i = 0; i < n; i++) {
        printf("%d ", p->v);
        p = p->next;
    }
    printf("\n");
    p = head->next;
    while (p != head) {
        struct Elem *ttt = p;
        p = p->next;
        free(ttt);
    }
    free(head);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
