package interview

// 面试题 01.05. 一次编辑
// 字符串有三种编辑操作:插入一个字符、删除一个字符或者替换一个字符。 给定两个字符串，编写一个函数判定它们是否只需要一次(或者零次)编辑。
//
//
//
// 示例 1:
//
// 输入:
// first = "pale"
// second = "ple"
// 输出: True
//
//
// 示例 2:
//
// 输入:
// first = "pales"
// second = "pal"
// 输出: False
func oneEditAway(first string, second string) bool {
	fLen := len(first)
	sLen := len(second)

	if diffAbs(fLen, sLen) > 1 {
		return false
	}

	fi, si, diff := 0, 0, 0
	for fi < fLen && si < sLen {
		if diff > 1 {
			return false
		}
		if first[fi] == second[si] {
			fi++
			si++
			continue
		}

		diff++

		switch {
		case fLen == sLen:
			fi++
			si++
		default:
			if fLen > sLen {
				fi++
			} else {
				si++
			}
		}
	}

	return diff <= 1
}

func diffAbs(a, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
