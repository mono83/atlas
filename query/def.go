package query

import "github.com/mono83/atlas/query/match"

// Named is interface for named entries (columns, schemas)
type Named interface {
	GetName() string
}

// Aliased is interface for entries (columns, schemas) that can be aliased.
type Aliased interface {
	GetAlias() string
}

// OrderDef contains order definition
type OrderDef interface {
	GetType() OrderType
	GetColumn() Named
}

// RuleDef contains rule definition
type RuleDef interface {
	GetLeft() interface{}
	GetRight() interface{}
	GetType() match.Type
}

// ConditionDef represents complex condition, that may include rules and other conditions
type ConditionDef interface {
	GetType() ConditionType
	GetRules() []RuleDef
	GetConditions() []ConditionDef
}

// SelectDef contains query definition for data acquiring
type SelectDef interface {
	GetSchema() Named
	GetColumns() []Named
	GetCondition() ConditionDef
	GetOrder() []OrderDef
	GetOffsetLimit() (offset int64, limit int)
}
