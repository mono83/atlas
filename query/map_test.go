package query

import (
	"fmt"
	"testing"

	"github.com/mono83/atlas/query/match"
	"github.com/stretchr/testify/assert"
)

var mapConditionsDataProvider = []struct {
	Expected Condition
	Provided Condition
}{
	{CommonCondition{Type: And}, CommonCondition{Type: And}},
	{CommonCondition{Type: Or}, CommonCondition{Type: Or}},
	{
		CommonCondition{Type: And, Rules: []Rule{CommonRule{L: "foo", R: "bar", Type: match.NotEquals}}},
		CommonCondition{Type: And, Rules: []Rule{CommonRule{L: "foo", R: "bar", Type: match.Equals}}},
	},
	{
		CommonCondition{
			Type:  Or,
			Rules: []Rule{CommonRule{L: 1, R: 2, Type: match.NotEquals}},
			Conditions: []Condition{
				CommonCondition{Type: And, Rules: []Rule{CommonRule{L: 3, R: 4, Type: match.Gte}}},
				CommonCondition{Type: And, Rules: []Rule{CommonRule{L: 5, R: 6, Type: match.NotEquals}}},
			},
		},
		CommonCondition{
			Type:  Or,
			Rules: []Rule{CommonRule{L: 1, R: 2, Type: match.Equals}},
			Conditions: []Condition{
				CommonCondition{Type: And, Rules: []Rule{CommonRule{L: 3, R: 4, Type: match.Gte}}},
				CommonCondition{Type: And, Rules: []Rule{CommonRule{L: 5, R: 6, Type: match.Equals}}},
			},
		},
	},
}

func TestMapCondition(t *testing.T) {

	mapFunc := func(r Rule) Rule {
		if r.GetType() == match.Equals {
			return CommonRule{L: r.GetLeft(), R: r.GetRight(), Type: match.NotEquals}
		}
		return r
	}

	for _, d := range mapConditionsDataProvider {
		t.Run(fmt.Sprintf("%v", d.Provided), func(t *testing.T) {
			assert.Equal(t, d.Expected, MapCondition(d.Provided, mapFunc))
		})
	}
}

func TestMapFilter(t *testing.T) {

	mapFunc := func(r Rule) Rule {
		if r.GetType() == match.Equals {
			return CommonRule{L: r.GetLeft(), R: r.GetRight(), Type: match.NotEquals}
		}
		return r
	}

	for _, d := range mapConditionsDataProvider {
		t.Run(fmt.Sprintf("%v", d.Provided), func(t *testing.T) {
			provided := CommonFilter{
				Type:       d.Provided.GetType(),
				Rules:      d.Provided.GetRules(),
				Conditions: d.Provided.GetConditions(),
				Limit:      88,
				Offset:     12,
				Sorting:    []Sorting{SimpleAsc("foo")},
			}
			expected := CommonFilter{
				Type:       d.Expected.GetType(),
				Rules:      d.Expected.GetRules(),
				Conditions: d.Expected.GetConditions(),
				Limit:      88,
				Offset:     12,
				Sorting:    []Sorting{SimpleAsc("foo")},
			}

			assert.Equal(t, expected, MapFilter(provided, mapFunc))
		})
	}
}
