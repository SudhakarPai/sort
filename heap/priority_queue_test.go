package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePriorityQueue(t *testing.T) {
	queue := queue([]int{9, 7, 8, 6, 3, 0, 2})
	cases := []struct {
		name          string
		slice         []int
		increaseIndex int
		increaseValue int
		maxvalue      int
		addValue      int
		queue         PriorityQueue
	}{
		{
			name:          "Sort_int_array",
			slice:         []int{1, 6, 2, 7, 3, 9, 0},
			increaseIndex: 5,
			increaseValue: 10,
			maxvalue:      10,
			addValue:      8,
			queue:         &queue,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			q := CreatePriorityQueue(c.slice)
			q.IncreasePriority(c.increaseIndex, c.increaseValue)
			maxvalue := q.ExtractMax()
			assert.Equal(t, c.maxvalue, maxvalue)
			q.AddNewItem(c.addValue)
			assert.Equal(t, c.queue, q)
		})
	}
}
