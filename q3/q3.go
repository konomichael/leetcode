package q3

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	l, r := 0, 0
	seen := make(map[byte]int)
	for r < len(s) {
		r++
		rc := s[r-1]
		seen[rc]++

		if seen[rc] > 1 {
			for l < r {
				lc := s[l]
				l++
				seen[lc]--
				if lc == rc {
					break
				}
			}
		}

		if length := r-l; length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
