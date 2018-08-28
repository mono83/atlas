package mysql

import (
	"errors"
	"github.com/mono83/atlas/query"
)

// WriteCondition converts condition into SQL and writes it into buffer
func (s *StatementBuilder) WriteCondition(cond query.Condition) error {
	if len(cond.GetConditions()) == 0 && len(cond.GetRules()) == 0 {
		return errors.New("empty condition - it has no rules and nested conditions")
	} else if len(cond.GetConditions()) == 0 && len(cond.GetRules()) == 1 {
		return s.WriteRule(cond.GetRules()[0])
	} else if len(cond.GetRules()) == 0 && len(cond.GetConditions()) == 1 {
		return s.WriteCondition(cond.GetConditions()[0])
	}

	sep := ""
	if cond.GetType() == query.Or {
		sep = " OR "
	} else if cond.GetType() == query.And {
		sep = " AND "
	} else {
		return errors.New("unsupported condition logic - it neither AND nor OR")
	}

	s.buf.WriteRune('(')
	i := 0

	for _, r := range cond.GetRules() {
		if i > 0 {
			s.buf.WriteString(sep)
		}
		err := s.WriteRule(r)
		if err != nil {
			return err
		}
		i++
	}

	for _, c := range cond.GetConditions() {
		if i > 0 {
			s.buf.WriteString(sep)
		}
		err := s.WriteCondition(c)
		if err != nil {
			return err
		}
		i++
	}

	s.buf.WriteRune(')')

	return nil
}
