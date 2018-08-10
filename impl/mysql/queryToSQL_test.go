package mysql

import (
	"github.com/mono83/atlas/query"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryToSQL(t *testing.T) {
	assert := assert.New(t)

	q := query.Select{Schema: query.String("tokens")}

	sql, ph, err := QueryToSQL(q)
	assert.NoError(err)
	assert.Len(ph, 0)
	assert.Equal("SELECT * FROM `tokens`", sql)

	q = query.Select{Schema: query.String("tokens"), Order: []query.Sorting{query.SimpleAsc("id"), query.SimpleDesc("time")}, Limit: 2}

	sql, ph, err = QueryToSQL(q)
	assert.NoError(err)
	assert.Len(ph, 0)
	assert.Equal("SELECT * FROM `tokens` ORDER BY `id` ASC,`time` DESC LIMIT 2", sql)

	// Building query
	q = query.Select{
		Schema:  query.String("users"),
		Columns: []query.Named{query.String("id"), query.AliasedName{Name: "user_name", Alias: "name"}},
		Limit:   10,
		Offset:  500,
	}

	sql, ph, err = QueryToSQL(q)
	assert.NoError(err)
	assert.Len(ph, 0)
	assert.Equal("SELECT `id`,`user_name` AS `name` FROM `users` LIMIT 500,10", sql)
}
