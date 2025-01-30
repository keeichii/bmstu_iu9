#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Word {
    int start;
    int end;
};

struct Node {
    struct Word word;
    struct Node *next;
    int index;
};

void csort(char *input, char *output) {
    struct Node *head = malloc(sizeof(struct Node));
    head->next = NULL;
    head->index = 0;
    struct Node *tail = head;
    int isSpace = 1, length = strlen(input);

    for (int i = 0; i < length; i++) {
        if (isSpace && input[i] != ' ') {
            tail->next = malloc(sizeof(struct Node));
            tail->next->index = tail->index + 1;
            tail = tail->next;
            tail->next = NULL;
            tail->word.start = i;
            tail->word.end = length - 1;
            isSpace = 0;
        }
        if (!isSpace && (input[i] == ' ' || input[i] == '\n' || input[i] == '\0')) {
            tail->word.end = i - 1;
            isSpace = 1;
        }
    }

    int wordCount = tail->index;
    int positions[wordCount];
    struct Node *current = head->next;
    
    while (current != NULL) {
        positions[current->index - 1] = 0;
        struct Node *temp = head->next;
        while (temp != current) {
            int currentWordLength = current->word.end - current->word.start + 1;
            int tempWordLength = temp->word.end - temp->word.start + 1;
            if (currentWordLength < tempWordLength) {
                positions[current->index - 1]++;
            } else {
                positions[temp->index - 1]++;
            }
            temp = temp->next;
        }
        current = current->next;
    }

    struct Word sortedWords[wordCount];
    current = head->next;
    while (current != NULL) {
        sortedWords[positions[current->index - 1]] = current->word;
        current = current->next;
    }

    int idx = 0;
    for (int i = wordCount - 1; i >= 0; i--) {
        for (int j = sortedWords[i].start; j <= sortedWords[i].end; j++) {
            output[idx++] = input[j];
        }
        output[idx++] = ' ';
    }
    output[idx - 1] = '\0';

    while (head != NULL) {
        struct Node *nextNode = head->next;
        free(head);
        head = nextNode;
    }
    free(head);
}

int main() {
    char *input = calloc(1000, sizeof(char));
    fgets(input, 1000, stdin);

    char *output = calloc(1000, sizeof(char));
    csort(input, output);

    printf("%s", output);

    free(output);
    free(input);
    return 0;
}

//Антиплагиат: Найдены очень похожие посылки
//Комментарий преподавателя:
