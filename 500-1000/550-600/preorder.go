package leetcode

// preorder  N 叉树的前序遍历
// 给定一个 n 叉树的根节点  root ，返回 其节点值的 前序遍历 。
// n 叉树 在输入中按层序遍历进行序列化表示，每组子节点由空值 null 分隔（请参见示例）。
//
// 提示：
// 节点总数在范围 [0, 104]内
// 0 <= Node.val <= 104
// n 叉树的高度小于或等于 1000
//
// 示例 1：
// 输入：root = [1,null,3,2,4,null,5,6]
// 输出：[1,3,5,6,2,4]
//
// 示例 2：
// 输入：root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
// 输出：[1,2,3,6,7,11,14,4,8,12,5,9,13,10]
//
func preorder(root *Node) []int {
	res := make([]int, 0)
	var recursion func(node *Node)
	recursion = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, val := range node.Children {
			recursion(val)
		}
	}
	return res
}

type Node struct {
	Val      int
	Children []*Node
}

func preorder1(root *Node) []int {
	res := make([]int, 0)
	nodes := []*Node{root}
	for len(nodes) > 1 {
		node := nodes[len(nodes)-1]
		res = append(res, node.Val)
		nodes = nodes[1 : len(nodes)-1]
		for i := len(node.Children) - 1; i >= 0; i++ {
			nodes = append(nodes, node.Children[i])
		}
	}
	return res
}
