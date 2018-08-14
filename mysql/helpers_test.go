package mysql

import (
	"github.com/mono83/atlas/query"
	"github.com/stretchr/testify/assert"
	"testing"
)

func assertBuilder(t *testing.T, b *StatementBuilder, sql string, ph ...interface{}) {
	stmt := b.Build()
	if assert.Equal(t, sql, stmt.GetSQL()) {
		if assert.Equal(t, len(ph), len(stmt.GetPlaceholders()), "Parameters count don't match") {
			for i, a := range stmt.GetPlaceholders() {
				assert.Equal(t, ph[i], a)
			}
		}
	}
}

func assertRule(t *testing.T, rule query.Rule, sql string, ph ...interface{}) {
	b := NewStatementBuilder()
	if assert.NoError(t, b.WriteRule(rule)) {
		assertBuilder(t, b, sql, ph...)
	}
}

func assertCondition(t *testing.T, c query.Condition, sql string, ph ...interface{}) {
	b := NewStatementBuilder()
	if assert.NoError(t, b.WriteCondition(c)) {
		assertBuilder(t, b, sql, ph...)
	}
}

func assertFilter(t *testing.T, f query.Filter, sql string, ph ...interface{}) {
	b := NewStatementBuilder()
	if assert.NoError(t, b.WriteFilter(f)) {
		assertBuilder(t, b, sql, ph...)
	}
}
