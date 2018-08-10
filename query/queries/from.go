package queries

import (
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
)

type Builder struct {
	schema    query.Named
	condition *query.Condition
	order     []query.OrderDef
	offset    int64
	limit     int
}

func From(schema string, alias ...string) *Builder {
	b := new(Builder)
	if len(alias) == 1 {
		b.schema = query.AliasedName{Name: schema, Alias: alias[0]}
	} else {
		b.schema = query.String(schema)
	}

	return b
}

func (b *Builder) FindById64(ID int64) *Builder {
	b.condition = &query.Condition{
		Rules: []query.RuleDef{query.Rule{Type: match.Eq, L: query.String("id"), R: ID}},
	}
	b.limit = 1
	return b
}

func (b *Builder) Or() *Builder {
	if b.condition != nil {
		b.condition.Type = query.Or
	}

	return b
}

func (b *Builder) WhereSimple(left interface{}, op match.Type, right interface{}) *Builder {
	if b.condition == nil {
		b.condition = &query.Condition{Type: query.And}
	}

	b.condition.Rules = append(b.condition.Rules, query.Rule{
		L:    left,
		R:    right,
		Type: op,
	})

	return b
}

func (b *Builder) Select(c ...interface{}) query.SelectDef {
	columns := []query.Named{}
	for _, v := range c {
		if x, ok := v.(query.Named); ok {
			columns = append(columns, x)
		} else if x, ok := v.(string); ok {
			columns = append(columns, query.String(x))
		}
	}

	sel := query.Select{
		Schema:  b.schema,
		Columns: columns,
		Order:   b.order,
		Limit:   b.limit,
		Offset:  b.offset,
	}

	if b.condition != nil {
		sel.Condition = b.condition
	}

	return sel
}
