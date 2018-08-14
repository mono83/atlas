package mysql

import (
	"errors"
	"github.com/mono83/atlas/query"
	"strconv"
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

// WriteFilter converts filter into SQL and writes it into buffer
func (s *StatementBuilder) WriteFilter(f query.Filter) error {
	if err := s.WriteCondition(f); err != nil {
		return err
	}

	if len(f.GetSorting()) > 0 {
		// Applying sorting
		s.buf.WriteString(" ORDER BY ")

		for i, sort := range f.GetSorting() {
			if i > 0 {
				s.buf.WriteRune(',')
			}
			s.WriteNamed(sort)
			s.buf.WriteRune(' ')
			if sort.GetType() == query.Desc {
				s.buf.WriteString("DESC")
			} else if sort.GetType() == query.Asc {
				s.buf.WriteString("ASC")
			} else {
				return errors.New("unknown sort type")
			}
		}
	}

	if lim := f.GetLimit(); lim > 0 {
		// Applying limit
		s.buf.WriteString(" LIMIT ")
		if off := f.GetOffset(); off > 0 {
			s.buf.WriteString(strconv.Itoa(off))
			s.buf.WriteRune(',')
		}
		s.buf.WriteString(strconv.Itoa(lim))
	}

	return nil
}
