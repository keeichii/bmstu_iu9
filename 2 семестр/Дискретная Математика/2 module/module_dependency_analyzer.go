package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func isL(c byte) bool { return c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' }
func isD(c byte) bool { return c >= '0' && c <= '9' }

func lex(src []byte) ([]string, bool) {
	var t []string
	add := func(s string) { t = append(t, s) }
	for i := 0; i < len(src); {
		if src[i] <= ' ' {
			i++
			continue
		}
		if isL(src[i]) {
			j := i + 1
			for j < len(src) && (isL(src[j]) || isD(src[j])) {
				j++
			}
			add(string(src[i:j]))
			i = j
			continue
		}
		if isD(src[i]) {
			j := i + 1
			for j < len(src) && isD(src[j]) {
				j++
			}
			if j < len(src) && isL(src[j]) {
				return nil, false
			}
			add(string(src[i:j]))
			i = j
			continue
		}
		if i+1 < len(src) {
			p := string(src[i : i+2])
			if p == ":=" || p == "<=" || p == ">=" || p == "<>" {
				add(p)
				i += 2
				continue
			}
		}
		switch src[i] {
		case '+', '-', '*', '/', '(', ')', ',', ';', '?', ':', '=', '<', '>':
			add(string(src[i]))
			i++
		default:
			return nil, false
		}
	}
	return t, true
}

type prs struct{ t []string; p int }

func (p *prs) peek() string {
	if p.p >= len(p.t) {
		return ""
	}
	return p.t[p.p]
}
func (p *prs) next() string { v := p.peek(); if v != "" { p.p++ }; return v }
func (p *prs) accept(s string) bool { if p.peek() == s { p.p++; return true }; return false }

func ident(s string) bool { return s != "" && unicode.IsLetter(rune(s[0])) }
func number(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return s != ""
}

type call struct{ name string; cnt int }
type graph struct{ id map[string]int; adj [][]int }

func (g *graph) idx(name string) int {
	if v, ok := g.id[name]; ok {
		return v
	}
	v := len(g.id)
	g.id[name] = v
	g.adj = append(g.adj, nil)
	return v
}

func expr(p *prs, par map[string]bool, dep map[string]struct{}) bool {
	if !comp(p, par, dep) {
		return false
	}
	if p.accept("?") {
		return comp(p, par, dep) && p.accept(":") && expr(p, par, dep)
	}
	return true
}
func comp(p *prs, par map[string]bool, dep map[string]struct{}) bool {
	if !add(p, par, dep) {
		return false
	}
	switch p.peek() {
	case "=", "<>", "<", ">", "<=", ">=":
		p.next()
		return add(p, par, dep)
	}
	return true
}
func add(p *prs, par map[string]bool, dep map[string]struct{}) bool {
	if !mul(p, par, dep) {
		return false
	}
	for p.peek() == "+" || p.peek() == "-" {
		p.next()
		if !mul(p, par, dep) {
			return false
		}
	}
	return true
}
func mul(p *prs, par map[string]bool, dep map[string]struct{}) bool {
	if !una(p, par, dep) {
		return false
	}
	for p.peek() == "*" || p.peek() == "/" {
		p.next()
		if !una(p, par, dep) {
			return false
		}
	}
	return true
}
func una(p *prs, par map[string]bool, dep map[string]struct{}) bool {
	if p.accept("-") {
		return una(p, par, dep)
	}
	return pri(p, par, dep)
}
func pri(p *prs, par map[string]bool, dep map[string]struct{}) bool {
	if p.accept("(") {
		return expr(p, par, dep) && p.accept(")")
	}
	t := p.peek()
	if ident(t) {
		p.next()
		if p.accept("(") {
			if !par[t] {
				dep[t] = struct{}{}
			}
			arg := 0
			if !p.accept(")") {
				arg++
				if !expr(p, par, dep) {
					return false
				}
				for p.accept(",") {
					arg++
					if !expr(p, par, dep) {
						return false
					}
				}
				if !p.accept(")") {
					return false
				}
			}
			dep[t+"#"+fmt.Sprint(arg)] = struct{}{}
			return true
		}
		if !par[t] {
			return false
		}
		return true
	}
	if number(t) {
		p.next()
		return true
	}
	return false
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024), 1<<20)
	var sb strings.Builder
	for sc.Scan() {
		sb.WriteString(sc.Text())
		sb.WriteByte('\n')
	}
	tok, ok := lex([]byte(sb.String()))
	if !ok {
		fmt.Println("error")
		return
	}
	p := &prs{t: tok}
	g := &graph{id: map[string]int{}}
	paramCnt := map[string]int{}
	calls := map[int][]call{}

	for p.peek() != "" {
		fn := p.next()
		if !ident(fn) || !p.accept("(") {
			fmt.Println("error")
			return
		}
		par := map[string]bool{}
		if !p.accept(")") {
			for {
				v := p.next()
				if !ident(v) {
					fmt.Println("error")
					return
				}
				par[v] = true
				if p.accept(")") {
					break
				}
				if !p.accept(",") {
					fmt.Println("error")
					return
				}
			}
		}
		if !p.accept(":=") {
			fmt.Println("error")
			return
		}
		if _, dup := paramCnt[fn]; dup {
			fmt.Println("error")
			return
		}
		paramCnt[fn] = len(par)
		id := g.idx(fn)

		dep := map[string]struct{}{}
		if !expr(p, par, dep) || !p.accept(";") {
			fmt.Println("error")
			return
		}
		for d := range dep {
			if i := strings.IndexByte(d, '#'); i != -1 {
				name := d[:i]
				cnt := 0
				fmt.Sscan(d[i+1:], &cnt)
				calls[id] = append(calls[id], call{name, cnt})
			}
		}
	}

	for _, lst := range calls {
		for _, c := range lst {
			if pc, ok := paramCnt[c.name]; !ok || pc != c.cnt {
				fmt.Println("error")
				return
			}
		}
	}

	for u, lst := range calls {
		for _, c := range lst {
			v := g.idx(c.name)
			g.adj[u] = append(g.adj[u], v)
		}
	}

	n := len(g.id)
	disc, low := make([]int, n), make([]int, n)
	var stack []int
	in := make([]bool, n)
	var time, scc int
	var dfs func(int)
	dfs = func(v int) {
		time++
		disc[v], low[v] = time, time
		stack = append(stack, v)
		in[v] = true
		for _, w := range g.adj[v] {
			if disc[w] == 0 {
				dfs(w)
				if low[w] < low[v] {
					low[v] = low[w]
				}
			} else if in[w] && disc[w] < low[v] {
				low[v] = disc[w]
			}
		}
		if low[v] == disc[v] {
			for {
				w := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				in[w] = false
				if w == v {
					break
				}
			}
			scc++
		}
	}
	for v := 0; v < n; v++ {
		if disc[v] == 0 {
			dfs(v)
		}
	}
	fmt.Println(scc)
}

//Антиплагиат: Похожих посылок не найдено
//Комментарий преподавателя:
