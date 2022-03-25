package leetcode

// 172. 阶乘后的零
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
// n足够大的时候总会出错
func trailingZeroes(n int) int {
	zeroNum := 0
	sum := 1
	for i := 1; i <= n; i++ {
		sum *= i
		for sum%10 == 0 {
			zeroNum++
			sum /= 10
		}
		sum %= 1000000
	}
	return zeroNum
}

func trailingZeroes1(n int) (ans int) {
	for i := 5; i <= n; i += 5 {
		for x := i; x%5 == 0; x /= 5 {
			ans++
		}
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
