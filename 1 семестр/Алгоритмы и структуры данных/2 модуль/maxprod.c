#include <stdio.h>
#include <stdlib.h>

#define MAX(a, b) ((a) > (b) ? (a) : (b))

int main() {
    int n;
    scanf("%d", &n);
    
    double *fractions = malloc(n * sizeof(double));
    
    for (int i = 0; i < n; i++) {
        double numerator, denominator;
        scanf("%lf/%lf", &numerator, &denominator);
        fractions[i] = numerator / denominator;
    }

    double maxProduct = fractions[0];
    int startIndex = 0, endIndex = 0;

    for (int start = 0; start < n; start++) {
        double currentProduct = 1.0;
        for (int end = start; end < n; end++) {
            currentProduct *= fractions[end];

            if (currentProduct > maxProduct) {
                maxProduct = currentProduct;
                startIndex = start;
                endIndex = end;
            }
        }
    }

    printf("%d %d\n", startIndex, endIndex);
    
    free(fractions);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Есть подозрительно похожие работы, но порог сходства не пересечён.
