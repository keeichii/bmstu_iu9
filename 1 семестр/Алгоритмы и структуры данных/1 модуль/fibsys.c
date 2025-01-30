
#include <stdio.h>

void generate_fibonacci(unsigned long long fib[], int max_fib) {
    fib[0] = 1;
    fib[1] = 2;
    for (int i = 2; i < max_fib; i++) {
        fib[i] = fib[i - 1] + fib[i - 2];
    }
}

void fib_sys(unsigned long long x) {
    if (x == 0) {
        printf("0\n");
        return;
    }

    int max_fib = 90;
    unsigned long long fib[max_fib];
    generate_fibonacci(fib, max_fib);

    int used[max_fib];
    for (int i = 0; i < max_fib; i++) {
        used[i] = 0;
    }

    int i = max_fib - 1;

    while (x > 0) {
        while (fib[i] > x) {
            i--;
        }
        used[i] = 1;
        x -= fib[i];
    }

    int started = 0;
    for (int j = max_fib - 1; j >= 0; j--) {
        if (used[j]) {
            started = 1;
        }
        if (started) {
            printf("%d", used[j]);
        }
    }
    printf("\n");
}

int main() {
    unsigned long long x;
    scanf("%llu", &x);
    fib_sys(x);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
