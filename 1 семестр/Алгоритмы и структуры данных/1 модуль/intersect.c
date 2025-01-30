
#include <stdio.h>

int main() {
    int nA, nB, element;
    unsigned int setA = 0, setB = 0, intersection;
    
    scanf("%d", &nA);
    for (int i = 0; i < nA; i++) {
        scanf("%d", &element);
        setA |= (1 << element); 
    }

    scanf("%d", &nB);
    for (int i = 0; i < nB; i++) {
        scanf("%d", &element);
        setB |= (1 << element); 
    }

    intersection = setA & setB; 

    for (int i = 0; i < 32; i++) {
        if (intersection & (1 << i)) {
            printf("%d ", i); 
        }
    }

    printf("\n");

    return 0;
}

//Антиплагиат: Найдены очень похожие посылки
//Комментарий преподавателя: В этой задаче бывает подозрительно высокое сходство.
