package test16_findKthLargest

/*
思路:快排可以解决问题，但是它需要确定数组中所有元素的正确位置，对于本题而言，
我们只需要确定第k大元素的位置pos,我们只需要确保pos左边的元素都比它小，pos右边的元素都比它大即可，
不需要关心其左边和右边的集合是否有序，所以，我们需要对快排进行改进(以下称分区)，
将目标值的位置pos与分区函数Partition求得的位置index进行比对，如果两值相等，
说明index对应的元素即为所求值，如果index<pos，则递归的在[index+1, right]范围求解；
否则则在[left, index-1]范围求解，如此便可大幅缩小求解范围。

*/

func findKthLargest2(nums []int, k int) int {
	TopKSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}

func Parition(nums []int, start, stop int) int {
	if start >= stop {
		return -1
	}
	pivot := nums[start]
	l, r := start, stop
	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] < pivot {
			l++
		}
		nums[r] = nums[l]
	}
	// 循环结束，l与r相等
	// 确定基准元素pivot在数组中的位置
	nums[l] = pivot
	return l
}

func TopKSplit(nums []int, k, start, stop int) {
	if start < stop {
		index := Parition(nums, start, stop)
		if index == k {
			return
		} else if index < k {
			TopKSplit(nums, k, index+1, stop)
		} else {
			TopKSplit(nums, k, start, index-1)
		}
	}
}

//
//作者：ybzdqhl
//链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array/solution/topklei-wen-ti-xiang-jie-by-ybzdqhl-w7lo/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
