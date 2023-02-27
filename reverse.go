package leetcode

import "math"

// reverse 整数反转
// 给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
//
// 假设环境不允许存储 64 位整数（有符号或无符号）。
//
// 提示：
//
// -2^31 <= x <= 2^31 - 1
//
// 示例 1：
//
// 输入：x = 123
// 输出：321
//
// 示例 2：
//
// 输入：x = -123
// 输出：-321
//
// 示例 3：
//
// 输入：x = 120
// 输出：21
//
// 示例 4：
//
// 输入：x = 0
// 输出：0
//
//
func reverse(x int) int {
	ans := int64(0)
	isNegative := x < 0
	if isNegative {
		x = 0 - x
	}

	for x > 0 {
		ans = ans*10 + int64(x%10)
		x /= 10
	}

	if isNegative {
		ans = 0 - ans
	}

	if ans < math.MinInt32 || ans > math.MaxInt32 {
		return 0
	}

	return int(ans)
}
