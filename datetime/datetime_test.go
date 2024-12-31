package datetime

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	layoutDateIn = "060102"
)

func TestDateOnly(t *testing.T) {

	{
		date, err := time.Parse(time.RFC3339, "2022-12-19T15:04:05Z")
		assert.NoError(t, err)
		date1Only := DateOnly(date)
		assert.Equal(t, "2022-12-19T00:00:00Z", date1Only.Format(time.RFC3339))
	}

	{
		date, err := time.Parse(time.RFC3339, "2022-12-19T00:00:00Z")
		assert.NoError(t, err)
		date2Only := DateOnly(date)
		assert.Equal(t, "2022-12-19T00:00:00Z", date2Only.Format(time.RFC3339))
	}

	{
		date, err := time.Parse(time.RFC3339, "2022-12-19T23:59:59Z")
		assert.NoError(t, err)
		date3Only := DateOnly(date)
		assert.Equal(t, "2022-12-19T00:00:00Z", date3Only.Format(time.RFC3339))
	}

}

func TestCreateTime(t *testing.T) {

	{
		date, err := time.Parse(time.RFC3339, "2022-12-19T00:00:00Z")
		assert.NoError(t, err)
		startTime, err := CreateTime(date, "15:04")
		assert.NoError(t, err)
		assert.Equal(t, "2022-12-19T15:04:00Z", startTime.Format(time.RFC3339))
	}

	{
		date, err := time.Parse(time.RFC3339, "2022-12-20T00:00:00Z")
		assert.NoError(t, err)
		startTime, err := CreateTime(date, "00:00")
		assert.NoError(t, err)
		assert.Equal(t, "2022-12-20T00:00:00Z", startTime.Format(time.RFC3339))
	}

	{
		date, err := time.Parse(time.RFC3339, "2022-12-21T00:00:00Z")
		assert.NoError(t, err)
		startTime, err := CreateTime(date, "23:59")
		assert.NoError(t, err)
		assert.Equal(t, "2022-12-21T23:59:00Z", startTime.Format(time.RFC3339))
	}
}
