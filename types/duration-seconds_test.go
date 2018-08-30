package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var durationSecondsDataProvider = []struct {
	Expected time.Duration
	Error    string
	Source   interface{}
}{
	{10 * time.Second, "", int(10)},
	{11 * time.Second, "", int32(11)},
	{12 * time.Second, "", int64(12)},
	{13 * time.Second, "", uint32(13)},
	{14 * time.Second, "", uint64(14)},
	{time.Duration(0), "unable to sql.Scan into DurationSeconds from <nil>", nil},
	{time.Duration(0), "unable to sql.Scan into DurationSeconds from uint8", byte(1)},
	{time.Duration(0), "unable to sql.Scan into DurationSeconds from slice of uint8", []byte{}},
	{time.Duration(0), "unable to sql.Scan into DurationSeconds from slice of uint8", []byte{1, 2, 3}},
	{time.Duration(0), "unable to sql.Scan into DurationSeconds from string", "123"},
}

func TestDurationSeconds_Scan(t *testing.T) {
	for _, d := range durationSecondsDataProvider {
		t.Run(d.Expected.String(), func(t *testing.T) {
			var ds DurationSeconds
			err := ds.Scan(d.Source)
			if err != nil {
				assert.Equal(t, d.Error, err.Error())
			} else {
				if assert.Equal(t, "", d.Error) {
					assert.Equal(t, d.Expected, ds.Duration)
				}
			}
		})
	}
}

func BenchmarkDurationSeconds_Scan(b *testing.B) {
	inInt64 := int64(1000)
	inInt32 := int32(1000)
	inUint64 := uint64(1000)
	inUint32 := uint32(1000)
	inInt := int(1000)
	var out DurationSeconds

	b.Run("FromInt64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := out.Scan(inInt64); err != nil {
				b.Error(b)
			}
		}
	})
	b.Run("FromInt32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := out.Scan(inInt32); err != nil {
				b.Error(b)
			}
		}
	})
	b.Run("FromUint64", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := out.Scan(inUint64); err != nil {
				b.Error(b)
			}
		}
	})
	b.Run("FromUint32", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := out.Scan(inUint32); err != nil {
				b.Error(b)
			}
		}
	})
	b.Run("FromInt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := out.Scan(inInt); err != nil {
				b.Error(b)
			}
		}
	})
}
