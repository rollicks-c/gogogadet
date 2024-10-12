package funk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetKeys(t *testing.T) {
	m := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	keys := GetKeys(m)
	assert.Equal(t, 3, len(keys))
	assert.Contains(t, []string{"key1", "key2", "key3"}, keys[0])
	assert.Contains(t, []string{"key1", "key2", "key3"}, keys[1])
	assert.Contains(t, []string{"key1", "key2", "key3"}, keys[2])
}

func TestPickOne(t *testing.T) {
	{
		m := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}
		value, ok := PickOne(m)
		assert.True(t, ok)
		assert.Contains(t, []string{"value1", "value2", "value3"}, value)
	}
	{
		m := map[string]string{}
		_, ok := PickOne(m)
		assert.False(t, ok)
	}
}
