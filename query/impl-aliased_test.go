package query

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAliasedName(t *testing.T) {
	assert := assert.New(t)

	a := AliasedName{Name: "foo", Alias: "bar"}
	assert.Equal("foo", a.GetName())
	assert.Equal("bar", a.GetAlias())
	assert.Equal("foo", a.String())

	a = AliasedName{Name: "foo"}
	assert.Equal("foo", a.GetName())
	assert.Equal("", a.GetAlias())
	assert.Equal("foo", a.String())

	a = AliasedName{Alias: "bar"}
	assert.Equal("", a.GetName())
	assert.Equal("bar", a.GetAlias())
	assert.Equal("", a.String())
}
