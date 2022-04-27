package leetcode

// pacificAtlantic 太平洋大西洋水流问题
// 有一个 m × n 的矩形岛屿，与 太平洋 和 大西洋 相邻。 “太平洋” 处于大陆的左边界和上边界，而 “大西洋” 处于大陆的右边界和下边界。
//
// 这个岛被分割成一个由若干方形单元格组成的网格。给定一个 m x n 的整数矩阵 heights ， heights[r][c] 表示坐标 (r, c) 上单元格 高于海平面的高度 。
//
// 岛上雨水较多，如果相邻单元格的高度 小于或等于 当前单元格的高度，雨水可以直接向北、南、东、西流向相邻单元格。水可以从海洋附近的任何单元格流入海洋。
//
// 返回 网格坐标 result 的 2D列表 ，其中 result[i] = [ri, ci] 表示雨水可以从单元格 (ri, ci) 流向 太平洋和大西洋 。
// 提示：
//
// m == heights.length
// n == heights[r].length
// 1 <= m, n <= 200
// 0 <= heights[r][c] <= 105
// 示例 1：
//
// 输入: heights = [[1,2,2,3,5],[3,2,3,4,4],[2,4,5,3,1],[6,7,1,4,5],[5,1,1,2,4]]
// 输出: [[0,4],[1,3],[1,4],[2,2],[3,0],[3,1],[4,0]]
//
// 示例 2：
//
// 输入: heights = [[2,1],[1,2]]
// 输出: [[0,0],[0,1],[1,0],[1,1]]
//
//
func pacificAtlantic(heights [][]int) [][]int {
	var ans [][]int
	pacific := followPacific(heights)
	atlantic := followAtlantic(heights)
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if pacific[i][j] && atlantic[i][j] {
				ans = append(ans, []int{i, j})
			}
		}
	}

	return ans
}

// followPacific heights 能流向太平洋
func followPacific(heights [][]int) [][]bool {
	var dfs func(i, j int, res [][]bool)
	dfs = func(i, j int, res [][]bool) {
		if res[i][j] { // 已访问过
			return
		}

		res[i][j] = true // 能进入表示能流入大西洋

		// 向四周扩散
		if i > 0 && heights[i][j] <= heights[i-1][j] { // 上
			dfs(i-1, j, res)
		}
		if i < len(heights)-1 && heights[i][j] <= heights[i+1][j] { // 下
			dfs(i+1, j, res)
		}

		if j > 0 && heights[i][j] <= heights[i][j-1] { // 左
			dfs(i, j-1, res)
		}
		if j < len(heights[i])-1 && heights[i][j] <= heights[i][j+1] { // 右
			dfs(i, j+1, res)
		}
	}

	res := make([][]bool, len(heights))
	for i := range res {
		res[i] = make([]bool, len(heights[i]))
	}

	for i := 0; i < len(heights[0]); i++ { // 第一行
		dfs(0, i, res)
	}
	for i := 0; i < len(heights); i++ { // 第一列
		dfs(i, 0, res)
	}
	return res
}

// followAtlantic heights 流向大西洋
func followAtlantic(heights [][]int) [][]bool {
	var dfs func(i, j int, res [][]bool)
	dfs = func(i, j int, res [][]bool) {
		if res[i][j] { // 已访问过
			return
		}

		res[i][j] = true // 能进入表示能流入大西洋

		// 向四周扩散
		if i > 0 && heights[i][j] <= heights[i-1][j] { // 上
			dfs(i-1, j, res)
		}
		if i < len(heights)-1 && heights[i][j] <= heights[i+1][j] { // 下
			dfs(i+1, j, res)
		}
		if j > 0 && heights[i][j] <= heights[i][j-1] { // 左
			dfs(i, j-1, res)
		}
		if j < len(heights[i])-1 && heights[i][j] <= heights[i][j+1] { // 右
			dfs(i, j+1, res)
		}
	}

	res := make([][]bool, len(heights))
	for i := range res {
		res[i] = make([]bool, len(heights[i]))
	}

	for i := 0; i < len(heights[len(heights)-1]); i++ { // 最后一行
		dfs(len(heights)-1, i, res)
	}
	for i := len(heights) - 1; i >= 0; i-- { // 最后一列
		dfs(i, len(heights[0])-1, res)
	}
	return res
}
