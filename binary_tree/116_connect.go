package binary_tree

import "leetcode/utils"

func connect(root *utils.Node) *utils.Node {
	if root == nil {
		return root
	}

	// 每次循环从该层的最左侧节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		// 通过 Next 遍历这一层节点，为下一层的节点更新 Next 指针
		for node := leftmost; node != nil; node = node.Next {
			// 左节点指向右节点
			node.Left.Next = node.Right

			// 右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	return root
}
