package multiply

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	answerArr := make([]int, len(num1)+len(num2))
	for index1, n1 := range num1 {
		for index2, n2 := range num2 {
			answerArr[index1+index2+1] += int(n1-'0') * int(n2-'0')
		}
	}
	for i := len(answerArr)-1; i > 0; i-- {
		if answerArr[i] >= 10 {
			answerArr[i-1] += answerArr[i] / 10
			answerArr[i] %= 10
		}
	}

	ans := ""
	for index, v := range answerArr {
		if index == 0 && v == 0{
			continue
		}
		ans += strconv.Itoa(v)
	}
	return ans
}
