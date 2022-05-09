package swordoffer

func numWays(n int) int {
	const mod = 1e9 + 7
	if n == 0 {
		return 1
	}
	if n < 3 {
		return n
	}
	p, q, r := 0, 1, 2
	for i := 3; i <= n; i++ {
		p = q
		q = r
		r = (p + q) % mod
	}
	return r
}
