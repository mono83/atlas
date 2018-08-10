package query

// SortOrder specifies collection ordering
type SortOrder byte

// List of defined order types
const (
	UnknownSort SortOrder = 0
	Asc         SortOrder = 1
	Desc        SortOrder = 2
)
