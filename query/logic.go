package query

// Logic describes inner relations inside conditions
// Can be AND or OR
type Logic byte

// List of defined condition relations
const (
	None Logic = 0
	And  Logic = 1
	Or   Logic = 2
)
