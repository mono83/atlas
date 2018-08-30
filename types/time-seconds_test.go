package types

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timeSecondsDataProvider = []struct {
	Expected time.Time
	Error    string
	Source   interface{}
}{
	{time.Unix(10, 0), "", int(10)},
	{time.Unix(11, 0), "", int32(11)},
	{time.Unix(12, 0), "", int64(12)},
	{time.Unix(13, 0), "", uint32(13)},
	{time.Unix(14, 0), "", uint64(14)},
	{time.Time{}, "unable to sql.Scan into TimeSeconds from <nil>", nil},
	{time.Time{}, "unable to sql.Scan into TimeSeconds from uint8", byte(1)},
	{time.Time{}, "unable to sql.Scan into TimeSeconds from slice of uint8", []byte{}},
	{time.Time{}, "unable to sql.Scan into TimeSeconds from slice of uint8", []byte{1, 2, 3}},
	{time.Time{}, "unable to sql.Scan into TimeSeconds from string", "123"},
}

func TestUnixSeconds_Scan(t *testing.T) {
	for _, d := range timeSecondsDataProvider {
		t.Run(d.Expected.String(), func(t *testing.T) {
			var ts TimeSeconds
			err := ts.Scan(d.Source)
			if err != nil {
				assert.Equal(t, d.Error, err.Error())
			} else {
				if assert.Equal(t, "", d.Error) {
					assert.Equal(t, d.Expected.UTC(), ts.Time.UTC())
				}
			}
		})
	}
}

func BenchmarkTimeSeconds_Scan(b *testing.B) {
	inInt64 := int64(1000)
	inInt32 := int32(1000)
	inUint64 := uint64(1000)
	inUint32 := uint32(1000)
	inInt := int(1000)
	var out TimeSeconds

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
