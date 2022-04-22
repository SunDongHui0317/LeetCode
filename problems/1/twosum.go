package twosum

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, item := range nums {
		hash[item] = i+1
	}
	for i, item := range nums {
		v := target - item
		if hash[v] > 0 && (hash[v] -1) != i{
			return []int{
				i,
				hash[v]-1,
			}
		}
	}
	return []int{}
}

