package leetcode

// construct 建立四叉树
// 给你一个 n * n 矩阵 grid ，矩阵由若干 0 和 1 组成。请你用四叉树表示该矩阵 grid 。
//
// 你需要返回能表示矩阵的 四叉树 的根结点。
//
// 注意，当 isLeaf 为 False 时，你可以把 True 或者 False 赋值给节点，两种值都会被判题机制 接受 。
//
// 四叉树数据结构中，每个内部节点只有四个子节点。此外，每个节点都有两个属性：
//
// val：储存叶子结点所代表的区域的值。1 对应 True，0 对应 False；
// isLeaf: 当这个节点是一个叶子结点时为 True，如果它有 4 个子节点则为 False 。
// class Node {
//    public boolean val;
//    public boolean isLeaf;
//    public Node topLeft;
//    public Node topRight;
//    public Node bottomLeft;
//    public Node bottomRight;
// }
// 我们可以按以下步骤为二维区域构建四叉树：
//
// 如果当前网格的值相同（即，全为 0 或者全为 1），将 isLeaf 设为 True ，将 val 设为网格相应的值，并将四个子节点都设为 Null 然后停止。
// 如果当前网格的值不同，将 isLeaf 设为 False， 将 val 设为任意值，然后如下图所示，将当前网格划分为四个子网格。
// 使用适当的子网格递归每个子节点。
//
//
// 如果你想了解更多关于四叉树的内容，可以参考 wiki 。
//
// 四叉树格式：
//
// 输出为使用层序遍历后四叉树的序列化形式，其中 null 表示路径终止符，其下面不存在节点。
//
// 它与二叉树的序列化非常相似。唯一的区别是节点以列表形式表示 [isLeaf, val] 。
//
// 如果 isLeaf 或者 val 的值为 True ，则表示它在列表 [isLeaf, val] 中的值为 1 ；如果 isLeaf 或者 val 的值为 False ，则表示值为 0 。
//
//
//
// 示例 1：
//
// 输入：grid = [[0,1],[1,0]]
// 输出：[[0,1],[1,0],[1,1],[1,1],[1,0]]
// 解释：此示例的解释如下：
// 请注意，在下面四叉树的图示中，0 表示 false，1 表示 True 。
//
// 示例 2：
//
// 输入：grid = [[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,1],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0],[1,1,1,1,0,0,0,0]]
// 输出：[[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
// 解释：网格中的所有值都不相同。我们将网格划分为四个子网格。
// topLeft，bottomLeft 和 bottomRight 均具有相同的值。
// topRight 具有不同的值，因此我们将其再分为 4 个子网格，这样每个子网格都具有相同的值。
// 解释如下图所示：
//
// 示例 3：
//
// 输入：grid = [[1,1],[1,1]]
// 输出：[[1,1]]
//
// 示例 4：
//
// 输入：grid = [[0]]
// 输出：[[1,0]]
//
// 示例 5：
//
// 输入：grid = [[1,1,0,0],[1,1,0,0],[0,0,1,1],[0,0,1,1]]
// 输出：[[0,1],[1,1],[1,0],[1,0],[1,1]]

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

type point struct {
	i, j int // i: 行 j: 列
}

func construct(grid [][]int) *Node {
	if len(grid) == 1 {
		return &Node{
			IsLeaf: true,
			Val:    false,
		}
	}

	prefixSumGrid := PrefixSum(grid)
	ans := &Node{
		IsLeaf: false,
		Val:    false,
	}
	var dfs func(left, right point, n *Node)
	dfs = func(left, right point, n *Node) {
		// 判断当前区域网格值是否一致
		if area := filedArea(prefixSumGrid, left, right); area == 0 ||
			area == (right.i-left.i+1)*(right.j-left.j+1) {
			n.IsLeaf = true
			n.Val = area != 0
			return
		}

		if right.i == left.i {
			n.IsLeaf = true
			n.Val = grid[left.i][left.j] != 0
			return
		}

		// 计算中心点
		p1 := point{left.i + (right.i-left.i)/2, left.j + (right.j-left.j)/2}
		p2 := point{p1.i + 1, p1.j + 1}

		// 左上部分
		n.TopLeft = new(Node)
		dfs(left, p1, n.TopLeft)

		// 右上部分
		n.TopRight = new(Node)
		dfs(point{left.i, p2.j}, point{p1.i, right.j}, n.TopRight)

		// 左下部分
		n.BottomLeft = new(Node)
		dfs(point{p2.i, left.j}, point{right.i, p1.j}, n.BottomLeft)

		// 右下部分
		n.BottomRight = new(Node)
		dfs(p2, right, n.BottomRight)
	}
	dfs(point{0, 0}, point{len(grid) - 1, len(grid) - 1}, ans)
	return ans
}

// filedArea 计算点 left，right 形成的四边形内数组节点值总和
func filedArea(prefixSum [][]int, left, right point) int {
	sum := prefixSum[right.i][right.j]
	if left.i > 0 {
		sum -= prefixSum[left.i-1][right.j]
	}
	if left.j > 0 {
		sum -= prefixSum[right.i][left.j-1]
	}
	if left.i > 0 && left.j > 0 {
		sum += prefixSum[left.i-1][left.j-1]
	}

	return sum
}

func PrefixSum(s [][]int) [][]int {
	if len(s) == 0 {
		return nil
	}

	row := len(s)
	ns := make([][]int, row)
	for r := 0; r < row; r++ {
		ns[r] = make([]int, len(s[r]))
	}

	// 计算前缀和
	for r := 0; r < row; r++ {
		for c := 0; c < len(s[r]); c++ {
			// 不考虑边界计算二维数组前缀和公式：
			// ns[r][c] = ns[r][c-1] + ns[r-1][c] + s[r][c] - ns[r-1][c-1]
			ns[r][c] = s[r][c]
			if c-1 >= 0 {
				ns[r][c] += ns[r][c-1]
			}
			if r-1 >= 0 {
				ns[r][c] += ns[r-1][c]
			}
			if r-1 >= 0 && c-1 >= 0 {
				ns[r][c] -= ns[r-1][c-1]
			}
		}
	}

	return ns
}
