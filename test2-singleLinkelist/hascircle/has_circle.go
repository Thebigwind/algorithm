package hascircle

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
复杂度分析

时间复杂度：O(N)O(N)，其中 NN 是链表中的节点数。最坏情况下我们需要遍历每个节点一次。

空间复杂度：O(N)O(N)，其中 NN 是链表中的节点数。主要为哈希表的开销，最坏情况下我们需要将每个节点插入到哈希表中一次。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/linked-list-cycle/solution/huan-xing-lian-biao-by-leetcode-solution/
来源：力扣（LeetCode）
*/
func hasCycle(head *ListNode) bool {
	seen := map[*ListNode]struct{}{} // 开一个map记录该节点是否已经遍历过，值记录节点索引
	for head != nil {
		if _, ok := seen[head]; ok { // 该节点遍历过，形成了环
			return true
		}
		seen[head] = struct{}{} /// 记录该节点已经遍历过
		head = head.Next
	}
	return false
}

/*
复杂度分析

时间复杂度：O(N)O(N)，其中 NN 是链表中的节点数。

当链表中不存在环时，快指针将先于慢指针到达链表尾部，链表中每个节点至多被访问两次。

当链表中存在环时，每一轮移动后，快慢指针的距离将减小一。而初始距离为环的长度，因此至多移动 NN 轮。

空间复杂度：O(1)O(1)。我们只使用了两个指针的额外空间。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/linked-list-cycle/solution/huan-xing-lian-biao-by-leetcode-solution/
来源：力扣（LeetCode）
*/
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next //慢指针一次走一步， 快指针，每次走两步
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next      //慢指针一次走一步
		fast = fast.Next.Next // 快指针，每次走两步
	}
	return true
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/linked-list-cycle/solution/huan-xing-lian-biao-by-leetcode-solution/

/*
方法二：快慢指针（12ms）

快指针一次走两步，慢指针一次走一步，如果链表有环，那么两个指针始终会相遇。

时间复杂度 O(n)，空间复杂度 O(1)
*/
func hasCycle3(head *ListNode) bool { // 快慢指针。假如爱有天意，那么快慢指针终会相遇
	if head == nil {
		return false
	}
	fastHead := head.Next // 快指针，每次走两步
	for fastHead != nil && head != nil && fastHead.Next != nil {
		if fastHead == head { // 快慢指针相遇，表示有环
			return true
		}
		fastHead = fastHead.Next.Next
		head = head.Next // 慢指针，每次走一步
	}
	return false
}

//作者：elliotxx
//链接：https://leetcode-cn.com/problems/linked-list-cycle/solution/8mskuai-man-zhi-zhen-hashsi-lu-de-go-shi-xian-by-e/
//来源：力扣（LeetCode）
