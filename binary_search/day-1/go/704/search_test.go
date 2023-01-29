package binary_search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	{
		tc1 := []int{-1, 0, 3, 5, 9, 12}
		target := 9
		expected := 4
		assert.Equal(t, expected, Search(tc1, target))
	}
	{
		tc1 := []int{-1, 0, 3, 5, 9, 12}
		target := 2
		expected := -1
		assert.Equal(t, expected, Search(tc1, target))
	}
}
