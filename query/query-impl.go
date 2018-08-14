package query

// CommonSelect is Query implementation to query database
type CommonSelect struct {
	Schema  Named
	Columns []Named
	Filter
}

// GetSchema returns schema name, used in query
func (c CommonSelect) GetSchema() Named { return c.Schema }

// GetColumns returns columns, used in query
func (c CommonSelect) GetColumns() []Named { return c.Columns }
