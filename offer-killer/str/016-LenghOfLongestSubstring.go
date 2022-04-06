package str

var LengthOfLongestSubstring = lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l < 2 {
		return l
	}

	counts := make([]int, 256)

	p1, p2 := 0, 0

	maxLen := 0
	var index byte
	var i int
	for ; p2 < l; p2++ {
		index = s[p2] - 'a'
		if counts[index] > 0 {
			if maxLen < p2-p1 {
				maxLen = p2 - p1
			}

			for i, p1 = p1, counts[index]; i < p1; i++ {
				counts[s[i]-'a'] = 0
			}
		}
		counts[index] = p2 + 1
	}

	return maxLen
}
