
#include <stdio.h>
#include <math.h>
#include <stdbool.h>
#include <limits.h>
#include <stdlib.h>

#define MAXN 70000

void sieve(bool prime[], int n) {
    for (int i = 2; i <= n; i++) {
        prime[i] = true;
    }
    for (int p = 2; p * p <= n; p++) {
        if (prime[p]) {
            for (int i = p * p; i <= n; i += p) {
                prime[i] = false;
            }
        }
    }
}

int largest_prime_divisor(unsigned long long x) {
    if (x == 0 || x == 1) {
        return -1;
    }

    bool prime[MAXN];
    sieve(prime, MAXN - 1);

    int largest_prime = -1;

    for (int p = 2; p < MAXN; p++) {
        if (prime[p] && x % p == 0) {
            while (x % p == 0) {
                x /= p;
            }
            largest_prime = p;
        }
    }

    if (x > 1) {
        largest_prime = x;
    }

    return largest_prime;
}

int main() {
    long long x;
    scanf("%lld", &x);

    if (x == INT_MIN) {
        printf("2\n");
        return 0;
    }

    x = llabs(x);

    if (x == 0) {
        printf("-1\n");
        return 0;
    }

    int result = largest_prime_divisor(x);

    if (result != -1) {
        printf("%d\n", result);
    } else {
        printf("-1\n");
    }

    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: А это уже другой вариант решения, не похожий на чужую неправильную посылку.
