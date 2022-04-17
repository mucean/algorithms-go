package str

var Permutation = permutation

func permutation(s string) []string {
	bts := []byte(s)
	l := len(bts)
	if l == 0 {
		return []string{}
	}
	ans := make([]string, 0)
	items := make([]byte, l)
	var recursive func(p int)
	recursive = func(p int) {
		if p == l {
			ans = append(ans, string(items))
			return
		}
		for i := p; i < l; i++ {
			bts[i], bts[p] = bts[p], bts[i]
			items[p] = bts[p]
			recursive(p + 1)
			bts[i], bts[p] = bts[p], bts[i]
		}
	}

	recursive(0)

	return ans
}
