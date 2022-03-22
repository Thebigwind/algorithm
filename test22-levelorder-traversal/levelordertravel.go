package test22_Inorder_traversal

import "container/list"

//二叉树中序遍历
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//迭代
func levelOrder(root *TreeNode) [][]int {
	// 层序遍历-迭代
	var res = [][]int{}
	if root == nil {
		return nil
	}
	// 双端队列实现层序遍历
	queue := list.New()
	// 开始将根节点入队
	queue.PushBack(root)
	// tmpArr数组作为临时数组，存放当层元素
	var tmpArr []int
	// 遍历每层的元素
	for queue.Len() > 0 {
		length := queue.Len()
		// 通过length控制遍历的数量
		for i := 0; i < length; i++ {
			// 队头出队
			node := queue.Remove(queue.Front()).(*TreeNode)
			// 节点的左节点入队
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			// 节点的右节点入队
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			// 当层元素先添加进tmpArr
			tmpArr = append(tmpArr, node.Val)
		}
		// 每层遍历后，tmpArr加入到res中
		res = append(res, tmpArr)
		// tmpArr重新初始化
		tmpArr = []int{}
	}
	return res
}
