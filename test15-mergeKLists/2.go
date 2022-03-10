package test15_mergeKLists

import "sort"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/*
解题思路
第一步：
遍历接收集的数据
第二步：
递归将所有数据放在数组当中
第三步：
利用golang的升序数组函数
第四步：
遍历降序数组，一个个插入到树中返回

作者：yu-qian-d
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists/solution/jian-dan-li-jie-liang-xiao-shi-yan-jiu-jie-guo-hao/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func mergeKLists2(lists []*ListNode) *ListNode {
	res := []int{}
	total := 0
	for _, v := range lists {
		if v != nil {
			res = Digui(res, v)
		} else {
			total += 1
		}
	}
	if total >= len(lists) || len(lists) == 0 {
		return nil
	}
	sort.Ints(res)
	result := &ListNode{}
	for i := len(res) - 1; i >= 0; i-- {
		new_list := &ListNode{}
		if i == len(res)-1 {
			new_list.Val = res[i]
			new_list.Next = nil
		} else {
			new_list.Val = res[i]
			new_list.Next = result
		}
		result = new_list
	}
	return result
}
func Digui(kwargs []int, args *ListNode) []int {
	kwargs = append(kwargs, args.Val)
	if args.Next != nil {
		return Digui(kwargs, args.Next)
	} else {
		return kwargs
	}
}
