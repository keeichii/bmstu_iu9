#include <stdio.h>
#include <string.h>
#include <limits.h>

#define MAXN 10
char s[MAXN][1001];
int overlap[MAXN][MAXN], dp[1 << MAXN][MAXN], len[MAXN];

int main() {
    int n;
    scanf("%d", &n);
    for (int i = 0; i < n; i++) {
        scanf("%s", s[i]);
    }
    for (int i = 0; i < n; i++) {
        len[i] = strlen(s[i]);
    }
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            overlap[i][j] = 0;
            for (int k = 1; k <= len[i] && k <= len[j]; k++) {
                if (strncmp(s[i] + len[i] - k, s[j], k) == 0) {
                    overlap[i][j] = k;
                }
            }
        }
    }
    for (int mask = 0; mask < (1 << n); mask++) {
        for (int i = 0; i < n; i++) {
            dp[mask][i] = INT_MAX;
        }
    }
    for (int i = 0; i < n; i++) {
        dp[1 << i][i] = len[i];
    }
    for (int mask = 0; mask < (1 << n); mask++) {
        for (int i = 0; i < n; i++) {
            if (dp[mask][i] == INT_MAX) continue;
            for (int j = 0; j < n; j++) {
                if (mask & (1 << j)) continue;
                int nm = mask | (1 << j);
                int cost = dp[mask][i] + len[j] - overlap[i][j];
                if (cost < dp[nm][j]) {
                    dp[nm][j] = cost;
                }
            }
        }
    }
    int ans = INT_MAX;
    for (int i = 0; i < n; i++) {
        if (dp[(1 << n) - 1][i] < ans) {
            ans = dp[(1 << n) - 1][i];
        }
    }
    printf("%d\n", ans);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
