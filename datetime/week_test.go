package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStartOfWeek(t *testing.T) {

	{
		date, err := time.Parse(layoutDateIn, "221219")
		assert.NoError(t, err)
		dateStart := BeginOfWeek(date)
		assert.Equal(t, "221219", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "221218")
		assert.NoError(t, err)
		dateStart := BeginOfWeek(date)
		assert.Equal(t, "221212", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "221216")
		assert.NoError(t, err)
		dateStart := BeginOfWeek(date)
		assert.Equal(t, "221212", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "221213")
		assert.NoError(t, err)
		dateStart := BeginOfWeek(date)
		assert.Equal(t, "221212", dateStart.Format(layoutDateIn))
	}

}

func TestEndOfWeek(t *testing.T) {

	{
		date, err := time.Parse(layoutDateIn, "221218")
		assert.NoError(t, err)
		dateStart := EndOfWeek(date)
		assert.Equal(t, "221218", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "221219")
		assert.NoError(t, err)
		dateStart := EndOfWeek(date)
		assert.Equal(t, "221225", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "221222")
		assert.NoError(t, err)
		dateStart := EndOfWeek(date)
		assert.Equal(t, "221225", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "221224")
		assert.NoError(t, err)
		dateStart := EndOfWeek(date)
		assert.Equal(t, "221225", dateStart.Format(layoutDateIn))
	}

}
