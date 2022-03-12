package binary_tree

import "leetcode/utils"

func connect2(root *utils.Node) *utils.Node {
	start := root
	for start != nil {
		var nextStart, last *utils.Node
		handle := func(cur *utils.Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}
