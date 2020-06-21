package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 2, 5, 7, 1}
	fmt.Println(fmt.Sprintf("%s%d", "one time: ", oneTimeProfit(arr)))
	fmt.Println(fmt.Sprintf("%s%d", "all time: ", allTimeProfit(arr)))
}

func allTimeProfit(arr []int) int {
	var allTimeProfit = 0
	for i := 1; i < len(arr); i++ {
		if arr[i - 1] < arr[i] {
			allTimeProfit += arr[i] - arr[i - 1]
		}
	}
	return allTimeProfit
}

func oneTimeProfit(arr []int) int {
	minResult := arr[0]
	profit := 0
	for i := 1; i < len(arr); i++ {
		if arr[i] - minResult > profit {
			profit = arr[i] - minResult
		}
		if arr[i] < minResult {
			minResult = arr[i]
		}
	}
	return profit
}