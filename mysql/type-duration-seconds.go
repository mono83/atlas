package mysql

import (
	"reflect"
	"strconv"
	"time"
)

// DurationSeconds is wrapper over time.Duration
type DurationSeconds struct {
	time.Duration
}

// Scan is sql.Scanner interface implementation
func (d *DurationSeconds) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		return d.Scan(string(src.([]byte)))
	case int64:
		d.Duration = time.Second * time.Duration(src.(int64))
		return nil
	case uint64:
		d.Duration = time.Second * time.Duration(int64(src.(uint64)))
		return nil
	case int:
		d.Duration = time.Second * time.Duration(int64(src.(int)))
		return nil
	case string:
		ui, err := strconv.ParseInt(src.(string), 10, 64)
		if err == nil {
			return d.Scan(ui)
		}
		return err
	default:
		return ScanError{Target: reflect.TypeOf(DurationSeconds{}), Source: reflect.TypeOf(src)}
	}
}
