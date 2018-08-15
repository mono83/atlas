package mysql

import (
	"reflect"
	"strconv"
	"time"
)

// UnixTimestamp is wrapper over unix timestamp stored in database in seconds
type UnixTimestamp struct {
	time.Time
}

// Scan is sql.Scanner interface implementation
func (u *UnixTimestamp) Scan(src interface{}) error {
	switch src.(type) {
	case []byte:
		return u.Scan(string(src.([]byte)))
	case int64:
		u.Time = time.Unix(src.(int64), 0).UTC()
		return nil
	case uint64:
		u.Time = time.Unix(int64(src.(uint64)), 0).UTC()
		return nil
	case int:
		u.Time = time.Unix(int64(src.(int)), 0).UTC()
		return nil
	case string:
		ui, err := strconv.ParseInt(src.(string), 10, 64)
		if err == nil {
			return u.Scan(ui)
		}
		return err
	default:
		return ScanError{Target: reflect.TypeOf(UnixTimestamp{}), Source: reflect.TypeOf(src)}
	}
}
