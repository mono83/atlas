package query

// Named is interface for named entries (columns, schemas)
type Named interface {
	GetName() string
}
