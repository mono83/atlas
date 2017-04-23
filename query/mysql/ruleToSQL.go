package mysql

import (
	"bytes"
	"errors"
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
)

// RuleToSQL converts rule struct into partial RuleSQL
func RuleToSQL(rule query.RuleDef) (*PartialSQL, error) {
	sb := bytes.NewBufferString("")
	placeholders := []interface{}{}

	if err := ruleToSQL(rule, sb, &placeholders); err != nil {
		return nil, err
	}

	return &PartialSQL{SQL: sb.String(), Placeholders: placeholders}, nil
}

func ruleToSQL(rule query.RuleDef, sb *bytes.Buffer, placeholders *[]interface{}) error {
	left := rule.GetLeft()
	right := rule.GetRight()

	if column, ok := left.(query.Named); ok {
		columnToSQL(column, sb)
	} else {
		return errors.New("Not a column on left side of query")
	}

	rightSideExpected := true
	switch rule.GetType() {
	case match.IsNull:
		sb.WriteString(" IS NULL")
		rightSideExpected = false
	case match.NotIsNull:
		sb.WriteString(" NOT IS NULL")
		rightSideExpected = false
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
		return errors.New("Unsupported match " + rule.GetType().String())
	}

	if rightSideExpected {
		if column, ok := right.(query.Named); ok {
			sb.WriteRune(' ')
			columnToSQL(column, sb)
		} else {
			sb.WriteRune(' ')
			sb.WriteString("?")
			*placeholders = append(*placeholders, right)
		}
	}

	return nil
}
