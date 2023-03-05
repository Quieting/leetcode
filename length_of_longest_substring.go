package leetcode

// lengthOfLongestSubstring 无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
//

// 提示：
//
// 0 <= s.length <= 5 * 10^4
// s 由英文字母、数字、符号和空格组成
//
// 示例 1:
//
// 输入: s = "abcabcbb"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
// 示例 2:
//
// 输入: s = "bbbbb"
// 输出: 1
// 解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
// 示例 3:
//
// 输入: s = "pwwkew"
// 输出: 3
// 解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
func lengthOfLongestSubstring(s string) int {
	ans := 0

	start := 0
	for i := range s {
		for j := start; j < i; j++ {
			if s[j] == s[i] {
				start = j + 1
				break
			}
		}
		ans = max(ans, i-start+1)
	}

	return ans
}
