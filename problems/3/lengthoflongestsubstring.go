package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	m := make(map[uint8]int)
	l, r, longest := 0, 0, 0
	for r < len(s) {
		vr := s[r]
		if times := m[vr]; times > 0 {
			vl := s[l]
			m[vl]--
			l++
			continue
		}
		if r-l+1 > longest {
			longest = r - l + 1
		}
		m[vr]++
		r++
	}
	return longest
}

func LengthOfLongestSubstring(s string) int {
	list := make([]int32, 0, len(s))
	l, r, longest := 0, 0, 0
	for r < len(s) {
		vr := int32(s[r])
		if isSliceHasValue(list, vr) {
			list = slicePop(list)
			l++
			continue
		}
		if r-l+1 > longest {
			longest = r - l + 1
		}
		list = append(list, vr)
		r++
	}
	return longest
}

func slicePop(slice []int32) []int32 {
	if len(slice) < 1 {
		return nil
	}
	return slice[1:]
}

func isSliceHasValue(slice []int32, v int32) bool {
	for _, s := range slice {
		if s == v {
			return true
		}
	}
	return false
}

func removeValueInSlice(slice []int32, s int32) []int32 {
	for i, v := range slice {
		if v == s {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
