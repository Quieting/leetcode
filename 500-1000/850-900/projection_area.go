package leetcode

// 883. 三维形体投影面积
// 在 n x n 的网格 grid 中，我们放置了一些与 x，y，z 三轴对齐的 1 x 1 x 1 立方体。
//
// 每个值 v = grid[i][j] 表示 v 个正方体叠放在单元格 (i, j) 上。
//
// 现在，我们查看这些立方体在 xy 、yz 和 zx 平面上的投影。
//
// 投影 就像影子，将 三维 形体映射到一个 二维 平面上。从顶部、前面和侧面看立方体时，我们会看到“影子”。
//
// 返回 所有三个投影的总面积 。
//
// 提示：
//
// n == grid.length == grid[i].length
// 1 <= n <= 50
// 0 <= grid[i][j] <= 50
//
//
// 示例 1：
//
// 输入：[[1,2],[3,4]]
// 输出：17
// 解释：这里有该形体在三个轴对齐平面上的三个投影(“阴影部分”)。
//
// 示例 2:
//
// 输入：grid = [[2]]
// 输出：5
//
// 示例 3：
//
// 输入：[[1,0],[0,2]]
// 输出：8
//
//
func projectionArea(grid [][]int) int {
	gLen := len(grid)
	ans := gLen * gLen           // 计算俯视阴影面积
	rows := make([]int, gLen)    // 计算侧视阴影面积
	columns := make([]int, gLen) // 计算正视阴影面积

	for i := 0; i < gLen; i++ {
		for j := 0; j < gLen; j++ {
			val := grid[i][j]
			rows[i] = max(rows[i], val)
			columns[j] = max(columns[j], val)
			if val == 0 {
				ans--
			}
		}
	}

	for _, val := range rows {
		ans += val
	}

	for _, val := range columns {
		ans += val
	}
	return ans
}

func projectionArea1(grid [][]int) int {
	gLen := len(grid)

	row := 0    // 当前行最大面积
	column := 0 // 当前列最大面积
	ans := 0
	for i := 0; i < gLen; i++ {
		row = 0
		column = 0
		for j := 0; j < gLen; j++ {
			row = max(row, grid[i][j])
			column = max(column, grid[j][i])
			if grid[i][j] > 0 {
				ans++
			}
		}
		ans = ans + column + row
	}
	return ans
}
