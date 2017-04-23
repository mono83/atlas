package query

// Order is simple OrderDef implementation
type Order struct {
	Column Named
	Type   OrderType
}

// GetType returns ordering type (ASC or DESC)
func (o Order) GetType() OrderType { return o.Type }

// GetColumn returns column name, used in ordering
func (o Order) GetColumn() Named { return o.Column }

// SimpleAsc is ASC-only OrderDef implementation
type SimpleAsc string

// GetName returns column name, used in ordering
func (s SimpleAsc) GetName() string { return string(s) }

// GetType always returns ASC
func (s SimpleAsc) GetType() OrderType { return Asc }

// GetColumn returns self, because structure implements Named
func (s SimpleAsc) GetColumn() Named { return s }

// SimpleDesc is DESC-only OrderDef implementation
type SimpleDesc string

// GetName returns column name, used in ordering
func (s SimpleDesc) GetName() string { return string(s) }

// GetType always returns DESC
func (s SimpleDesc) GetType() OrderType { return Desc }

// GetColumn returns self, because structure implements Named
func (s SimpleDesc) GetColumn() Named { return s }
