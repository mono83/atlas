package query

// MapCondition applies provided mapping func to each rule
// inside condition and returns new condition with remapped rules
func MapCondition(c Condition, mapFunc func(Rule) Rule) Condition {
	if c == nil || mapFunc == nil {
		return c
	}

	response := CommonCondition{Type: c.GetType()}

	if l := len(c.GetRules()); l > 0 {
		response.Rules = make([]Rule, l)
	}
	if l := len(c.GetConditions()); l > 0 {
		response.Conditions = make([]Condition, l)
	}

	// Mapping rules
	for i, ir := range c.GetRules() {
		response.Rules[i] = mapFunc(ir)
	}

	// Mapping inner conditions
	for i, ic := range c.GetConditions() {
		response.Conditions[i] = MapCondition(ic, mapFunc)
	}

	return response
}

// MapFilter applies provided mapping func to each rule
// inside condition within filder and returns new filter
// with remapped rules
func MapFilter(f Filter, mapFunc func(Rule) Rule) Filter {
	c := MapCondition(f, mapFunc)

	return CommonFilter{
		Type:       c.GetType(),
		Conditions: c.GetConditions(),
		Rules:      c.GetRules(),
		Sorting:    f.GetSorting(),
		Limit:      f.GetLimit(),
		Offset:     f.GetOffset(),
	}
}
