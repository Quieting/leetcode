package leetcode

// countNumbersWithUniqueDigits 统计各位数字都不同的数字个数
// 给你一个整数 n ，统计并返回各位数字都不同的数字 x 的个数，其中 0 <= x < 10n 。
//
// 提示：
//
// 0 <= n <= 8
//
// 示例 1：
//
// 输入：n = 2
// 输出：91
// 解释：答案应为除去 11、22、33、44、55、66、77、88、99 外，在 0 ≤ x < 100 范围内的所有数字。
//
// 示例 2：
//
// 输入：n = 0
// 输出：1
//
//
func countNumbersWithUniqueDigits(n int) int {
	ans := 0
	if n == 0 {
		return 1
	}

	for i := n; i > 0; i-- {
		ans += findNumbersWithUniqueDigits(i)
	}

	return ans
}

// findNumbersWithUniqueDigits 返回嘛n位数中各位数组都不同的数字个数
func findNumbersWithUniqueDigits(n int) int {
	if n == 1 {
		return 10
	}

	num := 9
	m := 9

	for ; n > 1; n-- {
		num *= m
		m--
	}
	return num
}

func countNumbersWithUniqueDigits1(n int) int {
	data := []int{1, 10, 91, 739, 5275, 32491, 168571, 712891, 2345851}
	return data[n]
}
