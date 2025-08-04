package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	val  string
	left *node
	right *node
}

var pos int
var memo map[string]int

func parse(s string) *node {
	if s[pos] == '(' {
		pos++
		op := string(s[pos])
		pos++
		left := parse(s)
		right := parse(s)
		pos++
		return &node{val: op, left: left, right: right}
	} else {
		ch := string(s[pos])
		pos++
		return &node{val: ch}
	}
}

func hash(t *node) string {
	if t.left == nil && t.right == nil {
		return t.val
	}
	return t.val + "(" + hash(t.left) + ")(" + hash(t.right) + ")"
}

func count(t *node) int {
	h := hash(t)
	if _, ok := memo[h]; ok {
		return 0
	}
	memo[h] = 1
	if t.left == nil && t.right == nil {
		return 0
	}
	return 1 + count(t.left) + count(t.right)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	expr := scanner.Text()
	pos = 0
	memo = make(map[string]int)
	tree := parse(expr)
	fmt.Println(count(tree))
}

//Антиплагиат: Похожих посылок не найдено
//Комментарий преподавателя:
