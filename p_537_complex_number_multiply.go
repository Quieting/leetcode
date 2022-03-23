package leetcode

import (
	"strconv"
	"strings"
)

// complexNumberMultiply  复数乘法
// 复数 可以用字符串表示，遵循 "实部+虚部i" 的形式，并满足下述条件：
// 实部 是一个整数，取值范围是 [-100, 100]
// 虚部 也是一个整数，取值范围是 [-100, 100]
// i*i == -1
//
// 给你两个字符串表示的复数 num1 和 num2 ，请你遵循复数表示形式，返回表示它们乘积的字符串。
// 提示：
// num1 和 num2 都是有效的复数表示。
//
// 示例 1：
// 输入：num1 = "1+1i", num2 = "1+1i"
// 输出："0+2i"
// 解释：(1 + i) * (1 + i) = 1 + i2 + 2 * i = 2i ，你需要将它转换为 0+2i 的形式。

// 示例 2：
// 输入：num1 = "1+-1i", num2 = "1+-1i"
// 输出："0+-2i"
// 解释：(1 - i) * (1 - i) = 1 + i2 - 2 * i = -2i ，你需要将它转换为 0+-2i 的形式。
//
//
func complexNumberMultiply(num1 string, num2 string) string {
	return multiply(newComplex(num1), newComplex(num2)).String()
}

type complex struct {
	real      int64
	imaginary int64
}

func newComplex(s string) complex {
	var r, ima int64
	i := strings.Index(s, "+")
	r, _ = strconv.ParseInt(string([]byte(s)[:i]), 10, 64)
	ima, _ = strconv.ParseInt(string([]byte(s)[i+1:len(s)-1]), 10, 64)
	return complex{
		real:      r,
		imaginary: ima,
	}
}

func (c complex) String() string {
	s := strconv.FormatInt(c.real, 10)
	s += "+"
	s += strconv.FormatInt(c.imaginary, 10)
	s += "i"
	return s
}

func multiply(a, b complex) complex {
	ima := a.real*b.imaginary + b.real*a.imaginary
	r := a.real*b.real - a.imaginary*b.imaginary
	return complex{
		real:      r,
		imaginary: ima,
	}
}
