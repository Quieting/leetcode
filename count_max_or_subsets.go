package leetcode

// countMaxOrSubsets 统计按位或能得到最大值的子集数目
// 给你一个整数数组 nums ，请你找出 nums 子集 按位或 可能得到的 最大值 ，并返回按位或能得到最大值的 不同非空子集的数目 。
// 如果数组 a 可以由数组 b 删除一些元素（或不删除）得到，则认为数组 a 是数组 b 的一个 子集 。如果选中的元素下标位置不一样，则认为两个子集 不同 。
// 对数组 a 执行 按位或，结果等于 a[0] OR a[1] OR ... OR a[a.length - 1]（下标从 0 开始）。
//
// 提示：
//
// 1 <= nums.length <= 16
// 1 <= nums[i] <= 105
//
//
// 示例 1：
// 输入：nums = [3,1]
// 输出：2
// 解释：子集按位或能得到的最大值是 3 。有 2 个子集按位或可以得到 3 ：
// - [3]
// - [3,1]
//
// 示例 2：
// 输入：nums = [2,2,2]
// 输出：7
// 解释：[2,2,2] 的所有非空子集的按位或都可以得到 2 。总共有 2^3 - 1 = 7 个子集。
//
// 示例 3：
// 输入：nums = [3,2,1,5]
// 输出：6
// 解释：子集按位或可能的最大值是 7 。有 6 个子集按位或可以得到 7 ：
// - [3,5]
// - [3,1,5]
// - [3,2,5]
// - [3,2,1,5]
// - [2,5]
// - [2,1,5]
//
func countMaxOrSubsets(nums []int) int {
	max := 0 // nums按位或最大值
	bi := 0  // 用位下标表示nums中的元素下标
	for i, n := range nums {
		max |= n
		bi += 1 << i
	}

	num := 0 // 最大数组合次数

	for ; bi > 0; bi-- {
		m := 0
		for _, n := range bitForOne(bi) {
			if n < len(nums) {
				m |= nums[n]
			}
		}
		if m == max {
			num++
		}
	}

	return num
}

// bitForOne 返回 bit 位为1的位下标
func bitForOne(n int) []int {
	res := make([]int, 0)
	for i := 0; (1 << i) <= n; i++ {
		if (n & (1 << i)) > 0 {
			res = append(res, i)
		}
	}
	return res
}

func countMaxOrSubsets1(nums []int) (ans int) {
	maxOr := 0
	for i := 1; i < 1<<len(nums); i++ {
		or := 0
		for j, num := range nums {
			if i>>j&1 == 1 {
				or |= num
			}
		}
		if or > maxOr {
			maxOr = or
			ans = 1
		} else if or == maxOr {
			ans++
		}
	}
	return
}
