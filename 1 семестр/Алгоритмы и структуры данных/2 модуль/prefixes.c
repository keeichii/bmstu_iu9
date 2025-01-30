#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX(a, b) ((a) > (b) ? (a) : (b))

int isPeriodicPrefix(char *str, int periodLength, int endPos) {
    for (int i = 0; i < periodLength; i++) {
        if (str[endPos - i] != str[periodLength - i - 1])
            return 0;
    }
    return 1;
}

int main(int argc, char **argv) {
    char *inputStr = argv[1];
    int strLen = strlen(inputStr);
    int result[strLen + 1];
    memset(result, 0, sizeof(result));

    for (int prefixLen = 1; prefixLen <= strLen / 2; prefixLen++) {
        int startPos = prefixLen + prefixLen - 1;
        while (isPeriodicPrefix(inputStr, prefixLen, startPos) && startPos < strLen) {
            result[startPos + 1] = MAX(result[startPos + 1], (startPos + 1) / prefixLen);
            startPos += prefixLen;
        }
    }

    for (int i = 0; i <= strLen; i++) {
        if (result[i]) {
            printf("%d %d\n", i, result[i]);
        }
    }

    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
