package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleAsc(t *testing.T) {
	assert := assert.New(t)

	var sa Sorting = SimpleAsc("foo")
	assert.Equal("foo", sa.GetName())
	assert.Equal(Asc, sa.GetType())
}

func TestSimpleDesc(t *testing.T) {
	assert := assert.New(t)

	var sd Sorting = SimpleDesc("bar")
	assert.Equal("bar", sd.GetName())
	assert.Equal(Desc, sd.GetType())
}

func TestCommonSorting(t *testing.T) {
	assert := assert.New(t)

	var cs Sorting = CommonSorting{Column: String("foobar"), Type: Asc}
	assert.Equal("foobar", cs.GetName())
	assert.Equal(Asc, cs.GetType())

	cs = CommonSorting{Column: String("barbaz"), Type: Desc}
	assert.Equal("barbaz", cs.GetName())
	assert.Equal(Desc, cs.GetType())

	cs = CommonSorting{Column: String("baz")}
	assert.Equal("baz", cs.GetName())
	assert.Equal(UnknownSort, cs.GetType())
}
