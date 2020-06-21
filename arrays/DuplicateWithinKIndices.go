package main

import "fmt"

func main() {
	elements := []int{1, 3, 4, 1}
	fmt.Println(duplicate(elements, 2))
}

func duplicate(arr[]int, k int) bool {
	var viewedElements = map[int]bool{5: true, 2: true}

	for i := 0; i < len(arr); i++ {
		_, ok := viewedElements[arr[i]]
		if ok {
			return true
		}
		if i >= k {
			delete(viewedElements, arr[i - k])
		}
		viewedElements[arr[i]] = true
	}
	return false
}