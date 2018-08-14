package mysql

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"github.com/mono83/atlas/query/rules"
	"testing"
)

func TestConditionToSQL(t *testing.T) {
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

	assertCondition(
		t,
		cond,
		"(`foo` IS NULL OR (`bar` = ? AND `baz` <= ?))",
		"5", 300,
	)
}

func TestFilterToSQL(t *testing.T) {
	filter := query.CommonFilter{
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

	assertFilter(
		t,
		filter,
		"(`foo` IS NULL OR (`bar` = ? AND `baz` <= ?))",
		"5", 300,
	)

	filter = query.CommonFilter{
		Type:  query.Or,
		Rules: []query.Rule{rules.False{}},
		Limit: 2,
	}

	assertFilter(
		t,
		filter,
		"1=0 LIMIT 2",
	)

	filter = query.CommonFilter{
		Type:    query.Or,
		Rules:   []query.Rule{rules.False{}},
		Limit:   2,
		Offset:  8,
		Sorting: []query.Sorting{query.SimpleAsc("id"), query.SimpleDesc("name")},
	}

	assertFilter(
		t,
		filter,
		"1=0 ORDER BY `id` ASC,`name` DESC LIMIT 8,2",
	)
}
