#include <stdio.h>
#include <string.h>

int main(int argc,char *argv[]){
    if(argc!=3)return 0;
    char *S=argv[1],*T=argv[2];
    int n=strlen(S),m=strlen(T);
    if(!n||!m) return 0;
    int bc[256];
    for(int i=0;i<256;i++){
        bc[i]=n;
    }
    for(int i=0;i<n-1;i++){
        bc[(unsigned char)S[i]]=n-1-i;
    }
    int i=0;
    while(i<=m-n){
        int j=n-1;
        while(j>=0&&T[i+j]==S[j])j--;
        if(j<0){
            printf("%d ",i);
            i++;
        } else {
            int shift=bc[(unsigned char)T[i+j]]-(n-1-j);
            if(shift<1) shift=1;
            i+=shift;
        }
    }
    printf("\n");
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Упрощённый алгоритм — 1 балл, не заблокировал.
