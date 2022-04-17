package p50

var MinDistance = minDistance

func minDistance(word1, word2 string) int {
	l1 := len(word1)
	l2 := len(word2)
	if l1 == 0 {
		return l2
	}
	if l2 == 0 {
		return l1
	}

	dp := make([][]int, l1+1)
	for i := 0; i <= l1; i++ {
		dp[i] = make([]int, l2+1)
		dp[i][0] = i
	}
	for i := 1; i <= l2; i++ {
		dp[0][i] = i
	}

	var min, tmp int
	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				min = dp[i][j-1] + 1
				if tmp = dp[i-1][j] + 1; tmp < min {
					min = tmp
				}
				if tmp = dp[i-1][j-1] + 1; tmp < min {
					min = tmp
				}
				dp[i][j] = min
			}
		}
	}

	return dp[l1][l2]
}
