
#include <stddef.h>

void revarray(void *base, size_t nel, size_t width) {
    char *arr_start = (char *)base;
    char *arr_end = arr_start + (nel - 1) * width;

    for (size_t offset = 0; arr_start + offset < arr_end - offset; offset += width) {
        for (size_t byte = 0; byte < width; byte++) {
            char temp = arr_start[offset + byte];
            arr_start[offset + byte] = arr_end[-offset + byte];
            arr_end[-offset + byte] = temp;
        }
    }
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: 👍Можно и так.
