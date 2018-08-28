package mysql

import (
	"fmt"
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/rules"
	"github.com/stretchr/testify/assert"
	"testing"
)

var filterToSQLDataProvider = []struct {
	SQL    string
	Filter query.Filter
}{
	// No sorting
	{"`id` NOT IS NULL", query.CommonFilter{Type: query.And, Rules: []query.Rule{rules.IsNotNull(query.String("id"))}}},
	{"`id` NOT IS NULL LIMIT 5", query.CommonFilter{Type: query.And, Limit: 5, Rules: []query.Rule{rules.IsNotNull(query.String("id"))}}},
	{"`id` NOT IS NULL", query.CommonFilter{Type: query.And, Offset: 5, Rules: []query.Rule{rules.IsNotNull(query.String("id"))}}},
	{"`id` NOT IS NULL LIMIT 5,2", query.CommonFilter{Type: query.And, Offset: 5, Limit: 2, Rules: []query.Rule{rules.IsNotNull(query.String("id"))}}},

	// With sorting
	{
		"`id` NOT IS NULL ORDER BY `id` ASC",
		query.CommonFilter{
			Type:    query.And,
			Rules:   []query.Rule{rules.IsNotNull(query.String("id"))},
			Sorting: []query.Sorting{query.SimpleAsc("id")},
		},
	},
	{
		"`id` NOT IS NULL ORDER BY `id` ASC,`name` DESC",
		query.CommonFilter{
			Type:    query.And,
			Rules:   []query.Rule{rules.IsNotNull(query.String("id"))},
			Sorting: []query.Sorting{query.SimpleAsc("id"), query.SimpleDesc("name")},
		},
	},
	{
		"`id` NOT IS NULL ORDER BY `id` ASC LIMIT 3",
		query.CommonFilter{
			Type:    query.And,
			Rules:   []query.Rule{rules.IsNotNull(query.String("id"))},
			Sorting: []query.Sorting{query.SimpleAsc("id")},
			Limit:   3,
		},
	},
}

func TestFilterToSQL(t *testing.T) {
	for _, d := range filterToSQLDataProvider {
		t.Run(fmt.Sprintf("%v", d.Filter), func(t *testing.T) {
			b := NewStatementBuilder()
			if assert.NoError(t, b.WriteFilter(d.Filter)) {
				stmt := b.Build()
				assert.Equal(t, d.SQL, stmt.GetSQL())
			}
		})
	}
}
