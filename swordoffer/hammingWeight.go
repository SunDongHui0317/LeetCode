package swordoffer

func hammingWeight(num uint32) int {
	r := 0
	for num > 0 {
		if h := num % 2; h == 1 {
			r++
		}

		num = num / 2
	}
	if num == 1 {
		r++
	}
	return r
}
