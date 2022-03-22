package test22_Inorder_traversal

func levelOrder2(root *TreeNode) [][]int {
	// 层序遍历-递归
	var res [][]int
	// 1.递归函数及参数
	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		// 2.递归终止条件
		if node == nil {
			return
		}
		// 3.递归逻辑
		// 初始化，depth个元素
		if len(res) == depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], node.Val)
		// 遍历节点的左右子节点，depth+1
		if node.Left != nil {
			traversal(node.Left, depth+1)
		}
		if node.Right != nil {
			traversal(node.Right, depth+1)
		}
	}
	// 从0开始遍历
	traversal(root, 0)
	return res
}
