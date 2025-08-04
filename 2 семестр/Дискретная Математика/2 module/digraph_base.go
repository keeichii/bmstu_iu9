package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	next := func() int { in.Scan(); v, _ := strconv.Atoi(in.Text()); return v }

	n, m := next(), next()
	g, gr := make([][]int, n), make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := next(), next()
		g[u] = append(g[u], v)
		gr[v] = append(gr[v], u)
	}

	used := make([]bool, n)
	order := make([]int, 0, n)
	var dfs1 func(int)
	dfs1 = func(u int) {
		stack := []struct{ v, i int }{{u, 0}}
		used[u] = true
		for len(stack) > 0 {
			top := &stack[len(stack)-1]
			v, i := top.v, top.i
			if i < len(g[v]) {
				top.i++
				to := g[v][i]
				if !used[to] {
					used[to] = true
					stack = append(stack, struct{ v, i int }{to, 0})
				}
			} else {
				order = append(order, v)
				stack = stack[:len(stack)-1]
			}
		}
	}
	for i := 0; i < n; i++ {
		if !used[i] {
			dfs1(i)
		}
	}

	comp := make([]int, n)
	for i := range comp {
		comp[i] = -1
	}
	cid := 0
	var dfs2 func(int)
	dfs2 = func(u int) {
		stack := []int{u}
		comp[u] = cid
		for len(stack) > 0 {
			v := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, to := range gr[v] {
				if comp[to] == -1 {
					comp[to] = cid
					stack = append(stack, to)
				}
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		v := order[i]
		if comp[v] == -1 {
			dfs2(v)
			cid++
		}
	}

	indeg := make([]int, cid)
	for u := 0; u < n; u++ {
		for _, v := range g[u] {
			if comp[u] != comp[v] {
				indeg[comp[v]]++
			}
		}
	}

	minv := make([]int, cid)
	for i := range minv {
		minv[i] = n
	}
	for v, c := range comp {
		if v < minv[c] {
			minv[c] = v
		}
	}

	var res []int
	for c, d := range indeg {
		if d == 0 {
			res = append(res, minv[c])
		}
	}
	sort.Ints(res)
	for i, v := range res {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(v)
	}
	fmt.Println()
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
