package compact

import (
	"testing"

	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"github.com/mono83/atlas/query/rules"
	"github.com/stretchr/testify/assert"
)

var ruleDataProvider = []struct {
	JSON string
	rule query.Rule
}{
	{`["is_null","foo"]`, rule{t: match.IsNull, l: "foo"}},
	{`["is_not_null","foo"]`, rule{t: match.NotIsNull, l: "foo"}},
	{`["equal","a","b"]`, rule{t: match.Equals, l: "a", r: "b"}},
	{`["equal","b","c"]`, rules.Eq("b", "c")},
	{`["not_equal","foo","bar"]`, rule{t: match.Neq, l: "foo", r: "bar"}},
	{`["in","zzz",[1,2,3]]`, rule{t: match.In, l: "zzz", r: []interface{}{1., 2., 3.}}},
	{`["not_in","yyy",[2]]`, rule{t: match.NotIn, l: "yyy", r: []interface{}{2.}}},
	{`["greater_than","foo",5]`, rule{t: match.Gt, l: "foo", r: float64(5)}},
	{`["greater_than_equals","foo",5]`, rule{t: match.Gte, l: "foo", r: float64(5)}},
	{`["lesser_than","bat",-9]`, rule{t: match.Lt, l: "bat", r: float64(-9)}},
	{`["lesser_than_equals","bat",-9]`, rule{t: match.Lte, l: "bat", r: float64(-9)}},
}

func TestRuleJson(t *testing.T) {
	for _, row := range ruleDataProvider {
		t.Run("To "+row.JSON, func(t *testing.T) {
			bts, err := mapRule(row.rule).MarshalJSON()
			if assert.NoError(t, err) {
				assert.Equal(t, row.JSON, string(bts))
			}
		})
		t.Run("From "+row.JSON, func(t *testing.T) {
			var rule rule
			if assert.NoError(t, rule.UnmarshalJSON([]byte(row.JSON))) {
				assert.Equal(t, mapRule(row.rule), rule)
			}
		})
	}
}
