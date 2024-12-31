package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStartOfYear(t *testing.T) {

	{
		date, err := time.Parse(layoutDateIn, "221201")
		assert.NoError(t, err)
		dateStart := BeginOfYear(date)
		assert.Equal(t, "220101", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "221130")
		assert.NoError(t, err)
		dateStart := BeginOfYear(date)
		assert.Equal(t, "220101", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "220606")
		assert.NoError(t, err)
		dateStart := BeginOfYear(date)
		assert.Equal(t, "220101", dateStart.Format(layoutDateIn))
	}

}

func TestEndOfYear(t *testing.T) {

	{
		date, err := time.Parse(layoutDateIn, "221231")
		assert.NoError(t, err)
		dateStart := EndOfYear(date)
		assert.Equal(t, "221231", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "220606")
		assert.NoError(t, err)
		dateStart := EndOfYear(date)
		assert.Equal(t, "221231", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "220122")
		assert.NoError(t, err)
		dateStart := EndOfYear(date)
		assert.Equal(t, "221231", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "220224")
		assert.NoError(t, err)
		dateStart := EndOfYear(date)
		assert.Equal(t, "221231", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "230224")
		assert.NoError(t, err)
		dateStart := EndOfYear(date)
		assert.Equal(t, "231231", dateStart.Format(layoutDateIn))
	}

	{
		date, err := time.Parse(layoutDateIn, "240224")
		assert.NoError(t, err)
		dateStart := EndOfYear(date)
		assert.Equal(t, "241231", dateStart.Format(layoutDateIn))
	}

}
