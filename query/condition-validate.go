package query

import (
	"strconv"
)

// ValidateCondition validates condition contents
// Empty conditions are valid
func ValidateCondition(c Condition) error {
	return validateConditionRecursive(1, c)
}

// validateConditionRecursive performs recursive validation
// of condition and preserves depth level for more verbose errors
func validateConditionRecursive(level int, c Condition) error {
	if c == nil {
		return onLevelError{level: level, err: "nil condition"}
	}
	if c.GetType() != And && c.GetType() != Or {
		return onLevelError{level: level, err: "unsupported logic type"}
	}

	// Validating rules
	for i, r := range c.GetRules() {
		if err := ValidateRule(r); err != nil {
			return onLevelError{
				level: level,
				err:   "rule #" + strconv.Itoa(i) + "(starting from 0): " + err.Error(),
			}
		}
	}

	// Validating inner conditions
	for _, ic := range c.GetConditions() {
		if err := validateConditionRecursive(level+1, ic); err != nil {
			return err
		}
	}

	return nil
}

type onLevelError struct {
	level int
	err   string
}

func (o onLevelError) Error() string {
	switch o.level {
	case 1:
		return "error on top condition level: " + o.err
	case 2:
		return "error on second condition level: " + o.err
	case 3:
		return "error on third condition level: " + o.err
	default:
		return "error on " + strconv.Itoa(o.level) + "th condition level: " + o.err
	}
}
