package leetcode

// convertZString 字符串Z型变换
// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
//
// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
//
// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
// 提示：
// 1 <= s.length <= 1000
// s 由英文字母（小写和大写）、',' 和 '.' 组成
// 1 <= numRows <= 1000
//
// 示例 1：
// 输入：s = "PAYPALISHIRING", numRows = 3
// 输出："PAHNAPLSIIGYIR"
//
// 示例 2：
// 输入：s = "PAYPALISHIRING", numRows = 4
// 输出："PINALSIGYAHRPI"
// 解释：
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
//
//
// 示例 3：
// 输入：s = "A", numRows = 1
// 输出："A"
//
// Z型字符的基本规律：i: 某行第一个字符，从 1 开始； rows：行的数量, 计算得到的是字符在字符串中点下标
// 第一行：0，2*rows-2，4*rows-4  规律 （i-1）*（2 * rows - 2）
// 第二行：1，2*rows-3，2*rows-1, 4*rows-5, 规律：奇数：(i/2) *（2 * rows - 2) + 1； 偶数：（i/2） *（2 * rows - 2）- 1
// 第三行：2，2*rows-4, 2*rows,   4*rows-4 规律： 奇数：(i/2) *（2 * rows - 2) + 2； 偶数：（i/2） *（2 * rows - 2）- 2
// ...
// 最后一行：rows-1， 3*rows-3，5*rows-5 规律：(i*rows)+(i-1)*(rows-2) = (2*i-1)*rows+1-2*i
//
// 总结：
// 第一行规律：（i-1）*（2 * rows - 2）
// 第n行规律：奇数：(i/2) *（2 * rows - 2) + n -1； 偶数：（i/2） *（2 * rows - 2）- n + 1 （1<n, n < rows）
// 最后一行规律： i*rows-1
//
// 延伸：
// 第一行和最后一行也可以使用第n行规律，只是有两个点重合，我们只需要去除重合的点就可以了
func convertZString(s string, rows int) string {
	if rows == 1 || len(s) == 1 || rows >= len(s) {
		return s
	}
	ns := make([]byte, len(s))
	ni := 0

	// 第一行元素
	for i := 1; ; i++ {
		n := (i - 1) * (2*rows - 2)
		if n > len(s)-1 {
			break
		}
		ns[ni] = s[n]
		ni++
	}

	// 中间行
	for n := 2; n < rows; n++ { // 行循环
		for i := 1; ; i++ { // 行内字符循环
			m := 0

			if i%2 > 0 {
				m = (i/2)*(2*rows-2) + n - 1
			} else {
				m = (i/2)*(2*rows-2) - n + 1
			}
			if m > len(s)-1 {
				break
			}

			ns[ni] = s[m]
			ni++
		}
	}

	// 最后一行
	for i := 1; ; i++ {
		n := (2*i-1)*rows + 1 - 2*i
		if n > len(s)-1 {
			break
		}
		ns[ni] = s[n]
		ni++
	}
	return string(ns)
}

func convertZString1(s string, rows int) string {
	if rows == 1 || len(s) == 1 || rows >= len(s) {
		return s
	}
	ns := make([]byte, len(s))
	ni := 0
	for n := 1; n < rows+1; n++ { // 行循环
		for i := 1; ; i++ { // 行内字符循环
			if n == 1 && i == 1 {
				ns[ni] = s[0]
				ni++
				continue
			}
			if (i/2)*(2*rows-2)-n+1 == ((i-1)/2)*(2*rows-2)+n-1 {
				continue
			}

			m := 0
			if i%2 > 0 {
				m = (i/2)*(2*rows-2) + n - 1
			} else {
				m = (i/2)*(2*rows-2) - n + 1
			}
			if m > len(s)-1 {
				break
			}

			ns[ni] = s[m]
			ni++
		}
	}
	return string(ns)
}

func convertZString2(s string, rows int) string {
	if rows < 2 {
		return s
	}
	ans := make([]string, rows)
	i, flag := 0, -1
	for _, val := range s {
		ans[i] += string(val)
		if i == rows-1 || i == 0 {
			flag = -flag
		}
		i += flag
	}

	res := ""
	for _, val := range ans {
		res += val
	}
	return res
}
