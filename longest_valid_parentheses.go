package leetcode

// longestValidParentheses 最长有效括号
// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
//
//
//
// 示例 1：
//
// 输入：s = "(()"
// 输出：2
// 解释：最长有效括号子串是 "()"
// 示例 2：
//
// 输入：s = ")()())"
// 输出：4
// 解释：最长有效括号子串是 "()()"
// 示例 3：
//
// 输入：s = ""
// 输出：0
func longestValidParentheses(s string) int {
	ans := 0
	const lp, rp = '(', ')'

	index := make([][]int, 0, len(s))
	curr := 0
	for i, val := range s {
		switch val {
		case lp:
			if i > 0 && s[i-1] == rp {
				index = append(index, []int{i, curr})
				curr = 0
			} else {
				index = append(index, []int{i, 0})
			}
		case rp:
			if len(index) == 0 {
				if curr > ans {
					ans = curr
				}
				curr = 0
				continue
			}
			last := index[len(index)-1]
			curr += last[1] + 2
			if curr > ans {
				ans = curr
			}

			index = index[:len(index)-1]
		}
	}
	return ans
}
