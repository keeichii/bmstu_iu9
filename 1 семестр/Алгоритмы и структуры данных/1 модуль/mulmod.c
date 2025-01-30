
#include <stdio.h>
#include <stdint.h>
#include <inttypes.h>

uint64_t add_mod(uint64_t x, uint64_t y, uint64_t m) {
    uint64_t sum = x + y;
    if (sum < x) {
        sum -= m;
    }
    if (sum >= m) {
        sum -= m;
    }
    return sum;
}

uint64_t mul_mod(uint64_t a, uint64_t b, uint64_t m) {
    uint64_t result = 0;
    a %= m;
    while (b > 0) {
        if (b & 1) {
            result = add_mod(result, a, m);
        }
        a = add_mod(a, a, m);
        b >>= 1;
    }
    return result;
}

int main(void) {
    uint64_t a, b, m;
    if (scanf("%" SCNu64 " %" SCNu64 " %" SCNu64, &a, &b, &m) != 3) {
        return 1;
    }
    printf("%" PRIu64 "\n", mul_mod(a, b, m));
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
