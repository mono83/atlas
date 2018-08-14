package queries

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/rules"
)

// FindByID64Simple returns query for matching entries by IDs
func FindByID64Simple(table string, id ...int64) query.Query {
	return id64SimpleQuery{
		table: table,
		ids:   id,
	}
}

type id64SimpleQuery struct {
	table string
	ids   []int64
}

func (id64SimpleQuery) GetType() query.Logic             { return query.And }
func (id64SimpleQuery) GetColumns() []query.Named        { return nil }
func (id64SimpleQuery) GetConditions() []query.Condition { return nil }
func (id64SimpleQuery) GetSorting() []query.Sorting      { return nil }
func (id64SimpleQuery) GetLimit() int                    { return 0 }
func (id64SimpleQuery) GetOffset() int                   { return 0 }
func (i id64SimpleQuery) GetSchema() query.Named         { return query.String(i.table) }
func (i id64SimpleQuery) GetRules() []query.Rule         { return []query.Rule{rules.MatchID64(i.ids...)} }
