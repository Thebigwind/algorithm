package main

import (
	"fmt"
)

//快速排序
func quickSort(array []int, left int, right int) {
	if left < right {
		key := array[left] //以最左侧第一个元素作为pivot
		low := left
		high := right
		for low < high { //如果头和尾没有相遇，就会一直触发交换
			for low < high && array[high] > key { //如果最右侧元素大于分界点pivot，下标向左移动一位
				high--
			}
			array[low] = array[high]
			for low < high && array[low] < key { //如果左侧
				low++
			}
			array[high] = array[low]
		}
		array[low] = key
		quickSort(array, left, low-1)
		quickSort(array, low+1, right)
	}
}

func main() {

	b := []int{0, 10, 19, 24, 61, 5, 121, 9, 11, 34, 21, 22}

	quickSort(b, 0, len(b)-1)
	fmt.Println(b) //[0 5 9 10 11 19 21 22 24 34 61 121]

	a := []int{0, 10, 19, 24, 61, 5, 121, 9, 11, 34, 21, 22}
	QuickSort(a)
	fmt.Println(a)

	c := []int{0, 10, 19, 24, 61, 5, 121, 9, 11, 34, 21, 22}

	fmt.Println(qsor(c))
	fmt.Println(bubbSort(c))
	fmt.Println(qs(c))
}

func QuickSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	pivot, i := arr[0], 1     // 取第一个元素作为分水岭，i下标初始为1，即分水岭右侧的第一个元素的下标
	head, tail := 0, length-1 // 头尾的下标

	for head < tail { // 如果头和尾没有相遇，就会一直触发交换
		if arr[i] > pivot { // 如果分水岭右侧的元素大于分水岭
			arr[i], arr[tail] = arr[tail], arr[i] //就将右侧的尾部元素和分水岭右侧元素交换
			tail--                                // 尾下标左移一位
		} else { // 如果分水岭右侧的元素小于等于分水岭
			arr[i], arr[head] = arr[head], arr[i] //就将分水岭右移一位
			head++                                // 头下标右移一位
			i++                                   // i下标右移一位
		}
	}

	// 分水岭左右的元素递归做同样处理
	QuickSort(arr[:head])
	QuickSort(arr[head+1:])
}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	//pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	//a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	//a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	//qsort(a[:left])
	//qsort(a[left+1:])

	return a
}


func qsor(arr []int)[]int{
	if len(arr)<2{
		return arr
	}
	left,right := 0,len(arr)-1
	for i := range arr{
		if arr[i]<arr[right]{
			arr[i],arr[left] = arr[left],arr[i]
			left++
		}
	}
	arr[right],arr[left] = arr[left],arr[right]

	qsor(arr[:left])
	qsor(arr[left+1:])
	return arr
}

func bubbSort(arr []int)[]int{
	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr)-i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return arr
}

func qs(a []int)[]int{
	if len(a)<2{
		return a
	}

	left,right := 0,len(a)-1
	for i := range a{
		if a[i]<a[right]{
			a[i],a[left] = a[left],a[i]
			left++
		}
	}
	a[right],a[left] = a[left],a[right]
	qs(a[:left])
	qs(a[left+1:])

	return a
}