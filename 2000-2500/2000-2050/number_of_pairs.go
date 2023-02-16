package leetcode

// 2341. 数组能形成多少数对
// 给你一个下标从 0 开始的整数数组 nums 。在一步操作中，你可以执行以下步骤：
// 从 nums 选出 两个 相等的 整数
// 从 nums 中移除这两个整数，形成一个 数对
// 请你在 nums 上多次执行此操作直到无法继续执行。
// 返回一个下标从 0 开始、长度为 2 的整数数组 answer 作为答案，其中 answer[0] 是形成的数对数目，answer[1] 是对 nums 尽可能执行上述操作后剩下的整数数目。
//
// 示例 1：
// 输入：nums = [1,3,2,1,3,2,2]
// 输出：[3,1]
// 解释：
// nums[0] 和 nums[3] 形成一个数对，并从 nums 中移除，nums = [3,2,3,2,2] 。
// nums[0] 和 nums[2] 形成一个数对，并从 nums 中移除，nums = [2,2,2] 。
// nums[0] 和 nums[1] 形成一个数对，并从 nums 中移除，nums = [2] 。
// 无法形成更多数对。总共形成 3 个数对，nums 中剩下 1 个数字。
//
// 示例 2：
// 输入：nums = [1,1]
// 输出：[1,0]
// 解释：nums[0] 和 nums[1] 形成一个数对，并从 nums 中移除，nums = [] 。
// 无法形成更多数对。总共形成 1 个数对，nums 中剩下 0 个数字。
//
// 示例 3：
// 输入：nums = [0]
// 输出：[0,1]
// 解释：无法形成数对，nums 中剩下 1 个数字。
//
// 提示：
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100
func numsOfPairs(nums []int) []int {
	pairs := 0
	sortNum := [101]int{} // 存储0～100没个数的个数
	for _, val := range nums {
		sortNum[val]++
		if sortNum[val]%2 == 0 {
			pairs++
		}
	}

	return []int{pairs, len(nums) - 2*pairs}
}

// numberOfPairs 观摩大神极致使用内存解法，每一bit都有作用
// 用1bit存储某位数是否成数对
//
func numberOfPairs(nums []int) []int {
	paris := 0
	freq := [16]uint8{}
	for _, num := range nums {
		if freq[num/8]&(0x80>>(num%8)) > 0 {
			paris++
			freq[num/8] ^= 0x80 >> (num % 8)
		} else {
			freq[num/8] |= 0x80 >> (num % 8)
		}
	}

	return []int{paris, len(nums) - 2*paris}

}
