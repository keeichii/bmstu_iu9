


#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int main(int argc, char *argv[]){
    if(argc!=4){
        printf("Usage: frame <height> <width> <text>\n");
        return 0;
    }
    int h=atoi(argv[1]);
    int w=atoi(argv[2]);
    char *text=argv[3];
    int len=strlen(text);
    if(h<3 || len>=(w-1)){
        printf("Error\n");
        return 0;
    }
    for(int i=0;i<w;i++){
        putchar('*');
    }
    putchar('\n');
    int textRow=(h-1)/2;
    for(int row=1;row<h-1;row++){
        putchar('*');
        if(row==textRow){
            int spaceLeft=(w-2-len)/2;
            int spaceRight=w-2-len-spaceLeft;
            for(int i=0;i<spaceLeft;i++){
                putchar(' ');
            }
            fputs(text,stdout);
            for(int i=0;i<spaceRight;i++){
                putchar(' ');
            }
        } else {
            for(int i=0;i<w-2;i++){
                putchar(' ');
            }
        }
        putchar('*');
        putchar('\n');
    }
    for(int i=0;i<w;i++){
        putchar('*');
    }
    putchar('\n');
    return 0;
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
