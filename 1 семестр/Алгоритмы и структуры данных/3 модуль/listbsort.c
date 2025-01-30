#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Node {
    struct Node *next;
    char *word;
} Node;

Node* bsort(Node *list) {
    int swapped;
    Node *ptr1;
    Node *lptr = NULL;

    if (list == NULL) return NULL;

    do {
        swapped = 0;
        ptr1 = list;

        while (ptr1->next != lptr) {
            if (strlen(ptr1->word) > strlen(ptr1->next->word)) {
                char *temp = ptr1->word;
                ptr1->word = ptr1->next->word;
                ptr1->next->word = temp;
                swapped = 1;
            }
            ptr1 = ptr1->next;
        }
        lptr = ptr1;
    } while (swapped);

    return list;
}

char* readNextWord() {
    char *buffer = calloc(1000, sizeof(char));
    int index = 0;
    char ch = getc(stdin);
    
    while (ch == ' ') ch = getc(stdin);
    
    while (ch != '\n' && ch != EOF && ch != ' ') {
        buffer[index++] = ch;
        ch = getc(stdin);
    }
    
    if (index == 0) {
        free(buffer);
        return NULL;
    }

    buffer[index] = '\0';
    return buffer;
}

int main() {
    Node *head = NULL;
    Node *tail = NULL;

    while (1) {
        char *word = readNextWord();
        
        if (word == NULL) break;

        Node *newNode = malloc(sizeof(Node));
        newNode->next = NULL;
        newNode->word = word;

        if (tail != NULL) {
            tail->next = newNode;
        } else {
            head = newNode;
        }
        
        tail = newNode;
    }

    head = bsort(head);
    
    Node *current = head;
    
    while (current != NULL) {
        printf("%s ", current->word);
        Node *temp = current;
        current = current->next;
        free(temp->word);
        free(temp);
    }

    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Можете ведь решать самостоятельно, зачем списывать?
