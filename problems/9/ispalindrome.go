package ispalindrome

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	list := []int{}
	for x >= 10 {
		list = append(list, x%10)
		x = x / 10
	}
	list = append(list, x)
	half := len(list) / 2
	for i := 0; i <= half; i++ {
		if list[i] != list[len(list)-1-i] {
			return false
		}
	}
	return true
}
