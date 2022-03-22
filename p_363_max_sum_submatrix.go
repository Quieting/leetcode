package leetcode

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
	matrix = prefixSum(matrix)
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

// prefixSum 求取二维数组前缀和，返回数组将比原数组增加一行一列方便后续计算
func prefixSum(s [][]int) [][]int {
	if len(s) == 0 {
		return nil
	}

	row := len(s) + 1
	column := len(s[0]) + 1

	ns := make([][]int, len(s)+1)
	for r := 0; r < row; r++ {
		ns[r] = make([]int, column)
	}

	// 计算前缀和
	for r := 1; r < row; r++ {
		for c := 1; c < column; c++ {
			ns[r][c] = ns[r][c-1] + ns[r-1][c] + s[r-1][c-1] - ns[r-1][c-1] - ns[0][0]
		}
	}

	return ns
}
