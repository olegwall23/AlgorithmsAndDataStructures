package main

import "fmt"

func main() {
	fmt.Println(findMax([]int{1, -1, 2, -2, 3}))
}

func findMax(arr[]int) int {
	currentMax, currentGlobal := arr[0], arr[0]
	for i := 1; i < len(arr); i++ {
		currentMax = getMax(arr[i], currentMax + arr[i])
		if currentMax > currentGlobal {
			currentGlobal = currentMax
		}
	}
	return currentGlobal
}

func getMax(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}