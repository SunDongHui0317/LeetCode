package pseudocode

// 插入排序
func insertionSort(list []int) []int {
	n := len(list)
	a := make([]int, n, n)
	copy(a, list) // 不改变原切片值
	if n <= 1 {
		return a
	}
	for i := 1; i < n; i++ {
		v := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > v {
				a[j+1] = a[j]
			} else {
				break
			}
		}
		a[j+1] = v
	}
	return a
}
