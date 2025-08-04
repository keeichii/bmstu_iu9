package main

import (
	"bufio"
	"fmt"
	"os"
)

type edge struct{ to, c int }

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(in, &n, &m)
	g := make([][]edge, n)
	for i := 0; i < m; i++ {
		var u, v, c int
		fmt.Fscan(in, &u, &v, &c)
		u--; v--
		g[u] = append(g[u], edge{v, c})
		g[v] = append(g[v], edge{u, c})
	}
	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	q := make([]int, 0, n)
	dist[n-1] = 0
	q = append(q, n-1)
	for i := 0; i < len(q); i++ {
		u := q[i]
		for _, e := range g[u] {
			if dist[e.to] < 0 {
				dist[e.to] = dist[u] + 1
				q = append(q, e.to)
			}
		}
	}
	d := dist[0]
	res := make([]int, 0, d)
	cur := []int{0}
	seen := make([]int, n)
	mrk := 1
	for dist[cur[0]] > 0 {
		minc := 1<<31 - 1
		for _, u := range cur {
			for _, e := range g[u] {
				if dist[e.to] == dist[u]-1 && e.c < minc {
					minc = e.c
				}
			}
		}
		res = append(res, minc)
		var nxt []int
		for _, u := range cur {
			for _, e := range g[u] {
				if dist[e.to] == dist[u]-1 && e.c == minc && seen[e.to] != mrk {
					seen[e.to] = mrk
					nxt = append(nxt, e.to)
				}
			}
		}
		mrk++
		cur = nxt
	}
	fmt.Println(len(res))
	for i, c := range res {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(c)
	}
	fmt.Println()
}

//Антиплагиат: Похожих посылок не найдено
//Комментарий преподавателя:
