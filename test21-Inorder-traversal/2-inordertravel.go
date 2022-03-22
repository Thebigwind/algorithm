package test21_Inorder_traversal

//二叉树中序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := make([]int, 0)
	stk := make([]*TreeNode, 0)

	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}

		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]

		ans = append(ans, root.Val)
		root = root.Right
	}

	return ans
}
