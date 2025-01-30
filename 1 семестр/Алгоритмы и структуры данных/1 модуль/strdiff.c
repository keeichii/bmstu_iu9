
#include <stdio.h>

int strdiff(char *a, char *b) {
    int bit_index = 0;
    int i = 0;
    while (1) {
        unsigned char aa = (unsigned char)a[i];
        unsigned char bb = (unsigned char)b[i];
        if (aa == 0 && bb == 0) {
            return -1;
        }
        if (aa != bb) {
            unsigned char x = (unsigned char)(aa ^ bb);
            int pos = 0;
            while ((x & 1) == 0) {
                x >>= 1;
                pos++;
            }
            return bit_index + pos;
        }
        bit_index += 8;
        i++;
    }
}

int main(void){
    char s1[1024], s2[1024];
    if(!fgets(s1,1024,stdin)) return 1;
    if(!fgets(s2,1024,stdin)) return 1;
    for(int i=0; s1[i]; i++){
        if(s1[i]=='\n'){
            s1[i]='\0';
            break;
        }
    }
    for(int i=0; s2[i]; i++){
        if(s2[i]=='\n'){
            s2[i]='\0';
            break;
        }
    }
    printf("%d\n",strdiff(s1,s2));
    return 0;
}

//Антиплагиат: Похожих посылок не найдено
//Комментарий преподавателя:
