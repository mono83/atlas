package query

// Condition represents complex condition, that may include rules and other conditions
type Condition interface {
	GetType() Logic
	GetRules() []Rule
	GetConditions() []Condition
}
