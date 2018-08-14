package conditions

import "github.com/mono83/atlas/query"

// New builds and returns new condition
func New(logic query.Logic, rules []query.Rule, conditions []query.Condition) query.Condition {
	return query.CommonCondition{
		Type:       logic,
		Rules:      rules,
		Conditions: conditions,
	}
}

// ForAllRules returns conditions built for all rules with logic AND
func ForAllRules(rules ...query.Rule) query.Condition {
	return query.CommonCondition{
		Type:  query.And,
		Rules: rules,
	}
}

// ForAnyRule returns conditions built for all rules with logic OR
func ForAnyRule(rules ...query.Rule) query.Condition {
	return query.CommonCondition{
		Type:  query.Or,
		Rules: rules,
	}
}
