package sort

import "testing"

func TestQuickSort(t *testing.T) {
	var nums []int

	nums = []int{8, 5, 2, 6, 9, 3, 1, 4, 0, 7}
	//nums = []int{2, 1}
	t.Logf("start nums=%v", nums)
	quickSort(nums)
	t.Logf("end nums=%v", nums)
}
