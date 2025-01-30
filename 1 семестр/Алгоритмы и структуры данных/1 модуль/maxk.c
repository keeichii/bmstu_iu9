
#include <stdio.h>

int main() {
    int n, k, i, sum = 0, max_sum = 0;

    scanf("%d", &n);
    int arr[n];

    for (i = 0; i < n; i++) {
        scanf("%d", &arr[i]);
    }

    scanf("%d", &k);
    
    if (k > n || n <= 0 || k <= 0) {
        printf("error\n");
        return 1;
    }

    for (i = 0; i < k; i++) {
        sum += arr[i];
    }

    max_sum = sum;

    for (i = k; i < n; i++) {
        sum = sum - arr[i - k] + arr[i];
        if (sum > max_sum) {
            max_sum = sum;
        }
    }

    printf("%d\n", max_sum);

    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
