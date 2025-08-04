package main

import "fmt"

func add(a, b []int32, p int32) []int32 {
    var result []int32
    var carry int32
    for i := 0; i < len(a) || i < len(b) || carry > 0; i++ {
        sum := carry
        if i < len(a) { sum += a[i] }
        if i < len(b) { sum += b[i] }
        result = append(result, sum%p)
        carry = sum/p
    }
    return result
}

func main() {
    fmt.Println(add([]int32{1}, []int32{1}, 10))
    fmt.Println(add([]int32{9,9}, []int32{1}, 10))
    fmt.Println(add([]int32{1,1}, []int32{1,1}, 2))
    fmt.Println(add([]int32{15}, []int32{15}, 16))
    fmt.Println(add([]int32{229}, []int32{1}, 230))
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Лаконично!
