#include <stdio.h>
#include <stdlib.h>

static long long gcd(long long a,long long b){
    while(b!=0){
        long long t=a%b;
        a=b;
        b=t;
    }
    return a>0?a:-a;
}

static int lg[300001];
static long long st[300001][19];

int main(){
    int n;
    scanf("%d",&n);
    for(int i=1;i<=n;i++){
        lg[i]=lg[i-1];
        if((1<<(lg[i]+1))<=i) lg[i]++;
    }
    long long *arr=malloc(n*sizeof(long long));
    for(int i=0;i<n;i++){
        scanf("%lld",&arr[i]);
        st[i][0]=arr[i];
    }
    for(int j=1;(1<<j)<=n;j++){
        for(int i=0;i+(1<<j)-1<n;i++){
            st[i][j]=gcd(st[i][j-1],st[i+(1<<(j-1))][j-1]);
        }
    }
    int m;
    scanf("%d",&m);
    while(m--){
        int l,r;
        scanf("%d%d",&l,&r);
        int k=lg[r-l+1];
        long long g=gcd(st[l][k],st[r-(1<<k)+1][k]);
        printf("%lld\n",g);
    }
    free(arr);
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
