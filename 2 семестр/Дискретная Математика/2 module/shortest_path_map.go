package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

type node struct{ d, i int }
type pq []node

func (h pq) Len() int           { return len(h) }
func (h pq) Less(i, j int) bool { return h[i].d < h[j].d }
func (h pq) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *pq) Push(x interface{}) { *h = append(*h, x.(node)) }
func (h *pq) Pop() interface{} {
	old := *h; n := len(old); x := old[n-1]; *h = old[:n-1]; return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	var N int
	fmt.Fscan(in, &N)
	grid := make([]int, N*N)
	for i := range grid {
		fmt.Fscan(in, &grid[i])
	}
	inf := 1_000_000_000
	dist := make([]int, N*N)
	for i := range dist {
		dist[i] = inf
	}
	start, end := 0, N*N-1
	dist[start] = grid[start]
	h := &pq{{dist[start], start}}
	heap.Init(h)
	dirs := [4]int{-1, 1, -N, N}
	for h.Len() > 0 {
		cur := heap.Pop(h).(node)
		if cur.d != dist[cur.i] {
			continue
		}
		if cur.i == end {
			fmt.Println(cur.d)
			return
		}
		for _, dir := range dirs {
			j := cur.i + dir
			if j < 0 || j >= N*N {
				continue
			}
			if dir == -1 && cur.i%N == 0 {
				continue
			}
			if dir == 1 && j%N == 0 {
				continue
			}
			nd := cur.d + grid[j]
			if nd < dist[j] {
				dist[j] = nd
				heap.Push(h, node{nd, j})
			}
		}
	}
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
