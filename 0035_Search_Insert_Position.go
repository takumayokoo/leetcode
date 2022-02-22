package main

func searchInsert(nums []int, target int) int {
	length := len(nums)
	if length == 1 {
		v := nums[0]
		if target <= v {
			return 0
		} else {
			return 1
		}
	}

	m := int((length) / 2)
	v := nums[m]
	if target < v {
		return searchInsert(nums[0:m], target)
	} else {
		return m + searchInsert(nums[m:], target)
	}
}