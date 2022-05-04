package swordoffer

func findRepeatNumber(nums []int) int {
	m := make(map[int]int)
	for _, v := range nums {
		if m[v] > 0 {
			return v
		}
		m[v]++
	}
	return 0
}
