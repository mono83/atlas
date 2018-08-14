package sql

// Modifier is an interface for components, able to perform modification operations
// using statements
type Modifier interface {
	Invoker

	Modify(stmt Statement) (int64, error)
}
