package mysql

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"testing"
)

func TestConditionToSQL(t *testing.T) {
	assert := sqla{t}

	cond := query.Condition{
		query.Or,
		[]query.RuleDef{query.Rule{query.String("foo"), match.IsNull, nil}},
		[]query.ConditionDef{
			query.Condition{
				query.And,
				[]query.RuleDef{
					query.Rule{query.String("bar"), match.Equals, "5"},
					query.Rule{query.String("baz"), match.Lte, 300},
				},
				nil,
			},
		},
	}

	assert.CondSQL(cond, "(`foo` IS NULL OR (`bar` = ? AND `baz` <= ?))")
	assert.CondPh(cond, "5", 300)
}
