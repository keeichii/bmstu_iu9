package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type e struct{ u, v int; w float64 }

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	read := func() int { in.Scan(); n, _ := strconv.Atoi(in.Text()); return n }

	n := read()
	x, y := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i] = read(), read()
	}
	var es []e
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			d := math.Hypot(float64(x[i]-x[j]), float64(y[i]-y[j]))
			es = append(es, e{i, j, d})
		}
	}
	sort.Slice(es, func(i, j int) bool { return es[i].w < es[j].w })
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	var find func(int) int
	find = func(a int) int {
		if p[a] != a {
			p[a] = find(p[a])
		}
		return p[a]
	}
	merge := func(a, b int) bool {
		a, b = find(a), find(b)
		if a == b {
			return false
		}
		p[b] = a
		return true
	}
	var res float64
	cnt := 0
	for _, ed := range es {
		if merge(ed.u, ed.v) {
			res += ed.w
			cnt++
			if cnt == n-1 {
				break
			}
		}
	}
	fmt.Printf("%.2f\n", res)
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
