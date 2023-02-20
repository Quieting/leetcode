package leetcode

// 2347. 最好的扑克手牌
// 给你一个整数数组 ranks 和一个字符数组 suit 。你有 5 张扑克牌，第 i 张牌大小为 ranks[i] ，花色为 suits[i] 。
//
// 下述是从好到坏你可能持有的 手牌类型 ：
//
// "Flush"：同花，五张相同花色的扑克牌。
// "Three of a Kind"：三条，有 3 张大小相同的扑克牌。
// "Pair"：对子，两张大小一样的扑克牌。
// "High Card"：高牌，五张大小互不相同的扑克牌。
// 请你返回一个字符串，表示给定的 5 张牌中，你能组成的 最好手牌类型 。
//
// 注意：返回的字符串 大小写 需与题目描述相同。
func bestHand(ranks []int, suits []byte) string {
	if suits[0] == suits[1] && suits[1] == suits[2] && suits[2] == suits[3] && suits[3] == suits[4] {
		return "Flush"
	}

	newRanks := make(map[int]int, len(ranks)) // 原 ranks 每个元素出现的次数
	maxNum := 0
	for _, r := range ranks {
		newRanks[r]++
		if newRanks[r] > maxNum {
			maxNum = newRanks[r]
		}
	}

	switch maxNum {
	case 5, 4, 3:
		return "Three of a Kind"
	case 2:
		return "Pair"
	case 1:
		return "High Card"
	default:
		return ""
	}
}
