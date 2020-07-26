package heap

type PriorityQueue interface {
	ExtractMax() int
	AddNewItem(value int)
	IncreasePriority(index, value int)
}

type queue []int

func (q *queue) ExtractMax() int {
	size := len(*q)
	(*q)[0], (*q)[size-1] = (*q)[size-1], (*q)[0]
	heapifyPQ(*q, 0, size-1)
	max := (*q)[size-1]
	*q = (*q)[:size-1]
	return max
}

func (q *queue) AddNewItem(value int) {
	*q = append(*q, value)
	q.IncreasePriority(len(*q)-1, value)
}

func (q *queue) IncreasePriority(index, value int) {
	size := len(*q)
	if size <= index {
		return
	}
	(*q)[index] = value
	parentHeapify(*q, index)
}

func parentHeapify(q queue, index int) {
	if index <= 0 {
		return
	}
	parentIndex := (index >> 1) + (index & 1) - 1
	if q[parentIndex] < q[index] {
		q[index], q[parentIndex] = q[parentIndex], q[index]
		parentHeapify(q, parentIndex)
	}
}

func heapifyPQ(q queue, index, size int) {
	maxIndex := index
	if leftIndex := index<<1 + 1; leftIndex < size && q[maxIndex] <= q[leftIndex] {
		maxIndex = leftIndex
	}
	if rightIndex := index<<1 + 2; rightIndex < size && q[maxIndex] <= q[rightIndex] {
		maxIndex = rightIndex
	}
	if index != maxIndex {
		q[index], q[maxIndex] = q[maxIndex], q[index]
		heapifyPQ(q, maxIndex, size)
	}
}

// CreatePriorityQueue accepts collection and returns the priority queue implementation of that
func CreatePriorityQueue(slice []int) PriorityQueue {
	var q queue
	q = slice
	size := len(q)
	for i := len(q) >> 1; i >= 0; i-- {
		heapifyPQ(q, i, size)
	}
	return &q
}
