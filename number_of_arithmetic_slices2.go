package leetcode

// numberOfArithmeticSlices 等差数列划分 II - 子序列
// 给你一个整数数组 nums ，返回 nums 中所有 等差子序列 的数目。
//
// 如果一个序列中 至少有三个元素 ，并且任意两个相邻元素之差相同，则称该序列为等差序列。
//
// 例如，[1, 3, 5, 7, 9]、[7, 7, 7, 7] 和 [3, -1, -5, -9] 都是等差序列。
// 再例如，[1, 1, 2, 5, 7] 不是等差序列。
// 数组中的子序列是从数组中删除一些元素（也可能不删除）得到的一个序列。
//
// 例如，[2,5,10] 是 [1,2,1,2,4,1,5,10] 的一个子序列。
// 题目数据保证答案是一个 32-bit 整数。
//
// 提示：
//
// 1  <= nums.length <= 1000
// -2^31 <= nums[i] <= 2^31 - 1
//
//
// 示例 1：
// 输入：nums = [2,4,6,8,10]
// 输出：7
// 解释：所有的等差子序列为：
// [2,4,6]
// [4,6,8]
// [6,8,10]
// [2,4,6,8]
// [4,6,8,10]
// [2,4,6,8,10]
// [2,6,10]
//
// 示例 2：
// 输入：nums = [7,7,7,7,7]
// 输出：16
// 解释：数组中的任意子序列都是等差子序列。
//
//
func numberOfArithmeticSlices2(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	// // dp[i]： nums[i] 为结尾的等差数列
	// // dp[2][3]：以 nums[i] 结尾的等差等于 3 的数列的个数
	dp := make([]map[int]int, len(nums))
	for i := range dp {
		dp[i] = make(map[int]int)
	}

	ant := 0
	for i := 0; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			d := nums[i] - nums[j]
			ant += dp[j][d]
			dp[i][d] += dp[j][d] + 1

			// todo：为什么这么不对？
			// d := nums[i] - nums[j]
			// dp[i][d] += dp[j][d] + 1
			// ant += dp[j][d]
		}
	}

	return ant
}
