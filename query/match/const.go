package match

// List of supported rule operations
const (
	Unknown Type = 0

	IsNull            Type = 1
	NotIsNull         Type = 2
	Eq                Type = 3 // Alias
	Equals            Type = 3
	Neq               Type = 4 // Alias
	NotEquals         Type = 4
	Gt                Type = 5 // Alias
	GreaterThan       Type = 5
	Lte               Type = 6 // Alias
	LowerThanEquals   Type = 6
	Gte               Type = 7 // Alias
	GreaterThanEquals Type = 7
	Lt                Type = 8 // Alias
	LowerThan         Type = 8
)

var lower = 1
var upper = 8

// All returns full list of rule operations, except Unknown
// Used in tests primarily
func All() []Type {
	all := []Type{}
	for i := lower; i <= upper; i++ {
		all = append(all, Type(i))
	}

	return all
}
