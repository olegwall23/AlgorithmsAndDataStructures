package main

import "fmt"

func main() {
	arr1 := []int{8, 8, 8, 8, 8}
	arr2 := []int{1, 0, 1, 1, 1}
	var result []int
	result = multiplier(arr1, arr2)

	fmt.Println(result)
}

func multiplier(arr1[]int, arr2[]int) []int {
	fmt.Println(arr1)
	fmt.Println(arr2)

	var result = make([]int, getMax(len(arr1), len(arr2)))

	var startFromForArray1 = len(arr1) - 1
	var startFromForArray2 = len(arr2) - 1

	var addToNextElement = 0
	var sum = 0
	var indexInResult = len(result) - 1

	for startFromForArray1 >= 0 && startFromForArray2 >= 0 {
		sum = arr1[startFromForArray1] + arr2[startFromForArray2] + addToNextElement
		addToNextElement = sum / 10
		result[indexInResult] = sum % 10

		indexInResult--
		startFromForArray1--
		startFromForArray2--
	}

	for startFromForArray1 >= 0 {
		sum = arr1[startFromForArray1] + addToNextElement
		addToNextElement = sum / 10
		result[indexInResult] = sum % 10

		startFromForArray1--
		indexInResult--
	}

	for startFromForArray2 >= 0 {
		sum = arr1[startFromForArray2] + addToNextElement
		addToNextElement = sum / 10
		result[indexInResult] = sum % 10

		startFromForArray2--
		indexInResult--
	}

	if addToNextElement > 0 {
		var resultTmp = make([]int, len(result) + 1)
		var v = 1
		for v <= len(result) {
			resultTmp[len(resultTmp) - v] = result[len(result) - v]
			v++
		}
		resultTmp[0] = addToNextElement
		return resultTmp
	}

	return result
}


func getMax(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
