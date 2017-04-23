package query

// OrderType specifies collection ordering
type OrderType byte

// List of defined order types
const (
	Asc  OrderType = 1
	Desc OrderType = 2
)

// ConditionType describes inner relations inside conditions
// Can be AND or OR
type ConditionType byte

// List of defined condition relations
const (
	And ConditionType = 1
	Or  ConditionType = 2
)
