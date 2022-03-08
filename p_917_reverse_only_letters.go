package leetcode

// reverseOnlyLetters 仅仅反转字符串
// 给你一个字符串 s ，根据下述规则反转字符串：
// 所有非英文字母保留在原有位置。
// 所有英文字母（小写或大写）位置反转。
// 返回反转后的 s 。
//
// 提示
// 1 <= s.length <= 100
// s 仅由 ASCII 值在范围 [33, 122] 的字符组成
// s 不含 '\"' 或 '\\'
//
//
//
//
// 示例 1：
// 输入：s = "ab-cd"
// 输出："dc-ba"
//
// 示例 2：
// 输入：s = "a-bC-dEf-ghIj"
// 输出："j-Ih-gfE-dCba"

// 示例 3：
// 输入：s = "Test1ng-Leet=code-Q!"
// 输出："Qedo1ct-eeLg=ntse-T!"
//
// 第一次执行失败：
// i,j 重复计算
func reverseOnlyLetters(s string) string {
	if len(s) == 1 {
		return s
	}
	bs := make([]byte, len(s))
	i, j := 0, len(s)-1
	for i <= j {
		if isLetter(s[i]) && isLetter(s[j]) {
			bs[i] = s[j]
			bs[j] = s[i]
			i++
			j--
			continue
		}
		if !isLetter(s[i]) {
			bs[i] = s[i]
			i++
		}
		if !isLetter(s[j]) {
			bs[j] = s[j]
			j--
		}

	}
	return string(bs)
}

func isLetter(s uint8) bool {
	return (s > 64 && s < 91) || (s < 123 && s > 96)
}
