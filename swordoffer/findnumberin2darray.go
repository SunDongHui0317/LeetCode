package swordoffer

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) < 1 {
		return false
	}
	for _, list := range matrix {
		if len(list) < 1 {
			return false
		}
		if list[0] > target {
			break
		}
		if binSearch(list, target, 0, len(list)) {
			return true
		}
	}
	return false
}

// 二分法
func binSearch(list []int, target, sIdx, eIdx int) bool {
	mid := (sIdx + eIdx) / 2
	for !(sIdx == eIdx) {
		if list[mid] < target {
			sIdx = mid
		} else if list[mid] > target {
			eIdx = mid
		} else {
			return true
		}
		mid = (sIdx + eIdx) / 2
		if sIdx == mid || eIdx == mid {
			return list[mid] == target
		}
	}
	return false
}
