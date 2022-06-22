package swordoffer

func exchange(nums []int) []int {
	oddList := make([]int, 0, len(nums))
	evenList := make([]int, 0, len(nums))
	for _, v := range nums {
		if v&1 == 1 {
			oddList = append(oddList, v)
		} else {
			evenList = append(evenList, v)
		}
	}
	oddList = append(oddList, evenList...)
	return oddList
}

func exchange2(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]&1 == 1 {
			l += 1
		}
		for l < r && nums[r]&1 == 0 {
			r -= 1
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	return nums
}
