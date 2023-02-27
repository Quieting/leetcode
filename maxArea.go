package leetcode

// 1139. 最大的以 1 为边界的正方形
// 给你一个由若干 0 和 1 组成的二维网格 grid，请你找出边界全部由 1 组成的最大 正方形 子网格，并返回该子网格中的元素数量。如果不存在，则返回 0。
//
// 示例 1：
// 输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
// 输出：9
//
// 示例 2：
// 输入：grid = [[1,1,0,0]]
// 输出：1
//
// 提示：
//
// 1 <= grid.length <= 100
// 1 <= grid[0].length <= 100
// grid[i][j] 为 0 或 1
func largest1BorderedSquare(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	gridSums := make([][][2]int, len(grid)+1)
	for i := 0; i < len(gridSums); i++ {
		gridSums[i] = make([][2]int, len(grid[0])+1)
	}
	for i := 1; i < len(gridSums); i++ {
		for j := 1; j < len(gridSums[i]); j++ {
			gridSums[i][j] = [2]int{grid[i-1][j-1] + gridSums[i][j-1][0], grid[i-1][j-1] + gridSums[i-1][j][1]}
		}
	}

	maxSquare := 0
	for i := 1; i < len(gridSums); i++ {
		for j := 1; j < len(gridSums[i]); j++ {
			m, n := i, j
			for m < len(gridSums) && n < len(gridSums[m]) {
				if gridSums[i][n][0]-gridSums[i][j-1][0] != m-i+1 || // 上边
					gridSums[m][j][1]-gridSums[i-1][j][1] != m-i+1 || // 右边
					gridSums[m][n][0]-gridSums[m][j-1][0] != m-i+1 || // 下边
					gridSums[m][n][1]-gridSums[i-1][n][1] != m-i+1 { // 左边
				} else {
					if s := (m - i + 1) * (m - i + 1); s > maxSquare {
						maxSquare = s
					}
				}

				m++
				n++
			}
		}
	}

	return maxSquare
}
