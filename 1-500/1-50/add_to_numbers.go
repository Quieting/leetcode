package leetcode

// addTwoNumbers 两数相加
// 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//
// 请你将两个数相加，并以相同形式返回一个表示和的链表。
//
// 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//
//
// 示例 1：
//
//
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]
// 解释：342 + 465 = 807.
// 示例 2：
//
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
// 示例 3：
//
// 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// 输出：[8,9,9,9,0,0,0,1]
//
//
// 提示：
//
// 每个链表中的节点数在范围 [1, 100] 内
// 0 <= Node.val <= 9
// 题目数据保证列表表示的数字不含前导零
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := new(ListNode)

	carry := 0
	val := 0
	tmp := ans
	for l1 != nil || l2 != nil || carry != 0 {
		val = carry
		// 计算当前节点相加的值
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}

		tmp.Val = val % 10
		carry = val / 10

		if l1 != nil || l2 != nil || carry != 0 {
			tmp.Next = new(ListNode)
		}
		tmp = tmp.Next

	}
	return ans
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	return numberToList(listToNumber(l1) + listToNumber(l2))
}
func listToNumber(l *ListNode) (num int) {
	num = l.Val
	l = l.Next
	index := 10
	for ; l != nil; l = l.Next {
		num += l.Val * index
		index *= 10
	}
	return
}

func numberToList(num int) (l *ListNode) {
	l = new(ListNode)
	ans := l
	for num > 0 {
		l.Val = num % 10
		if (num / 10) > 0 {
			l.Next = new(ListNode)
			l = l.Next
		}
		num /= 10
	}
	return ans
}
