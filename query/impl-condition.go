package query

// CommonCondition is simple Condition implementation
type CommonCondition struct {
	Type       Logic
	Rules      []Rule
	Conditions []Condition
}

// GetType return condition relation type
func (c CommonCondition) GetType() Logic { return c.Type }

// GetRules returns rules, used in condition
func (c CommonCondition) GetRules() []Rule { return c.Rules }

// GetConditions return inner conditions list
func (c CommonCondition) GetConditions() []Condition { return c.Conditions }
