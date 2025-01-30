
#include <stdio.h>

#define SIZE 8

int are_equal(int a[], int b[], int size) {
    for (int i = 0; i < size; i++) {
        if (a[i] != b[i]) {
            return 0;
        }
    }
    return 1;
}

void swap(int *x, int *y) {
    int temp = *x;
    *x = *y;
    *y = temp;
}

int permute(int a[], int b[], int l, int r) {
    if (l == r) {
        if (are_equal(a, b, SIZE)) {
            return 1;
        }
    } else {
        for (int i = l; i <= r; i++) {
            swap(&b[l], &b[i]);
            if (permute(a, b, l + 1, r)) {
                return 1;
            }
            swap(&b[l], &b[i]); 
        }
    }
    return 0;
}

int main() {
    int a[SIZE], b[SIZE];

    for (int i = 0; i < SIZE; i++) {
        scanf("%d", &a[i]);
    }
    for (int i = 0; i < SIZE; i++) {
        scanf("%d", &b[i]);
    }

    if (permute(a, b, 0, SIZE - 1)) {
        printf("yes\n");
    } else {
        printf("no\n");
    }

    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Подозрительно высокое сходство с другой работой, но можно зачесть.
