package rules

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
)

// IsNull returns rule with IS NULL matcher for provided field
func IsNull(field interface{}) query.Rule {
	return query.CommonRule{L: field, R: nil, Type: match.IsNull}
}

// IsNotNull returns rule with IS NOT NULL matcher for provided field
func IsNotNull(field interface{}) query.Rule {
	return query.CommonRule{L: field, R: nil, Type: match.NotIsNull}
}

// Eq returns rule built with EQUALS matcher
func Eq(left, right interface{}) query.Rule {
	return query.CommonRule{L: left, R: right, Type: match.Equals}
}
