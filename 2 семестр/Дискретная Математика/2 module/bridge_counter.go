package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type pair struct{ to, id int }

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	nextInt := func() int {
		in.Scan()
		n, _ := strconv.Atoi(in.Text())
		return n
	}

	n, m := nextInt(), nextInt()
	adj := make([][]pair, n)
	for i := 0; i < m; i++ {
		u, v := nextInt(), nextInt()
		adj[u] = append(adj[u], pair{v, i})
		adj[v] = append(adj[v], pair{u, i})
	}

	used := make([]bool, n)
	tin := make([]int, n)
	low := make([]int, n)
	isBridge := make([]bool, m)
	timer := 0

	var dfs func(v, p int)
	dfs = func(v, p int) {
		used[v] = true
		tin[v] = timer
		low[v] = timer
		timer++
		for _, e := range adj[v] {
			if e.to == p {
				continue
			}
			if used[e.to] {
				if tin[e.to] < low[v] {
					low[v] = tin[e.to]
				}
			} else {
				dfs(e.to, v)
				if low[e.to] > tin[v] {
					isBridge[e.id] = true
				}
				if low[e.to] < low[v] {
					low[v] = low[e.to]
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if !used[i] {
			dfs(i, -1)
		}
	}

	count := 0
	for _, b := range isBridge {
		if b {
			count++
		}
	}
	fmt.Println(count)
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
