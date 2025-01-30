#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int compare(const void *a, const void *b) {
    const char *str1 = *(const char **)a;
    const char *str2 = *(const char **)b;

    int count_a1 = 0, count_a2 = 0;
    while (*str1) {
        if (*str1 == 'a') count_a1++;
        str1++;
    }
    while (*str2) {
        if (*str2 == 'a') count_a2++;
        str2++;
    }

    if (count_a1 != count_a2) {
        return count_a1 - count_a2;
    }
    
    return strcmp(*(const char **)a, *(const char **)b);
}

void hsort(void *base, size_t nel, size_t width, int (*compare)(const void *a, const void *b)) {
    size_t last_parent = nel / 2 - 1;

    for (size_t i = last_parent; i < nel; i--) {
        size_t parent = i;
        while (1) {
            size_t left = 2 * parent + 1;
            size_t right = 2 * parent + 2;
            size_t largest = parent;

            if (left < nel && compare((char *)base + left * width, (char *)base + largest * width) > 0) {
                largest = left;
            }
            if (right < nel && compare((char *)base + right * width, (char *)base + largest * width) > 0) {
                largest = right;
            }

            if (largest == parent) {
                break;
            }

            void *temp = malloc(width);
            memcpy(temp, (char *)base + parent * width, width);
            memcpy((char *)base + parent * width, (char *)base + largest * width, width);
            memcpy((char *)base + largest * width, temp, width);
            free(temp);

            parent = largest;
        }
    }

    for (size_t i = nel; i > 1; i--) {
        void *temp = malloc(width);
        memcpy(temp, base, width);
        memcpy(base, (char *)base + (i - 1) * width, width);
        memcpy((char *)base + (i - 1) * width, temp, width);
        free(temp);

        size_t parent = 0;
        size_t remaining = i - 1;
        while (1) {
            size_t left = 2 * parent + 1;
            size_t right = 2 * parent + 2;
            size_t largest = parent;

            if (left < remaining && compare((char *)base + left * width, (char *)base + largest * width) > 0) {
                largest = left;
            }
            if (right < remaining && compare((char *)base + right * width, (char *)base + largest * width) > 0) {
                largest = right;
            }

            if (largest == parent) {
                break;
            }

            temp = malloc(width);
            memcpy(temp, (char *)base + parent * width, width);
            memcpy((char *)base + parent * width, (char *)base + largest * width, width);
            memcpy((char *)base + largest * width, temp, width);
            free(temp);

            parent = largest;
        }
    }
}

int main() {
    size_t word_count;
    scanf("%zu", &word_count);
    getchar();

    char **words = malloc(word_count * sizeof(char *));
    for (size_t i = 0; i < word_count; i++) {
        words[i] = malloc(1000);
        fgets(words[i], 1000, stdin);
        words[i][strcspn(words[i], "\n")] = '\0';
    }

    hsort(words, word_count, sizeof(char *), compare);

    for (size_t i = 0; i < word_count; i++) {
        puts(words[i]);
        free(words[i]);
    }

    free(words);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
