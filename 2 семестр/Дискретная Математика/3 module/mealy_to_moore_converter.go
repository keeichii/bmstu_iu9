package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type pair struct {
	q   int
	out string
}

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(make([]byte, 1024), 1<<20)
	in.Split(bufio.ScanWords)
	next := func() string { in.Scan(); return in.Text() }
	atoi := func() int { v, _ := strconv.Atoi(next()); return v }

	kin := atoi()
	inAlpha := make([]string, kin)
	for i := range inAlpha {
		inAlpha[i] = next()
	}

	kout := atoi()
	outAlpha := make([]string, kout)
	for i := range outAlpha {
		outAlpha[i] = next()
	}

	n := atoi()
	delta := make([][]int, n)
	for i := 0; i < n; i++ {
		delta[i] = make([]int, kin)
		for j := 0; j < kin; j++ {
			delta[i][j] = atoi()
		}
	}

	phi := make([][]string, n)
	translate := func(tok string) string {
		if idx, err := strconv.Atoi(tok); err == nil && idx >= 0 && idx < kout {
			return outAlpha[idx]
		}
		return tok
	}
	for i := 0; i < n; i++ {
		phi[i] = make([]string, kin)
		for j := 0; j < kin; j++ {
			phi[i][j] = translate(next())
		}
	}

	set := map[pair]struct{}{}
	for s := 0; s < n; s++ {
		for j := 0; j < kin; j++ {
			p := pair{delta[s][j], phi[s][j]}
			set[p] = struct{}{}
		}
	}

	states := make([]pair, 0, len(set))
	for p := range set {
		states = append(states, p)
	}
	sort.Slice(states, func(i, j int) bool {
		if states[i].q != states[j].q {
			return states[i].q < states[j].q
		}
		return states[i].out < states[j].out
	})

	id := make(map[pair]int, len(states))
	for i, p := range states {
		id[p] = i
	}

	out := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(out, "digraph {")
	fmt.Fprintln(out, "    rankdir = LR")
	for i, p := range states {
		fmt.Fprintf(out, "    %d [label = \"(%d,%s)\"]\n", i, p.q, p.out)
	}
	for i, p := range states {
		for j := 0; j < kin; j++ {
			dp := pair{delta[p.q][j], phi[p.q][j]}
			fmt.Fprintf(out, "    %d -> %d [label = \"%s\"]\n", i, id[dp], inAlpha[j])
		}
	}
	fmt.Fprintln(out, "}")
	out.Flush()
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
