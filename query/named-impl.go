package query

// String is a string, that implements Named interface
// Can be used for columns and schemas
type String string

// GetName returns name of structure
func (n String) GetName() string { return string(n) }
func (n String) String() string  { return string(n) }
