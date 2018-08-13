package rules

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEq(t *testing.T) {
	assert := assert.New(t)

	r := Eq("foo", "bar")
	assert.Equal("foo", r.GetLeft())
	assert.Equal("bar", r.GetRight())
	assert.Equal(match.Eq, r.GetType())
}

func TestIsNull(t *testing.T) {
	assert := assert.New(t)

	r := IsNull("foo")
	assert.Equal("foo", r.GetLeft())
	assert.Nil(r.GetRight())
	assert.Equal(match.IsNull, r.GetType())
}

func TestIsNotNull(t *testing.T) {
	assert := assert.New(t)

	r := IsNotNull("foo")
	assert.Equal("foo", r.GetLeft())
	assert.Nil(r.GetRight())
	assert.Equal(match.NotIsNull, r.GetType())
}

func TestMatchID64(t *testing.T) {
	assert := assert.New(t)
	var r query.Rule

	r = MatchID64()
	assert.IsType(False{}, r)

	r = MatchID64(10)
	assert.Equal(match.Eq, r.GetType())
	assert.Equal("id", r.GetLeft())
	assert.Equal(int64(10), r.GetRight())

	r = MatchID64(10, 11, 12)
	assert.Equal(match.In, r.GetType())
	assert.Equal("id", r.GetLeft())
	v, ok := r.GetRight().([]int64)
	if assert.True(ok) {
		assert.Len(r.GetRight(), 3)
		assert.Equal(int64(10), v[0])
		assert.Equal(int64(11), v[1])
		assert.Equal(int64(12), v[2])
	}
}
