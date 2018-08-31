package sql

import "github.com/mono83/xray"

// Modifier is an interface for components, able to perform modification operations
// using statements
type Modifier interface {
	Invoker

	Modify(stmt Statement) (int64, error)
}

// ModifierX is a Modifier, but with xray support
type ModifierX interface {
	InvokerX

	ModifyX(ray xray.Ray, stmt Statement) (int64, error)
}
