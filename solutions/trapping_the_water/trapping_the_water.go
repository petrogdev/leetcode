package trapping_the_water

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	if height == nil {
		return 0
	}

	leftMax := make([]int, len(height))
	if len(leftMax) != 0 {
		leftMax[0] = height[0]
	}

	rightMax := make([]int, len(height))
	if len(height) > 1 {
		rightMax[len(height)-1] = height[len(height)-1]
	}

	for i:=1; i<len(height); i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}

	for i:=len(height)-2; i>=0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}

	totalWater := 0
	for i:=0; i < len(height)-1; i++ {
		totalWater += min(leftMax[i], rightMax[i]) - height[i]
	}

	return totalWater
}
