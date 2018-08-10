package query

import "github.com/mono83/atlas/query/match"

// CommonRule is simple Rule implementation
type CommonRule struct {
	L    interface{}
	Type match.Type
	R    interface{}
}

// GetLeft returns left part of rule condition
func (c CommonRule) GetLeft() interface{} { return c.L }

// GetRight returns right part of rule condition
func (c CommonRule) GetRight() interface{} { return c.R }

// GetType return operation, used in CommonRule
func (c CommonRule) GetType() match.Type { return c.Type }
