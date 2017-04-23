package query

// String is a string, that implements Named interface
// Can be used for columns and schemas
type String string

// GetName returns name of structure
func (n String) GetName() string { return string(n) }
func (n String) String() string  { return string(n) }

// AliasedName is structure, that contains name with its alias
// Can be used for columns and schemas
type AliasedName struct {
	Name, Alias string
}

// GetAlias returns alias name of structure
func (a AliasedName) GetAlias() string { return a.Alias }

// GetName returns name of structure
func (a AliasedName) GetName() string { return a.Name }
func (a AliasedName) String() string  { return a.Name }
