package query

import "github.com/mono83/atlas/query/match"

// Rule contains rule definition
type Rule interface {
	GetLeft() interface{}
	GetRight() interface{}
	GetType() match.Type
}

// RulesAreEqual returns true if provided rules contains same logic
func RulesAreEqual(a, b Rule) bool {
	return a.GetType() == b.GetType() &&
		a.GetLeft() == b.GetLeft() &&
		a.GetRight() == b.GetRight()
}
