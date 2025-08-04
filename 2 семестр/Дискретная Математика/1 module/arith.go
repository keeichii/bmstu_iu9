package main

import (
	"fmt"
	"os"
	"strconv"
)

type Tag int

const (
	ERROR Tag = 1 << iota
	NUMBER
	VAR
	PLUS
	MINUS
	MUL
	DIV
	LPAREN
	RPAREN
)

type Lexem struct {
	Tag   Tag
	Image string
}

var cur Lexem
var lexs = make(chan Lexem)
var varsMap = map[string]int64{}

func lexer(expr string, lexems chan Lexem) {
	for i := 0; i < len(expr); {
		switch c := expr[i]; {
		case c == ' ':
			i++
		case c >= '0' && c <= '9':
			j := i
			for j < len(expr) && expr[j] >= '0' && expr[j] <= '9' {
				j++
			}
			lexems <- Lexem{NUMBER, expr[i:j]}
			i = j
		case (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'):
			j := i
			for j < len(expr) && ((expr[j] >= 'a' && expr[j] <= 'z') ||
				(expr[j] >= 'A' && expr[j] <= 'Z') ||
				(expr[j] >= '0' && expr[j] <= '9')) {
				j++
			}
			lexems <- Lexem{VAR, expr[i:j]}
			i = j
		case c == '+':
			lexems <- Lexem{PLUS, "+"}
			i++
		case c == '-':
			lexems <- Lexem{MINUS, "-"}
			i++
		case c == '*':
			lexems <- Lexem{MUL, "*"}
			i++
		case c == '/':
			lexems <- Lexem{DIV, "/"}
			i++
		case c == '(':
			lexems <- Lexem{LPAREN, "("}
			i++
		case c == ')':
			lexems <- Lexem{RPAREN, ")"}
			i++
		default:
			lexems <- Lexem{ERROR, string(c)}
			i++
		}
	}
	close(lexems)
}

func next() {
	cur = <-lexs
}

func parseE() (int64, error) {
	v, err := parseT()
	if err != nil {
		return 0, err
	}
	return parseEPrime(v)
}

func parseEPrime(acc int64) (int64, error) {
	if cur.Tag == PLUS {
		next()
		t, err := parseT()
		if err != nil {
			return 0, err
		}
		return parseEPrime(acc + t)
	}
	if cur.Tag == MINUS {
		next()
		t, err := parseT()
		if err != nil {
			return 0, err
		}
		return parseEPrime(acc - t)
	}
	return acc, nil
}

func parseT() (int64, error) {
	v, err := parseF()
	if err != nil {
		return 0, err
	}
	return parseTPrime(v)
}

func parseTPrime(acc int64) (int64, error) {
	if cur.Tag == MUL {
		next()
		f, err := parseF()
		if err != nil {
			return 0, err
		}
		return parseTPrime(acc * f)
	}
	if cur.Tag == DIV {
		next()
		f, err := parseF()
		if err != nil {
			return 0, err
		}
		return parseTPrime(acc / f)
	}
	return acc, nil
}

func parseF() (int64, error) {
	if cur.Tag == MINUS {
		next()
		v, err := parseF()
		return -v, err
	}
	if cur.Tag == NUMBER {
		v, _ := strconv.ParseInt(cur.Image, 10, 64)
		next()
		return v, nil
	}
	if cur.Tag == VAR {
		name := cur.Image
		next()
		if val, ok := varsMap[name]; ok {
			return val, nil
		}
		var v int64
		fmt.Fscan(os.Stdin, &v)
		varsMap[name] = v
		return v, nil
	}
	if cur.Tag == LPAREN {
		next()
		v, err := parseE()
		if err != nil || cur.Tag != RPAREN {
			return 0, fmt.Errorf("error")
		}
		next()
		return v, nil
	}
	return 0, fmt.Errorf("error")
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	expr := os.Args[1]
	go lexer(expr, lexs)
	next()
	res, err := parseE()
	if err != nil || cur.Tag != 0 {
		fmt.Fprint(os.Stdout, "error")
		return
	}
	fmt.Fprint(os.Stdout, res)
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
