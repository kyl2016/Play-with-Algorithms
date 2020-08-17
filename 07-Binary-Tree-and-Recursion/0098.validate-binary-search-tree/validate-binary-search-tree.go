package validate_binary_search_tree

import (
	"fmt"
	"math"
)

func isValidBST(root *TreeNode) bool {
	// init stack (use slice), lastValue
	var stack []*TreeNode
	lastVal := math.MinInt64

	// recursion
	for len(stack) >0 || root != nil {
		// find left node recursion
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// pop from stack
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// compare with lastVal
		if root.Val <= lastVal {
			fmt.Println(root.Val, lastVal)
			return false
		}
		// update lastVal
		lastVal = root.Val
		// move to Right child
		root = root.Right
	}

	return true
}
