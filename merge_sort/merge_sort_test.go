package mergesort

import "testing"

func TestMerge(t *testing.T) {
	var arr = []int{7, 8, 9, 1, 2, 3}
	Merge(arr, 0, 2, 5)
	t.Log(arr)
}

func TestMergeSort(t *testing.T) {
	var arr = []int{3, 1}
	MergeSort(arr, 0, len(arr)-1)
	t.Log(arr)
}
