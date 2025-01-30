#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static void buildExtendedStopTable(const char *S, int lenS, int *d1) {
    for(int i=0; i<lenS; i++){
        for(int c=0; c<256; c++){
            d1[i*256 + c] = i + 1;
        }
        for(int j=0; j<i; j++){
            unsigned char cc = (unsigned char)S[j];
            d1[i*256 + cc] = i - j;
        }
    }
}

int main(int argc, char *argv[]){
    if(argc != 3) return 0;
    char *S = argv[1], *T = argv[2];
    int lenS = (int)strlen(S), lenT = (int)strlen(T);
    if(lenS == 0){
        printf("0\n");
        return 0;
    }
    if(lenT == 0){
        printf("%d\n", lenT);
        return 0;
    }
    int *d1 = malloc(lenS * 256 * sizeof(int));
    if(!d1) return 0;
    buildExtendedStopTable(S, lenS, d1);
    int i = 0;
    while(i <= lenT - lenS){
        int j = lenS - 1;
        while(j >= 0 && T[i + j] == S[j]){
            j--;
        }
        if(j < 0){
            printf("%d\n", i);
            free(d1);
            return 0;
        } else {
            unsigned char c = (unsigned char)T[i + j];
            int shift = d1[j*256 + c];
            if(shift < 1) shift = 1;
            i += shift;
        }
    }
    printf("%d\n", lenT);
    free(d1);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
