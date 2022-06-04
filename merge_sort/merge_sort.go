package mergesort

func Merge(a []int, low, mid, high int) {
	var (
		i   = low
		j   = mid + 1
		tmp []int
	)
	for i <= mid && j <= high {
		if a[i] < a[j] {
			tmp = append(tmp, a[i])
			i++
		} else {
			tmp = append(tmp, a[j])
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp = append(tmp, a[i])

	}
	for ; j <= high; j++ {
		tmp = append(tmp)
	}

	for i := 0; i <= len(tmp)-1; i++ {
		a[low+i] = tmp[i]
	}
}

func MergeSort(arr []int, low, high int) {

	if low < high {
		mid := (low + high) / 2
		MergeSort(arr, low, mid)
		MergeSort(arr, mid+1, high)
		Merge(arr, low, mid, high)
	}
}
