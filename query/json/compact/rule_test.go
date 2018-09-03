package compact

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"github.com/mono83/atlas/query/rules"
	"github.com/stretchr/testify/assert"
	"testing"
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
