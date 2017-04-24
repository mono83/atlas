package queries

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
)

// FindById64 builds query, used to find entry by ID in int64 format
func FindById64(schema string, ID int64) query.SelectDef {
	return query.Select{
		Schema: query.String(schema),
		Condition: query.Condition{
			Rules: []query.RuleDef{query.Rule{Type: match.Eq, L: query.String("id"), R: ID}},
		},
		Limit: 1,
	}
}
