package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStartOfMonth(t *testing.T) {

	{
		date, err := time.Parse(layoutDateIn, "221201")
		assert.NoError(t, err)
		dateStart := BeginOfMonth(date)
		assert.Equal(t, "221201", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "221231")
		assert.NoError(t, err)
		dateStart := BeginOfMonth(date)
		assert.Equal(t, "221201", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "221216")
		assert.NoError(t, err)
		dateStart := BeginOfMonth(date)
		assert.Equal(t, "221201", dateStart.Format(layoutDateIn))
	}

}

func TestEndOfMonth(t *testing.T) {

	{
		date, err := time.Parse(layoutDateIn, "221231")
		assert.NoError(t, err)
		dateStart := EndOfMonth(date)
		assert.Equal(t, "221231", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "221219")
		assert.NoError(t, err)
		dateStart := EndOfMonth(date)
		assert.Equal(t, "221231", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "220122")
		assert.NoError(t, err)
		dateStart := EndOfMonth(date)
		assert.Equal(t, "220131", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "220224")
		assert.NoError(t, err)
		dateStart := EndOfMonth(date)
		assert.Equal(t, "220228", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "230224")
		assert.NoError(t, err)
		dateStart := EndOfMonth(date)
		assert.Equal(t, "230228", dateStart.Format(layoutDateIn))

	}
	{
		date, err := time.Parse(layoutDateIn, "240224")
		assert.NoError(t, err)
		dateStart := EndOfMonth(date)
		assert.Equal(t, "240229", dateStart.Format(layoutDateIn))
	}

}
