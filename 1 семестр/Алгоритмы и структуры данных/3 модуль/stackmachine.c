#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct StackNode {
    long value;
    struct StackNode *previous;
};

struct StackNode *createStackNode(long value, struct StackNode *previous) {
    struct StackNode *newNode = malloc(sizeof(struct StackNode));
    newNode->value = value;
    newNode->previous = previous;
    return newNode;
}

int compareStrings(char *a, char *b) {
    return (a[0] == b[0]) && (a[1] == b[1]);
}

void stackmachine() {
    struct StackNode *stack = NULL;
    char *instruction = calloc(6, sizeof(char));
    long operand;
    scanf("%s", instruction);

    while (!compareStrings(instruction, "END")) {
        if (compareStrings(instruction, "CONST")) {
            scanf("%ld", &operand);
            stack = createStackNode(operand, stack);
            scanf("%s", instruction);
            continue;
        }
        if (compareStrings(instruction, "ADD")) {
            struct StackNode *first = stack;
            struct StackNode *second = stack->previous;
            long result = first->value + second->value;
            stack = second;
            second->value = result;
            stack->previous = second->previous;
            free(first);
            scanf("%s", instruction);
            continue;
        }
        if (compareStrings(instruction, "SUB")) {
            struct StackNode *first = stack;
            struct StackNode *second = stack->previous;
            long result = first->value - second->value;
            stack = second;
            second->value = result;
            stack->previous = second->previous;
            free(first);
            scanf("%s", instruction);
            continue;
        }
        if (compareStrings(instruction, "MUL")) {
            struct StackNode *first = stack;
            struct StackNode *second = stack->previous;
            long result = first->value * second->value;
            stack = second;
            second->value = result;
            stack->previous = second->previous;
            free(first);
            scanf("%s", instruction);
            continue;
        }
        if (compareStrings(instruction, "DIV")) {
            struct StackNode *first = stack;
            struct StackNode *second = stack->previous;
            long result = first->value / second->value;
            stack = second;
            second->value = result;
            stack->previous = second->previous;
            free(first);
            scanf("%s", instruction);
            continue;
        }
        if (compareStrings(instruction, "MAX")) {
            struct StackNode *first = stack;
            struct StackNode *second = stack->previous;
            long result = (first->value > second->value) ? first->value : second->value;
            stack = second;
            second->value = result;
            stack->previous = second->previous;
            free(first);
            scanf("%s", instruction);
            continue;
        }
        if (compareStrings(instruction, "MIN")) {
            struct StackNode *first = stack;
            struct StackNode *second = stack->previous;
            long result = (first->value < second->value) ? first->value : second->value;
            stack = second;
            second->value = result;
            stack->previous = second->previous;
            free(first);
            scanf("%s", instruction);
            continue;
        }
        if (compareStrings(instruction, "NEG")) {
            stack->value = -stack->value;
            scanf("%s", instruction);
            continue;
        }
        if (compareStrings(instruction, "DUP")) {
            stack = createStackNode(stack->value, stack);
            scanf("%s", instruction);
            continue;
        }
        if (compareStrings(instruction, "SWAP")) {
            struct StackNode *first = stack;
            struct StackNode *second = stack->previous;
            first->previous = second->previous;
            second->previous = first;
            stack = second;
            scanf("%s", instruction);
            continue;
        }
    }

    printf("%ld", stack->value);
    while (stack != NULL) {
        struct StackNode *previousNode = stack->previous;
        free(stack);
        stack = previousNode;
    }
    free(instruction);
}

int main() {
    stackmachine();
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Засчитываю, как и обещал.
