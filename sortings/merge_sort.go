package sortings

func MergeSort(nums []int) []int {

	if len(nums) < 2 {
		return nums
	}
	first := MergeSort(nums[:len(nums)/2])
	second := MergeSort(nums[len(nums)/2:])

	return Merge(first, second)
}

func Merge(a []int, b []int) []int {
	nums1 := []int{}
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			nums1 = append(nums1, a[i])
			i++
		} else {
			nums1 = append(nums1, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		nums1 = append(nums1, a[i])

	}
	for ; j < len(b); j++ {
		nums1 = append(nums1, b[j])

	}
	return nums1
}
