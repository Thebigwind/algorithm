package main

import (
	"fmt"
)

func main() {
	slice1 := []int{3, 2, 7, 1, 0, 4, 9, 8}
	BubbleSort(slice1)

	arr := []int{6, 2, 4, 8, 9, 1, 4, 0, 10, 5, 2}
	selectionSort(arr, 0)
	fmt.Println(arr)
}
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println(arr)

}

/*
选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理如下。首先在未排序序列中找到最小（大）元素，
存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。以此类推，
直到所有元素均排序完毕。
选择排序的主要优点与数据移动有关。如果某个元素位于正确的最终位置上，则它不会被移动。
选择排序每次交换一对元素，它们当中至少有一个将被移到其最终位置上，因此对n个元素的表进行排序总共进行至多n-1次交换。
在所有的完全依靠交换去移动元素的排序方法中，选择排序属于非常好的一种。

*/
/**
选择排序法：在未排序的切片中选取最小的值放在首位，然后在未排序的切片中选取最小的值放在第二位，以此类推。。。
*/
func selectionSort(arr []int, start int) {
	if start == len(arr) {
		return
	}
	minIdx := start
	minVal := arr[start]
	for i := start + 1; i < len(arr); i++ {
		if arr[i] < minVal {
			minIdx, minVal = i, arr[i]
		}
	}
	arr[start], arr[minIdx] = arr[minIdx], arr[start]
	selectionSort(arr, start+1)
}

/*
复杂度分析： 选择排序的交换操作介于 0和 (n-1)次之间。选择排序的比较操作为n(n-1)/2次。选择排序的赋值操作介于 0和 3(n-1)次之间。
比较次数 O(n^2)，比较次数与关键字的初始状态无关，总的比较次数 N=(n-1)+(n-2)+…+1=n (n-1)/2。
交换次数 O(n)，最好情况是，已经有序，交换0次；最坏情况是，逆序，交换n-1次。交换次数比冒泡排序较少，
由于交换所需CPU时间比比较所需的CPU时间多，n值较小时，选择排序比冒泡排序快。
原地操作几乎是选择排序的唯一优点，当空间复杂度要求较高时，可以考虑选择排序；实际适用的场合非常罕见。

*/

//选择排序（排序10000个随机整数，用时约45ms）
func selectSort(nums []int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		maxIndex := 0
		//寻找最大的一个数，保存索引值
		for j := 1; j < length-i; j++ {
			if nums[j] > nums[maxIndex] {
				maxIndex = j
			}
		}
		nums[length-i-1], nums[maxIndex] = nums[maxIndex], nums[length-i-1]
	}
}
