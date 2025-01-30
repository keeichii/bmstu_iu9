#include <stdio.h>
#include <stdlib.h>

union Int32 {
    int x;
    unsigned char bytes[4];
};

void countingSort(union Int32 *arr,int n,int pass){
    int count[256]={0};
    union Int32 *output=malloc(sizeof(union Int32)*n);
    for(int i=0;i<n;i++){
        unsigned char key=arr[i].bytes[pass];
        if(pass==3) key^=0x80;
        count[key]++;
    }
    for(int i=1;i<256;i++){
        count[i]+=count[i-1];
    }
    for(int i=n-1;i>=0;i--){
        unsigned char key=arr[i].bytes[pass];
        if(pass==3) key^=0x80;
        output[--count[key]]=arr[i];
    }
    for(int i=0;i<n;i++){
        arr[i]=output[i];
    }
    free(output);
}

int main(){
    int n;
    scanf("%d",&n);
    union Int32 *arr=malloc(n*sizeof(union Int32));
    for(int i=0;i<n;i++){
        scanf("%d",&arr[i].x);
    }
    for(int pass=0;pass<4;pass++){
        countingSort(arr,n,pass);
    }
    for(int i=0;i<n;i++){
        printf("%d ",arr[i].x);
    }
    printf("\n");
    free(arr);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
