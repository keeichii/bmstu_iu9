package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(make([]byte, 1024), 1<<20)
	in.Split(bufio.ScanWords)
	next := func() string { in.Scan(); return in.Text() }
	nextInt := func() int { v, _ := strconv.Atoi(next()); return v }
	n, m, q0 := nextInt(), nextInt(), nextInt()
	trans := make([][]int, n)
	for i := 0; i < n; i++ {
		trans[i] = make([]int, m)
		for j := 0; j < m; j++ {
			trans[i][j] = nextInt()
		}
	}
	outs := make([][]string, n)
	for i := 0; i < n; i++ {
		outs[i] = make([]string, m)
		for j := 0; j < m; j++ {
			outs[i][j] = next()
		}
	}
	mapOld := make([]int, n)
	for i := range mapOld {
		mapOld[i] = -1
	}
	newToOld := make([]int, 0, n)
	type frame struct{ s, idx int }
	stack := []frame{{q0, 0}}
	mapOld[q0] = 0
	newToOld = append(newToOld, q0)
	for len(stack) > 0 {
		top := &stack[len(stack)-1]
		if top.idx < m {
			u := trans[top.s][top.idx]
			top.idx++
			if mapOld[u] == -1 {
				mapOld[u] = len(newToOld)
				newToOld = append(newToOld, u)
				stack = append(stack, frame{u, 0})
			}
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	k := len(newToOld)
	newTrans := make([][]int, k)
	newOuts := make([][]string, k)
	for i := 0; i < k; i++ {
		newTrans[i] = make([]int, m)
		newOuts[i] = make([]string, m)
		old := newToOld[i]
		for j := 0; j < m; j++ {
			newTrans[i][j] = mapOld[trans[old][j]]
			newOuts[i][j] = outs[old][j]
		}
	}
	out := bufio.NewWriterSize(os.Stdout, 1<<20)
	fmt.Fprintf(out, "%d %d %d\n", k, m, 0)
	for i := 0; i < k; i++ {
		for j := 0; j < m; j++ {
			if j+1 < m {
				fmt.Fprintf(out, "%d ", newTrans[i][j])
			} else {
				fmt.Fprintf(out, "%d", newTrans[i][j])
			}
		}
		fmt.Fprintln(out)
	}
	for i := 0; i < k; i++ {
		for j := 0; j < m; j++ {
			if j+1 < m {
				fmt.Fprintf(out, "%s ", newOuts[i][j])
			} else {
				fmt.Fprintf(out, "%s", newOuts[i][j])
			}
		}
		if i+1 < k {
			fmt.Fprintln(out)
		}
	}
	out.Flush()
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
