package query

// CommonFilter is plain Filter interface implementation
type CommonFilter struct {
	Type       Logic
	Rules      []Rule
	Conditions []Condition
	Sorting    []Sorting
	Offset     int
	Limit      int
}

// GetType return condition relation type
func (c CommonFilter) GetType() Logic { return c.Type }

// GetRules returns rules, used in condition
func (c CommonFilter) GetRules() []Rule { return c.Rules }

// GetConditions returns inner conditions list
func (c CommonFilter) GetConditions() []Condition { return c.Conditions }

// GetSorting returns sorting configuration
func (c CommonFilter) GetSorting() []Sorting { return c.Sorting }

// GetOffset return filtering offset
func (c CommonFilter) GetOffset() int { return c.Offset }

// GetLimit return filtering limit
func (c CommonFilter) GetLimit() int { return c.Limit }
