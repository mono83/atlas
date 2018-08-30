package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

func BenchmarkOptString_Scan(b *testing.B) {
	inString := "This is string"
	inBytes := []byte(inString)
	var out OptString

	b.Run("FromString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := out.Scan(inString); err != nil {
				b.Error(b)
			}
		}
	})
	b.Run("FromBytes", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := out.Scan(inBytes); err != nil {
				b.Error(b)
			}
		}
	})
	b.Run("FromNil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := out.Scan(nil); err != nil {
				b.Error(b)
			}
		}
	})
}

func BenchmarkOptString_MapFilterOrElse(b *testing.B) {
	v := "foo"
	for i := 0; i < b.N; i++ {
		r := OptString{value: &v}.Map(func(string) string { return "bar" }).Filter(func(x string) bool { return x == "foo" }).OrElse("baz")
		if r != "baz" {
			b.Error("baz expected")
		}
	}
}
