package rules

import "github.com/mono83/atlas/query/match"

// False is special rule instance, that is always false
type False struct{}

// GetLeft is query.Rule interface implementation
func (False) GetLeft() interface{} { return 1 }

// GetRight is query.Rule interface implementation
func (False) GetRight() interface{} { return 2 }

// GetType is query.Rule interface implementation
func (False) GetType() match.Type { return match.Eq }
