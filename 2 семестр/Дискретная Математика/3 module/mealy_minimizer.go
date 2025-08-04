package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024), 1<<20)
	sc.Split(bufio.ScanWords)
	next := func() string { sc.Scan(); return sc.Text() }
	atoi := func() int { v, _ := strconv.Atoi(next()); return v }

	n, m, q0 := atoi(), atoi(), atoi()
	tr := make([][]int, n)
	for i := 0; i < n; i++ {
		tr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			tr[i][j] = atoi()
		}
	}
	out := make([][]string, n)
	for i := 0; i < n; i++ {
		out[i] = make([]string, m)
		for j := 0; j < m; j++ {
			out[i][j] = next()
		}
	}

	visited := make([]bool, n)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = -1
	}
	queue := []int{q0}
	visited[q0] = true
	order := []int{}
	for h := 0; h < len(queue); h++ {
		s := queue[h]
		idx[s] = len(order)
		order = append(order, s)
		for j := 0; j < m; j++ {
			u := tr[s][j]
			if !visited[u] {
				visited[u] = true
				queue = append(queue, u)
			}
		}
	}

	r := len(order)
	part := make([]int, r)
	key := map[string]int{}
	var sb strings.Builder
	for i, s := range order {
		sb.Reset()
		for j := 0; j < m; j++ {
			sb.WriteString(out[s][j])
			sb.WriteByte(0)
		}
		k := sb.String()
		if id, ok := key[k]; ok {
			part[i] = id
		} else {
			id := len(key)
			key[k] = id
			part[i] = id
		}
	}

	for {
		changed := false
		sig := map[string]int{}
		newPart := make([]int, r)
		for i, s := range order {
			sb.Reset()
			for j := 0; j < m; j++ {
				sb.WriteString(out[s][j])
				sb.WriteByte(1)
				sb.WriteString(strconv.Itoa(part[idx[tr[s][j]]]))
				sb.WriteByte(2)
			}
			s := sb.String()
			id, ok := sig[s]
			if !ok {
				id = len(sig)
				sig[s] = id
			}
			newPart[i] = id
			if newPart[i] != part[i] {
				changed = true
			}
		}
		part = newPart
		if !changed {
			break
		}
	}

	groups := 0
	for _, p := range part {
		if p+1 > groups {
			groups = p + 1
		}
	}
	rep := make([]int, groups)
	for i := range rep {
		rep[i] = -1
	}
	for i, s := range order {
		if rep[part[i]] == -1 {
			rep[part[i]] = s
		}
	}

	gtr := make([][]int, groups)
	gout := make([][]string, groups)
	for g := 0; g < groups; g++ {
		gtr[g] = make([]int, m)
		gout[g] = make([]string, m)
		rs := rep[g]
		for j := 0; j < m; j++ {
			gtr[g][j] = part[idx[tr[rs][j]]]
			gout[g][j] = out[rs][j]
		}
	}

	num := make([]int, groups)
	for i := range num {
		num[i] = -1
	}
	canon := []int{}
	start := part[idx[q0]]
	num[start] = 0
	canon = append(canon, start)
	type frame struct{ g, i int }
	stack := []frame{{start, 0}}
	for len(stack) > 0 {
		top := &stack[len(stack)-1]
		if top.i < m {
			dst := gtr[top.g][top.i]
			top.i++
			if num[dst] == -1 {
				num[dst] = len(canon)
				canon = append(canon, dst)
				stack = append(stack, frame{dst, 0})
			}
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	w := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(w, "digraph {")
	fmt.Fprintln(w, "    rankdir = LR")
	for _, g := range canon {
		for j := 0; j < m; j++ {
			fmt.Fprintf(w, "    %d -> %d [label = \"%c(%s)\"]\n",
				num[g], num[gtr[g][j]], 'a'+rune(j), gout[g][j])
		}
	}
	fmt.Fprintln(w, "}")
	w.Flush()
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
