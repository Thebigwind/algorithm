package test20_treereverse

import "container/list"

//给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。
//即左右兄弟结点互换

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	 DFS-前序-递归
func invertTree(root *TreeNode) *TreeNode {
	// 基于前序的递归遍历
	// 1.递归函数与逻辑如上
	// 2.递归终止条件
	if root == nil {
		return root
	}
	// 3.递归逻辑
	// 先交换子树节点，再分别递归左右子树
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

//  DFS-前序-迭代
func invertTree2(root *TreeNode) *TreeNode {
	// 基于DFS的前序迭代
	if root == nil {
		return root
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		// 交换左右节点, 其余和前序遍历的迭代法无二
		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return root
}

//迭代
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//BFS-层序-迭代
func invertTree3(root *TreeNode) *TreeNode {
	// 基于BFS的层序遍历
	queue := list.New()
	if root != nil {
		queue.PushBack(root)
	}
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			// 调换左右子节点，其他无二
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return root
}
