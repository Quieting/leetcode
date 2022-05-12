package leetcode

// numIsLands 岛屿数量
// 给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
//
// 岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
//
// 此外，你可以假设该网格的四条边均被水包围。
//
// 提示:
//  m == grid.length
//  n == grid[i].length
//  1 <= m, n <= 300
//  grid[i][j] 的值为 '0' 或 '1'
//
// 示例 1：
//
// 输入：grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
// ]
// 输出：1
// 示例 2：
//
// 输入：grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
// ]
// 输出：3
func numIsLands(grid [][]byte) int {
	ans := 0

	row := len(grid)
	col := len(grid[0])

	var findLand func(r, c int)

	findLand = func(r, c int) {
		if grid[r][c] != '1' {
			return
		}

		grid[r][c] = '2'

		if c+1 < col {
			findLand(r, c+1)
		}
		if c > 0 {
			findLand(r, c-1)
		}
		if r+1 < row {
			findLand(r+1, c)
		}
		if r > 0 {
			findLand(r-1, c)
		}
	}

	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] != '1' {
				continue
			}
			findLand(r, c)
			ans++
		}
	}
	return ans
}
