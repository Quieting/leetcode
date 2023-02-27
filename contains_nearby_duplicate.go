package leetcode

// containsNearbyDuplicate 存在重复元素
// 给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足
// nums[i] == nums[j] 且 abs(i - j) <= k 。
// 如果存在，返回 true ；否则，返回 false 。
//
// 提示：
//
// 1 <= nums.length <= 105
// -109 <= nums[i] <= 109
// 0 <= k <= 105
//
//
// 示例 1：
// 输入：nums = [1,2,3,1], k = 3
// 输出：true
//
// 示例 2：
// 输入：nums = [1,0,1,1], k = 1
// 输出：true
//
// 示例 3：
// 输入：nums = [1,2,3,1,2,3], k = 2
// 输出：false
//
// 本题提交三次才通过，主要原因是轻视态度，因为轻视犯了以下错误
// 没有写测试文件，莽撞提交
// 忽略了数组边界（第一次提交失败）
// 低级错误，应该是 j-i <= k 写成 j < k（第二次提交）
func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j-i <= k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
