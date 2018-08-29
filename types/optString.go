package types

// OptString is optional string
type OptString struct {
	value *string
}

// OrElse returns value or provided string if value is nil
func (o OptString) OrElse(s string) string {
	if o.value == nil {
		return s
	}

	return *o.value
}

// IsPresent returns true if OptString contains value that is not nil
func (o OptString) IsPresent() bool {
	return o.value != nil
}

// IsNotEmpty return true if value is present and is not empty
func (o OptString) IsNotEmpty() bool {
	return o.value != nil && len(*o.value) > 0
}

// Filter method applies filter function on OptString contents
func (o OptString) Filter(f func(string) bool) OptString {
	if f != nil && o.value != nil && !f(*o.value) {
		return OptString{value: nil}
	}
	return o
}

// IfPresent invokes consumer function if OptString contains value that is not nil
func (o OptString) IfPresent(f func(string)) {
	if f != nil && o.value != nil {
		f(*o.value)
	}
}

// IfNotEmpty invokes consumer function if OptString contains value that is not nil
// and not empty
func (o OptString) IfNotEmpty(f func(string)) {
	if f != nil && o.IsNotEmpty() {
		f(*o.value)
	}
}

// Scan is sql.Scanner interface implementation
func (o *OptString) Scan(src interface{}) error {
	if src == nil {
		*o = OptString{value: nil}
		return nil
	}

	switch src.(type) {
	case []byte:
		str := string(src.([]byte))
		*o = OptString{value: &str}
		return nil
	case string:
		str := src.(string)
		*o = OptString{value: &str}
		return nil
	case OptString:
		*o = src.(OptString)
		return nil
	case *OptString:
		p := src.(*OptString)
		*o = *p
		return nil
	default:
		return newScanError(OptString{}, src)
	}
}
