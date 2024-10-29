package funk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiff(t *testing.T) {
	{
		a := []int{1, 2, 3, 4}
		b := []int{3, 4, 5, 6}
		res := Diff(a, b)
		assert.Equal(t, []int{1, 2}, res)
	}

	{
		a := []string{"a", "b", "c"}
		b := []string{"b", "c", "d"}
		res := Diff(a, b)
		assert.Equal(t, []string{"a"}, res)
	}

	{
		a := []int{}
		b := []int{1, 2, 3}
		res := Diff(a, b)
		assert.Equal(t, []int{}, res)
	}

	{
		a := []int{1, 2, 3}
		b := []int{}
		res := Diff(a, b)
		assert.Equal(t, []int{1, 2, 3}, res)
	}

	{
		a := []int{1, 1, 2, 2}
		b := []int{2}
		res := Diff(a, b)
		assert.Equal(t, []int{1, 1}, res)
	}
}
