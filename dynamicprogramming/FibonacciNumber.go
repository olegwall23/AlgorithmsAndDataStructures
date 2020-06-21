package main

import "fmt"

func main() {
	fmt.Println(findNumber(15))
}

func findNumber(n int) int {
	numbers := make([]int, 2)
	numbers[0] = 1
	numbers[1] = 1
	for i := 2; i < n; i++ {
		fn := numbers[0] + numbers[1]
		numbers[0] = numbers[1]
		numbers[1] = fn
	}
	return numbers[1]
}