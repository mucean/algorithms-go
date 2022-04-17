package str

import (
	"sort"
)

var Permutation = permutationByNext

func permutationRecursive(s string) []string {
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

func permutationByNext(s string) []string {
	ans := make([]string, 0)
	l := len(s)
	if l == 0 {
		return ans
	}
	bts := []byte(s)
	sort.Slice(bts, func(i, j int) bool {
		return bts[i] < bts[j]
	})

	for {
		ans = append(ans, string(bts))
		if nextPermutation(bts, l) == false {
			return ans
		}
	}
}

func nextPermutation(bts []byte, l int) bool {
	pos := l - 2
	for ; pos >= 0; pos-- {
		if bts[pos] < bts[pos+1] {
			break
		}
	}
	if pos < 0 {
		return false
	}
	j := pos + 1
	for ; j < l; j++ {
		if bts[pos] >= bts[j] {
			break
		}
	}
	j--
	bts[pos], bts[j] = bts[j], bts[pos]

	s, e := pos+1, l-1
	for s < e {
		bts[s], bts[e] = bts[e], bts[s]
		s++
		e--
	}
	return true
}
