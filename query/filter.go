package query

// Filter extends condition with limits and sorting
type Filter interface {
	Condition

	GetSorting() []Sorting
	GetLimit() int
	GetOffset() int
}
