package quick

import "github.com/SudhakarPai/sort"

// Sort sorts the given slice in the ascending order using quick sort
func Sort(slice sort.Collection) {
	quickSort(slice, 0, slice.Len()-1)
}

func quickSort(slice sort.Collection, start, end int) {
	q := partition(slice, start, end)
	if q-1 > start {
		quickSort(slice, start, q-1)
	}
	if q+1 < end {
		quickSort(slice, q+1, end)
	}
}

func partition(slice sort.Collection, start, end int) int {
	i := start
	for j := start; j < end; j++ {
		if slice.Greater(end, j) {
			slice.Swap(i, j)
			i++
		}
	}
	slice.Swap(i, end)
	return i
}
