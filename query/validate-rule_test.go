package query

import (
	"testing"

	"github.com/mono83/atlas/query/match"
)

var ruleValidationDataProvider = []struct {
	Error       string
	Type        match.Type
	Left, Right interface{}
}{
	// Unknown
	{"unknown rule match type", match.Unknown, nil, nil},

	// Nulls
	{"null check rules must not have right part", match.IsNull, 1, 2},
	{"null check rules must not have right part", match.NotIsNull, 1, 2},
	{"", match.IsNull, 1, nil},
	{"", match.NotIsNull, "foo", nil},

	// Equality
	{"nil found in equality check", match.Equals, nil, nil},
	{"", match.Equals, 1, 2},

	// >
	{"nil found in numeric rule on the right side", match.GreaterThan, 1, nil},
	{"nil found in numeric rule on the left side", match.GreaterThan, nil, 2},
	{"only numbers and nameds should be in numeric rule on right side", match.GreaterThan, 1, "foo"},
	{"only numbers and nameds should be in numeric rule on left side", match.GreaterThan, "foo", 2},
	{"", match.GreaterThan, 1, String("foo")},
	{"", match.GreaterThan, 1, 2},

	// <
	{"nil found in numeric rule on the right side", match.LowerThan, 1, nil},
	{"nil found in numeric rule on the left side", match.LowerThan, nil, 2},
	{"only numbers and nameds should be in numeric rule on right side", match.LowerThan, 1, "foo"},
	{"only numbers and nameds should be in numeric rule on left side", match.LowerThan, "foo", 2},
	{"", match.LowerThan, 1, String("foo")},
	{"", match.LowerThan, 1, 2},

	// >=
	{"nil found in numeric rule on the right side", match.GreaterThanEquals, 1, nil},
	{"nil found in numeric rule on the left side", match.GreaterThanEquals, nil, 2},
	{"only numbers and nameds should be in numeric rule on right side", match.GreaterThanEquals, 1, "foo"},
	{"only numbers and nameds should be in numeric rule on left side", match.GreaterThanEquals, "foo", 2},
	{"", match.GreaterThanEquals, 1, String("foo")},
	{"", match.GreaterThanEquals, 1, 2},

	// <
	{"nil found in numeric rule on the right side", match.LowerThanEquals, 1, nil},
	{"nil found in numeric rule on the left side", match.LowerThanEquals, nil, 2},
	{"only numbers and nameds should be in numeric rule on right side", match.LowerThanEquals, 1, "foo"},
	{"only numbers and nameds should be in numeric rule on left side", match.LowerThanEquals, "foo", 2},
	{"", match.LowerThanEquals, 1, String("foo")},
	{"", match.LowerThanEquals, 1, 2},

	// IN
	{"nil found in IN/NOT IN rule on the right side", match.In, "foo", nil},
	{"nil found in IN/NOT IN rule on the right side", match.NotIn, "foo", nil},
	{"not a slice found in IN/NOT IN rule on the right side", match.In, "foo", "foo"},
	{"not a slice found in IN/NOT IN rule on the right side", match.NotIn, "foo", "foo"},
	{"empty slice found in IN/NOT IN rule on the right side", match.In, "foo", []int{}},
	{"empty slice found in IN/NOT IN rule on the right side", match.NotIn, "foo", []int{}},
	{"", match.In, "foo", []int{10}},
	{"", match.NotIn, "foo", []int{10}},
}

func TestRuleValidate(t *testing.T) {
	for _, d := range ruleValidationDataProvider {
		rule := CommonRule{L: d.Left, R: d.Right, Type: d.Type}
		t.Run(rule.String(), func(t *testing.T) {
			err := ValidateRule(rule)
			if err == nil {
				if len(d.Error) != 0 {
					t.Errorf(`expected error "%s" but got nothing`, d.Error)
				}
			} else {
				if len(d.Error) == 0 {
					t.Errorf(`expected no error, but got "%s"`, err.Error())
				}
			}
		})
	}
}

func TestRuleValidateNil(t *testing.T) {
	if err := ValidateRule(nil); err == nil {
		t.Errorf("expected error on nil rule")
	}
}
