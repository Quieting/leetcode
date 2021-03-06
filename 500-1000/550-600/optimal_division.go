package leetcode

import (
	"strconv"
)

// optimalDivision  553. 最优除法
// 给定一组正整数，相邻的整数之间将会进行浮点除法操作。例如， [2,3,4] -> 2 / 3 / 4 。
// 你可以在任意位置添加任意数目的括号，来改变算数的优先级。你需要找出怎么添加括号，才能得到最大的结果，并且返回相应的字符串格式的表达式。
// 你的表达式不应该含有冗余的括号。

// 说明:
// 输入数组的长度在 [1, 10] 之间。
// 数组中每个元素的大小都在 [2, 1000] 之间。
// 每个测试用例只有一个最优除法解。

// 中等

// 示例：
// 输入: [1000,100,10,2]
// 输出: "1000/(100/10/2)"
// 解释:
// 1000/(100/10/2) = 1000/((100/10)/2) = 200
// 但是，以下加粗的括号 "1000/((100/10)/2)" 是冗余的，
// 因为他们并不影响操作的优先级，所以你需要返回 "1000/(100/10/2)"。

// 思路：
// 两数相除分子越大分母越小所得值越大
func optimalDivision(nums []int) string {
	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}
	if len(nums) == 2 {
		return strconv.Itoa(nums[0]) + "/" + strconv.Itoa(nums[1])
	}

	s := strconv.Itoa(nums[0]) + "/(" + strconv.Itoa(nums[1])
	for i := 2; i < len(nums); i++ {
		s += "/" + strconv.Itoa(nums[i])
	}
	s += ")"

	return s
}
