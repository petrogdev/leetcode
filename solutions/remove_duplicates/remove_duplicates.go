package remove_duplicates

func removeDuplicates(nums []int) int {
	for index := 1 ;index < len(nums); index ++ {
		if nums[index-1] == nums[index]  {
			nums = append(nums[:index], nums[index+1:]...)
			index --
		}
	}
	return len(nums)
}
