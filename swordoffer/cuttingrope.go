package swordoffer

import "math"

func cuttingRope(n int) int {
	if n <= 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	// 能分成几份3
	parts := n / 3
	another := n % 3
	var result float64

	switch another {
	case 2:
		result = math.Pow(3, float64(parts))
		result *= 2
	case 1:
		result = math.Pow(3, float64(parts-1))
		result *= 4
	default:
		result = math.Pow(3, float64(parts))
	}
	return int(result)
}

func cuttingRope1(n int) int {
	dp := make(map[int]int)
	dp[1] = 1 // 首项
	for i := 2; i <= n; i++ {
		j, k := 1, i-1
		res := 0
		for j <= k {
			res = max(res, max(j, dp[j])*max(k, dp[k])) // 递推公式
			j++
			k--
		}
		dp[i] = res
	}
	return dp[n]
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	r := 1
	for n > 4 {
		r = r * 3 % 1000000007
		n -= 3
	}
	return r * n % 1000000007
}
