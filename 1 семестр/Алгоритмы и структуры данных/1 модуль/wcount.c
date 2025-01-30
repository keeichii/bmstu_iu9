


#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int wcount(char *s) {
    int count = 0;
    int in_word = 0;

    for (int i = 0; s[i] != '\0'; i++) {
        if (s[i] != ' ') {
            if (!in_word) {
                in_word = 1;
                count++;
            }
        } else {
            in_word = 0;
        }
    }

    return count;
}

int main() {
    char buffer[1000];

    // Пример с пословицей
    const char *proverb = "Без труда не вытащишь и рыбку из пруда.";
    //printf("пословица: %s\n", proverb);
    int word_count = wcount((char *)proverb);
    //printf("кол-во слов: %d\n", word_count);

    if (fgets(buffer, sizeof(buffer), stdin) != NULL) {
        size_t len = strlen(buffer);
        if (len > 0 && buffer[len - 1] == '\n') {
            buffer[len - 1] = '\0';
        }

        int input_word_count = wcount(buffer);
        printf("%d\n", input_word_count);
    }

    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
