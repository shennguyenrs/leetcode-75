package first

import (
	"log"
)

func runningSum(nums []int) []int {
	arrLength := len(nums)
	result := make([]int, arrLength)
	result[0] = nums[0]

	for i := 1; i < arrLength; i += 1 {
		result[i] = nums[i] + result[i-1]
	}

	return result
}

func pivotIndex(nums []int) int {
	arrLength := len(nums)
	var total int = 0
	var left int = 0

	for _, n := range nums {
		total += n
	}

	// Check if first place is the mid
	if total == nums[0] {
		return 0
	}

	// Check the next value is the mid
	for i := 1; i < arrLength; i += 1 {
		left += nums[i-1]

		if left*2 == total-nums[i] {
			return i
		}
	}

	return -1
}

func Run() {
	testFirst := []int{1, 1, 1, 1, 1}
	first := runningSum(testFirst)

	testSecond := []int{1, 2, 3}
	second := pivotIndex(testSecond)

	log.Println("Testing first task", first)
	log.Println("Testing second task", second)
}
