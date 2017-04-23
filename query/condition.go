package query

// Condition is simple ConditionDef implementation
type Condition struct {
	Type       ConditionType
	Rules      []RuleDef
	Conditions []ConditionDef
}

// GetType return condition relation type
func (c Condition) GetType() ConditionType { return c.Type }

// GetRules returns rules, used in condition
func (c Condition) GetRules() []RuleDef { return c.Rules }

// GetConditions return inner conditions list
func (c Condition) GetConditions() []ConditionDef { return c.Conditions }
