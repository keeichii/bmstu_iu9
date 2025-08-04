package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type fast struct{ s *bufio.Scanner }

func newFast() *fast {
	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 1024), 1<<20)
	sc.Split(bufio.ScanWords)
	return &fast{s: sc}
}
func (f *fast) str() string { f.s.Scan(); return f.s.Text() }
func (f *fast) int() int    { v, _ := strconv.Atoi(f.str()); return v }

func read(f *fast) (int, int, int, [][]int, [][]string) {
	n, m, q0 := f.int(), f.int(), f.int()
	tr := make([][]int, n)
	for i := 0; i < n; i++ {
		tr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			tr[i][j] = f.int()
		}
	}
	out := make([][]string, n)
	for i := 0; i < n; i++ {
		out[i] = make([]string, m)
		for j := 0; j < m; j++ {
			out[i][j] = f.str()
		}
	}
	return n, m, q0, tr, out
}

func main() {
	f := newFast()
	n1, m1, q1, tr1, out1 := read(f)
	n2, m2, q2, tr2, out2 := read(f)
	if m1 != m2 {
		fmt.Println("NOT EQUAL")
		return
	}
	m := m1
	vis := make([]bool, n1*n2)
	qa, qb := []int{q1}, []int{q2}
	vis[q1*n2+q2] = true
	for len(qa) > 0 {
		a, b := qa[0], qb[0]
		qa, qb = qa[1:], qb[1:]
		for i := 0; i < m; i++ {
			if out1[a][i] != out2[b][i] {
				fmt.Println("NOT EQUAL")
				return
			}
			na, nb := tr1[a][i], tr2[b][i]
			id := na*n2 + nb
			if !vis[id] {
				vis[id] = true
				qa = append(qa, na)
				qb = append(qb, nb)
			}
		}
	}
	fmt.Println("EQUAL")
}

//Антиплагиат: Похожих посылок не найдено
//Комментарий преподавателя: Ачивка!
