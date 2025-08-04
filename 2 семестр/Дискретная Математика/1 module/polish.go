package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func eval(s string, i *int) int {
	for *i < len(s) && s[*i] == ' ' {
		*i++
	}
	if *i < len(s) && s[*i] >= '0' && s[*i] <= '9' {
		v := int(s[*i] - '0')
		*i++
		return v
	}
	if *i < len(s) && s[*i] == '(' {
		*i++
		for *i < len(s) && s[*i] == ' ' {
			*i++
		}
		op := s[*i]
		*i++
		left := eval(s, i)
		right := eval(s, i)
		for *i < len(s) && s[*i] == ' ' {
			*i++
		}
		*i++
		switch op {
		case '+':
			return left + right
		case '-':
			return left - right
		default:
			return left * right
		}
	}
	return 0
}

func main() {
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	s = strings.TrimSpace(s)
	idx := 0
	fmt.Fprint(os.Stdout, eval(s, &idx))
}

//Антиплагиат: Похожих посылок не найдено
//Комментарий преподавателя:
