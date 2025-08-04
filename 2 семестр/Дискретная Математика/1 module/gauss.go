package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)

	m := make([][]*big.Rat, n)
	for i := range m {
		m[i] = make([]*big.Rat, n+1)
		for j := range m[i] {
			var num int
			fmt.Scan(&num)
			m[i][j] = big.NewRat(int64(num), 1)
		}
	}

	for k := 0; k < n; k++ {
		if m[k][k].Num().Int64() == 0 {
			for i := k + 1; i < n; i++ {
				if m[i][k].Num().Int64() != 0 {
					m[k], m[i] = m[i], m[k]
					break
				}
			}
			if m[k][k].Num().Int64() == 0 {
				if m[k][n].Num().Int64() == 0 {
					continue
				} else {
					fmt.Println("No solution")
					return
				}
			}
		}

		for i := k + 1; i < n; i++ {
			f := new(big.Rat).Quo(m[i][k], m[k][k])
			for j := k; j < n+1; j++ {
				m[i][j].Sub(m[i][j], new(big.Rat).Mul(f, m[k][j]))
			}
		}
	}

	for i := n - 1; i >= 0; i-- {
		if m[i][i].Num().Int64() == 0 {
			if m[i][n].Num().Int64() == 0 {
				continue
			} else {
				fmt.Println("No solution")
				return
			}
		}

		for j := i + 1; j < n; j++ {
			m[i][n].Sub(m[i][n], new(big.Rat).Mul(m[i][j], m[j][n]))
		}
		m[i][n].Quo(m[i][n], m[i][i])
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%s/%s\n", m[i][n].Num(), m[i][n].Denom())
	}
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя: 🐀
