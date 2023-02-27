package leetcode

// 1238. 循环码排列
// 给你两个整数 n 和 start。你的任务是返回任意 (0,1,2,,...,2^n-1) 的排列 p，并且满足：
// p[0] = start
// p[i] 和 p[i+1] 的二进制表示形式只有一位不同
// p[0] 和 p[2^n -1] 的二进制表示形式也只有一位不同
//
//
// 示例 1：
// 输入：n = 2, start = 3
// 输出：[3,2,0,1]
// 解释：这个排列的二进制表示是 (11,10,00,01)
// 所有的相邻元素都有一位是不同的，另一个有效的排列是 [3,1,0,2]
//
// 示例 2：
// 输出：n = 3, start = 2
// 输出：[2,6,7,5,4,0,1,3]
// 解释：这个排列的二进制表示是 (010,110,111,101,100,000,001,011)
//
// 提示：
// 1 <= n <= 16
// 0 <= start < 2^n
func circularPermutation(n int, start int) []int {
	count := 1 << n
	status := make(Status, (count+7)/8)

	res := make([]int, 0, count)
	res = append(res, start)
	status.used(start)

	for i := 0; i < count-1; i++ {
		for j := 0; j < n; j++ {
			n := changeOneBit(res[i], j)
			if status.isUse(n) {
				continue
			}
			res = append(res, n)
			status.used(n)
			break
		}
	}

	return res
}

func changeOneBit(n int, i int) int {
	return n ^ (1 << i)
}

type Status []uint8

// isUse true：已使用 false: 未使用
func (s Status) isUse(n int) bool {
	i := n / 8
	return (s[i]>>(n%8))&1 > 0
}

func (s Status) used(n int) {
	i := n / 8
	s[i] |= 1 << (n % 8)
}
