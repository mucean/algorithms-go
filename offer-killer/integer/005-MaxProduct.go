package integer

var MaxProduct = myMaxProduct

func myMaxProduct(words []string) int {
	l := len(words)
	nums := make([]int, l)

	var left byte
	for i, word := range words {
		for _, c := range []byte(word) {
			left = c - 'a'
			if left == 0 {
				continue
			}
			nums[i] ^= 1 << left
		}
	}

	res := 0
	for i := 1; i < l; i++ {
		for j := i; j < l; j++ {
			if nums[i-1]&nums[j] == 0 {
				n := len(words[i-1]) * len(words[j])
				if n > res {
					res = n
				}
			}
		}
	}
	return res
}
