package query

// SelectDef contains query definition for data acquiring
type SelectDef interface {
	GetSchema() Named
	GetColumns() []Named
	GetCondition() Condition
	GetOrder() []Sorting
	GetOffsetLimit() (offset int64, limit int)

	Apply(dao ReadOnlyDAO, target interface{}) error
}

// ReadOnlyDAO describes components, used to invoke select statements
type ReadOnlyDAO interface {
	Select(def SelectDef, target interface{}) error
}
