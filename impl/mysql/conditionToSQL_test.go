package mysql

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"github.com/mono83/atlas/query/rules"
	"testing"
)

func TestConditionToSQL(t *testing.T) {
	assert := sqla{t}

	cond := query.CommonCondition{
		Type:  query.Or,
		Rules: []query.Rule{rules.New(query.String("foo"), match.IsNull, nil)},
		Conditions: []query.Condition{
			query.CommonCondition{
				Type: query.And,
				Rules: []query.Rule{
					rules.New(query.String("bar"), match.Equals, "5"),
					rules.New(query.String("baz"), match.Lte, 300),
				},
			},
		},
	}

	assert.CondSQL(cond, "(`foo` IS NULL OR (`bar` = ? AND `baz` <= ?))")
	assert.CondPh(cond, "5", 300)
}
