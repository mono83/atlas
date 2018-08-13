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

// MatchID64 returns rule for matching IDs
func MatchID64(id ...int64) query.Rule {
	switch len(id) {
	case 0:
		return False{}
	case 1:
		return Eq("id", id[0])
	default:
		return query.CommonRule{L: "id", R: id, Type: match.In}
	}
}
