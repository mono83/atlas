package match

// List of supported rule operations
const (
	Unknown Type = 0

	IsNull            Type = 1
	NotIsNull         Type = 2
	Eq                Type = 3 // Alias
	Equals            Type = 3
	In                Type = 4
	NotIn             Type = 5
	Neq               Type = 6 // Alias
	NotEquals         Type = 6
	Gt                Type = 7 // Alias
	GreaterThan       Type = 7
	Lte               Type = 8 // Alias
	LowerThanEquals   Type = 8
	Gte               Type = 9 // Alias
	GreaterThanEquals Type = 9
	Lt                Type = 10 // Alias
	LowerThan         Type = 10
)

var lower = 1
var upper = 10

// All returns full list of rule operations, except Unknown
// Used in tests primarily
func All() []Type {
	all := []Type{}
	for i := lower; i <= upper; i++ {
		all = append(all, Type(i))
	}

	return all
}
