package mysql

import (
	"bytes"
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/sql"
	"strings"
)

// StatementBuilder is component, used to create statements from conditions and filters
type StatementBuilder struct {
	buf          *bytes.Buffer
	placeholders []interface{}
}

// NewStatementBuilder returns new StatementBuilder struct
func NewStatementBuilder() *StatementBuilder {
	return &StatementBuilder{buf: bytes.NewBuffer(nil)}
}

// Build returns statement
func (s *StatementBuilder) Build() sql.Statement {
	if len(s.placeholders) == 0 {
		return sql.StringStatement(s.buf.String())
	}

	return sql.CommonStatement{
		SQL:          s.buf.String(),
		Placeholders: s.placeholders,
	}
}

// WriteKey writes table or column name
func (s *StatementBuilder) WriteKey(key string) *StatementBuilder {
	if l := len(key); l > 2 {
		if key[0] == '`' && key[l-1] == '`' && strings.Count(key, "`") == 2 {
			s.buf.WriteString(key)
			return s
		}
	}

	s.buf.WriteRune('`')
	s.buf.WriteString(key)
	s.buf.WriteRune('`')
	return s
}

// WriteNamed writes named entity (column or schema)
// Aliases are also supported
func (s *StatementBuilder) WriteNamed(n query.Named) *StatementBuilder {
	if n != nil {
		s.WriteKey(n.GetName())

		if a, ok := n.(query.Aliased); ok {
			s.buf.WriteRune(' ')
			s.buf.WriteString(a.GetAlias())
		}
	}

	return s
}
