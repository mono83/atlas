package mysql

import (
	"fmt"
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/conditions"
	"github.com/mono83/atlas/query/match"
	"github.com/mono83/atlas/query/rules"
	"github.com/stretchr/testify/assert"
	"testing"
)

var conditionToSQLDataProvider = []struct {
	SQL          string
	Placeholders []interface{}
	Condition    query.Condition
}{
	{"`name` = ?", []interface{}{"bar"}, conditions.ForAllRules(rules.Eq(query.String("name"), "bar"))},
	{
		"(`id` > ? AND `name` = ?)",
		[]interface{}{10, "bar"},
		conditions.ForAllRules(
			rules.New(query.String("id"), match.GreaterThan, 10),
			rules.Eq(query.String("name"), "bar"),
		),
	},
	{
		"(`id` <= ? OR `name` = ?)",
		[]interface{}{3, "bar"},
		conditions.ForAnyRule(
			rules.New(query.String("id"), match.LesserThanEquals, 3),
			rules.Eq(query.String("name"), "bar"),
		),
	},
	{
		"(`type` = ? OR (`lastLoginAt` = `firstLoginAt` OR `blockedAt` IS NULL) OR (`scope` = ? AND `type` = ?))",
		[]interface{}{"admin", "current", "user"},
		conditions.New(
			query.Or,
			[]query.Rule{rules.Eq(query.String("type"), "admin")},
			[]query.Condition{
				conditions.ForAnyRule(
					rules.Eq(query.String("lastLoginAt"), query.String("firstLoginAt")),
					rules.IsNull(query.String("blockedAt")),
				),
				conditions.ForAllRules(
					rules.Eq(query.String("scope"), "current"),
					rules.Eq(query.String("type"), "user"),
				),
			},
		),
	},
}

func TestConditionToSQL(t *testing.T) {
	for _, d := range conditionToSQLDataProvider {
		t.Run(fmt.Sprintf("%v", d.Condition), func(t *testing.T) {
			b := NewStatementBuilder()
			if assert.NoError(t, b.WriteCondition(d.Condition)) {
				stmt := b.Build()
				if assert.Equal(t, d.SQL, stmt.GetSQL()) {
					if assert.Equal(t, len(d.Placeholders), len(stmt.GetPlaceholders()), "parameters count don't match") {
						for i, a := range stmt.GetPlaceholders() {
							assert.Equal(t, d.Placeholders[i], a)
						}
					}
				}
			}
		})
	}
}
