package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type edge struct{ to, w int }
type item struct{ to, w int }
type pq []item

func (h pq) Len() int            { return len(h) }
func (h pq) Less(i, j int) bool  { return h[i].w < h[j].w }
func (h pq) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *pq) Push(x interface{}) { *h = append(*h, x.(item)) }
func (h *pq) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	read := func() int { in.Scan(); v, _ := strconv.Atoi(in.Text()); return v }

	n, m := read(), read()
	g := make([][]edge, n)
	for i := 0; i < m; i++ {
		u, v, w := read(), read(), read()
		g[u] = append(g[u], edge{v, w})
		g[v] = append(g[v], edge{u, w})
	}

	used := make([]bool, n)
	h := &pq{}
	heap.Init(h)
	used[0] = true
	for _, e := range g[0] {
		heap.Push(h, item{e.to, e.w})
	}

	total := 0
	for cnt := 1; cnt < n && h.Len() > 0; {
		it := heap.Pop(h).(item)
		if used[it.to] {
			continue
		}
		used[it.to] = true
		total += it.w
		cnt++
		for _, e := range g[it.to] {
			if !used[e.to] {
				heap.Push(h, item{e.to, e.w})
			}
		}
	}

	fmt.Println(total)
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
