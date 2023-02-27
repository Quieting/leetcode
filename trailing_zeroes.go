package leetcode

// trailingZeroes. 阶乘后的零
// 给定一个整数 n ，返回 n! 结果中尾随零的数量。
//
// 提示 n! = n * (n - 1) * (n - 2) * ... * 3 * 2 * 1
//
//
//
// 示例 1：
//
// 输入：n = 3
// 输出：0
// 解释：3! = 6 ，不含尾随 0
// 示例 2：
//
// 输入：n = 5
// 输出：1
// 解释：5! = 120 ，有一个尾随 0
// 示例 3：
//
// 输入：n = 0
// 输出：0
//
// trailingZeroes1 官方解析
// 寻找质因数 2 和 5 的对数，因为 质因数 5 的个数小于等于质因数 2 的个数，因此只需要寻找质因数 5 的个数
func trailingZeroes1(n int) (ans int) {
	for i := 5; i <= n; i += 5 {
		for x := i; x%5 == 0; x /= 5 {
			ans++
		}
	}
	return
}

// trailingZeroes2 官方解析
// 对于 n! 质因数 5 的个数计算规律:
// 每隔 5 个数，出现一个 5，每隔 25 个数，出现 2 个 5，每隔 125 个数，出现 3 个 5... 以此类推。
func trailingZeroes2(n int) (ans int) {
	for n > 0 {
		n /= 5
		ans += n
	}
	return
}

// factorial 计算数的阶乘
func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	sum := int64(1)
	for i := int64(1); i <= n; i++ {
		sum *= i
	}
	return sum
}
