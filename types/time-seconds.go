package types

import (
	"reflect"
	"time"
)

// TimeSeconds is wrapper over unix timestamp stored in database in seconds
type TimeSeconds struct {
	time.Time
}

// Scan is sql.Scanner interface implementation
func (t *TimeSeconds) Scan(src interface{}) error {
	switch src.(type) {
	case int64:
		t.Time = time.Unix(src.(int64), 0).UTC()
		return nil
	case uint64:
		t.Time = time.Unix(int64(src.(uint64)), 0).UTC()
		return nil
	case int32:
		t.Time = time.Unix(int64(src.(int32)), 0).UTC()
		return nil
	case uint32:
		t.Time = time.Unix(int64(src.(uint32)), 0).UTC()
		return nil
	case int:
		t.Time = time.Unix(int64(src.(int)), 0).UTC()
		return nil
	default:
		return ScanError{Target: reflect.TypeOf(TimeSeconds{}), Source: reflect.TypeOf(src)}
	}
}
