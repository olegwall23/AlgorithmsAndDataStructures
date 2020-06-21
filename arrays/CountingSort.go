package main

import "fmt"

func main() {
	var arr = []int{1, 0, 5, 3, 1, 3, 1}
	fmt.Println(sort(arr, 5))
}

func sort(arr []int , maxNumber int) []int {
	var arrPosMap = make([]int, maxNumber + 1)

	for i := 0; i < len(arr); i++ {
		arrPosMap[arr[i]]++
	}

	for i := 1; i < len(arrPosMap); i++ {
		arrPosMap[i] = arrPosMap[i - 1] + arrPosMap[i]
	}

	for i := len(arrPosMap) - 1; i > 0; i-- {
		arrPosMap[i] = arrPosMap[i - 1]
	}
	arrPosMap[0] = 0

	var sortedArray = make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		position := arrPosMap[arr[i]]
		sortedArray[position] = arr[i]
		arrPosMap[arr[i]]++
	}

	return sortedArray
}