
#define _POSIX_C_SOURCE 200809L
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static char *mystrdup(const char *s){
    char *p=malloc(strlen(s)+1);
    if(p) strcpy(p,s);
    return p;
}

char *concat(char **s, int n) {
    int i, total=0;
    for(i=0;i<n;i++){
        total+=strlen(s[i]);
    }
    char *res=malloc(total+1);
    char *p=res;
    for(i=0;i<n;i++){
        strcpy(p,s[i]);
        p+=strlen(s[i]);
    }
    *p='\0';
    return res;
}

int main(void){
    int n;
    scanf("%d",&n);
    getchar();
    char **arr=malloc(n*sizeof(char*));
    for(int i=0;i<n;i++){
        char buf[1024];
        if(!fgets(buf,1024,stdin)) return 1;
        int len=strlen(buf);
        if(len>0 && buf[len-1]=='\n') buf[len-1]='\0';
        arr[i]=mystrdup(buf);
    }
    char *res=concat(arr,n);
    printf("%s\n",res);
    free(res);
    for(int i=0;i<n;i++){
        free(arr[i]);
    }
    free(arr);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
