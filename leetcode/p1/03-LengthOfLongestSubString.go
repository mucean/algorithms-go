package p1

var LengthOfLongestSubString = lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l <= 1 {
		return l
	}

	p1, p2 := 0, 0
	e := make(map[byte]int)
	res := 0

	for ; p2 < l; p2++ {
		if p, ok := e[s[p2]]; ok {
			for p1 <= p {
				delete(e, s[p1])
				p1++
			}
		}
		e[s[p2]] = p2
		if newRes := p2 - p1 + 1; newRes > res {
			res = newRes
		}
	}

	return res
}
