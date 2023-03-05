package leetcode

import (
	"sort"
	"strconv"
)

// findKthNumber 字典序的第K小数字
// 给定整数 n 和 k，返回  [1, n] 中字典序第 k 小的数字。
//
// 提示:
// 1 <= k <= n <= 10^9
//
//
// 示例 1:
// 输入: n = 13, k = 2
// 输出: 10
// 解释: 字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10
//
// 示例 2:
// 输入: n = 1, k = 1
// 输出: 1
//
//
func findKthNumber(n, k int) int {
	cur := 1
	k--
	for k > 0 {
		steps := getSteps(cur, n)
		if steps <= k { // 目标数不在当前节点子节点中
			k -= steps
			cur++ // 当前节点下一个节点
		} else { // 目标数在当前节点子节点中
			cur *= 10 // 当前节点第一个子节点
			k--
		}
	}
	return cur
}

// [1, n]字典序排序后，转化为前缀树形如：
//
//      1
// /    |    \
// 0 1 2 ...  9
//
// cur的节点下子节点数量
func getSteps(cur, n int) (steps int) {
	first, last := cur, cur
	for first <= n {
		steps += min(last, n) - first + 1
		first *= 10
		last = last*10 + 9
	}
	return
}

func findKthNumber1(n int, k int) int {
	strs := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		strs = append(strs, strconv.Itoa(i))
	}
	sort.Strings(strs)

	k, _ = strconv.Atoi(strs[k-1])
	return k
}
