package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type node struct {
	s int
	w string
	l int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024), 1<<20)
	var lines []string
	for sc.Scan() {
		t := strings.TrimSpace(sc.Text())
		if t != "" {
			lines = append(lines, t)
		}
	}
	if len(lines) == 0 {
		return
	}
	idx := 0
	N, _ := strconv.Atoi(lines[idx])
	idx++
	delta := make([][]int, N)
	var m int
	for i := 0; i < N; i++ {
		row := strings.Fields(lines[idx])
		if i == 0 {
			m = len(row)
		}
		dr := make([]int, m)
		for j, v := range row {
			dr[j], _ = strconv.Atoi(v)
		}
		delta[i] = dr
		idx++
	}
	outs := make([][]string, N)
	for i := 0; i < N; i++ {
		row := strings.Fields(lines[idx])
		or := make([]string, m)
		copy(or, row)
		outs[i] = or
		idx++
	}
	q0, _ := strconv.Atoi(lines[idx])
	idx++
	M, _ := strconv.Atoi(lines[idx])
	vis := make(map[string]struct{})
	res := make(map[string]struct{})
	queue := []node{{q0, "", 0}}
	vis[fmt.Sprintf("%d|", q0)] = struct{}{}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for j := 0; j < m; j++ {
			ns := delta[cur.s][j]
			out := outs[cur.s][j]
			nw, nl := cur.w, cur.l
			if out != "-" {
				if nl == M {
					continue
				}
				nw += out
				nl++
				res[nw] = struct{}{}
			}
			key := fmt.Sprintf("%d|%s", ns, nw)
			if _, ok := vis[key]; ok {
				continue
			}
			vis[key] = struct{}{}
			queue = append(queue, node{ns, nw, nl})
		}
	}
	words := make([]string, 0, len(res))
	for w := range res {
		words = append(words, w)
	}
	sort.Strings(words)
	bw := bufio.NewWriter(os.Stdout)
	for i, w := range words {
		if i > 0 {
			fmt.Fprint(bw, " ")
		}
		fmt.Fprint(bw, w)
	}
	fmt.Fprintln(bw)
	bw.Flush()
}


//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
