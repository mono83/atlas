package types

import "reflect"

func newScanError(target, provided interface{}) ScanError {
	se := ScanError{}
	if target != nil {
		se.Target = reflect.TypeOf(target)
	}
	if provided != nil {
		se.Source = reflect.TypeOf(provided)
	}

	return se
}

// ScanError is error, emitted on scan error (when data from MySQL writes into structs)
type ScanError struct {
	Target, Source reflect.Type
}

func (s ScanError) Error() string {
	if s.Source == nil && s.Target == nil {
		return "unable to sql.Scan"
	} else if s.Source == nil {
		return "unable to sql.Scan into " + s.Target.Name() + " from <nil>"
	}

	return "unable to sql.Scan into " + s.Target.Name() + " from " + s.Source.Name()
}
