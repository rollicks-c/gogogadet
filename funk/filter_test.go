package funk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNonEmpty(t *testing.T) {

	pool := []string{
		"a",
		"B",
		"",
		"d",
	}
	res := NonEmpty(pool...)

	assert.Equal(t, 3, len(res))
	assert.Equal(t, "a", res[0])
	assert.Equal(t, "B", res[1])
	assert.Equal(t, "d", res[2])

}

func TestFirstNonEmpty(t *testing.T) {
	{
		pool := []string{
			"",
			"",
			"c",
		}
		res, ok := FirstNonEmpty(pool...)
		assert.True(t, ok)
		assert.Equal(t, "c", res)
	}

	{
		pool := []string{
			"",
			"",
			"",
		}
		res, ok := FirstNonEmpty(pool...)
		assert.False(t, ok)
		assert.Equal(t, "", res)
	}
}
