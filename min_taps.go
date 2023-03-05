package leetcode

// 1326. 灌溉花园的最少水龙头数目
// 在 x 轴上有一个一维的花园。花园长度为 n，从点 0 开始，到点 n 结束。
// 花园里总共有 n + 1 个水龙头，分别位于 [0, 1, ..., n] 。
// 给你一个整数 n 和一个长度为 n + 1 的整数数组 ranges ，其中 ranges[i] （下标从 0 开始）表示：
// 如果打开点 i 处的水龙头，可以灌溉的区域为 [i -  ranges[i], i + ranges[i]] 。
// 请你返回可以灌溉整个花园的 最少水龙头数目 。如果花园始终存在无法灌溉到的地方，请你返回 -1 。
func minTaps(n int, ranges []int) int {
	var link = &tapNode{
		start: 0,
		end:   0,
		count: 0,
	}

	for i, val := range ranges {
		if val == 0 {
			continue
		}

		start, end := i-val, i+val
		if start < 0 {
			start = 0
		}
		if end > n {
			end = n
		}
		link = addTap(link, start, end)
	}

	if link.end < n {
		return -1
	}

	return link.count
}

func addTap(link *tapNode, start, end int) *tapNode {
	// 前一个节点覆盖新加入节点或者两个节点范围无交集
	if link.end >= end || start > link.end {
		return link
	}

	// 当前节点覆盖上一个节点
	if start <= link.start {
		if link.last == nil { // 首节点
			return &tapNode{
				start: start,
				end:   end,
				last:  link,
				count: 1,
			}
		}
		return addTap(link.last, start, end)
	}

	return &tapNode{
		start: link.end,
		end:   end,
		last:  link,
		count: link.count + 1,
	}
}

type tapNode struct {
	// 当前水龙头有效区间
	start int
	end   int
	count int // 当前节点数量

	last *tapNode
}
