#include <stdio.h>
#include <stdlib.h>

struct Node {
    long value;
    struct Node *next;
};

long isPowerOfTwo(long num) {
    return (num > 0) && ((num & (num - 1)) == 0);
}

long countPower2Combinations(struct Node *current, long sum) {
    if (current == NULL) {
        return 0;
    }

    long count = 0;
    if (isPowerOfTwo(current->value + sum)) {
        count = 1;
    }

    count += countPower2Combinations(current->next, sum + current->value);
    count += countPower2Combinations(current->next, sum);

    return count;
}

int main() {
    int n;
    scanf("%d", &n);

    struct Node *head = malloc(sizeof(struct Node));
    struct Node *temp = head;
    for (int i = 0; i < n; i++) {
        temp->next = malloc(sizeof(struct Node));
        temp = temp->next;
        scanf("%ld", &temp->value);
    }
    temp->next = NULL;

    long result = countPower2Combinations(head->next, 0);
    printf("%ld\n", result);

    while (head != NULL) {
        struct Node *nextNode = head->next;
        free(head);
        head = nextNode;
    }

    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Вдохновлялись работой Ивана Кужеля, но здесь зачесть можно.
