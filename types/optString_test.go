package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var optStringScanDataProvider = []struct {
	IsPresent  bool
	IsNotEmpty bool
	Value      string
	Error      string
	Source     interface {
	}
}{
	{false, false, "", "", nil},
	{true, false, "", "", ""},
	{true, false, "", "", []byte{}},
	{true, true, "foo", "", "foo"},
	{true, true, "bar", "", []byte("bar")},
	{false, false, "", "unable to sql.Scan into OptString from uint8", byte(1)},
	{false, false, "", "unable to sql.Scan into OptString from int16", int16(2)},
	{false, false, "", "unable to sql.Scan into OptString from int32", int32(3)},
	{false, false, "", "unable to sql.Scan into OptString from int64", int64(4)},
	{false, false, "", "unable to sql.Scan into OptString from uint16", uint16(5)},
	{false, false, "", "unable to sql.Scan into OptString from uint32", uint32(6)},
	{false, false, "", "unable to sql.Scan into OptString from uint64", uint64(7)},
	{false, false, "", "unable to sql.Scan into OptString from float32", float32(8)},
	{false, false, "", "unable to sql.Scan into OptString from float64", float64(9)},
	{false, false, "", "", OptString{}},
	{false, false, "", "", &OptString{}},
}

func TestOptStringScan(t *testing.T) {
	for _, d := range optStringScanDataProvider {
		t.Run("", func(t *testing.T) {
			var os OptString
			err := os.Scan(d.Source)
			if err != nil {
				assert.Equal(t, d.Error, err.Error())
			} else {
				if assert.Equal(t, d.Error, "") {
					assert.Equal(t, d.IsPresent, os.IsPresent())
					assert.Equal(t, d.IsNotEmpty, os.IsNotEmpty())
					assert.Equal(t, d.Value, os.OrElse(""))
				}
			}
		})
	}
}
