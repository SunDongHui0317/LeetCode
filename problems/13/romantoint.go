package romantoint

func romanToInt(s string) int {
	mapRomanInt := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}
	r := 0
	preNum := mapRomanInt[s[0:1]]
	for i := 1; i < len(s); i++ {
		num := mapRomanInt[s[i:i+1]]
		if num > preNum {
			r -= preNum
		} else {
			r += preNum
		}
		preNum = num
	}
	r += preNum
	return r
}
