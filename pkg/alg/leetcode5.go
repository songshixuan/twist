package alg

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	max := 1 // every single char
	res := s[:1]
	dp := make([][]bool, len(s))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s))
	}
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {

				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}

			}
			if dp[i][j] {
				if max < j-i+1 {
					max = j - i + 1
					res = s[i : j+1]
				}
			}
		}

	}

	return res
}
