package query

import "github.com/mono83/atlas/query/match"

// Rule is simple RuleDef implementation
type Rule struct {
	L    interface{}
	Type match.Type
	R    interface{}
}

// GetLeft returns left part of rule condition
func (r Rule) GetLeft() interface{} { return r.L }

// GetRight returns right part of rule condition
func (r Rule) GetRight() interface{} { return r.R }

// GetType return operation, used in Rule
func (r Rule) GetType() match.Type { return r.Type }
