package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBeginOfDay(t *testing.T) {
	date := time.Date(2023, 10, 5, 15, 30, 45, 123456789, time.UTC)
	expected := time.Date(2023, 10, 5, 0, 0, 0, 0, time.UTC)
	result := BeginOfDay(date)
	assert.Equal(t, expected, result)
}

func TestEndOfDay(t *testing.T) {
	date := time.Date(2023, 10, 5, 15, 30, 45, 123456789, time.UTC)
	expected := time.Date(2023, 10, 5, 23, 59, 59, 999999999, time.UTC)
	result := EndOfDay(date)
	assert.Equal(t, expected, result)
}
