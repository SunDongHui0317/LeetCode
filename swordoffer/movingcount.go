package swordoffer

type mov struct {
	x, y int
}

var movingDirections = []mov{
	{
		0, 1,
	}, {
		0, -1,
	}, {
		1, 0,
	}, {
		-1, 0,
	},
}

func movingCount(m int, n int, k int) int {
	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	return dfs(dp, m, n, k, 0, 0)
}

func dfs(dp [][]bool, m, n, k, i, j int) int {
	if i < 0 || j < 0 || i >= m || j >= n || dp[i][j] || (sumOps(i)+sumOps(j)) > k {
		return 0
	}
	dp[i][j] = true
	sum := 1
	for _, v := range movingDirections {
		newI := i + v.x
		newJ := j + v.y
		sum += dfs(dp, m, n, k, newI, newJ)
	}
	return sum

}

func sumOps(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
