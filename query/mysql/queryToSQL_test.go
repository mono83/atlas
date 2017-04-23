package mysql

import (
	"github.com/mono83/atlas/query"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueryToSQL(t *testing.T) {
	assert := assert.New(t)

	q := query.Select{query.String("tokens"), nil, nil, nil, 0, 0}

	sql, ph, err := QueryToSQL(q)
	assert.NoError(err)
	assert.Len(ph, 0)
	assert.Equal("SELECT * FROM `tokens`", sql)

	q = query.Select{query.String("tokens"), nil, nil, []query.OrderDef{query.SimpleAsc("id"), query.SimpleDesc("time")}, 2, 0}

	sql, ph, err = QueryToSQL(q)
	assert.NoError(err)
	assert.Len(ph, 0)
	assert.Equal("SELECT * FROM `tokens` ORDER BY `id` ASC,`time` DESC LIMIT 2", sql)

	// Building query
	q = query.Select{
		query.String("users"),
		[]query.Named{query.String("id"), query.AliasedName{"user_name", "name"}},
		nil,
		nil,
		10,
		500,
	}

	sql, ph, err = QueryToSQL(q)
	assert.NoError(err)
	assert.Len(ph, 0)
	assert.Equal("SELECT `id`,`user_name` AS `name` FROM `users` LIMIT 500,10", sql)
}
