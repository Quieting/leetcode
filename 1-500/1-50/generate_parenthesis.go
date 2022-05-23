package leetcode

// generateParenthesis 括号生成
// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例 1：
//
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]
// 示例 2：
//
// 输入：n = 1
// 输出：["()"]
func generateParenthesis(n int) []string {
	add := func(d []string) []string {
		if len(d) == 0 {
			return []string{"()"}
		}
		var ans []string
		isExsit := make(map[string]struct{})
		for _, val := range d {
			s := make([]byte, 0, len(val)+2)
			for i := 0; i < len(val)+1; i++ {
				s = make([]byte, 0, len(val)+2)
				s = append(s, val[:i]...)
				s = append(s, "()"...)
				s = append(s, val[i:]...)
				if _, ok := isExsit[string(s)]; !ok {
					ans = append(ans, string(s))
					isExsit[string(s)] = struct{}{}
				}
			}
		}

		return ans
	}

	var ans []string
	for i := 0; i < n; i++ {
		ans = add(ans)
	}

	return ans
}
