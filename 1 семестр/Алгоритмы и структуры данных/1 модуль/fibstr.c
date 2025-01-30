
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *fibstr(int n){
    if(n==1){
        char *r=malloc(2);
        strcpy(r,"a");
        return r;
    }
    if(n==2){
        char *r=malloc(2);
        strcpy(r,"b");
        return r;
    }
    char *s1=malloc(2);
    strcpy(s1,"a");
    char *s2=malloc(2);
    strcpy(s2,"b");
    for(int i=3;i<=n;i++){
        int len=strlen(s1)+strlen(s2)+1;
        char *s3=malloc(len);
        strcpy(s3,s1);
        strcat(s3,s2);
        free(s1);
        s1=s2;
        s2=s3;
    }
    free(s1);
    return s2;
}

int main(void){
    int n;
    scanf("%d",&n);
    char *res=fibstr(n);
    printf("%s\n",res);
    free(res);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
