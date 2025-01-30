#include <stdio.h>
#include <stdlib.h>

void swap_elements(int *arr, int idx1, int idx2) {
    int temp = arr[idx1];
    arr[idx1] = arr[idx2];
    arr[idx2] = temp;
}

void simple_sort(int *arr, int left, int right) {
    for (int i = left; i < right; i++) {
        int min_idx = i;
        for (int j = i + 1; j <= right; j++) {
            if (arr[j] < arr[min_idx]) {
                min_idx = j;
            }
        }
        if (min_idx != i) {
            swap_elements(arr, i, min_idx);
        }
    }
}

int partition(int *arr, int left, int right) {
    int pivot = arr[(left + right) / 2];
    int i = left - 1, j = right + 1;
    while (1) {
        do { i++; } while (arr[i] < pivot);
        do { j--; } while (arr[j] > pivot);
        if (i >= j) return j;
        swap_elements(arr, i, j);
    }
}

void quick_sort_recursive(int *arr, int left, int right, int threshold) {
    if (left >= right) return;

    if (right - left + 1 < threshold) {
        simple_sort(arr, left, right);
        return;
    }

    int pivot_index = partition(arr, left, right);
    quick_sort_recursive(arr, left, pivot_index, threshold);
    quick_sort_recursive(arr, pivot_index + 1, right, threshold);
}

int main() {
    int n, m;
    scanf("%d %d", &n, &m);
    int *array = (int *)calloc(n, sizeof(int));
    for (int i = 0; i < n; i++) {
        scanf("%d", &array[i]);
    }
    quick_sort_recursive(array, 0, n - 1, m);
    for (int i = 0; i < n; i++) {
        printf("%d ", array[i]);
    }
    free(array);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Не выполнено условие на логарифмическую глубину стека (слайд 77) — 1 балл из 2, не заблокировал.
