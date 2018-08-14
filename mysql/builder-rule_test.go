package mysql

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"github.com/mono83/atlas/query/rules"
	"testing"
)

func TestRuleToSQL(t *testing.T) {
	assertRule(t, rules.New(query.String("foo"), match.IsNull, nil), "`foo` IS NULL")
	assertRule(t, rules.New(query.String("bar"), match.NotIsNull, nil), "`bar` NOT IS NULL")
	assertRule(t, rules.New(query.String("foo"), match.Equals, query.String("bar")), "`foo` = `bar`")
	assertRule(t, rules.New(query.String("foo"), match.NotEquals, query.String("bar")), "`foo` <> `bar`")
	assertRule(t, rules.New(query.String("foo"), match.Equals, 5), "`foo` = ?", 5)
	assertRule(t, rules.New(query.String("foo"), match.NotEquals, "7"), "`foo` <> ?", "7")
	assertRule(t, rules.New(query.String("foo"), match.Gt, "7"), "`foo` > ?", "7")
	assertRule(t, rules.New(query.String("foo"), match.Gte, 7), "`foo` >= ?", 7)
	assertRule(t, rules.New(query.String("foo"), match.Lt, "7"), "`foo` < ?", "7")
	assertRule(t, rules.New(query.String("foo"), match.Lte, "7"), "`foo` <= ?", "7")
	assertRule(t, rules.New(query.String("bar"), match.In, []interface{}{5, int64(6), 7}), "`bar` IN (?,?,?)", 5, int64(6), 7)
	assertRule(t, rules.New(query.String("bar"), match.NotIn, []interface{}{3, "false"}), "`bar` NOT IN (?,?)", 3, "false")
	assertRule(t, rules.New(query.String("bar"), match.In, []interface{}{"true"}), "`bar` IN (?)", "true")
}
