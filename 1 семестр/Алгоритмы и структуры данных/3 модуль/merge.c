#include <stdio.h>
#include <stdlib.h>

typedef struct {
    int val;
    int arrIndex;
    int idxInArr;
} Node;

static int cmp(const void *a,const void *b){
    const Node *x=(const Node*)a;
    const Node *y=(const Node*)b;
    return x->val - y->val;
}

int main(){
    int k;
    scanf("%d",&k);
    int *sizes=malloc(k*sizeof(int));
    for(int i=0;i<k;i++){
        scanf("%d",&sizes[i]);
    }
    int **arrays=malloc(k*sizeof(int*));
    for(int i=0;i<k;i++){
        arrays[i]=malloc(sizes[i]*sizeof(int));
        for(int j=0;j<sizes[i];j++){
            scanf("%d",&arrays[i][j]);
        }
    }
    int total=0;
    for(int i=0;i<k;i++){
        total+=sizes[i];
    }
    Node *heap=malloc(k*sizeof(Node));
    int heapSize=0;

    #define PARENT(i) (((i)-1)/2)
    #define LEFT(i) (2*(i)+1)
    #define RIGHT(i) (2*(i)+2)
    void swap(Node *x,Node *y){
        Node t=*x;*x=*y;*y=t;
    }
    void push(Node nd){
        heap[heapSize]=nd;
        int cur=heapSize++;
        while(cur>0){
            int par=PARENT(cur);
            if(heap[cur].val<heap[par].val){
                swap(&heap[cur],&heap[par]);
                cur=par;
            } else break;
        }
    }
    void heapify(int idx){
        while(1){
            int l=LEFT(idx),r=RIGHT(idx),small=idx;
            if(l<heapSize && heap[l].val<heap[small].val) small=l;
            if(r<heapSize && heap[r].val<heap[small].val) small=r;
            if(small==idx) break;
            swap(&heap[idx],&heap[small]);
            idx=small;
        }
    }
    Node pop(){
        Node ret=heap[0];
        heap[0]=heap[--heapSize];
        heapify(0);
        return ret;
    }

    for(int i=0;i<k;i++){
        if(sizes[i]>0){
            Node nd={arrays[i][0],i,0};
            push(nd);
        }
    }
    for(int i=0;i<total;i++){
        Node top=pop();
        printf("%d ",top.val);
        if(top.idxInArr+1<sizes[top.arrIndex]){
            Node nxt={arrays[top.arrIndex][top.idxInArr+1],top.arrIndex,top.idxInArr+1};
            push(nxt);
        }
    }
    printf("\n");

    for(int i=0;i<k;i++){
        free(arrays[i]);
    }
    free(arrays);
    free(heap);
    free(sizes);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
