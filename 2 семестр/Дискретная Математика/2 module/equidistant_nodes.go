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
	read := func() int {
		in.Scan()
		n, _ := strconv.Atoi(in.Text())
		return n
	}

	n, m := read(), read()
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		u, v := read(), read()
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	k := read()
	src := make([]int, k)
	for i := range src {
		src[i] = read()
	}

	dist := make([][]int, k)
	for i := range src {
		d := make([]int, n)
		for j := range d {
			d[j] = -1
		}
		q := []int{src[i]}
		d[src[i]] = 0
		for len(q) > 0 {
			v := q[0]
			q = q[1:]
			for _, u := range g[v] {
				if d[u] == -1 {
					d[u] = d[v] + 1
					q = append(q, u)
				}
			}
		}
		dist[i] = d
	}

	var res []int
	for v := 0; v < n; v++ {
		ok := true
		d0 := dist[0][v]
		if d0 == -1 {
			continue
		}
		for i := 1; i < k; i++ {
			if dist[i][v] != d0 {
				ok = false
				break
			}
		}
		if ok {
			res = append(res, v)
		}
	}

	if len(res) == 0 {
		fmt.Println("-")
	} else {
		sort.Ints(res)
		for _, v := range res {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
