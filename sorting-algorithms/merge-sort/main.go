package main

import "fmt"

func merge(left []int, right []int) []int {
	// MAKING THE RESULT ARRAY
	result := make([]int, len(right)+len(left))

	leftIndex, rightIndex := 0, 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}

	// APPEND ANY REMAINING ELEMENTS

	result = append(result, left[leftIndex:]...)
	result = append(result, right[rightIndex:]...)

	return result
}

func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	mid := len(numbers) / 2

	left := mergeSort(numbers[:mid])
	right := mergeSort(numbers[mid:])

	return merge(left, right)
}

func main() {
	numbers := []int{5, 4, 3, 2, 1}
	sortedArray := mergeSort(numbers)

	fmt.Println("SORTED ARRAY: ", sortedArray)
}
