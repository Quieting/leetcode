package leetcode

import (
	"sort"
	"strconv"
)

// lexicalOrder 字典序排数
// 给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
//
// 你必须设计一个时间复杂度为 O(n) 且使用 O(1) 额外空间的算法。
//
// 提示：
//
// 1 <= n <= 5 * 10^4
//
// 示例 1：
//
// 输入：n = 13
// 输出：[1,10,11,12,13,2,3,4,5,6,7,8,9]
//
// 示例 2：
//
// 输入：n = 2
// 输出：[1,2]
//
//
// 思路：构造树前序遍历输出即为结果
func lexicalOrder(n int) []int {
	ans := make([]int, 0, n)
	var dfs func(num int)
	dfs = func(num int) {
		if num > n {
			return
		}
		ans = append(ans, num)
		num *= 10
		for i := 0; i < 10; i++ {
			dfs(num + i)
		}
	}
	for i := 1; i < 10; i++ {
		dfs(i)
	}
	return ans
}

func lexicalOrder1(n int) []int {
	strs := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		strs = append(strs, strconv.Itoa(i))
	}
	sort.Strings(strs)

	ans := make([]int, 0, n)
	for _, val := range strs {
		v, _ := strconv.Atoi(val)
		ans = append(ans, v)
	}
	return ans
}
