package mysql

import (
	"bytes"
	"errors"
	"github.com/mono83/atlas/query"
	"strconv"
)

// QueryToSQL converts provided query to SQL
func QueryToSQL(q query.SelectDef) (string, []interface{}, error) {
	sb := bytes.NewBufferString("SELECT ")
	placeholders := []interface{}{}

	// Printing columns
	if columns := q.GetColumns(); len(columns) > 0 {
		for i, c := range columns {
			if i > 0 {
				sb.WriteRune(',')
			}

			columnToSQL(c, sb)
		}

	} else {
		sb.WriteString("*")
	}

	// Printing schema
	sb.WriteString(" FROM ")
	schemaToSQL(q.GetSchema(), sb)

	// Printing condition
	if condition := q.GetCondition(); condition != nil {
		sb.WriteString(" WHERE ")
		if err := conditionToSQL(condition, sb, &placeholders); err != nil {
			return "", nil, err
		}
	}

	// Printing order
	if order := q.GetOrder(); order != nil {
		sb.WriteString(" ORDER BY ")
		for i, o := range order {
			if i > 0 {
				sb.WriteRune(',')
			}

			columnToSQL(o.GetColumn(), sb)
			if o.GetType() == query.Asc {
				sb.WriteString(" ASC")
			} else if o.GetType() == query.Desc {
				sb.WriteString(" DESC")
			} else {
				return "", nil, errors.New("Unsupported order type")
			}
		}
	}

	// Printing limits
	if offset, limit := q.GetOffsetLimit(); limit > 0 {
		sb.WriteString(" LIMIT ")
		if offset != 0 {
			sb.WriteString(strconv.FormatInt(offset, 10))
			sb.WriteRune(',')
		}

		sb.WriteString(strconv.Itoa(limit))
	}

	return sb.String(), placeholders, nil
}

func schemaToSQL(s query.Named, sb *bytes.Buffer) {
	sb.WriteString(escape(s.GetName()))
	if a, ok := s.(query.Aliased); ok {
		sb.WriteRune(' ')
		sb.WriteString(a.GetAlias())
	}
}

func columnToSQL(c query.Named, sb *bytes.Buffer) {
	sb.WriteString(escape(c.GetName()))

	if a, ok := c.(query.Aliased); ok {
		sb.WriteString(" AS `")
		sb.WriteString(a.GetAlias())
		sb.WriteRune('`')
	}
}
