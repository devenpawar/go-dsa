package main

import "fmt"

func twoSum(nums []int, target int) []int {
	complement := 0
	seen := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		complement = target - nums[i]
		if prevIndex, found := seen[complement]; found {
			return []int{prevIndex, i}
		}
		// Store current number and it's index
		seen[nums[i]] = i
	}

	return nil
}

func main() {
	result := twoSum([]int{5, 3, 4, 1}, 9)
	fmt.Println("The result is: ", result)
}
