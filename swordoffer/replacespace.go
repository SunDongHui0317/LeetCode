package swordoffer

func replaceSpace(s string) string {
	count := 0
	for _, v := range s {
		if v == ' ' {
			count++
		}
	}

	ans := make([]byte,len(s)+2*count)
	ansIdx := 0
	for _, v := range s {
		if v != ' ' {
			ans[ansIdx] = byte(v)
		} else {
			ans[ansIdx] = '%'
			ansIdx++
			ans[ansIdx] = '2'
			ansIdx++
			ans[ansIdx] = '0'
		}
		ansIdx++
	}
	return string(ans)
}
