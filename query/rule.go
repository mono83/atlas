package query

import "github.com/mono83/atlas/query/match"

// Rule contains rule definition
type Rule interface {
	GetLeft() interface{}
	GetRight() interface{}
	GetType() match.Type
}
