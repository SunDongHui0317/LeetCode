package swordoffer

func myPow(x float64, n int) float64 {
	if n > 0 {
		return quickMul(x, n)
	}
	return 1 / quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	ans := 1.0
	for n > 0 {
		if d := n % 2; d == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}
	return ans
}

// myPow1 超时了
func myPow1(x float64, n int) float64 {
	r := 1.0
	if n > 0 {
		for n > 0 {
			n -= 1
			r = r * x
		}
		return r
	}

	for n < 0 {
		n += 1
		r = r * x
	}

	return 1 / r
}
