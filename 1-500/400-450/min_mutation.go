package leetcode

// minMutation 最小基因变化
// 基因序列可以表示为一条由 8 个字符组成的字符串，其中每个字符都是 'A'、'C'、'G' 和 'T' 之一。
//
// 假设我们需要调查从基因序列 start 变为 end 所发生的基因变化。一次基因变化就意味着这个基因序列中的一个字符发生了变化。
//
// 例如，"AACCGGTT" --> "AACCGGTA" 就是一次基因变化。
// 另有一个基因库 bank 记录了所有有效的基因变化，只有基因库中的基因才是有效的基因序列。
//
// 给你两个基因序列 start 和 end ，以及一个基因库 bank ，请你找出并返回能够使 start 变化为 end 所需的最少变化次数。如果无法完成此基因变化，返回 -1 。
//
// 注意：起始基因序列 start 默认是有效的，但是它并不一定会出现在基因库中。
//
// 提示：
//
// start.length == 8
// end.length == 8
// 0 <= bank.length <= 10
// bank[i].length == 8
// start、end 和 bank[i] 仅由字符 ['A', 'C', 'G', 'T'] 组成
//
// 示例 1：
//
// 输入：start = "AACCGGTT", end = "AACCGGTA", bank = ["AACCGGTA"]
// 输出：1
//
// 示例 2：
//
// 输入：start = "AACCGGTT", end = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
// 输出：2
//
// 示例 3：
//
// 输入：start = "AAAAACCC", end = "AACCCCCC", bank = ["AAAACCCC","AAACCCCC","AACCCCCC"]
// 输出：3
//
//
func minMutation(start string, end string, bank []string) int {
	bankSet := map[string]struct{}{}
	for _, val := range bank {
		bankSet[val] = struct{}{}
	}

	if _, ok := bankSet[end]; !ok {
		return -1
	}
	if start == end {
		return 0
	}

	var changes = []string{start}
	for step := 0; len(changes) > 0; step++ {
		tmp := changes
		changes = nil
		for _, c := range tmp {
			for i, val := range c {
				for _, b := range []rune{'A', 'C', 'G', 'T'} {
					if b == val {
						continue
					}
					nxt := c[:i] + string(b) + c[i+1:]
					if _, ok := bankSet[nxt]; ok {
						if nxt == end {
							return step + 1
						}
						delete(bankSet, nxt)
						changes = append(changes, nxt)
					}
				}
			}
		}
	}
	return -1
}
