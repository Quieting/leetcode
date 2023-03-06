package leetcode

import (
	"strings"
)

// 1653. 使字符串平衡的最少删除次数
// 给你一个字符串 s ，它仅包含字符 'a' 和 'b' 。
// 你可以删除 s 中任意数目的字符，使得 s 平衡 。当不存在下标对 (i,j) 满足 i < j ，且 s[i] = 'b' 的同时 s[j]= 'a' ，此时认为 s 是 平衡 的。
// 请你返回使 s 平衡 的 最少 删除次数。
//
// 示例 1：
// 输入：s = "aababbab"
// 输出：2
// 解释：你可以选择以下任意一种方案：
// 下标从 0 开始，删除第 2 和第 6 个字符（"aababbab" -> "aaabbb"），
// 下标从 0 开始，删除第 3 和第 6 个字符（"aababbab" -> "aabbbb"）。
//
// 示例 2：
// 输入：s = "bbaaaaabb"
// 输出：2
// 解释：唯一的最优解是删除最前面两个字符。
//
// 提示：
// 1 <= s.length <= 105
// s[i] 要么是 'a' 要么是 'b'
func minimumDeletions(s string) int {
	count := make([][2]int, len(s)+1) // [2]int 第一个元素存储前i个元素中元素'b'第数量, 第二个元素存储后 len(s) - i 个元素中元素'a'的数量
	for i := 0; i < len(s); i++ {
		if s[i] == 'b' {
			count[i+1][0] = count[i][0] + 1
		} else {
			count[i+1][0] = count[i][0]
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 'a' {
			count[i][1] = count[i+1][1] + 1
		} else {
			count[i][1] = count[i+1][1]
		}
	}

	minDelNum := len(s)
	for i := 0; i < len(count); i++ {
		num := count[i][0] + count[i][1]
		minDelNum = min(minDelNum, num)
	}

	return minDelNum
}

// 最佳示例
func minimumDeletions1(s string) int {
	lb, ra := 0, strings.Count(s, "a")
	ans := ra
	for _, c := range s {
		if c == 'a' {
			ra--
		}
		if t := lb + ra; ans > t {
			ans = t
		}
		if c == 'b' {
			lb++
		}
	}
	return ans
}
