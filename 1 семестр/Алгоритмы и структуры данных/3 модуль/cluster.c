#include <stdio.h>
#include <stdlib.h>

static void swaplonglong(long long *a,long long *b){
    long long t=*a;*a=*b;*b=t;
}

#define PARENT(i) (((i)-1)>>1)
#define LEFT(i) (((i)<<1)+1)
#define RIGHT(i) (((i)<<1)+2)

static void heapifyDown(long long *heap,int size,int idx){
    for(;;){
        int l=LEFT(idx),r=RIGHT(idx),m=idx;
        if(l<size && heap[l]<heap[m]) m=l;
        if(r<size && heap[r]<heap[m]) m=r;
        if(m==idx) break;
        swaplonglong(&heap[idx],&heap[m]);
        idx=m;
    }
}

static void heapifyUp(long long *heap,int idx){
    while(idx>0){
        int p=PARENT(idx);
        if(heap[idx]<heap[p]){
            swaplonglong(&heap[idx],&heap[p]);
            idx=p;
        } else break;
    }
}

int main(){
    int N,M;
    scanf("%d%d",&N,&M);
    long long *heap=malloc(N*sizeof(long long));
    for(int i=0;i<N;i++){
        heap[i]=0;
    }
    int heapSize=N;
    for(int i=(N-1)/2;i>=0;i--){
        heapifyDown(heap,heapSize,i);
    }
    long long ans=0;
    for(int i=0;i<M;i++){
        long long t1,t2;
        scanf("%lld%lld",&t1,&t2);
        long long earliest=heap[0];
        if(earliest<t1) earliest=t1;
        long long finish=earliest+t2;
        heap[0]=finish;
        heapifyDown(heap,heapSize,0);
        if(finish>ans) ans=finish;
    }
    printf("%lld\n",ans);
    free(heap);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
