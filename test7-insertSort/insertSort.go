package main

import (
	"fmt"
)

func main() {
	slice1 := []int{3, 2, 7, 1, 5, 4, 8, 0}
	InsertSort(slice1)
}
func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}

	fmt.Println(arr)
}
