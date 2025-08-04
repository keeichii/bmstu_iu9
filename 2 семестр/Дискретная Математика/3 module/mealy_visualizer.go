package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m, q0 int
	fmt.Fscan(in, &n, &m, &q0)
	T := make([][]int, n)
	for i := 0; i < n; i++ {
		T[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &T[i][j])
		}
	}
	O := make([][]string, n)
	for i := 0; i < n; i++ {
		O[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(in, &O[i][j])
		}
	}
	fmt.Println("digraph {")
	fmt.Println("    rankdir = LR")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			letter := string('a' + rune(j))
			fmt.Printf("    %d -> %d [label = \"%s(%s)\"]\n",
				i, T[i][j], letter, O[i][j])
		}
	}
	fmt.Println("}")
}


//Антиплагиат: Найдены очень похожие посылки
//Комментарий преподавателя: Тут часто очень похожие.
