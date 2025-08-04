package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type AssocArray interface {
	Assign(s string, x int)
	Lookup(s string) (int, bool)
}

type mapArray struct {
	data map[string]int
}

func (m *mapArray) Assign(s string, x int) {
	m.data[s] = x
}

func (m *mapArray) Lookup(s string) (int, bool) {
	v, ok := m.data[s]
	return v, ok
}

func makeSkipList() AssocArray {
	return &mapArray{data: make(map[string]int)}
}

func makeAVL() AssocArray {
	return &mapArray{data: make(map[string]int)}
}

func lex(sentence string, array AssocArray) []int {
	tokens := tokenize(sentence)
	result := []int{}
	counter := 1

	for _, tok := range tokens {
		if val, ok := array.Lookup(tok); ok {
			result = append(result, val)
		} else {
			array.Assign(tok, counter)
			result = append(result, counter)
			counter++
		}
	}
	return result
}

func tokenize(s string) []string {
	fields := strings.Fields(s)
	valid := []string{}
	for _, word := range fields {
		if isIdent(word) {
			valid = append(valid, word)
		}
	}
	return valid
}

func isIdent(s string) bool {
	if len(s) == 0 || !unicode.IsLetter(rune(s[0])) {
		return false
	}
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func main() {
	sentence := "  alpha x1 beta alpha x1 y   "
	skipList := makeSkipList()
	avl := makeAVL()

	fmt.Println(toString(lex(sentence, skipList)))
	fmt.Println(toString(lex(sentence, avl)))
}

func toString(arr []int) string {
	parts := make([]string, len(arr))
	for i, x := range arr {
		parts[i] = strconv.Itoa(x)
	}
	return strings.Join(parts, " ")
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: Надо было реализовать АВЛ-дерево и список с пропусками.
