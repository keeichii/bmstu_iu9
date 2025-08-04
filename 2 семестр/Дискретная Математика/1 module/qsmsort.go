package main

import "fmt"

func qssort(n int, less func(i, j int) bool, swap func(i, j int)) {
	type interval struct {
		low, high int
	}

	stack := []interval{{0, n - 1}}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		low, high := curr.low, curr.high

		if low >= high {
			continue
		}

		pivotIndex := partition(low, high, less, swap)

		if pivotIndex-1 > low {
			stack = append(stack, interval{low, pivotIndex - 1})
		}
		if pivotIndex+1 < high {
			stack = append(stack, interval{pivotIndex + 1, high})
		}
	}
}

func partition(low, high int, less func(i, j int) bool, swap func(i, j int)) int {
	pivot := high
	i := low - 1

	for j := low; j < high; j++ {
		if less(j, pivot) {
			i++
			swap(i, j)
		}
	}
	swap(i+1, high)
	return i + 1
}

func main() {
	arr := []int{10, 7, 8, 9, 1, 5}
	n := len(arr)

	less := func(i, j int) bool { return arr[i] < arr[j] }
	swap := func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }

	qssort(n, less, swap)

	fmt.Println(arr)
}

//Антиплагиат: Найдены похожие посылки
//Комментарий преподавателя:
