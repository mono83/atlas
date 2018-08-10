package mysql

import "github.com/mono83/atlas/query/match"

// UnsupportedOperation is error, returned when unsupported operation requested
type UnsupportedOperation match.Type

func (u UnsupportedOperation) Error() string {
	return "unsupported operation " + match.Type(u).String()
}

// LeftIsNotColumn is error, returned when no column definition found on left side of rule
type LeftIsNotColumn struct {
	Real interface{}
}

func (LeftIsNotColumn) Error() string {
	return "no column definition on left side of rule"
}
