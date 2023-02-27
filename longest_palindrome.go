package leetcode

// longestPalindrome 最长回文子串
// 给你一个字符串 s，找到 s 中最长的回文子串。
//
// 提示：
//
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成
//
// 示例 1：
//
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
//
// 示例 2：
//
// 输入：s = "cbbd"
// 输出："bb"
//
//
func longestPalindrome(s string) string {
	var ans []byte
	subStr := make([][]byte, 0)
	for i := 0; i < len(s); i++ {
		for index, str := range subStr {
			str = append(str, s[i])
			subStr[index] = str
			if isPalindrome(str) {
				ans = longest(str, ans)
			}
		}
		subStr = append(subStr, []byte{s[i]})
	}
	if len(ans) == 0 {
		return string(s[0])
	}
	return string(ans)
}

func isPalindrome(s []byte) bool {
	sLen := len(s)
	for i := 0; i < sLen/2; i++ {
		if s[i] != s[sLen-i-1] {
			return false
		}
	}
	return true
}

func longest(a, b []byte) []byte {
	if len(a) > len(b) {
		return a
	}
	return b
}

// 中心扩散算法
func longestPalindrome1(s string) string {
	var ans []byte
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		left, right := i, i
		for right < sLen-1 {
			if s[right+1] != s[i] {
				break
			}
			right++
		}
		for left > 0 && right < sLen-1 {
			if s[left-1] != s[right+1] {
				break
			}

			left--
			right++
		}
		if right < sLen-1 {
			ans = longest(ans, []byte(s[left:right+1]))
		} else {
			ans = longest(ans, []byte(s[left:]))
		}

	}
	return string(ans)
}
