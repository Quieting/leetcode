package leetcode

// pushDominoes 推多米诺
// n 张多米诺骨牌排成一行，将每张多米诺骨牌垂直竖立。在开始时，同时把一些多米诺骨牌向左或向右推。
//
// 每过一秒，倒向左边的多米诺骨牌会推动其左侧相邻的多米诺骨牌。同样地，倒向右边的多米诺骨牌也会推动竖立在其右侧的相邻多米诺骨牌。
//
// 如果一张垂直竖立的多米诺骨牌的两侧同时有多米诺骨牌倒下时，由于受力平衡， 该骨牌仍然保持不变。
//
// 就这个问题而言，我们会认为一张正在倒下的多米诺骨牌不会对其它正在倒下或已经倒下的多米诺骨牌施加额外的力。
//
// 给你一个字符串 dominoes 表示这一行多米诺骨牌的初始状态，其中：
//
// dominoes[i] = 'L'，表示第 i 张多米诺骨牌被推向左侧，
// dominoes[i] = 'R'，表示第 i 张多米诺骨牌被推向右侧，
// dominoes[i] = '.'，表示没有推动第 i 张多米诺骨牌。
// 返回表示最终状态的字符串。
//
//
// 示例 1：
//
// 输入：dominoes = "RR.L"
// 输出："RR.L"
// 解释：第一张多米诺骨牌没有给第二张施加额外的力。
//
// 示例 2：
// 输入：dominoes = ".L.R...LR..L.."
// 输出："LL.RR.LLRRLL.."
//
func pushDominoes(dominoes string) string {
	ans := dominoes
	tmp := ""
	for {
		tmp = nextDominoes(ans)
		if ans == tmp {
			break
		}
		ans = tmp
	}

	return ans
}

// nextDominoes 当前状态的多米诺经过一秒后多米诺状态
func nextDominoes(dominoes string) string {
	res := []byte(dominoes)
	last, next := uint8('.'), uint8('.')
	for i := 0; i < len(res); i++ {
		if i == len(res)-1 {
			next = 'R'
		} else {
			next = dominoes[i+1]
		}

		switch res[i] {
		case 'L', 'R':
		case '.':
			if last != 'R' && next == 'L' {
				res[i] = 'L'
			} else if last == 'R' && next != 'L' {
				res[i] = 'R'
			}
		}
		last = dominoes[i]
	}
	return string(res)
}

// pushDominoes1 官方题解
func pushDominoes1(dominoes string) string {
	s := []byte(dominoes)
	i, n, left := 0, len(s), byte('L')
	for i < n {
		j := i
		for j < n && s[j] == '.' { // 找到一段连续的没有被推动的骨牌
			j++
		}
		right := byte('R')
		if j < n {
			right = s[j]
		}
		if left == right { // 方向相同，那么这些竖立骨牌也会倒向同一方向
			for i < j {
				s[i] = right
				i++
			}
		} else if left == 'R' && right == 'L' { // 方向相对，那么就从两侧向中间倒
			k := j - 1
			for i < k {
				s[i] = 'R'
				s[k] = 'L'
				i++
				k--
			}
		}
		left = right
		i = j + 1
	}
	return string(s)
}
