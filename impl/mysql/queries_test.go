package mysql

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/queries"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectAll(t *testing.T) {
	assert := assert.New(t)

	sql, ph, err := QueryToSQL(queries.SelectAll("foo"))
	assert.NoError(err)
	assert.Len(ph, 0)
	assert.Equal("SELECT * FROM `foo`", sql)

	sql, ph, err = QueryToSQL(queries.SelectAll("foo", "id", query.AliasedName{"UserName", "name"}))
	assert.NoError(err)
	assert.Len(ph, 0)
	assert.Equal("SELECT `id`,`UserName` AS `name` FROM `foo`", sql)

	sql, ph, err = QueryToSQL(queries.FindById64("token", 15))
	assert.NoError(err)
	if assert.Len(ph, 1) {
		assert.Equal(int64(15), ph[0])
	}
	assert.Equal("SELECT * FROM `token` WHERE `id` = ? LIMIT 1", sql)
}
