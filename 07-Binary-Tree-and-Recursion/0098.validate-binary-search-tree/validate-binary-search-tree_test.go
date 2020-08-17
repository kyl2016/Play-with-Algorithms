package validate_binary_search_tree

import (
	"fmt"
	"testing"
)

func TestInitTree(t *testing.T) {
	root := InitTree(1,2,3,-1,4)
	fmt.Println(root)
}

func TestIsValidBST_nil(t *testing.T) {
	if !isValidBST(nil) {
		t.Error("nil tree should be valid")
	}
}

func TestIsValidBST_OneNode(t *testing.T) {
	if !isValidBST(InitTree(0)) {
		t.Error("nil tree should be valid")
	}
}

func TestIsValidBST_OK(t *testing.T) {
	if !isValidBST(InitTree(3,2,4,1)) {
		t.Error("nil tree should be valid")
	}
}

func TestIsValidBST_Wrong(t *testing.T) {
	if isValidBST(InitTree(3,1,4,2)) {
		t.Error("nil tree should be not valid")
	}
}

func TestIsValidBST_SameValue(t *testing.T) {
	if isValidBST(InitTree(3,2,4,2)) {
		t.Error("nil tree node value should not be same")
	}
}
