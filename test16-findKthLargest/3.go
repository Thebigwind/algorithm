package test16_findKthLargest

import "sort"

//占用空间更大
func findKthLargest3(nums []int, k int) int {
	var res map[int]int = make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		res[nums[i]] = res[nums[i]] + 1
	}
	result := []int{}
	for key, _ := range res {
		result = append(result, key)
	}
	sort.Ints(result)
	for i := len(result) - 1; i >= 0; i-- {
		if k-res[result[i]] <= 0 {
			return result[i]
		} else {
			k = k - res[result[i]]
		}
	}
	return 0
}
