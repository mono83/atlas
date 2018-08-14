package mysql

import (
	"errors"
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"github.com/mono83/atlas/query/rules"
)

// WriteRule writes rule into builder
func (s *StatementBuilder) WriteRule(rule query.Rule) error {
	if rule == nil {
		return errors.New("empty rule")
	}
	if _, ok := rule.(rules.False); ok {
		s.buf.WriteString("1=0")
		return nil
	}

	left := rule.GetLeft()
	right := rule.GetRight()

	switch rule.GetType() {
	case match.IsNull, match.NotIsNull:
		return s.ruleToSQLNulls(left, rule.GetType())
	case match.Equals,
		match.NotEquals,
		match.GreaterThan,
		match.LowerThan,
		match.GreaterThanEquals,
		match.LowerThanEquals:
		return s.ruleToSQLSimpleOps(left, right, rule.GetType())
	case match.In, match.NotIn:
		return s.ruleToSQLIN(left, right, rule.GetType())
	default:
		return UnsupportedOperation(rule.GetType())
	}
}

// ruleToSQLNulls handles IS NULL and NOT IS NULL cases
func (s *StatementBuilder) ruleToSQLNulls(left interface{}, t match.Type) error {
	if column, ok := left.(query.Named); ok {
		s.WriteNamed(column)
	} else {
		return LeftIsNotColumn{Real: left}
	}

	if t == match.IsNull {
		s.buf.WriteString(" IS NULL")
	} else if t == match.NotIsNull {
		s.buf.WriteString(" NOT IS NULL")
	} else {
		return UnsupportedOperation(t)
	}

	return nil
}

// ruleToSQLSimpleOps handles simple operations
func (s *StatementBuilder) ruleToSQLSimpleOps(left, right interface{}, t match.Type) error {
	if column, ok := left.(query.Named); ok {
		s.WriteNamed(column)
	} else {
		return LeftIsNotColumn{Real: left}
	}

	switch t {
	case match.Equals:
		s.buf.WriteString(" =")
	case match.NotEquals:
		s.buf.WriteString(" <>")
	case match.GreaterThan:
		s.buf.WriteString(" >")
	case match.GreaterThanEquals:
		s.buf.WriteString(" >=")
	case match.LowerThan:
		s.buf.WriteString(" <")
	case match.LowerThanEquals:
		s.buf.WriteString(" <=")
	default:
		return UnsupportedOperation(t)
	}

	if column, ok := right.(query.Named); ok {
		s.buf.WriteRune(' ')
		s.WriteNamed(column)
	} else {
		s.buf.WriteRune(' ')
		s.buf.WriteString("?")
		s.placeholders = append(s.placeholders, right)
	}

	return nil
}

// ruleToSQLIN used to handle IN and NOT IN clauses
func (s *StatementBuilder) ruleToSQLIN(left, right interface{}, t match.Type) error {
	l := 0
	if x, ok := right.([]string); ok {
		// String slice
		l = len(x)
		for _, v := range x {
			s.placeholders = append(s.placeholders, v)
		}
	} else if x, ok := right.([]int); ok {
		// Integer slice
		l = len(x)
		for _, v := range x {
			s.placeholders = append(s.placeholders, v)
		}
	} else if x, ok := right.([]int64); ok {
		// Long slice
		l = len(x)
		for _, v := range x {
			s.placeholders = append(s.placeholders, v)
		}
	} else if x, ok := right.([]interface{}); ok {
		// Slice of some values
		l = len(x)
		s.placeholders = append(s.placeholders, x...)
	} else {
		return errors.New("only []int, []int64, []string and []interface{} are allowed for IN operations")
	}

	if l == 0 {
		return errors.New("missing data for IN operations - empty values slice received")
	}

	if column, ok := left.(query.Named); ok {
		s.WriteNamed(column)
	} else {
		return LeftIsNotColumn{Real: left}
	}

	switch t {
	case match.In:
		s.buf.WriteString(" IN (")
	case match.NotIn:
		s.buf.WriteString(" NOT IN (")
	default:
		return UnsupportedOperation(t)
	}

	for i := 0; i < l; i++ {
		if i > 0 {
			s.buf.WriteRune(',')
		}
		s.buf.WriteRune('?')
	}

	s.buf.WriteRune(')')
	return nil
}
