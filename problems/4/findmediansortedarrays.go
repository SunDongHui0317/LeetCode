package findmediansortedarrays

/*查找中位数*/
/*Runtime: 19 ms, faster than 47.55% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 5.1 MB, less than 81.02% of Go online submissions for Median of Two Sorted Arrays.*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	sum := len(nums1) + len(nums2)
	if sum%2 == 1 {
		return float64(getDivLine(nums1, nums2, sum/2+1))
	}
	return float64(getDivLine(nums1, nums2, sum/2+1)+getDivLine(nums1, nums2, sum/2)) / 2.0

}

func getDivLine(nums1, nums2 []int, need int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+need-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+need-1]
		}
		if need == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := need / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			need -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			need -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

/*Runtime: 26 ms, faster than 22.48% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 5.1 MB, less than 81.02% of Go online submissions for Median of Two Sorted Arrays.*/
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	if (m+n)%2 == 0 {
		val1 := findKth(nums1, nums2, (m+n)/2)
		val2 := findKth(nums1, nums2, (m+n)/2+1)
		return (float64(val1) + float64(val2)) * 0.5
	} else {
		val := findKth(nums1, nums2, (m+n+1)/2)
		return float64(val)
	}
}

func findKth(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if len(nums1) > len(nums2) {
		return findKth(nums2, nums1, k)
	}
	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}
	mi := k >> 1
	mi1, mi2 := mi-1, mi-1
	if mi1 >= len(nums1) {
		mi1 = len(nums1) - 1
	}
	if nums1[mi1] <= nums2[mi2] {
		nums1 := nums1[mi1+1:]
		k = k - (mi1 + 1)
		return findKth(nums1, nums2, k)
	} else {
		nums2 := nums2[mi2+1:]
		k = k - (mi2 + 1)
		return findKth(nums1, nums2, k)
	}
}

/* 二分查找 */
/*Runtime: 10 ms, faster than 88.38% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 5 MB, less than 90.27% of Go online submissions for Median of Two Sorted Arrays.*/
func findMedianSortedArrays3(nums1, nums2 []int) float64 {
	Infinity := 100000000000 // 范围
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1 // 无论合适，确保 nums1 为小数组，nums2 为大数组
	}
	len1, len2 := len(nums1), len(nums2)
	low := 0
	high := len1
	for low <= high {
		i := low + (high-low)/2 //
		j := (len1+len2+1)/2 - i

		maxLeftNums1, minRightNums1, maxLeftNums2, minRightNums2 := 0, 0, 0, 0
		if i == 0 {
			maxLeftNums1 = -Infinity
		} else {
			maxLeftNums1 = nums1[i-1]
		}

		if i == len1 {
			minRightNums1 = Infinity
		} else {
			minRightNums1 = nums1[i]
		}

		if j == 0 {
			maxLeftNums2 = -Infinity
		} else {
			maxLeftNums2 = nums2[j-1]
		}

		if j == len2 {
			minRightNums2 = Infinity
		} else {
			minRightNums2 = nums2[j]
		}

		if (maxLeftNums1 <= minRightNums2) && (minRightNums1 >= maxLeftNums2) {
			if (len1+len2)%2 == 1 {
				return float64(max(maxLeftNums1, maxLeftNums2))
			}
			return float64(max(maxLeftNums1, maxLeftNums2)+min(minRightNums1, minRightNums2)) / 2
		} else if maxLeftNums1 > minRightNums2 {
			high = i - 1
		} else {
			low++
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}