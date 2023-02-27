package leetcode

import (
	"math"
	"math/rand"

	"github.com/Quieting/leetcode/slice"
)

// maxSumSubmatrix 矩形区域不超过 K 的最大数值和
// 给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。
// 题目数据保证总会存在一个数值和不超过 k 的矩形区域。
//
// 提示：
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -100 <= matrix[i][j] <= 100
// -105 <= k <= 105
//
// 示例1：
// 输入：matrix = [[1,0,1],[0,-2,3]], k = 2
// 输出：2
// 解释：蓝色边框圈出来的矩形区域[[0, 1], [-2, 3]]的数值和是 2，且 2 是不超过 k 的最大数字（k = 2）。

// 示例 2：
// 输入：matrix = [[2,2,-1]], k = 3
// 输出：3
//
//
func maxSumSubmatrix(matrix [][]int, k int) int {
	matrix = slice.PrefixSum(slice.AddRow(slice.AddColumn(matrix, 0), 0))
	maxK := k + 1
	points := []struct {
		row    int
		column int
	}{
		{0, 0},
		{0, 0},
	}
	for r1 := 1; r1 < len(matrix); r1++ {
		for c1 := 1; c1 < len(matrix[r1]); c1++ {
			for r2 := r1; r2 < len(matrix); r2++ {
				for c2 := c1; c2 < len(matrix[r2]); c2++ {
					n := matrix[r2][c2] - matrix[r2][c1-1] - matrix[r1-1][c2] + matrix[r1-1][c1-1]

					if n <= k && (maxK == k+1 || (maxK != k+1 && n > maxK)) {
						points[0].row = r1 - 1
						points[0].column = c1 - 1

						points[1].row = r2 - 1
						points[1].column = c2 - 1

						maxK = n
					}
				}
			}
		}
	}
	// fmt.Printf("%+v\n", points)

	return maxK
}

// // prefixSum 求取二维数组前缀和，返回数组将比原数组增加一行一列方便后续计算
// func prefixSum(s [][]int) [][]int {
// 	if len(s) == 0 {
// 		return nil
// 	}
//
// 	row := len(s) + 1
// 	column := len(s[0]) + 1
//
// 	ns := make([][]int, len(s)+1)
// 	for r := 0; r < row; r++ {
// 		ns[r] = make([]int, column)
// 	}
//
// 	// 计算前缀和
// 	for r := 1; r < row; r++ {
// 		for c := 1; c < column; c++ {
// 			ns[r][c] = ns[r][c-1] + ns[r-1][c] + s[r-1][c-1] - ns[r-1][c-1]
// 		}
// 	}
//
// 	return ns
// }

type node struct {
	ch       [2]*node // 一个元素值比第二个元素值小
	priority int      // 此变量意义不懂
	val      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.val:
		return 0
	case b > o.val:
		return 1
	default:
		return -1
	}
}

func (o *node) rotate(d int) *node {
	x := o.ch[d^1]
	o.ch[d^1] = x.ch[d]
	x.ch[d] = o
	return x
}

type treap struct {
	root *node
}

func (t *treap) _put(o *node, val int) *node {
	if o == nil {
		return &node{priority: rand.Int(), val: val}
	}
	if d := o.cmp(val); d >= 0 { // val 值等于节点值
		o.ch[d] = t._put(o.ch[d], val)
		if o.ch[d].priority > o.priority {
			o = o.rotate(d ^ 1)
		}
	}
	return o
}

func (t *treap) put(val int) {
	t.root = t._put(t.root, val)
}

// lowerBound 返回大于等于 val值 的最近节点
func (t *treap) lowerBound(val int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0: // 节点值大于val
			lb = o
			o = o.ch[0]
		case c > 0: // 节点值小于val
			o = o.ch[1]
		default: // 节点值等于val
			return o
		}
	}
	return
}

// todo：官方题解，没看懂。不理解 treap 的作用和意义
func maxSumSubmatrix1(matrix [][]int, k int) int {
	ans := math.MinInt64
	for i := range matrix { // 枚举上边界
		sum := make([]int, len(matrix[0]))
		for _, row := range matrix[i:] { // 枚举下边界
			for c, v := range row {
				sum[c] += v // 更新每列的元素和
			}
			t := &treap{}
			t.put(0)
			s := 0
			for _, v := range sum {
				s += v
				if lb := t.lowerBound(s - k); lb != nil {
					ans = max(ans, s-lb.val)
				}
				t.put(s)
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
