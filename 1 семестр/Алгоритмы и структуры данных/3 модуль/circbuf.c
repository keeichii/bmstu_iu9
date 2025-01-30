#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    long front, length, capacity, *buffer;
} CircularQueue;

int isEqual(const char *a, const char *b) {
    return strcmp(a, b) == 0;
}

CircularQueue* initQueue() {
    CircularQueue *q = malloc(sizeof(CircularQueue));
    q->length = 0;
    q->front = 0;
    q->capacity = 4;
    q->buffer = malloc(q->capacity * sizeof(long));
    return q;
}

void resizeQueue(CircularQueue *q) {
    long *newBuffer = malloc(q->capacity * 2 * sizeof(long));
    for (int i = 0; i < q->length; i++) {
        newBuffer[i] = q->buffer[(q->front + i) % q->capacity];
    }
    free(q->buffer);
    q->buffer = newBuffer;
    q->front = 0;
    q->capacity *= 2;
}

void enqueue(CircularQueue *q, long value) {
    if (q->length == q->capacity) {
        resizeQueue(q);
    }
    q->buffer[(q->front + q->length) % q->capacity] = value;
    q->length++;
}

long dequeue(CircularQueue *q) {
    long value = q->buffer[q->front];
    q->front = (q->front + 1) % q->capacity;
    q->length--;
    return value;
}

int isEmpty(CircularQueue *q) {
    return q->length == 0;
}

void freeQueue(CircularQueue *q) {
    free(q->buffer);
    free(q);
}

int main() {
    CircularQueue *queue = initQueue();
    
    char command[6];
    
    while (1) {
        scanf("%s", command);
        
        if (isEqual(command, "END")) break;

        if (isEqual(command, "ENQ")) {
            long value;
            scanf("%ld", &value);
            enqueue(queue, value);
        } else if (isEqual(command, "DEQ")) {
            printf("%ld\n", dequeue(queue));
        } else if (isEqual(command, "EMPTY")) {
            printf(isEmpty(queue) ? "true\n" : "false\n");
        }
    }

    freeQueue(queue);
    
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Минус балл за подозрительное сходство с другими решениями.
