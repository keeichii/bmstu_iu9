
#include <stdio.h>

int main() {
    int d;
    long long x;
    scanf("%d", &d);
    scanf("%lld", &x);
    
    long long coef[d + 1];
    
    for (int i = 0; i <= d; i++) {
        scanf("%lld", &coef[i]);
    }
    
    long long poly = coef[0];
    long long derivative = 0;

    for (int i = 1; i <= d; i++) {
        derivative = derivative * x + poly;
        poly = poly * x + coef[i];
    }

    printf("%lld\n", poly);
    printf("%lld\n", derivative);
    
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
