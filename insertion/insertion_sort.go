package insertion

import "github.com/SudhakarPai/sort"

// Sort implements the insertion sort on the collection.
func Sort(slice sort.Collection) {
	for i := 1; i < slice.Len(); i++ {
		elmIndex := i
		for j := i - 1; j >= 0; j-- {
			if slice.Greater(j, elmIndex) {
				slice.Swap(j, j+1)
				elmIndex = j
			}
		}
	}
}
