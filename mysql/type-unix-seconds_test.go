package mysql

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestUnixTimestamp_Scan(t *testing.T) {
	assert := assert.New(t)

	var ts UnixTimestamp

	if assert.NoError(ts.Scan(int(10))) {
		assert.Equal(time.Unix(10, 0).UnixNano(), ts.UnixNano())
	}
	if assert.NoError(ts.Scan(int64(-33))) {
		assert.Equal(time.Unix(-33, 0).UnixNano(), ts.UnixNano())
	}
	if assert.NoError(ts.Scan(uint64(88))) {
		assert.Equal(time.Unix(88, 0).UnixNano(), ts.UnixNano())
	}
	if assert.NoError(ts.Scan("4567")) {
		assert.Equal(time.Unix(4567, 0).UnixNano(), ts.UnixNano())
	}
	if assert.NoError(ts.Scan([]byte("-989"))) {
		assert.Equal(time.Unix(-989, 0).UnixNano(), ts.UnixNano())
	}

	assert.Error(ts.Scan(int16(10)))
	assert.Error(ts.Scan(int8(10)))
	assert.Error(ts.Scan(byte(10)))
	assert.Error(ts.Scan(float64(10)))
	assert.Error(ts.Scan(float32(10)))
}
