package merge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	cases := []struct {
		name   string
		input  []int
		output []int
	}{
		{
			name:   "Test_ascending_sort",
			input:  []int{13, 18, 10, 9, 7, 16, 1, 10, 80},
			output: []int{1, 7, 9, 10, 10, 13, 16, 18, 80},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			Sort(c.input)
			assert.Equal(t, c.output, c.input)
		})
	}
}
