package interview

// 剑指 Offer 47. 礼物的最大价值
// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
// 你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
// 给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？
//
//
// 示例 1:
// 输入:
// [
//  [1,3,1],
//  [1,5,1],
//  [4,2,1]
// ]
// 输出: 12
// 解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物
//
// 提示：
// 0 < grid.length <= 200
// 0 < grid[0].length <= 200
func maxValue(grid [][]int) int {
	maxGrid := make([][]int, 0, len(grid))
	for i := 0; i < len(grid); i++ {
		maxGrid = append(maxGrid, make([]int, len(grid[i])))
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if i == 0 && j == 0 {
				maxGrid[i][j] = grid[i][j]
				continue
			}

			if i == 0 {
				maxGrid[i][j] = grid[i][j] + maxGrid[i][j-1]
				continue
			}

			if j == 0 {
				maxGrid[i][j] = grid[i][j] + maxGrid[i-1][j]
				continue
			}

			maxGrid[i][j] = max(maxGrid[i][j-1], maxGrid[i-1][j]) + grid[i][j]
		}
	}

	return maxGrid[len(maxGrid)-1][len(maxGrid[0])-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
