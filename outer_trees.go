package leetcode

// outerTrees 安装栅栏
// 在一个二维的花园中，有一些用 (x, y) 坐标表示的树。由于安装费用十分昂贵，你的任务是先用最短的绳子围起所有的树。只有当所有的树都被绳子包围时，花园才能围好栅栏。你需要找到正好位于栅栏边界上的树的坐标。
//
//
//
// 注意:
//
// 所有的树应当被围在一起。你不能剪断绳子来包围树或者把树分成一组以上。
// 输入的整数在 0 到 100 之间。
// 花园至少有一棵树。
// 所有树的坐标都是不同的。
// 输入的点没有顺序。输出顺序也没有要求
// 示例 1:
//
// 输入: [[1,1],[2,2],[2,0],[2,4],[3,3],[4,2]]
// 输出: [[1,1],[2,0],[4,2],[3,3],[2,4]]
// 解释:
//
// 示例 2:
//
// 输入: [[1,2],[2,2],[4,2]]
// 输出: [[1,2],[2,2],[4,2]]
// 解释:
//
// 即使树都在一条直线上，你也需要先用绳子包围它们。
//
// 阅读官方题解后解出
// 关键词：凸包 向量叉积
// 相关知识点:
//  1.两向量的数量积(向量叉积)定义为其中一条向量在另一条向量方向上的正投影的长度与被投影向量的长度之积，若投影出的向量与被投影向量方向一致则此值为正，若相反则此值为负。
func outerTrees(trees [][]int) [][]int {
	tLen := len(trees)
	if tLen < 4 {
		return trees
	}

	// 查找最左边的点
	var leftest int
	for i, val := range trees {
		if val[0] < trees[leftest][0] {
			leftest = i
		}
	}

	var ans [][]int
	vis := make([]bool, tLen) // 已添加的点
	p := leftest

	for {
		// 寻找最左边的点
		q := (p + 1) % tLen
		for r := range trees {
			if cross(trees[p], trees[q], trees[r]) < 0 {
				q = r
			}
		}

		// 寻找与最左边在同一条线上的点
		for r, b := range vis {
			if !b && r != q && r != p && cross(trees[p], trees[q], trees[r]) == 0 {
				ans = append(ans, trees[r])
				vis[r] = true
			}
		}
		if !vis[q] {
			ans = append(ans, trees[q])
			vis[q] = true
		}
		p = q
		if p == leftest {
			break
		}

	}

	return ans
}

// 计算向量 pq, qr 的向量积
func cross(p, q, r []int) int {
	return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
}

type Points [][]int

func (p Points) Compare(p1 Points) bool {
	if len(p) != len(p1) {
		return false
	}
	return false
}
