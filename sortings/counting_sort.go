package sortings

import "fmt"

func CountingSort(nums []int) {
	l := len(nums) + 1
	count := make([]int, l)

	for i := 0; i < l-1; i++ {
		n := nums[i]
		count[n]++
	}
	for i, j := 0, 0; i < l-1; {
		if count[j] == 0 {
			j++

		} else {
			m := j

			for k := 0; k < count[m]; k++ {
				nums[i] = m
				i++
				j++
			}

		}

	}
	fmt.Println(nums)
}
