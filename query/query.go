package query

// Query describes common select query
type Query interface {
	GetSchema() Named
	GetColumns() []Named
	Filter
}
