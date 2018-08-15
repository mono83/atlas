package mysql

import (
	"github.com/mono83/atlas/query/match"
	"reflect"
)

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

// ScanError is error, emitted on scan error (when data from MySQL writes into structs)
type ScanError struct {
	Target, Source reflect.Type
}

func (s ScanError) Error() string {
	if s.Source == nil && s.Target == nil {
		return "unable to Scan"
	} else if s.Source == nil {
		return "unable to Scan into " + s.Target.Name() + " from <nil>"
	}

	return "unable to Scan into " + s.Target.Name() + " from " + s.Source.Name()
}
