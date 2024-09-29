package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	layoutDateIn = "060102"
)

func TestStartOfWeek(t *testing.T) {

	date1, err := time.Parse(layoutDateIn, "221219")
	assert.NoError(t, err)
	date1Start := BeginOfWeek(date1)
	assert.Equal(t, "221219", date1Start.Format(layoutDateIn))

	date2, err := time.Parse(layoutDateIn, "221218")
	assert.NoError(t, err)
	date2Start := BeginOfWeek(date2)
	assert.Equal(t, "221212", date2Start.Format(layoutDateIn))

	date3, err := time.Parse(layoutDateIn, "221216")
	assert.NoError(t, err)
	date3Start := BeginOfWeek(date3)
	assert.Equal(t, "221212", date3Start.Format(layoutDateIn))

	date4, err := time.Parse(layoutDateIn, "221213")
	assert.NoError(t, err)
	date4Start := BeginOfWeek(date4)
	assert.Equal(t, "221212", date4Start.Format(layoutDateIn))

}

func TestEndOfWeek(t *testing.T) {

	date1, err := time.Parse(layoutDateIn, "221218")
	assert.NoError(t, err)
	date1Start := EndOfWeek(date1)
	assert.Equal(t, "221218", date1Start.Format(layoutDateIn))

	date2, err := time.Parse(layoutDateIn, "221219")
	assert.NoError(t, err)
	date2Start := EndOfWeek(date2)
	assert.Equal(t, "221225", date2Start.Format(layoutDateIn))

	date3, err := time.Parse(layoutDateIn, "221222")
	assert.NoError(t, err)
	date3Start := EndOfWeek(date3)
	assert.Equal(t, "221225", date3Start.Format(layoutDateIn))

	date4, err := time.Parse(layoutDateIn, "221224")
	assert.NoError(t, err)
	date4Start := EndOfWeek(date4)
	assert.Equal(t, "221225", date4Start.Format(layoutDateIn))

}

func TestStartOfMonth(t *testing.T) {

	date1, err := time.Parse(layoutDateIn, "221201")
	assert.NoError(t, err)
	date1Start := BeginOfMonth(date1)
	assert.Equal(t, "221201", date1Start.Format(layoutDateIn))

	date2, err := time.Parse(layoutDateIn, "221231")
	assert.NoError(t, err)
	date2Start := BeginOfMonth(date2)
	assert.Equal(t, "221201", date2Start.Format(layoutDateIn))

	date3, err := time.Parse(layoutDateIn, "221216")
	assert.NoError(t, err)
	date3Start := BeginOfMonth(date3)
	assert.Equal(t, "221201", date3Start.Format(layoutDateIn))

}

func TestEndOfMonth(t *testing.T) {

	date1, err := time.Parse(layoutDateIn, "221231")
	assert.NoError(t, err)
	date1Start := EndOfMonth(date1)
	assert.Equal(t, "221231", date1Start.Format(layoutDateIn))

	date2, err := time.Parse(layoutDateIn, "221219")
	assert.NoError(t, err)
	date2Start := EndOfMonth(date2)
	assert.Equal(t, "221231", date2Start.Format(layoutDateIn))

	date3, err := time.Parse(layoutDateIn, "220122")
	assert.NoError(t, err)
	date3Start := EndOfMonth(date3)
	assert.Equal(t, "220131", date3Start.Format(layoutDateIn))

	date4, err := time.Parse(layoutDateIn, "220224")
	assert.NoError(t, err)
	date4Start := EndOfMonth(date4)
	assert.Equal(t, "220228", date4Start.Format(layoutDateIn))

	date5, err := time.Parse(layoutDateIn, "230224")
	assert.NoError(t, err)
	date5Start := EndOfMonth(date5)
	assert.Equal(t, "230228", date5Start.Format(layoutDateIn))

	date6, err := time.Parse(layoutDateIn, "240224")
	assert.NoError(t, err)
	date6Start := EndOfMonth(date6)
	assert.Equal(t, "240229", date6Start.Format(layoutDateIn))

}

func TestStartOfYear(t *testing.T) {

	date1, err := time.Parse(layoutDateIn, "221201")
	assert.NoError(t, err)
	date1Start := BeginOfYear(date1)
	assert.Equal(t, "220101", date1Start.Format(layoutDateIn))

	date2, err := time.Parse(layoutDateIn, "221130")
	assert.NoError(t, err)
	date2Start := BeginOfYear(date2)
	assert.Equal(t, "220101", date2Start.Format(layoutDateIn))

	date3, err := time.Parse(layoutDateIn, "220606")
	assert.NoError(t, err)
	date3Start := BeginOfYear(date3)
	assert.Equal(t, "220101", date3Start.Format(layoutDateIn))

}

func TestEndOfYear(t *testing.T) {

	date1, err := time.Parse(layoutDateIn, "221231")
	assert.NoError(t, err)
	date1Start := EndOfYear(date1)
	assert.Equal(t, "221231", date1Start.Format(layoutDateIn))

	date2, err := time.Parse(layoutDateIn, "220606")
	assert.NoError(t, err)
	date2Start := EndOfYear(date2)
	assert.Equal(t, "221231", date2Start.Format(layoutDateIn))

	date3, err := time.Parse(layoutDateIn, "220122")
	assert.NoError(t, err)
	date3Start := EndOfYear(date3)
	assert.Equal(t, "221231", date3Start.Format(layoutDateIn))

	date4, err := time.Parse(layoutDateIn, "220224")
	assert.NoError(t, err)
	date4Start := EndOfYear(date4)
	assert.Equal(t, "221231", date4Start.Format(layoutDateIn))

	date5, err := time.Parse(layoutDateIn, "230224")
	assert.NoError(t, err)
	date5Start := EndOfYear(date5)
	assert.Equal(t, "231231", date5Start.Format(layoutDateIn))

	date6, err := time.Parse(layoutDateIn, "240224")
	assert.NoError(t, err)
	date6Start := EndOfYear(date6)
	assert.Equal(t, "241231", date6Start.Format(layoutDateIn))

}
