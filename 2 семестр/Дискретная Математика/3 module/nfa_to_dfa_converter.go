package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var N, M int
	fmt.Fscan(in, &N, &M)
	nfa := make([]map[string][]int, N)
	for i := range nfa {
		nfa[i] = make(map[string][]int)
	}
	var symOrder []string
	symSeen := map[string]bool{}
	for i := 0; i < M; i++ {
		var u, v int
		var sym string
		fmt.Fscan(in, &u, &v, &sym)
		nfa[u][sym] = append(nfa[u][sym], v)
		if sym != "lambda" && !symSeen[sym] {
			symSeen[sym] = true
			symOrder = append(symOrder, sym)
		}
	}
	final := make([]bool, N)
	for i := 0; i < N; i++ {
		var f int
		fmt.Fscan(in, &f)
		final[i] = f == 1
	}
	var start int
	fmt.Fscan(in, &start)
	closure := func(ss []int) []int {
		vis := make([]bool, N)
		q := append([]int{}, ss...)
		for _, u := range ss {
			vis[u] = true
		}
		for i := 0; i < len(q); i++ {
			for _, v := range nfa[q[i]]["lambda"] {
				if !vis[v] {
					vis[v] = true
					q = append(q, v)
				}
			}
		}
		sort.Ints(q)
		return q
	}
	key := func(ss []int) string {
		var b strings.Builder
		for i, v := range ss {
			if i > 0 {
				b.WriteByte(' ')
			}
			b.WriteString(strconv.Itoa(v))
		}
		return b.String()
	}
	dfaMap := map[string]int{}
	var dfaStates [][]int
	var dfaAccept []bool
	dfaTrans := map[int]map[int][]string{}
	startSet := closure([]int{start})
	k0 := key(startSet)
	dfaMap[k0] = 0
	dfaStates = append(dfaStates, startSet)
	a0 := false
	for _, u := range startSet {
		if final[u] {
			a0 = true
			break
		}
	}
	dfaAccept = append(dfaAccept, a0)
	dfaTrans[0] = make(map[int][]string)
	queue := []int{0}
	for qi := 0; qi < len(queue); qi++ {
		u := queue[qi]
		st := dfaStates[u]
		for _, sym := range symOrder {
			var mv []int
			for _, x := range st {
				mv = append(mv, nfa[x][sym]...)
			}
			cl := closure(mv)
			k := key(cl)
			v, ok := dfaMap[k]
			if !ok {
				v = len(dfaStates)
				dfaMap[k] = v
				dfaStates = append(dfaStates, cl)
				dfaTrans[v] = make(map[int][]string)
				af := false
				for _, x := range cl {
					if final[x] {
						af = true
						break
					}
				}
				dfaAccept = append(dfaAccept, af)
				queue = append(queue, v)
			}
			dfaTrans[u][v] = append(dfaTrans[u][v], sym)
		}
	}
	type E struct{ u, v int; sy []string }
	var edges []E
	for u, m := range dfaTrans {
		for v, sy := range m {
			edges = append(edges, E{u, v, sy})
		}
	}
	sort.Slice(edges, func(i, j int) bool {
		e1, e2 := edges[i], edges[j]
		if e1.u != e2.u {
			return e1.u < e2.u
		}
		return e1.v < e2.v
	})
	fmt.Println("digraph {")
	fmt.Println("\trankdir = LR")
	for i, ss := range dfaStates {
		label := "[" + key(ss) + "]"
		shape := "circle"
		if dfaAccept[i] {
			shape = "doublecircle"
		}
		fmt.Printf("\t%d [label = \"%s\", shape = %s]\n", i, label, shape)
	}
	for _, e := range edges {
		fmt.Printf("\t%d -> %d [label = \"%s\"]\n", e.u, e.v, strings.Join(e.sy, ", "))
	}
	fmt.Println("}")
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
