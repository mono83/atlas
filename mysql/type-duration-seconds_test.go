package mysql

import (
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestDurationSeconds_Scan(t *testing.T) {
	assert := assert.New(t)

	var ts DurationSeconds

	if assert.NoError(ts.Scan(int(10))) {
		assert.Equal(time.Duration(10 * time.Second).Seconds(), ts.Seconds())
	}
	if assert.NoError(ts.Scan(int64(-33))) {
		assert.Equal(time.Duration(-33 * time.Second).Seconds(), ts.Seconds())
	}
	if assert.NoError(ts.Scan(uint64(88))) {
		assert.Equal(time.Duration(88 * time.Second).Seconds(), ts.Seconds())
	}
	if assert.NoError(ts.Scan("4567")) {
		assert.Equal(time.Duration(4567 * time.Second).Seconds(), ts.Seconds())
	}
	if assert.NoError(ts.Scan([]byte("-989"))) {
		assert.Equal(time.Duration(-989 * time.Second).Seconds(), ts.Seconds())
	}

	assert.Error(ts.Scan(int16(10)))
	assert.Error(ts.Scan(int8(10)))
	assert.Error(ts.Scan(byte(10)))
	assert.Error(ts.Scan(float64(10)))
	assert.Error(ts.Scan(float32(10)))
}
