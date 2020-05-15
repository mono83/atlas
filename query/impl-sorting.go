package query

// CommonSorting is simple Sorting implementation
type CommonSorting struct {
	Column Named
	Type   SortOrder
}

// GetType returns ordering type (ASC or DESC)
func (c CommonSorting) GetType() SortOrder { return c.Type }

// GetName returns column name, used in ordering
func (c CommonSorting) GetName() string { return c.Column.GetName() }

// SimpleAsc is ASC-only Sorting implementation
type SimpleAsc string

// GetName returns column name, used in ordering
func (s SimpleAsc) GetName() string { return string(s) }

// GetType always returns ASC
func (s SimpleAsc) GetType() SortOrder { return Asc }

// GetColumn returns self, because structure implements Named
func (s SimpleAsc) GetColumn() Named { return s }

// SimpleDesc is DESC-only Sorting implementation
type SimpleDesc string

// GetName returns column name, used in ordering
func (s SimpleDesc) GetName() string { return string(s) }

// GetType always returns DESC
func (s SimpleDesc) GetType() SortOrder { return Desc }

// GetColumn returns self, because structure implements Named
func (s SimpleDesc) GetColumn() Named { return s }
