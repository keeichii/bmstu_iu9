#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int low, high;
} Task;

static void swap(int *a,int *b){
    int t=*a;*a=*b;*b=t;
}

int main(){
    int n;
    scanf("%d",&n);
    int *arr=malloc(n*sizeof(int));
    for(int i=0;i<n;i++){
        scanf("%d",&arr[i]);
    }
    if(n>1){
        Task *stack=malloc(n*sizeof(Task));
        int top=-1;
        stack[++top]=(Task){0,n-1};
        while(top>=0){
            Task t=stack[top--];
            int l=t.low,r=t.high;
            while(l<r){
                int m=(l+r)/2;
                int pivot=arr[m];
                int i=l,j=r;
                while(i<=j){
                    while(arr[i]<pivot)i++;
                    while(arr[j]>pivot)j--;
                    if(i<=j){
                        swap(&arr[i],&arr[j]);
                        i++;j--;
                    }
                }
                if(j-l<r-i){
                    if(l<j) stack[++top]=(Task){l,j};
                    l=i;
                } else {
                    if(i<r) stack[++top]=(Task){i,r};
                    r=j;
                }
            }
        }
        free(stack);
    }
    for(int i=0;i<n;i++){
        printf("%d ",arr[i]);
    }
    printf("\n");
    free(arr);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Всё красиво.
