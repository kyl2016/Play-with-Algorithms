package samples

import (
	"fmt"
	"testing"
	"time"
)

func TestPreOrder(t *testing.T) {
	tree := BinaryTree {
		Data:       1,
		LeftChild:  &BinaryTree{
			Data:       11,
			LeftChild:  &BinaryTree{
				Data:       111,
			},
			RightChild:  &BinaryTree{
				Data:       112,
			},
		},
		RightChild: &BinaryTree{
			Data:       12,
			LeftChild:  &BinaryTree{
				Data:       121,
				LeftChild:  nil,
				RightChild: nil,
			},
			RightChild: &BinaryTree{
				Data:       122,
				LeftChild:  nil,
				RightChild: nil,
			},
		},
	}

	fmt.Println("preorder")
	preOrder(&tree)
	fmt.Println("inorder")
	inOrder(&tree)
	fmt.Println("postOrder")
	postOrder(&tree)
}

func preOrder(t *BinaryTree) {
	if t == nil {
		return
	}

	time.Sleep(time.Millisecond * 10)
	fmt.Println(t.Data)
	preOrder(t.LeftChild)
	preOrder(t.RightChild)
}

func inOrder(t *BinaryTree) {
	if t == nil {
		return
	}

	inOrder(t.LeftChild)
	fmt.Println(t.Data)
	inOrder(t.RightChild)
}

func postOrder(t *BinaryTree) {
	if t == nil {
		return
	}

	postOrder(t.LeftChild)
	postOrder(t.RightChild)
	fmt.Println(t.Data)
}