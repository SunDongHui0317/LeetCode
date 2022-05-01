package maxarea

/*暴力破解，超出输出限制*/
func maxArea1(height []int) int {
	max := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			h := height[i]
			if height[i] > height[j] {
				h = height[j]
			}
			if h*(j-i) > max {
				max = h * (j - i)
			}
		}
	}
	return max
}

func maxArea(h []int) int {
	max := 0

	l := 0
	r := len(h) - 1

	for l < r {
		min := h[l]
		if h[l] > h[r] {
			min = h[r]
		}

		area := min * (r - l)
		if area > max {
			max = area
		}

		if h[l] > h[r] {
			r--
		} else {
			l++
		}
	}

	return max
}
