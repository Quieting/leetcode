package leetcode

// 1234. 替换子串得到平衡字符串
// 有一个只含有 'Q', 'W', 'E', 'R' 四种字符，且长度为 n 的字符串。
// 假如在该字符串中，这四个字符都恰好出现 n/4 次，那么它就是一个「平衡字符串」。
// 给你一个这样的字符串 s，请通过「替换一个子串」的方式，使原字符串 s 变成一个「平衡字符串」。
// 你可以用和「待替换子串」长度相同的 任何 其他字符串来完成替换。
// 请返回待替换子串的最小可能长度。
// 如果原字符串自身就是一个平衡字符串，则返回 0。
//
// 示例 1：
// 输入：s = "QWER"
// 输出：0
// 解释：s 已经是平衡的了。
//
// 示例 2：
// 输入：s = "QQWE"
// 输出：1
// 解释：我们需要把一个 'Q' 替换成 'R'，这样得到的 "RQWE" (或 "QRWE") 是平衡的。
//
// 示例 3：
// 输入：s = "QQQW"
// 输出：2
// 解释：我们可以把前面的 "QQ" 替换成 "ER"。
//
// 示例 4：
// 输入：s = "QQQQ"
// 输出：3
// 解释：我们可以替换后 3 个 'Q'，使 s = "QWER"。
//
//提示：
// 1 <= s.length <= 10^5
// s.length 是 4 的倍数
// s 中只含有 'Q', 'W', 'E', 'R' 四种字符
func balanceString(s string) int {
	cnt := [4]int{}
	for _, b := range s {
		switch b {
		case 'Q':
			cnt[0]++
		case 'W':
			cnt[1]++
		case 'E':
			cnt[2]++
		case 'R':
			cnt[3]++
		}
	}

	for i := 0; i < len(cnt); i++ {
		cnt[i] = cnt[i] - (len(s) / 4)
	}

	start := 0
	count := len(s)

	for i, b := range s {
		switch b {
		case 'Q':
			cnt[0]--
		case 'W':
			cnt[1]--
		case 'E':
			cnt[2]--
		case 'R':
			cnt[3]--
		}
		for cnt[0] <= 0 && cnt[1] <= 0 && cnt[2] <= 0 && cnt[3] <= 0 { // 符合要求
			count = min(count, i-start+1)
			switch s[start] {
			case 'Q':
				cnt[0]++
			case 'W':
				cnt[1]++
			case 'E':
				cnt[2]++
			case 'R':
				cnt[3]++
			}
			start++
		}

	}

	return count
}

// balanceStringOfficial 官方解答
func balanceStringOfficial(s string) int {
	cnt := map[byte]int{}
	for _, c := range s {
		cnt[byte(c)]++
	}
	partial := len(s) / 4
	check := func() bool {
		if cnt['Q'] > partial ||
			cnt['W'] > partial ||
			cnt['E'] > partial ||
			cnt['R'] > partial {
			return false
		}
		return true
	}

	if check() {
		return 0
	}

	res := len(s)
	r := 0
	for l, c := range s {
		for r < len(s) && !check() {
			cnt[s[r]]--
			r += 1
		}
		if !check() {
			break
		}
		res = min(res, r-l)
		cnt[byte(c)]++
	}
	return res
}
