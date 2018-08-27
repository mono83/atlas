package query

import (
	"errors"

	"github.com/mono83/atlas/query/match"
)

// ValidateRule performs validation on provided rule
// It is not silver bullet, but can help to avoid common problems
func ValidateRule(r Rule) error {
	// Simple checks
	if r == nil {
		return errors.New("nil rule")
	}
	if r.GetType() == match.Unknown {
		return errors.New("unknown rule match type")
	}

	switch r.GetType() {
	case match.IsNull, match.NotIsNull:
		if r.GetRight() != nil {
			return errors.New("null check rules must not have right part")
		}
	case match.Equals, match.NotEquals:
		if r.GetLeft() == nil || r.GetRight() == nil {
			return errors.New("nil found in equality check")
		}
	case match.Gt, match.Lt, match.Gte, match.Lte:
		if r.GetLeft() == nil {
			return errors.New("nil found in numeric rule on the left side")
		} else if r.GetRight() == nil {
			return errors.New("nil found in numeric rule on the right side")
		} else if !isNamedOrNumber(r.GetLeft()) {
			return errors.New("only numbers and nameds should be in numeric rule on left side")
		} else if !isNamedOrNumber(r.GetRight()) {
			return errors.New("only numbers and nameds should be in numeric rule on right side")
		}
	}

	return nil
}

// isNamedOrNumber return true if provided value is number
// or implements query.Named interface
func isNamedOrNumber(x interface{}) bool {
	if x == nil {
		return false
	}
	if _, ok := x.(Named); ok {
		return true
	}

	return isNumber(x)
}

// isNumber returns true if provided value is number
func isNumber(x interface{}) bool {
	if x == nil {
		return false
	}

	switch x.(type) {
	case int, int16, int32, int64, uint16, uint32, uint64, float32, float64:
		return true
	default:
		return false
	}
}
