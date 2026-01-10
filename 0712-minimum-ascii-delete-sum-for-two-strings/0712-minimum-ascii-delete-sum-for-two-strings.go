func minimumDeleteSum(s1 string, s2 string) int {
	n, m := len(s1), len(s2)

	if n < m {
		return minimumDeleteSum(s2, s1)
	}

	prev := make([]int, m+1)
	curr := make([]int, m+1)

	for j := 1; j <= m; j++ {
		prev[j] = prev[j-1] + int(s2[j-1])
	}

	for i := 1; i <= n; i++ {
		curr[0] = prev[0] + int(s1[i-1])

		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				curr[j] = prev[j-1]
			} else {
				deleteS1 := prev[j] + int(s1[i-1])
				deleteS2 := curr[j-1] + int(s2[j-1])
				
				if deleteS1 < deleteS2 {
					curr[j] = deleteS1
				} else {
					curr[j] = deleteS2
				}
			}
		}
		copy(prev, curr)
	}

	return prev[m]
}