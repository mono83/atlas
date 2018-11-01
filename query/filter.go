package query

// Filter extends condition with limits and sorting
type Filter interface {
	Condition

	GetSorting() []Sorting
	GetLimit() int
	GetOffset() int
}

// FilterAddAndRule places new rule on top level with AND logic
func FilterAddAndRule(source Filter, rule Rule) Filter {
	if source.GetType() == And {
		for _, r := range source.GetRules() {
			if RulesAreEqual(r, rule) {
				// Rule already present in condition on top level
				return source
			}
		}
	}

	resp := CommonFilter{
		Type:       And,
		Rules:      []Rule{},
		Conditions: []Condition{},
		Sorting:    []Sorting{},
		Offset:     source.GetOffset(),
		Limit:      source.GetLimit(),
	}

	resp.Sorting = append(resp.Sorting, source.GetSorting()...)

	if source.GetType() == And {
		resp.Rules = append(resp.Rules, source.GetRules()...)
		resp.Conditions = append(resp.Conditions, source.GetConditions()...)
	} else {
		resp.Conditions = append(resp.Conditions, source)
	}

	resp.Rules = append(resp.Rules, rule)
	return resp
}
