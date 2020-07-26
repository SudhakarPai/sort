package merge

func Sort(slice []int) {
	mergeSort(slice, 0, len(slice)-1)
}

func mergeSort(slice []int, start int, end int) {
	if start < end {
		mid := (start + end) / 2
		mergeSort(slice, start, mid)
		mergeSort(slice, mid+1, end)
		merge(slice, start, mid, end)
	}
}

func merge(slice []int, start int, mid int, end int) {
	var left = make([]int, mid-start+1)
	var right = make([]int, end-mid)
	_ = copy(left, slice[start:mid+1])
	_ = copy(right, slice[mid+1:end+1])
	for k, i, j := start, 0, 0; k < end+1; k++ {
		if i < len(left) && (j >= len(right) || left[i] <= right[j]) {
			slice[k] = left[i]
			i++
		} else if j < len(right) {
			slice[k] = right[j]
			j++
		}
	}
}
