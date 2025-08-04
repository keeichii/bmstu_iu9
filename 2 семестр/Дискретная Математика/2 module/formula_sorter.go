package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var s string
var pos int
var used map[string]struct{}

func parseExpr() error {
	if err := parseTerm(); err != nil {
		return err
	}
	for {
		skip()
		if pos < len(s) && (s[pos] == '+' || s[pos] == '-') {
			pos++
			if err := parseTerm(); err != nil {
				return err
			}
		} else {
			break
		}
	}
	return nil
}

func parseTerm() error {
	if err := parseFactor(); err != nil {
		return err
	}
	for {
		skip()
		if pos < len(s) && (s[pos] == '*' || s[pos] == '/') {
			pos++
			if err := parseFactor(); err != nil {
				return err
			}
		} else {
			break
		}
	}
	return nil
}

func parseFactor() error {
	skip()
	if pos >= len(s) {
		return fmt.Errorf("err")
	}
	c := s[pos]
	if c == '(' {
		pos++
		if err := parseExpr(); err != nil {
			return err
		}
		skip()
		if pos >= len(s) || s[pos] != ')' {
			return fmt.Errorf("err")
		}
		pos++
		return nil
	}
	if c == '-' {
		pos++
		return parseFactor()
	}
	if unicode.IsLetter(rune(c)) {
		start := pos
		pos++
		for pos < len(s) && (unicode.IsLetter(rune(s[pos])) || unicode.IsDigit(rune(s[pos]))) {
			pos++
		}
		used[s[start:pos]] = struct{}{}
		return nil
	}
	if unicode.IsDigit(rune(c)) {
		for pos < len(s) && unicode.IsDigit(rune(s[pos])) {
			pos++
		}
		return nil
	}
	return fmt.Errorf("err")
}

func skip() {
	for pos < len(s) && s[pos] == ' ' {
		pos++
	}
}

func main() {
	lines := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		t := strings.TrimSpace(scanner.Text())
		if t != "" {
			lines = append(lines, t)
		}
	}
	n := len(lines)
	leftVars := make([][]string, n)
	rightExprs := make([][]string, n)
	def := make(map[string]int)
	for i, line := range lines {
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			fmt.Println("syntax error")
			return
		}
		L := strings.Split(parts[0], ",")
		R := strings.Split(parts[1], ",")
		if len(L) != len(R) {
			fmt.Println("syntax error")
			return
		}
		for j := range L {
			v := strings.TrimSpace(L[j])
			if v == "" || !unicode.IsLetter(rune(v[0])) {
				fmt.Println("syntax error")
				return
			}
			for k, ch := range v {
				if k > 0 && !(unicode.IsLetter(ch) || unicode.IsDigit(ch)) {
					fmt.Println("syntax error")
					return
				}
			}
			if _, ok := def[v]; ok {
				fmt.Println("syntax error")
				return
			}
			def[v] = i
			leftVars[i] = append(leftVars[i], v)
		}
		for j := range R {
			e := strings.TrimSpace(R[j])
			if e == "" {
				fmt.Println("syntax error")
				return
			}
			rightExprs[i] = append(rightExprs[i], e)
		}
	}
	deps := make([][]int, n)
	for i := 0; i < n; i++ {
		used = make(map[string]struct{})
		for _, expr := range rightExprs[i] {
			s = expr
			pos = 0
			if err := parseExpr(); err != nil {
				fmt.Println("syntax error")
				return
			}
			skip()
			if pos != len(s) {
				fmt.Println("syntax error")
				return
			}
		}
		for v := range used {
			if j, ok := def[v]; ok {
				deps[i] = append(deps[i], j)
			} else {
				fmt.Println("syntax error")
				return
			}
		}
	}
	adj := make([][]int, n)
	indeg := make([]int, n)
	for i := 0; i < n; i++ {
		indeg[i] = len(deps[i])
		for _, j := range deps[i] {
			adj[j] = append(adj[j], i)
		}
	}
	h := &MaxHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		if indeg[i] == 0 {
			heap.Push(h, i)
		}
	}
	res := []int{}
	for h.Len() > 0 {
		u := heap.Pop(h).(int)
		res = append(res, u)
		for _, v := range adj[u] {
			indeg[v]--
			if indeg[v] == 0 {
				heap.Push(h, v)
			}
		}
	}
	if len(res) != n {
		fmt.Println("cycle")
		return
	}
	for _, idx := range res {
		fmt.Println(lines[idx])
	}
}

//Антиплагиат: Похожих посылок не найдено
//Комментарий преподавателя:
