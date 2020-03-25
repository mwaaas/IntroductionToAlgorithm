package sorting

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSorting(t *testing.T) {
	scenarios := []struct {
		listToSort         []int
		expectedSortedList []int
	}{
		{
			[]int{},
			[]int{},
		},
		{
			[]int{4, 5, 2, 0},
			[]int{0, 2, 4, 5},
		},
		{
			[]int{2, 9, -1, 3},
			[]int{-1, 2, 3, 9},
		},
	}

	for _, i := range scenarios {
		assert.Equal(t, i.expectedSortedList, insertion(i.listToSort))
	}

}
