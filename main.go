package main

import (
	"fmt"

	"alg/sortings"
)

func main() {
	// num := []int{2, 4, 5, 1, 7, 6, 2}
	// sortings.BubbleSort(num)
	// fmt.Println(num)
	// num = []int{2, 4, 5, 1, 7, 6, 2}
	// sortings.CountingSort(num)
	nums := []int{2, 4, 0, 1, 7, 6, 2}
	fmt.Println(nums)
	final := sortings.MergeSort(nums)
	fmt.Println(final)
}
