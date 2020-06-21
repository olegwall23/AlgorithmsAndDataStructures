package main

import "fmt"

func main() {
	var arr = [][]int {
		{-4, -3, -2},
		{-1, -1, -1},
		{-1, 0, 0},
	}

	cN := count(arr, len(arr[0]), len(arr))
	fmt.Println(cN)
}

func count(matrix [][]int, n int, m int) int {
	count := 0
	x := m - 1
	y := 0

	for x >= 0 && y < n {
		if matrix[y][x] < 0 {
			count += x + 1
			y++
		} else {
			x--
		}
	}

	return count
}