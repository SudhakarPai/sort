package insertion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	cases := []struct {
		name        string
		slice       arrayToBeSorted
		sortedSlice arrayToBeSorted
	}{
		{
			name:        "Sort_int_array",
			slice:       arrayToBeSorted([]int{1, 6, 2, 7, 3, 9, 0}),
			sortedSlice: arrayToBeSorted([]int{0, 1, 2, 3, 6, 7, 9}),
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			Sort(&c.slice)
			assert.Equal(t, c.sortedSlice, c.slice)
		})
	}
}

type arrayToBeSorted []int

func (ats *arrayToBeSorted) Len() int {
	return len(*ats)
}

func (ats *arrayToBeSorted) Greater(i, j int) bool {
	return (*ats)[i] > (*ats)[j]
}

func (ats *arrayToBeSorted) Swap(i, j int) {
	(*ats)[i], (*ats)[j] = (*ats)[j], (*ats)[i]
}
