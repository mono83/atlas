package query

// Sorting contains sort order definition
type Sorting interface {
	Named

	GetType() SortOrder
}
