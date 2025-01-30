void shellsort(unsigned long nel,
        int (*compare)(unsigned long i, unsigned long j),
        void (*swap)(unsigned long i, unsigned long j))
{
    if (nel < 2) return;
    unsigned long fib[nel];
    unsigned long prev = 1, cur = 1, count = 0;
    while (prev < nel) {
        fib[count++] = prev;
        unsigned long next = prev + cur;
        prev = cur;
        cur = next;
    }
    for (long k = count - 1; k >= 0; k--) {
        unsigned long d = fib[k];
        for (unsigned long i = d; i < nel; i++) {
            for (unsigned long j = i; j >= d && compare(j - d, j) > 0; j -= d) {
                swap(j - d, j);
            }
        }
    }
}

//Антиплагиат: Найдены очень похожие посылки
//Комментарий преподавателя: 👍
