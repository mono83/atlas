package queries

import (
	"github.com/mono83/atlas/query"
)

// SelectAll returns query definitions, used to obtain all entries from
// required schema without any conditions.
// Optionally, columns may be specified - as plain string or query.Named
func SelectAll(schema string, columns ...interface{}) query.SelectDef {
	cs := []query.Named{}
	for _, c := range columns {
		if n, ok := c.(query.Named); ok {
			cs = append(cs, n)
		} else if s, ok := c.(string); ok {
			cs = append(cs, query.String(s))
		}
	}

	return query.Select{
		Schema:  query.String(schema),
		Columns: cs,
	}
}
