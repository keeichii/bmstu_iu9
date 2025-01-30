void bubblesort(unsigned long nel,
    int (*compare)(unsigned long i, unsigned long j),
    void (*swap)(unsigned long i, unsigned long j))
{
    if(nel<2) return;
    unsigned long left=0,right=nel-1;
    int swapped=1;
    while(swapped){
        swapped=0;
        for(unsigned long i=left;i<right;i++){
            if(compare(i,i+1)==1){
                swap(i,i+1);
                swapped=1;
            }
        }
        if(!swapped) break;
        swapped=0;
        right--;
        for(unsigned long i=right;i>left;i--){
            if(compare(i-1,i)==1){
                swap(i-1,i);
                swapped=1;
            }
        }
        left++;
    }
}

//Антиплагиат: Найдены очень похожие посылки
//Комментарий преподавателя:
