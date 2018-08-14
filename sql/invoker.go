package sql

// Invoker represents components that are able to invoke statements
// and write results into provided target
type Invoker interface {
	Invoke(stmt Statement, target interface{}) error
}
