package query

// Condition represents complex condition, that may include rules and other conditions
type Condition interface {
	GetType() Logic
	GetRules() []Rule
	GetConditions() []Condition
}

// ConditionAddAndRule places new rule on top level with AND logic
func ConditionAddAndRule(source Condition, rule Rule) Condition {
	if source.GetType() == And {
		for _, r := range source.GetRules() {
			if RulesAreEqual(r, rule) {
				// Rule already present in condition on top level
				return source
			}
		}
	}

	resp := CommonCondition{
		Type:       And,
		Rules:      []Rule{},
		Conditions: []Condition{},
	}

	if source.GetType() == And {
		resp.Rules = append(resp.Rules, source.GetRules()...)
		resp.Conditions = append(resp.Conditions, source.GetConditions()...)
	} else {
		resp.Conditions = append(resp.Conditions, source)
	}

	resp.Rules = append(resp.Rules, rule)
	return resp
}
