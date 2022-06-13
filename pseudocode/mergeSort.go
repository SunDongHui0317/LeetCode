package pseudocode

func mergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	mid := len(list) / 2
	left := mergeSort(list[0:mid])
	right := mergeSort(list[mid:])
	res := merge(left, right)
	return res
}

func merge(list1, list2 []int) []int {
	res := make([]int, 0, len(list1)+len(list2))
	i, j := 0, 0
	l, r := len(list1), len(list2)
	for i < l && j < r {
		if list1[i] > list2[j] {
			res = append(res, list2[j])
			j++
		} else {
			res = append(res, list1[i])
			i++
		}
	}
	res = append(res, list2[j:]...)
	res = append(res, list1[i:]...)
	return res
}
