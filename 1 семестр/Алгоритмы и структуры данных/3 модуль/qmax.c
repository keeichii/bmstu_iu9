#include <stdio.h>
#include <string.h>

#define CAP 2000000

typedef struct {
    long long val, mx;
} Item;

static Item stPush[CAP], stPop[CAP];
static int topPush = 0, topPop = 0;

void pushPush(long long x) {
    if (topPush == 0) {
        stPush[topPush].val = x;
        stPush[topPush].mx = x;
    } else {
        stPush[topPush].val = x;
        long long m = stPush[topPush - 1].mx;
        stPush[topPush].mx = (x > m ? x : m);
    }
    topPush++;
}

long long popPush() {
    topPush--;
    return stPush[topPush].val;
}

long long maxPush() {
    return (topPush > 0 ? stPush[topPush - 1].mx : -2000000000LL);
}

void pushPop(long long x) {
    if (topPop == 0) {
        stPop[topPop].val = x;
        stPop[topPop].mx = x;
    } else {
        stPop[topPop].val = x;
        long long m = stPop[topPop - 1].mx;
        stPop[topPop].mx = (x > m ? x : m);
    }
    topPop++;
}

long long popPop() {
    topPop--;
    return stPop[topPop].val;
}

long long maxPop() {
    return (topPop > 0 ? stPop[topPop - 1].mx : -2000000000LL);
}

long long queueMax() {
    long long mp = -2000000000LL, mp2 = -2000000000LL;
    if (topPush > 0) mp = maxPush();
    if (topPop > 0) mp2 = maxPop();
    return (mp > mp2 ? mp : mp2);
}

int main() {
    while (1) {
        char cmd[16];
        if (scanf("%s", cmd) != 1) break;
        if (strcmp(cmd, "END") == 0) break;
        if (strcmp(cmd, "InitQueue") == 0) {
            topPush = 0;
            topPop = 0;
        } else if (strcmp(cmd, "ENQ") == 0) {
            long long x;
            scanf("%lld", &x);
            pushPush(x);
        } else if (strcmp(cmd, "DEQ") == 0) {
            if (topPop == 0) {
                while (topPush > 0) {
                    pushPop(popPush());
                }
            }
            long long val = popPop();
            printf("%lld\n", val);
        } else if (strcmp(cmd, "MAX") == 0) {
            printf("%lld\n", queueMax());
        } else if (strcmp(cmd, "EMPTY") == 0) {
            if (topPush == 0 && topPop == 0) printf("true\n");
            else printf("false\n");
        }
    }
    return 0;
}

//Антиплагиат: Похожих посылок не найдено
//Комментарий преподавателя:
