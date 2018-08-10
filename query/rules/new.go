package rules

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
)

// New builds new rule
func New(left interface{}, op match.Type, right interface{}) query.Rule {
	return query.CommonRule{L: left, R: right, Type: op}
}
