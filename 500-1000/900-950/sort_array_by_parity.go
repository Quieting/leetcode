package leetcode

// sortArrayByParity 按奇偶排序数组
// 给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。
//
// 返回满足此条件的 任一数组 作为答案。
//
// 提示：
//
// 1 <= nums.length <= 5000
// 0 <= nums[i] <= 5000
//
//
// 示例 1：
//
// 输入：nums = [3,1,2,4]
// 输出：[2,4,3,1]
// 解释：[4,2,3,1]、[2,4,1,3] 和 [4,2,1,3] 也会被视作正确答案。
//
// 示例 2：
//
// 输入：nums = [0]
// 输出：[0]
//
//
func sortArrayByParity(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; {
		if nums[i]&1 == 0 { // 偶数
			continue
		}

		if nums[i]&1 > 0 && nums[j]&1 == 0 { // 前奇数
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		} else if nums[i]&1 == 0 && nums[j]&1 == 0 { // 同为奇数
			i++
		} else if nums[i]&1 > 0 && nums[j]&1 > 0 { // 同为偶数
			j--
		} else {
			i++
			j--
		}

		for ; j > i; j-- {
			if nums[j]&1 > 0 { // 奇数
				continue
			}

			break
		}
	}

	return nums
}
