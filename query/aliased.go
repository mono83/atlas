package query

// Aliased is interface for entries (columns, schemas) that can be aliased.
type Aliased interface {
	GetAlias() string
}
