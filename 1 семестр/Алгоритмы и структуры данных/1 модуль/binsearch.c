
unsigned long binsearch(unsigned long nel, int (*compare)(unsigned long i)) {
    unsigned long left=0,right=nel;
    while(left<right){
        unsigned long mid=(left+right)/2;
        int res=compare(mid);
        if(res<0){
            left=mid+1;
        } else if(res>0){
            right=mid;
        } else {
            return mid;
        }
    }
    return nel;
}

//Антиплагиат: Найдены очень похожие посылки
//Комментарий преподавателя:
