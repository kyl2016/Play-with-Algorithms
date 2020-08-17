package validate_binary_search_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func InitTree(s ...int) *TreeNode {
	if len(s) == 0 {
		return nil
	}

	var nodes = make([]*TreeNode, len(s))
	for i, val := range s {
		if val >= 0 {
			nodes[i] = &TreeNode{Val: val}
		}
	}

	get := func(i int) *TreeNode {
		if len(s) > i {
			return nodes[i]
		}
		return nil
	}

	for i, _ := range s {
		if nodes[i] != nil {
		nodes[i].Left = get((i+1)*2-1)
		nodes[i].Right = get((i+1)*2)
		}
	}
	return nodes[0]
}
