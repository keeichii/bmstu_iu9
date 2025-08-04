package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type edge struct{ u, v int }

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	nextInt := func() int {
		in.Scan()
		n, _ := strconv.Atoi(in.Text())
		return n
	}

	n, m := nextInt(), nextInt()
	adj := make([][]int, n)
	edges := make([]edge, m)
	for i := 0; i < m; i++ {
		u, v := nextInt(), nextInt()
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
		edges[i] = edge{u, v}
	}

	visited := make([]bool, n)
	comp := make([]int, n)
	compSize := make(map[int]int)
	compEdges := make(map[int]int)
	compMin := make(map[int]int)
	id := 0

	var dfs func(v, cid int)
	dfs = func(v, cid int) {
		stack := []int{v}
		comp[v] = cid
		visited[v] = true
		compSize[cid]++
		compMin[cid] = v
		for len(stack) > 0 {
			u := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, w := range adj[u] {
				compEdges[cid]++
				if !visited[w] {
					visited[w] = true
					comp[w] = cid
					compSize[cid]++
					if w < compMin[cid] {
						compMin[cid] = w
					}
					stack = append(stack, w)
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, id)
			compEdges[id] /= 2
			id++
		}
	}

	bestID := -1
	for i := 0; i < id; i++ {
		if bestID == -1 ||
			compSize[i] > compSize[bestID] ||
			(compSize[i] == compSize[bestID] &&
				compEdges[i] > compEdges[bestID]) ||
			(compSize[i] == compSize[bestID] &&
				compEdges[i] == compEdges[bestID] &&
				compMin[i] < compMin[bestID]) {
			bestID = i
		}
	}

	fmt.Println("graph G {")
	for i := 0; i < n; i++ {
		if comp[i] == bestID {
			fmt.Printf("  %d [color=red];\n", i)
		} else {
			fmt.Printf("  %d;\n", i)
		}
	}
	for _, e := range edges {
		if comp[e.u] == bestID && comp[e.v] == bestID {
			fmt.Printf("  %d -- %d [color=red];\n", e.u, e.v)
		} else {
			fmt.Printf("  %d -- %d;\n", e.u, e.v)
		}
	}
	fmt.Println("}")
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
