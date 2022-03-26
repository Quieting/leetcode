package leetcode

import "strconv"

// nearestPalindromic 寻找最近的回文数
// 给定一个表示整数的字符串 n ，返回与它最近的回文整数（不包括自身）。如果不止一个，返回较小的那个。
// “最近的”定义为两个整数差的绝对值最小。
//
//
//
// 示例 1:
// 输入: n = "123"
// 输出: "121"
//
// 示例 2:
// 输入: n = "1"
// 输出: "0"
// 解释: 0 和 2是最近的回文，但我们返回最小的，也就是 0。
//
//
// 提示:
//
// 1 <= n.length <= 18
// n 只由数字组成
// n 不含前导 0
// n 代表在 [1, 10<<18 - 1] 范围内的整数
//
//
func nearestPalindromic(s string) string {
	// 有可能最接近的回文数字
	nums := findPlindromic(s)

	n, _ := strconv.ParseInt(s, 10, 64)
	min, minAbs := n, n // 最小数，最小差
	for _, val := range nums {
		if n == val {
			continue
		}
		abs := n - val
		if abs < 0 {
			abs = -abs
		}

		if abs < minAbs {
			min = val
			minAbs = abs
		}

		if abs == minAbs && min > val {
			min = val
		}
	}
	return strconv.FormatInt(min, 10)
}

func findPlindromic(s string) []int64 {
	ns := make([]byte, len(s))
	// 有可能最接近的回文数字
	nums := make([]int64, 0, 5)

	// 可能是由9组成的回文数
	n := int64(0)
	for i := 0; i < len(s)-1; i++ {
		n *= 10
		n += 9
	}
	nums = append(nums, n)

	// 原回文数可能由9组成
	nums = append(nums, n*10+11)

	i, j := 0, len(s)-1
	for i < len(s) {
		ns[i] = s[i]
		ns[j] = s[i]

		switch j - i {
		case 0:
			// 用前半部分替换后半部分的回文数字
			n, _ := strconv.ParseInt(string(ns), 10, 64)
			nums = append(nums, n)

			ni := ns[i]
			// 用前半部分替换后半部分的回文数字，中间部分减1的回文数字
			if ni == '0' {
				ns[i] = '1'
			} else {
				ns[i] = ni - 1
			}
			n, _ = strconv.ParseInt(string(ns), 10, 64)
			nums = append(nums, n)

			// 用前半部分替换后半部分的回文数字，中间部分加1的回文数字
			if ni == '9' {
				ns[i] = '0'
			} else {
				ns[i] = ni + 1
			}
			n, _ = strconv.ParseInt(string(ns), 10, 64)
			nums = append(nums, n)
			return nums
		case 1:
			// 用前半部分替换后半部分的回文数字, 有可能是它本身
			n, _ := strconv.ParseInt(string(ns), 10, 64)
			nums = append(nums, n)

			ni := ns[i]
			// 用前半部分替换后半部分的回文数字，中间部分减1的回文数字
			if ni == '0' {
				ns[i] = '1'
			} else {
				ns[i] = ni - 1
			}
			ns[j] = ns[i]
			n, _ = strconv.ParseInt(string(ns), 10, 64)
			nums = append(nums, n)

			// 用前半部分替换后半部分的回文数字，中间部分加1的回文数字
			if ns[i] == '9' {
				ns[i] = '0'
			} else {
				ns[i] = ni + 1
			}
			ns[j] = ns[i]
			n, _ = strconv.ParseInt(string(ns), 10, 64)
			nums = append(nums, n)
			return nums
		default:
			i++
			j--
		}
	}

	return nums
}
