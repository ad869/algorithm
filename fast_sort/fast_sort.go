package fastsort

func Partition(arr []int, low, high int) int {
	var (
		i     = low
		j     = high
		pivot = arr[low]
	)
	for i < j {
		for i < j && arr[j] > pivot {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

}
