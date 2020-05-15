package query

import (
	"testing"

	"github.com/mono83/atlas/query/match"
	"github.com/stretchr/testify/assert"
)

func TestCommonRule(t *testing.T) {
	assert := assert.New(t)

	r := CommonRule{L: 1, R: "foo", Type: match.Equals}
	assert.Equal(1, r.GetLeft())
	assert.Equal("foo", r.GetRight())
	assert.Equal(match.Equals, r.GetType())
	assert.Equal("{CommonRule {1 (int)} Equal {foo (string)}}", r.String())
}
