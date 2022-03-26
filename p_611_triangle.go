package leetcode

import "sort"

// triangleNumber 有效三角形的个数
// 给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。
//
// 提示:
//
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 1000
//
// 示例 1:
//
// 输入: nums = [2,2,3,4]
// 输出: 3
// 解释:有效的组合是:
// 2,3,4 (使用第一个 2)
// 2,3,4 (使用第二个 2)
// 2,2,3
//
// 示例 2:
//
// 输入: nums = [4,2,3,4]
// 输出: 4
//
//
func triangleNumber(nums []int) int {
	sort.Ints(nums)

	ans := 0
	nLen := len(nums)

	for i := 0; i < nLen; i++ {
		for j := i + 1; j < nLen; j++ {
			ans += sort.SearchInts(nums[j+1:], nums[i]+nums[j])
		}
	}
	return ans
}

func find(nums []int, side int) int {
	count := 0
	i, j := 0, 1
	nLen := len(nums)
	for i < nLen-1 {
		if j == nLen-1 && nums[j] < nums[i]+side {
			count += j - i
			i++
			continue
		}
		if nums[j] >= nums[i]+side {
			count += j - i - 1
			i++
			continue
		}

		if j < nLen-1 {
			j++
		}
	}
	return count
}

func triangleNumber1(nums []int) int {
	sort.Ints(nums)

	ans := 0
	nLen := len(nums)

	for i := 0; i < nLen-2; i++ {
		if nums[i] == 0 {
			continue
		}
		ans += find(nums[i+1:], nums[i])
	}
	return ans
}
