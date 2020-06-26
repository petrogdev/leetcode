package two_sums

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	input := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := input[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		input[nums[i]] = i
	}
	result := []int{-1, -1}
	return result
}

func main() {
	nums := []int{11, 15, 2, 7}
	target := 9

	fmt.Printf("result", twoSum(nums, target))
}