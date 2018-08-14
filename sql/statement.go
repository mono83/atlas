package sql

// Statement is an interface to SQL with placeholders
type Statement interface {
	GetSQL() string
	GetPlaceholders() []interface{}
}

// CommonStatement is Statement interface implementation
type CommonStatement struct {
	SQL          string
	Placeholders []interface{}
}

// GetSQL is Statement interface implementation
func (c CommonStatement) GetSQL() string { return c.SQL }

// GetPlaceholders is Statement interface implementation
func (c CommonStatement) GetPlaceholders() []interface{} { return c.Placeholders }

// StringStatement is implementation of statement interface without placeholders
type StringStatement string

// GetSQL is Statement interface implementation
func (s StringStatement) GetSQL() string { return string(s) }

// GetPlaceholders is Statement interface implementation
func (s StringStatement) GetPlaceholders() []interface{} { return nil }
