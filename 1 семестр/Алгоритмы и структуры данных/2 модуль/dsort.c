#include <stdio.h>

int main() {
    int counts[26] = {0};
    char c;

    while ((c = getchar()) != '\n' && c != EOF) {
        counts[c - 'a']++;
    }

    for (int i = 0; i < 26; i++) {
        while (counts[i]-- > 0) {
            putchar('a' + i);
        }
    }

    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
