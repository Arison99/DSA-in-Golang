package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Initialize the slow pointer
	slow := 0

	// Iterate through the array with the fast pointer
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	// The number of unique elements is slow + 1
	return slow + 1
}

func main() {
	// Example usage
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums)
	fmt.Println(k)        // Output: 5
	fmt.Println(nums[:k]) // Output: [0, 1, 2, 3, 4]
}
