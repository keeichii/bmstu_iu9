package main

import (
	"fmt"
	"sort"
)

func main() {
	var x uint32
	fmt.Scan(&x)

	divs := []uint32{}
	for i := uint32(1); i*i <= x; i++ {
		if x%i == 0 {
			divs = append(divs, i)
			if i != x/i {
				divs = append(divs, x/i)
			}
		}
	}

	sort.Slice(divs, func(i, j int) bool { return divs[i] < divs[j] })

	fmt.Println("graph {")

	for _, v := range divs {
		fmt.Printf("    %d\n", v)
	}

	for i, u := range divs {
		for j := i + 1; j < len(divs); j++ {
			v := divs[j]
			if v%u != 0 {
				continue
			}
			// проверяем, нет ли делителя между u и v
			edge := true
			for k := i + 1; k < j; k++ {
				w := divs[k]
				if v%w == 0 && w%u == 0 {
					edge = false
					break
				}
			}
			if edge {
				fmt.Printf("    %d--%d\n", u, v)
			}
		}
	}

	fmt.Println("}")
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
