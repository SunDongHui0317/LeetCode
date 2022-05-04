package inttoroman

func intToRoman(num int) string {
	// 数组哈希减法拼接
	intArr := []int{1000,900,500,400,100,90,50,40,10,9,5,4,1}
	romanArr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	roman := ""
	// 从大往小比较，通过减去最大的，每次拼接该字符串得到ans
	for i:=0;i<len(intArr);i++ {
		for num >= intArr[i] {
			num -= intArr[i]
			roman += romanArr[i]
		}
		if num == 0 {
			break
		}
	}
	return roman
}