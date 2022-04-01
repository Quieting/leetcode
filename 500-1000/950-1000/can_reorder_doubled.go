package leetcode

import "sort"

// canReorderDoubled 二倍数对数组
// 给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <= i < len(arr) / 2，都有 arr[2 * i + 1] = 2 * arr[2 * i]” 时，返回 true；否则，返回 false。
// 提示：
//
// 0 <= arr.length <= 3 * 104
// arr.length 是偶数
// -105 <= arr[i] <= 105
//
// 示例 1：
//
// 输入：arr = [3,1,3,6]
// 输出：false
//
// 示例 2：
//
// 输入：arr = [2,1,2,6]
// 输出：false
//
// 示例 3：
//
// 输入：arr = [4,-2,2,-4]
// 输出：true
// 解释：可以用 [-2,-4] 和 [2,4] 这两组组成 [-2,-4,2,4] 或是 [2,4,-2,-4]
//
//
func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)

	aLen := len(arr)
	exist := make(map[int]int, aLen)

	for i := 0; i < aLen; i++ {
		if _, ok := exist[arr[i]]; ok {
			exist[arr[i]] += 1
		} else {
			exist[arr[i]] = 1
		}
	}

	for i := 0; i < aLen; i++ {
		num, _ := exist[arr[i]]
		if num == 0 {
			continue
		}

		double := 0
		if arr[i] < 0 {
			double = arr[i] / 2
			if arr[i]%2 != 0 {
				return false
			}
		} else {
			double = arr[i] * 2
		}

		doubleNum, ok := exist[double]
		if num > doubleNum || !ok {
			return false
		}

		exist[double] -= num
		exist[arr[i]] -= num
	}

	return true
}
