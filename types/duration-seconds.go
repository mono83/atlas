package types

import (
	"reflect"
	"time"
)

// DurationSeconds is wrapper over time.Duration
type DurationSeconds struct {
	time.Duration
}

// Scan is sql.Scanner interface implementation
func (d *DurationSeconds) Scan(src interface{}) error {
	switch src.(type) {
	case int64:
		d.Duration = time.Second * time.Duration(src.(int64))
		return nil
	case int32:
		d.Duration = time.Second * time.Duration(src.(int32))
		return nil
	case uint64:
		d.Duration = time.Second * time.Duration(int64(src.(uint64)))
		return nil
	case uint32:
		d.Duration = time.Second * time.Duration(int64(src.(uint32)))
		return nil
	case int:
		d.Duration = time.Second * time.Duration(int64(src.(int)))
		return nil
	default:
		return ScanError{Target: reflect.TypeOf(DurationSeconds{}), Source: reflect.TypeOf(src)}
	}
}
