#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
    int Year,Month,Day;
} Date;

void stableCountingSort(Date *arr,int n,int range,int (*getKey)(const Date*)){
    int *count=calloc(range,sizeof(int));
    Date *output=malloc(n*sizeof(Date));
    for(int i=0;i<n;i++){
        count[getKey(&arr[i])]++;
    }
    for(int i=1;i<range;i++){
        count[i]+=count[i-1];
    }
    for(int i=n-1;i>=0;i--){
        int idx=getKey(&arr[i]);
        output[count[idx]-1]=arr[i];
        count[idx]--;
    }
    for(int i=0;i<n;i++){
        arr[i]=output[i];
    }
    free(count);
    free(output);
}

int getDay(const Date *d){
    return d->Day-1;
}

int getMonth(const Date *d){
    return d->Month-1;
}

int getYear(const Date *d){
    return d->Year-1970;
}

int main(){
    int n;
    scanf("%d",&n);
    Date *arr=malloc(n*sizeof(Date));
    for(int i=0;i<n;i++){
        scanf("%d%d%d",&arr[i].Year,&arr[i].Month,&arr[i].Day);
    }
    stableCountingSort(arr,n,31,getDay);
    stableCountingSort(arr,n,12,getMonth);
    stableCountingSort(arr,n,61,getYear);
    for(int i=0;i<n;i++){
        printf("%04d %02d %02d\n",arr[i].Year,arr[i].Month,arr[i].Day);
    }
    free(arr);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
