package leetcode

import "strconv"

// 给定一个字符串 s 表示一个整数嵌套列表，实现一个解析它的语法分析器并返回解析的结果 NestedInteger 。
//
// 列表中的每个元素只可能是整数或整数嵌套列表
//
// 提示：
//
// 1 <= s.length <= 5 * 104
// s 由数字、方括号"[]"、负号'-'、逗号','组成
// 用例保证 s 是可解析的 NestedInteger
// 输入中的所有值的范围是 [-106, 106]
//
// 示例 1：
//
// 输入：s = "324",
// 输出：324
// 解释：你应该返回一个 NestedInteger 对象，其中只包含整数值 324。
//
// 示例 2：
//
// 输入：s = "[123,[456,[789]]]",
// 输出：[123,[456,[789]]]
// 解释：返回一个 NestedInteger 对象包含一个有两个元素的嵌套列表：
// 1. 一个 integer 包含值 123
// 2. 一个包含两个元素的嵌套列表：
//    i.  一个 integer 包含值 456
//    ii. 一个包含一个元素的嵌套列表
//         a. 一个 integer 包含值 789
//

// NestedInteger This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	vals []int
}

// IsInteger Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return len(n.vals) == 1
}

// GetInteger Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.vals[0]
}

// SetInteger  this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	if len(n.vals) == 0 {
		n.vals = append(n.vals, value)
	} else {
		n.vals[0] = value
	}
}

// Add Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.vals = append(n.vals, elem.vals...)
}

// GetList Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	ans := make([]*NestedInteger, 0, len(n.vals))
	for _, val := range n.vals {
		nest := &NestedInteger{}
		nest.SetInteger(val)
		ans = append(ans, nest)
	}
	return ans
}

func deserialize(s string) *NestedInteger {
	if s[0] != '[' {
		val, _ := strconv.Atoi(s)
		return NestedIntegerWithSingle(val)
	}

	// "[123,456,[788,799,833],[[]],10,[]]"

	stack := make([]*NestedInteger, 0) // 已处理生成的 NestedInteger
	num := 0
	var negative bool
	for i, letter := range s {
		switch letter {
		case '-':
			negative = true
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = num*10 + int(letter-'0')
		case ',': // 把以解析的数添加到栈顶 NestedInteger 对象中
			if s[i-1] != ']' {
				if negative {
					num = 0 - num
				}
				stack[len(stack)-1].Add(*NestedIntegerWithSingle(num))
			}

			negative = false
			num = 0
		case '[': // 向栈顶添加成一个 NestedInteger 对象
			stack = append(stack, &NestedInteger{})
		case ']': // 将栈顶对象添加到栈顶位置-1对象中，并删除栈顶对象
			if '0' <= s[i-1] && s[i-1] <= '9' {
				if negative {
					num = 0 - num
				}
				stack[len(stack)-1].Add(*NestedIntegerWithSingle(num))
			}
			if len(stack) > 1 {
				stack[len(stack)-2].Add(*stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}

			negative = false
			num = 0
		}
	}

	return stack[0]
}

func NestedIntegerWithSingle(n int) *NestedInteger {
	nest := NestedInteger{}
	nest.SetInteger(n)
	return &nest
}
