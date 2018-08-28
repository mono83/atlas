package match

import "strconv"

// List of supported rule operations
const (
	Unknown Type = 0

	IsNull    Type = 1
	NotIsNull Type = 2

	Eq        Type = 3 // Alias
	Equals    Type = 3
	Neq       Type = 4 // Alias
	NotEquals Type = 4

	In    Type = 5
	NotIn Type = 6

	Gt               Type = 7 // Alias
	GreaterThan      Type = 7
	Lte              Type = 8 // Alias
	LesserThanEquals Type = 8 // Alias
	LowerThanEquals  Type = 8

	Gte               Type = 9 // Alias
	GreaterThanEquals Type = 9
	Lt                Type = 10 // Alias
	LesserThan        Type = 10 // Alias
	LowerThan         Type = 10
)

func (t Type) String() string {
	switch t {
	case IsNull:
		return "IS NULL"
	case NotIsNull:
		return "NOT IS NULL"
	case Equals:
		return "EQUALS"
	case NotEquals:
		return "NOT EQUALS"
	case In:
		return "IN"
	case NotIn:
		return "NOT IN"
	case GreaterThan:
		return "GREATER THAN"
	case GreaterThanEquals:
		return "GREATER THAN EQUALS"
	case LowerThan:
		return "LOWER THAN"
	case LowerThanEquals:
		return "LOWER THAN EQUALS"
	case Unknown:
		return "UNKNOWN"
	default:
		if t.IsCustom() {
			return "CUSTOM #" + strconv.Itoa(int(t))
		}
		return "UNSUPPORTED #" + strconv.Itoa(int(t))
	}
}

// Theese constants describes top and bottom boundaries for registered
// matcher type constants
var (
	lower byte = 1
	upper byte = 10
)

// All returns full list of rule operations, except Unknown
// Used in tests primarily
func All() []Type {
	all := []Type{}
	for i := lower; i <= upper; i++ {
		all = append(all, Type(i))
	}

	return all
}
