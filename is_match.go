package leetcode

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	f := make([][]bool, m+1)
	for i := range f {
		f[i] = make([]bool, n+1)
	}

	var match = func(s, p byte) bool {
		if s == p || p == '.' {
			return true
		}
		return false
	}

	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 0 {
				continue
			}

			if p[j-1] != '*' {
				f[i][j] = f[i-1][j-1] && match(s[i-1], p[j-1])
				continue
			}

			if p[j-2] != s[i-1] {
				f[i][j] = f[i][j-2]
			} else {
				f[i][j] = f[i-1][j] && f[i][j-2]
			}
		}
	}
	return f[m][n]
}
