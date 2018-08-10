package mysql

import (
	"bytes"
	"errors"
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
)

// RuleToSQL converts rule struct into partial RuleSQL
func RuleToSQL(rule query.Rule) (*PartialSQL, error) {
	sb := bytes.NewBufferString("")
	var placeholders []interface{}

	if err := ruleToSQL(rule, sb, &placeholders); err != nil {
		return nil, err
	}

	return &PartialSQL{SQL: sb.String(), Placeholders: placeholders}, nil
}

func ruleToSQL(rule query.Rule, sb *bytes.Buffer, placeholders *[]interface{}) error {
	left := rule.GetLeft()
	right := rule.GetRight()

	switch rule.GetType() {
	case match.IsNull, match.NotIsNull:
		return ruleToSQLNulls(left, rule.GetType(), sb)
	case match.Equals,
		match.NotEquals,
		match.GreaterThan,
		match.LowerThan,
		match.GreaterThanEquals,
		match.LowerThanEquals:
		return ruleToSQLSimpleOps(left, right, rule.GetType(), sb, placeholders)
	case match.In, match.NotIn:
		return ruleToSQLIN(left, right, rule.GetType(), sb, placeholders)
	default:
		return UnsupportedOperation(rule.GetType())
	}
}

// ruleToSQLNulls handles IS NULL and NOT IS NULL cases
func ruleToSQLNulls(left interface{}, t match.Type, sb *bytes.Buffer) error {
	if column, ok := left.(query.Named); ok {
		columnToSQL(column, sb)
	} else {
		return LeftIsNotColumn{Real: left}
	}

	if t == match.IsNull {
		sb.WriteString(" IS NULL")
	} else if t == match.NotIsNull {
		sb.WriteString(" NOT IS NULL")
	} else {
		return UnsupportedOperation(t)
	}

	return nil
}

// ruleToSQLSimpleOps handles simple operations
func ruleToSQLSimpleOps(left, right interface{}, t match.Type, sb *bytes.Buffer, placeholders *[]interface{}) error {
	if column, ok := left.(query.Named); ok {
		columnToSQL(column, sb)
	} else {
		return LeftIsNotColumn{Real: left}
	}

	switch t {
	case match.Equals:
		sb.WriteString(" =")
	case match.NotEquals:
		sb.WriteString(" <>")
	case match.GreaterThan:
		sb.WriteString(" >")
	case match.GreaterThanEquals:
		sb.WriteString(" >=")
	case match.LowerThan:
		sb.WriteString(" <")
	case match.LowerThanEquals:
		sb.WriteString(" <=")
	default:
		return UnsupportedOperation(t)
	}

	if column, ok := right.(query.Named); ok {
		sb.WriteRune(' ')
		columnToSQL(column, sb)
	} else {
		sb.WriteRune(' ')
		sb.WriteString("?")
		*placeholders = append(*placeholders, right)
	}

	return nil
}

// ruleToSQLIN used to handle IN and NOT IN clauses
func ruleToSQLIN(left, right interface{}, t match.Type, sb *bytes.Buffer, placeholders *[]interface{}) error {
	l := 0
	if s, ok := right.([]string); ok {
		// String slice
		l = len(s)
		for _, v := range s {
			*placeholders = append(*placeholders, v)
		}
	} else if s, ok := right.([]int); ok {
		// Integer slice
		l = len(s)
		for _, v := range s {
			*placeholders = append(*placeholders, v)
		}
	} else if s, ok := right.([]int64); ok {
		// Long slice
		l = len(s)
		for _, v := range s {
			*placeholders = append(*placeholders, v)
		}
	} else if s, ok := right.([]interface{}); ok {
		// Slice of some values
		l = len(s)
		*placeholders = append(*placeholders, s...)
	} else {
		return errors.New("only []int, []int64, []string and []interface{} are allowed for IN operations")
	}

	if l == 0 {
		return errors.New("missing data for IN operations - empty values slice received")
	}

	if column, ok := left.(query.Named); ok {
		columnToSQL(column, sb)
	} else {
		return LeftIsNotColumn{Real: left}
	}

	switch t {
	case match.In:
		sb.WriteString(" IN (")
	case match.NotIn:
		sb.WriteString(" NOT IN (")
	default:
		return UnsupportedOperation(t)
	}

	for i := 0; i < l; i++ {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteRune('?')
	}

	sb.WriteRune(')')
	return nil
}
