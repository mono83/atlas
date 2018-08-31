package sql

import "github.com/mono83/xray"

// Invoker represents components that are able to invoke statements
// and write results into provided target
type Invoker interface {
	Invoke(stmt Statement, target interface{}) error
}

// InvokerX is an Invoker, but with xray support
type InvokerX interface {
	InvokeX(ray xray.Ray, stmt Statement, target interface{}) error
}
