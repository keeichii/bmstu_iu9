#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[]){
    if(argc!=3){
        return 0;
    }
    char *S=argv[1], *T=argv[2];
    int n=strlen(S), m=strlen(T);
    int *pi=malloc(n*sizeof(int));
    pi[0]=0;
    for(int i=1;i<n;i++){
        int k=pi[i-1];
        while(k>0 && S[i]!=S[k]){
            k=pi[k-1];
        }
        if(S[i]==S[k]){
            k++;
        }
        pi[i]=k;
    }
    int j=0;
    for(int i=0;i<m;i++){
        while(j>0 && T[i]!=S[j]){
            j=pi[j-1];
        }
        if(T[i]==S[j]){
            j++;
        }
        if(j==n){
            printf("%d ",i-n+1);
            j=pi[j-1];
        }
    }
    printf("\n");
    free(pi);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
