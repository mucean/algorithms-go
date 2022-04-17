package p50

var TotalNQueens = totalNQueens

func totalNQueens(n int) int {
	queens := make([]int, n)
	for i := 1; i < n; i++ {
		queens[i] = i
	}

	total := 0
	var recursive func(p int)
	recursive = func(p int) {
		if p == n {
			for i := 0; i < n-1; i++ {
				for j := i + 1; j < n; j++ {
					t1, t2 := j-i, queens[j]-queens[i]
					if t1 == t2 || t1 == -t2 {
						return
					}
				}
			}
			total++
			return
		}
		for i := p; i < n; i++ {
			queens[i], queens[p] = queens[p], queens[i]
			recursive(p + 1)
			queens[i], queens[p] = queens[p], queens[i]
		}
	}

	recursive(0)

	return total
}
