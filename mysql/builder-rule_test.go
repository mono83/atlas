package mysql

import (
	"fmt"
	"testing"

	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"github.com/mono83/atlas/query/rules"
	"github.com/stretchr/testify/assert"
)

var ruleToSQLDataProvider = []struct {
	SQL          string
	Placeholders []interface{}
	Rule         query.Rule
}{
	{"1=0", nil, rules.False{}},
	{"`foo` IS NULL", nil, rules.New(query.String("foo"), match.IsNull, nil)},
	{"`bar` IS NULL", nil, rules.IsNull(query.String("bar"))},
	{"`bar` NOT IS NULL", nil, rules.New(query.String("bar"), match.NotIsNull, nil)},
	{"`foo` = `bar`", nil, rules.New(query.String("foo"), match.Equals, query.String("bar"))},
	{"`foo` <> `bar`", nil, rules.New(query.String("foo"), match.NotEquals, query.String("bar"))},
	{"`foo` = ?", []interface{}{5}, rules.New(query.String("foo"), match.Equals, 5)},
	{"`baz` = ?", []interface{}{3}, rules.Eq(query.String("baz"), 3)},
	{"`foo` <> ?", []interface{}{"bar"}, rules.New(query.String("foo"), match.NotEquals, "bar")},
	{"`foo` > ?", []interface{}{"7"}, rules.New(query.String("foo"), match.Gt, "7")},
	{"`foo` >= ?", []interface{}{7}, rules.New(query.String("foo"), match.Gte, 7)},
	{"`foo` < ?", []interface{}{"7"}, rules.New(query.String("foo"), match.Lt, "7")},
	{"`foo` <= ?", []interface{}{"7"}, rules.New(query.String("foo"), match.Lte, "7")},
	{"`bar` IN (?,?,?)", []interface{}{5, int64(6), 7}, rules.New(query.String("bar"), match.In, []interface{}{5, int64(6), 7})},
	{"`bar` NOT IN (?,?)", []interface{}{3, "false"}, rules.New(query.String("bar"), match.NotIn, []interface{}{3, "false"})},
	{"`bar` IN (?)", []interface{}{true}, rules.New(query.String("bar"), match.In, []interface{}{true})},
}

func TestRuleToSQL(t *testing.T) {
	for _, d := range ruleToSQLDataProvider {
		t.Run(fmt.Sprintf("%v", d.Rule), func(t *testing.T) {
			b := NewStatementBuilder()
			if assert.NoError(t, b.WriteRule(d.Rule)) {
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

var ruleToSQLErrorsDataProvder = []struct {
	Error string
	Rule  query.Rule
}{
	{"nil rule", nil},
	{"no column definition on left side of rule", rules.IsNull("foo")},
	{"no column definition on left side of rule", rules.IsNotNull("foo")},
	{"no column definition on left side of rule", rules.New("foo", match.Equals, "bar")},
	{"no column definition on left side of rule", rules.New("foo", match.NotEquals, "bar")},
	{"no column definition on left side of rule", rules.New("foo", match.GreaterThan, "bar")},
	{"no column definition on left side of rule", rules.New("foo", match.GreaterThanEquals, "bar")},
	{"no column definition on left side of rule", rules.New("foo", match.LowerThan, "bar")},
	{"no column definition on left side of rule", rules.New("foo", match.LowerThanEquals, "bar")},
	{"nil provided in right side of IN/NOT IN operation", rules.New("foo", match.In, nil)},
	{"nil provided in right side of IN/NOT IN operation", rules.New("foo", match.NotIn, nil)},
	{"no column definition on left side of rule", rules.New("foo", match.In, []int{1})},
	{"no column definition on left side of rule", rules.New("foo", match.NotIn, []int{1})},
	{"missing data for IN operations - empty values slice received", rules.New("foo", match.In, []int{})},
	{"missing data for IN operations - empty values slice received", rules.New("foo", match.NotIn, []int{})},
	{"only []int, []int64, []string and []interface{} are allowed for IN operations", rules.New("foo", match.In, []float32{})},
	{"only []int, []int64, []string and []interface{} are allowed for IN operations", rules.New("foo", match.NotIn, []float32{})},
}

func TestRuleToSQLErrors(t *testing.T) {
	for _, d := range ruleToSQLErrorsDataProvder {
		t.Run(fmt.Sprintf("%v", d.Rule), func(t *testing.T) {
			err := NewStatementBuilder().WriteRule(d.Rule)
			if err == nil {
				t.Errorf(`expected error "%s" but got nothing`, d.Error)
			} else if err.Error() != d.Error {
				t.Errorf(`expected error "%s" but got "%s"`, d.Error, err.Error())
			}
		})
	}
}
