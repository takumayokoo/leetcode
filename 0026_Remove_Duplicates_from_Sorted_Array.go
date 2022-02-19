package main

func removeDuplicates(nums []int) int {
	val := nums[0]
	index := 0

	for i := 1; i < len(nums); i++ {
		if val != nums[i] {
			index = index + 1
			nums[index] = nums[i]
			val = nums[i]
		}
	}

	return index + 1
}
