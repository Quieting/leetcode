package leetcode

// 913. 猫和老鼠
// 两位玩家分别扮演猫和老鼠，在一张 无向 图上进行游戏，两人轮流行动。
//
// 图的形式是：graph[a] 是一个列表，由满足 ab 是图中的一条边的所有节点 b 组成。
//
// 老鼠从节点 1 开始，第一个出发；猫从节点 2 开始，第二个出发。在节点 0 处有一个洞。
//
// 在每个玩家的行动中，他们 必须 沿着图中与所在当前位置连通的一条边移动。例如，如果老鼠在节点 1 ，那么它必须移动到 graph[1] 中的任一节点。
//
// 此外，猫无法移动到洞中（节点 0）。
//
// 然后，游戏在出现以下三种情形之一时结束：
//
// 如果猫和老鼠出现在同一个节点，猫获胜。
// 如果老鼠到达洞中，老鼠获胜。
// 如果某一位置重复出现（即，玩家的位置和移动顺序都与上一次行动相同），游戏平局。
// 给你一张图 graph ，并假设两位玩家都都以最佳状态参与游戏：
//
// 如果老鼠获胜，则返回 1；
// 如果猫获胜，则返回 2；
// 如果平局，则返回 0 。
func catMouseGame(graph [][]int) int {
	const (
		draw     = 0
		mouseWin = 1
		catWin   = 2
	)

	n := len(graph)
	dp := make([][][]int, n) // dps[cat][mouse][turns] 表示猫在节点 cat，老鼠在节点 mouse 第 turns 轮的状态
	for i := range graph {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2*n*(n-1))
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}

	var getRes, getNextRest func(mouse, cat, turn int) int
	getRes = func(mouse, cat, turn int) int {
		if turn == 2*n*(n-1) {
			return draw
		}

		res := dp[mouse][cat][turn]
		if res != -1 {
			return res
		}

		if mouse == 0 {
			res = mouseWin
		} else if cat == mouse {
			res = catWin
		} else {
			res = getNextRest(mouse, cat, turn)
		}

		dp[mouse][cat][turn] = res
		return res
	}

	getNextRest = func(mouse, cat, turn int) int {
		curMove := mouse
		if turn&1 > 0 { // 猫移动
			curMove = cat
		}

		defRes := mouseWin
		if curMove == mouse {
			defRes = catWin
		}

		res := defRes
		for _, next := range graph[curMove] {
			if curMove == cat && next == 0 {
				continue
			}

			nextMouse, nextCat := mouse, cat
			if curMove == mouse {
				nextMouse = next
			} else if curMove == cat {
				nextCat = next
			}

			nextRes := getRes(nextMouse, nextCat, turn+1)
			if nextRes != defRes {
				res = nextRes
				if res != draw {
					break
				}
			}
		}
		return res
	}

	return getRes(1, 2, 0)
}
