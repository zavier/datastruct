package tree

import (
	"fmt"
	"testing"
)

var tree *TreeNode

func init() {
	datas := []int{3, 2, 9, 0, 0, 10, 0, 0, 8, 0, 4}
	tree = CreateBinaryTree(datas)
}

func TestPreOrderTraveral(t *testing.T) {
	fmt.Println("前序遍历")
	PreOrderTraveral(tree)
}

func TestInOrderTraveral(t *testing.T) {
	fmt.Println("中序遍历")
	InOrderTraveral(tree)
}

func TestPostOrderTraveral(t *testing.T) {
	fmt.Println("后序遍历")
	PostOrderTraveral(tree)
}

func TestPreOrderTraveralWithStack(t *testing.T) {
	fmt.Println("非递归前序遍历")
	PreOrderTraveralWithStack(tree)
}

func TestInOrderTraveralWithStack(t *testing.T) {
	fmt.Println("非递归中序遍历")
	InOrderTraveralWithStack(tree)
}

func TestPostOrderTraveralWithStack(t *testing.T) {
	fmt.Println("非递归后序遍历")
	PostOrderTraveralWithStack(tree)
}
