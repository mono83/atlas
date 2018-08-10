package query

import "errors"

// Select is Query implementation to query database
type Select struct {
	Schema    Named
	Columns   []Named
	Condition Condition
	Order     []Sorting
	Limit     int
	Offset    int64
}

// GetSchema returns schema name, used in query
func (q Select) GetSchema() Named { return q.Schema }

// GetCondition return condition, used in query
func (q Select) GetCondition() Condition { return q.Condition }

// GetColumns returns columns, used in query
func (q Select) GetColumns() []Named { return q.Columns }

// GetOffsetLimit returns offset and limit, used in query
func (q Select) GetOffsetLimit() (int64, int) { return q.Offset, q.Limit }

// GetOrder return ordering, used in query
func (q Select) GetOrder() []Sorting { return q.Order }

// Apply applies query on invoker
func (q Select) Apply(dao ReadOnlyDAO, target interface{}) error {
	if dao == nil {
		return errors.New("Empty DAO")
	}

	return dao.Select(q, target)
}
