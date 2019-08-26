package tree

import "fmt"

// 二叉树节点
type TreeNode struct {
	data       int
	leftChild  *TreeNode
	rightChild *TreeNode
}

// 前序遍历创建二叉树
func CreateBinaryTree(inputs []int) *TreeNode {
	if len(inputs) == 0 {
		return nil
	}
	var node *TreeNode = nil
	data := inputs[0]
	copy(inputs[0:], inputs[1:])
	inputs[len(inputs)-1] = 0
	if data != 0 {
		node = &TreeNode{
			data: data,
		}
		node.leftChild = CreateBinaryTree(inputs)
		node.rightChild = CreateBinaryTree(inputs)
	}
	return node
}

// 二叉树前序遍历
func PreOrderTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	PreOrderTraveral(node.leftChild)
	PreOrderTraveral(node.rightChild)
}

// 二叉树中序遍历
func InOrderTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	InOrderTraveral(node.leftChild)
	fmt.Println(node.data)
	InOrderTraveral(node.rightChild)
}

// 二叉树后序遍历
func PostOrderTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	PostOrderTraveral(node.leftChild)
	PostOrderTraveral(node.rightChild)
	fmt.Println(node.data)
}

// 二叉树非递归前序遍历
func PreOrderTraveralWithStack(root *TreeNode) {
	var stack []*TreeNode
	treeNode := root
	for treeNode != nil || len(stack) > 0 {
		for treeNode != nil {
			fmt.Println(treeNode.data)
			stack = append(stack, treeNode)
			treeNode = treeNode.leftChild
		}
		if len(stack) > 0 {
			treeNode = stack[len(stack)-1]
			treeNode = treeNode.rightChild
			stack = stack[:len(stack)-1]
		}
	}
}

// 二叉树非递归中序遍历
func InOrderTraveralWithStack(root *TreeNode) {
	var stack []*TreeNode
	treeNode := root
	for treeNode != nil || len(stack) > 0 {
		for treeNode != nil {
			stack = append(stack, treeNode)
			treeNode = treeNode.leftChild
		}
		if len(stack) > 0 {
			treeNode = stack[len(stack)-1]
			fmt.Println(treeNode.data)
			treeNode = treeNode.rightChild
			stack = stack[:len(stack)-1]
		}
	}
}

// 二叉树非递归后序遍历
func PostOrderTraveralWithStack(root *TreeNode) {
	var stack []*TreeNode
	nodeViewStatusMap := make(map[*TreeNode]bool)
	treeNode := root
	for treeNode != nil || len(stack) > 0 {
		for treeNode != nil {
			stack = append(stack, treeNode)
			print(stack)
			nodeViewStatusMap[treeNode] = false
			treeNode = treeNode.leftChild
		}
		if len(stack) > 0 {
			treeNode = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			print(stack)
			// 此时当前节点还不能访问，因为它的右子树还没有访问
			// 这种情况下所有节点都会入栈两次
			if !nodeViewStatusMap[treeNode] {
				// 标记后将其放回
				nodeViewStatusMap[treeNode] = true
				stack = append(stack, treeNode)
				print(stack)
				// 继续入栈其右子树
				treeNode = treeNode.rightChild
			} else {
				fmt.Println(treeNode.data)
				// 此节点关联元素已经处理过，直接置空即可
				treeNode = nil
			}
		}
	}
}

func print(stack []*TreeNode) {
	var ints []int
	for _, item := range stack {
		data := item.data
		ints = append(ints, data)
	}
	fmt.Println(ints)
}

// 广度优先遍历
func WidthFirstTraveral(root *TreeNode) {
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		fmt.Println(node.data)
		if node.leftChild != nil {
			queue = append(queue, node.leftChild)
		}
		if node.rightChild != nil {
			queue = append(queue, node.rightChild)
		}
	}
}
