package mysql

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"github.com/stretchr/testify/assert"
	"testing"
)

type sqla struct {
	t *testing.T
}

func (s sqla) RuleSQL(r query.RuleDef, sql string) {
	p, e := RuleToSQL(r)
	if assert.NoError(s.t, e) {
		assert.Equal(s.t, sql, p.SQL)
	}
}
func (s sqla) RulePh(r query.RuleDef, ph ...interface{}) {
	p, e := RuleToSQL(r)
	if assert.NoError(s.t, e) {
		if assert.Len(s.t, p.Placeholders, len(ph)) {
			for i, v := range ph {
				assert.Equal(s.t, v, p.Placeholders[i])
			}
		}
	}
}
func (s sqla) CondSQL(c query.ConditionDef, sql string) {
	p, e := ConditionToSQL(c)
	if assert.NoError(s.t, e) {
		assert.Equal(s.t, sql, p.SQL)
	}
}
func (s sqla) CondPh(c query.ConditionDef, ph ...interface{}) {
	p, e := ConditionToSQL(c)
	if assert.NoError(s.t, e) {
		if assert.Len(s.t, p.Placeholders, len(ph)) {
			for i, v := range ph {
				assert.Equal(s.t, v, p.Placeholders[i])
			}
		}
	}
}

func TestRuleToSQL(t *testing.T) {
	assert := sqla{t}

	assert.RuleSQL(query.Rule{query.String("foo"), match.IsNull, nil}, "`foo` IS NULL")
	assert.RuleSQL(query.Rule{query.String("bar"), match.NotIsNull, nil}, "`bar` NOT IS NULL")
	assert.RuleSQL(query.Rule{query.String("foo"), match.Equals, query.String("bar")}, "`foo` = `bar`")
	assert.RuleSQL(query.Rule{query.String("foo"), match.NotEquals, query.String("bar")}, "`foo` <> `bar`")
	assert.RuleSQL(query.Rule{query.String("foo"), match.Equals, 5}, "`foo` = ?")
	assert.RuleSQL(query.Rule{query.String("foo"), match.NotEquals, "7"}, "`foo` <> ?")
	assert.RuleSQL(query.Rule{query.String("foo"), match.Gt, "7"}, "`foo` > ?")
	assert.RuleSQL(query.Rule{query.String("foo"), match.Gte, "7"}, "`foo` >= ?")
	assert.RuleSQL(query.Rule{query.String("foo"), match.Lt, "7"}, "`foo` < ?")
	assert.RuleSQL(query.Rule{query.String("foo"), match.Lte, "7"}, "`foo` <= ?")
	assert.RuleSQL(query.Rule{query.String("bar"), match.In, []interface{}{5, int64(6), 7}}, "`bar` IN (?,?,?)")
	assert.RuleSQL(query.Rule{query.String("bar"), match.NotIn, []interface{}{3, "false"}}, "`bar` NOT IN (?,?)")
	assert.RuleSQL(query.Rule{query.String("bar"), match.In, []interface{}{"true"}}, "`bar` IN (?)")

	assert.RulePh(query.Rule{query.String("foo"), match.Lte, "7"}, "7")
	assert.RulePh(query.Rule{query.String("foo"), match.Lte, 7}, 7)
	assert.RulePh(query.Rule{query.String("foo"), match.Lte, 0.1}, 0.1)
	assert.RulePh(query.Rule{query.String("bar"), match.In, []interface{}{5, int64(6), 7}}, 5, int64(6), 7)
	assert.RulePh(query.Rule{query.String("bar"), match.NotIn, []interface{}{3, "false"}}, 3, "false")
	assert.RulePh(query.Rule{query.String("bar"), match.In, []interface{}{"true"}}, "true")
}
