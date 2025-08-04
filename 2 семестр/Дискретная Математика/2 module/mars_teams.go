package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N int
	fmt.Fscan(os.Stdin, &N)
	graph := make([][]bool, N)
	for i := range graph {
		graph[i] = make([]bool, N)
	}
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			var s string
			fmt.Fscan(reader, &s)
			if s == "+" {
				graph[i][j] = true
			}
		}
	}
	vis := make([]bool, N)
	col := make([]int, N)
	A := [][]int{}
	B := [][]int{}
	compOf := make([]int, N)
	for i := 0; i < N; i++ {
		if !vis[i] {
			stack := []int{i}
			col[i] = 0
			vis[i] = true
			compOf[i] = len(A)
			compA := []int{}
			compB := []int{}
			for len(stack) > 0 {
				u := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if col[u] == 0 {
					compA = append(compA, u)
				} else {
					compB = append(compB, u)
				}
				for v := 0; v < N; v++ {
					if graph[u][v] {
						if !vis[v] {
							vis[v] = true
							col[v] = 1 - col[u]
							compOf[v] = len(A)
							stack = append(stack, v)
						} else if col[v] == col[u] {
							fmt.Fprint(os.Stdout, "No solution")
							return
						}
					}
				}
			}
			A = append(A, compA)
			B = append(B, compB)
		}
	}
	C := len(A)
	a := make([]int, C)
	b := make([]int, C)
	for i := 0; i < C; i++ {
		a[i] = len(A[i])
		b[i] = len(B[i])
	}
	total := N
	target := total / 2
	dp := make([][]bool, C+1)
	for i := range dp {
		dp[i] = make([]bool, total+1)
	}
	dp[C][0] = true
	for i := C - 1; i >= 0; i-- {
		for s := 0; s <= total; s++ {
			if dp[i+1][s] {
				dp[i][s+a[i]] = true
				dp[i][s+b[i]] = true
			}
		}
	}
	S := -1
	for s := target; s >= 0; s-- {
		if dp[0][s] {
			S = s
			break
		}
	}
	choice := make([]int, C)
	sum := 0
	for i := 0; i < C; i++ {
		minA := N + 1
		for _, v := range A[i] {
			if v < minA {
				minA = v
			}
		}
		minB := N + 1
		for _, v := range B[i] {
			if v < minB {
				minB = v
			}
		}
		inColor := 0
		szIn := a[i]
		szOut := b[i]
		if minB < minA {
			inColor = 1
			szIn = b[i]
			szOut = a[i]
		}
		if sum+szIn <= S && dp[i+1][S-(sum+szIn)] {
			choice[i] = inColor
			sum += szIn
		} else {
			choice[i] = 1 - inColor
			sum += szOut
		}
	}
	res := []int{}
	for i := 0; i < C; i++ {
		ch := choice[i]
		set := A[i]
		if ch == 1 {
			set = B[i]
		}
		for _, v := range set {
			res = append(res, v+1)
		}
	}
	for i := 0; i < len(res); i++ {
		for j := i + 1; j < len(res); j++ {
			if res[j] < res[i] {
				res[i], res[j] = res[j], res[i]
			}
		}
	}
	for _, v := range res {
		fmt.Fprint(os.Stdout, v, " ")
	}
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
