package myatoi

import (
	"math"
)

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	i, res, flag := 0, 0, 1
	for s[i] == ' ' {
		i++
		if i >= len(s) {
			return 0
		}
	}
	if s[i] == '-' {
		flag= -1
		i++
	} else if s[i] == '+'{
		i++
	}

	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		res = res*10 + int(s[i]-'0')
		i++
		if res > math.MaxInt32 {
			if flag > 0 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
	}
	return res * flag
}

func MyAtoi(str string) int {
	var ans int64 = 0
	var sign int64 = 1
	start := false

	for _, b := range str {
		if b <= '9' && b >= '0' {
			if !start {start = true}
			ans = ans * 10 + int64(b) - int64('0')
			if ans * sign > math.MaxInt32 {
				return math.MaxInt32
			} else if ans * sign < math.MinInt32 {
				return math.MinInt32
			}
		} else if !start && b == '+' {
			start = true
		} else if !start && b == '-' {
			start = true
			sign = -1
		} else if !start && b == ' '{
			continue
		} else {
			break
		}
	}
	return int(ans * sign)
}