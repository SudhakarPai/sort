package heap

import "github.com/SudhakarPai/sort"

// Sort sorts the collection in ascending order using the heap sort
func Sort(slice sort.Collection) {
	buildHeap(slice)
	for i := slice.Len() - 1; i >= 0; i-- {
		slice.Swap(i, 0)
		heapify(slice, 0, i)
	}
}

func buildHeap(slice sort.Collection) {
	for branch := slice.Len() >> 1; branch >= 0; branch-- {
		heapify(slice, branch, slice.Len())
	}
}

func heapify(slice sort.Collection, branch int, size int) {
	maxIndex := branch
	if leftChild := branch<<1 + 1; leftChild < size && slice.Greater(leftChild, maxIndex) {
		maxIndex = leftChild
	}
	if rightChild := branch<<1 + 2; rightChild < size && slice.Greater(rightChild, maxIndex) {
		maxIndex = rightChild
	}
	if maxIndex != branch {
		slice.Swap(branch, maxIndex)
		heapify(slice, maxIndex, size)
	}
}
