package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	var all strings.Builder
	for reader.Scan() {
		all.WriteString(reader.Text())
		all.WriteString(" ")
	}
	text := all.String()
	parts := strings.Split(text, ";")
	nameToID := map[string]int{}
	var names []string
	var dur []int
	var edges [][2]int
	for _, part := range parts {
		sent := strings.TrimSpace(part)
		if sent == "" {
			continue
		}
		toks := strings.Split(sent, "<")
		var seq []int
		for _, tok := range toks {
			t := strings.TrimSpace(tok)
			var name string
			var d int
			if i := strings.Index(t, "("); i >= 0 && strings.HasSuffix(t, ")") {
				name = strings.TrimSpace(t[:i])
				num := t[i+1 : len(t)-1]
				d = 0
				for _, ch := range num {
					d = d*10 + int(ch-'0')
				}
			} else {
				name = t
				d = -1
			}
			if name == "" || !unicode.IsLetter(rune(name[0])) {
				continue
			}
			id, ok := nameToID[name]
			if !ok {
				id = len(names)
				nameToID[name] = id
				names = append(names, name)
				if d >= 0 {
					dur = append(dur, d)
				} else {
					dur = append(dur, 0)
				}
			} else if d >= 0 {
				dur[id] = d
			}
			seq = append(seq, id)
		}
		for i := 0; i+1 < len(seq); i++ {
			edges = append(edges, [2]int{seq[i], seq[i+1]})
		}
	}
	n := len(names)
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
	}
	idx := make([]int, n)
	low := make([]int, n)
	on := make([]bool, n)
	for i := 0; i < n; i++ {
		idx[i] = -1
	}
	var st []int
	var ids int
	sccID := make([]int, n)
	var sccCnt int
	var sccSize []int
	var dfs func(int)
	dfs = func(v int) {
		idx[v] = ids
		low[v] = ids
		ids++
		st = append(st, v)
		on[v] = true
		for _, w := range adj[v] {
			if idx[w] == -1 {
				dfs(w)
				if low[w] < low[v] {
					low[v] = low[w]
				}
			} else if on[w] && idx[w] < low[v] {
				low[v] = idx[w]
			}
		}
		if low[v] == idx[v] {
			size := 0
			for {
				w := st[len(st)-1]
				st = st[:len(st)-1]
				on[w] = false
				sccID[w] = sccCnt
				size++
				if w == v {
					break
				}
			}
			sccSize = append(sccSize, size)
			sccCnt++
		}
	}
	for i := 0; i < n; i++ {
		if idx[i] == -1 {
			dfs(i)
		}
	}
	bad := make([]bool, n)
	for v := 0; v < n; v++ {
		id := sccID[v]
		if sccSize[id] > 1 {
			bad[v] = true
		}
		for _, w := range adj[v] {
			if v == w {
				bad[v] = true
			}
		}
	}
	var queue []int
	for i := 0; i < n; i++ {
		if bad[i] {
			queue = append(queue, i)
		}
	}
	for qi := 0; qi < len(queue); qi++ {
		u := queue[qi]
		for _, v := range adj[u] {
			if !bad[v] {
				bad[v] = true
				queue = append(queue, v)
			}
		}
	}
	rev := make([][]int, n)
	for u := 0; u < n; u++ {
		for _, v := range adj[u] {
			rev[v] = append(rev[v], u)
		}
	}
	indeg := make([]int, n)
	for v := 0; v < n; v++ {
		if bad[v] {
			continue
		}
		for _, u := range rev[v] {
			if !bad[u] {
				indeg[v]++
			}
		}
	}
	var topo []int
	var q []int
	for i := 0; i < n; i++ {
		if !bad[i] && indeg[i] == 0 {
			q = append(q, i)
		}
	}
	for qi := 0; qi < len(q); qi++ {
		u := q[qi]
		topo = append(topo, u)
		for _, v := range adj[u] {
			if bad[v] {
				continue
			}
			indeg[v]--
			if indeg[v] == 0 {
				q = append(q, v)
			}
		}
	}
	maxT := make([]int, n)
	pred := make([][]int, n)
	for _, v := range topo {
		best := -1
		for _, u := range rev[v] {
			if bad[u] {
				continue
			}
			cand := maxT[u] + dur[v]
			if cand > best {
				best = cand
				pred[v] = []int{u}
			} else if cand == best {
				pred[v] = append(pred[v], u)
			}
		}
		if best >= 0 {
			maxT[v] = best
		} else {
			maxT[v] = dur[v]
		}
	}
	global := 0
	var ends []int
	for i := 0; i < n; i++ {
		if bad[i] {
			continue
		}
		if maxT[i] > global {
			global = maxT[i]
			ends = []int{i}
		} else if maxT[i] == global {
			ends = append(ends, i)
		}
	}
	critNode := make([]bool, n)
	critEdge := map[[2]int]bool{}
	var mark func(int)
	mark = func(v int) {
		if critNode[v] {
			return
		}
		critNode[v] = true
		for _, u := range pred[v] {
			critEdge[[2]int{u, v}] = true
			mark(u)
		}
	}
	for _, v := range ends {
		mark(v)
	}
	sorted := make([]int, n)
	for i := range sorted {
		sorted[i] = i
	}
	sort.Slice(sorted, func(i, j int) bool {
		return names[sorted[i]] < names[sorted[j]]
	})
	fmt.Println("digraph {")
	for _, i := range sorted {
		name := names[i]
		label := fmt.Sprintf("%s(%d)", name, dur[i])
		if bad[i] {
			fmt.Printf("  %s [label = \"%s\", color = blue]\n", name, label)
		} else if critNode[i] {
			fmt.Printf("  %s [label = \"%s\", color = red]\n", name, label)
		} else {
			fmt.Printf("  %s [label = \"%s\"]\n", name, label)
		}
	}
	type E struct{ u, v int }
	var edgeList []E
	for _, e := range edges {
		edgeList = append(edgeList, E{e[0], e[1]})
	}
	sort.Slice(edgeList, func(i, j int) bool {
		ui, vi := edgeList[i].u, edgeList[i].v
		uj, vj := edgeList[j].u, edgeList[j].v
		if names[ui] != names[uj] {
			return names[ui] < names[uj]
		}
		return names[vi] < names[vj]
	})
	for _, e := range edgeList {
		u, v := e.u, e.v
		if bad[u] {
			fmt.Printf("  %s -> %s [color = blue]\n", names[u], names[v])
		} else if critEdge[[2]int{u, v}] {
			fmt.Printf("  %s -> %s [color = red]\n", names[u], names[v])
		} else {
			fmt.Printf("  %s -> %s\n", names[u], names[v])
		}
	}
	fmt.Println("}")
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
