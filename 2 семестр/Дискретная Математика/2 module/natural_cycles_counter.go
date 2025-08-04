package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(in, &n)
	labels := make([]int, n)
	cmd := make([]string, n)
	op := make([]int, n)
	lt := make(map[int]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &labels[i], &cmd[i])
		lt[labels[i]] = i
		if cmd[i] != "ACTION" {
			fmt.Fscan(in, &op[i])
		}
	}
	succ := make([][]int, n)
	for i := 0; i < n; i++ {
		switch cmd[i] {
		case "ACTION":
			if i+1 < n {
				succ[i] = append(succ[i], i+1)
			}
		case "JUMP":
			succ[i] = append(succ[i], lt[op[i]])
		case "BRANCH":
			succ[i] = append(succ[i], lt[op[i]])
			if i+1 < n {
				succ[i] = append(succ[i], i+1)
			}
		}
	}
	dfsn := make([]int, n)
	vertex := make([]int, n+1)
	parent := make([]int, n+1)
	var dfsCnt int
	var dfs func(int)
	dfs = func(u int) {
		dfsCnt++
		dfsn[u] = dfsCnt
		vertex[dfsCnt] = u
		for _, v := range succ[u] {
			if dfsn[v] == 0 {
				dfs(v)
				parent[dfsn[v]] = dfsn[u]
			}
		}
	}
	dfs(0)
	size := dfsCnt
	pred := make([][]int, size+1)
	for u := 0; u < n; u++ {
		if dfsn[u] == 0 {
			continue
		}
		du := dfsn[u]
		for _, v := range succ[u] {
			if dfsn[v] == 0 {
				continue
			}
			dv := dfsn[v]
			pred[dv] = append(pred[dv], du)
		}
	}
	semi := make([]int, size+1)
	idom := make([]int, size+1)
	ancestor := make([]int, size+1)
	label := make([]int, size+1)
	bucket := make([][]int, size+1)
	for i := 1; i <= size; i++ {
		semi[i] = i
		label[i] = i
	}
	var compress func(int)
	compress = func(v int) {
		if ancestor[ancestor[v]] != 0 {
			compress(ancestor[v])
			if semi[label[ancestor[v]]] < semi[label[v]] {
				label[v] = label[ancestor[v]]
			}
			ancestor[v] = ancestor[ancestor[v]]
		}
	}
	eval := func(v int) int {
		if ancestor[v] == 0 {
			return v
		}
		compress(v)
		if semi[label[ancestor[v]]] < semi[label[v]] {
			return label[ancestor[v]]
		}
		return label[v]
	}
	link := func(v, w int) {
		ancestor[w] = v
	}
	for i := size; i >= 2; i-- {
		for _, v := range pred[i] {
			u := eval(v)
			if semi[u] < semi[i] {
				semi[i] = semi[u]
			}
		}
		bucket[semi[i]] = append(bucket[semi[i]], i)
		link(parent[i], i)
		for _, v := range bucket[parent[i]] {
			u := eval(v)
			if semi[u] < semi[v] {
				idom[v] = u
			} else {
				idom[v] = parent[i]
			}
		}
	}
	for i := 2; i <= size; i++ {
		if idom[i] != semi[i] {
			idom[i] = idom[idom[i]]
		}
	}
	idom[1] = 0
	children := make([][]int, size+1)
	for i := 2; i <= size; i++ {
		children[idom[i]] = append(children[idom[i]], i)
	}
	tin := make([]int, size+1)
	tout := make([]int, size+1)
	var timer int
	var dfs2 func(int)
	dfs2 = func(u int) {
		timer++
		tin[u] = timer
		for _, v := range children[u] {
			dfs2(v)
		}
		tout[u] = timer
	}
	dfs2(1)
	header := make([]bool, size+1)
	for u := 0; u < n; u++ {
		if dfsn[u] == 0 {
			continue
		}
		for _, v := range succ[u] {
			if dfsn[v] == 0 {
				continue
			}
			du := dfsn[u]
			dv := dfsn[v]
			if tin[dv] <= tin[du] && tout[du] <= tout[dv] {
				header[dv] = true
			}
		}
	}
	var count int
	for i := 1; i <= size; i++ {
		if header[i] {
			count++
		}
	}
	fmt.Println(count)
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Решение основано на решении Ильи Резника.
