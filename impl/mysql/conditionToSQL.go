package mysql

import (
	"bytes"
	"errors"
	"github.com/mono83/atlas/query"
)

// ConditionToSQL converts provided condition to MySQL query
func ConditionToSQL(cond query.Condition) (*PartialSQL, error) {
	var placeholders []interface{}
	sb := bytes.NewBufferString("")
	if err := conditionToSQL(cond, sb, &placeholders); err != nil {
		return nil, err
	}

	return &PartialSQL{SQL: sb.String(), Placeholders: placeholders}, nil
}

func conditionToSQL(cond query.Condition, sb *bytes.Buffer, placeholders *[]interface{}) error {
	if len(cond.GetConditions()) == 0 && len(cond.GetRules()) == 0 {
		return errors.New("empty condition - it has no rules and nested conditions")
	} else if len(cond.GetConditions()) == 0 && len(cond.GetRules()) == 1 {
		return ruleToSQL(cond.GetRules()[0], sb, placeholders)
	} else if len(cond.GetRules()) == 0 && len(cond.GetConditions()) == 1 {
		return conditionToSQL(cond.GetConditions()[0], sb, placeholders)
	}

	sep := ""
	if cond.GetType() == query.Or {
		sep = " OR "
	} else if cond.GetType() == query.And {
		sep = " AND "
	} else {
		return errors.New("unsupported condition logic - it neither AND nor OR")
	}

	sb.WriteRune('(')
	i := 0

	for _, r := range cond.GetRules() {
		if i > 0 {
			sb.WriteString(sep)
		}
		err := ruleToSQL(r, sb, placeholders)
		if err != nil {
			return err
		}
		i++
	}

	for _, c := range cond.GetConditions() {
		if i > 0 {
			sb.WriteString(sep)
		}
		err := conditionToSQL(c, sb, placeholders)
		if err != nil {
			return err
		}
		i++
	}

	sb.WriteRune(')')

	return nil
}
