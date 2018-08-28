package mysql

import (
	"github.com/mono83/atlas/query"
	"github.com/stretchr/testify/assert"
	"testing"
)

var keysDataProvider = []struct {
	Expected string
	Provided string
}{
	{"`id`", "id"},
	{"`id`", "`id`"},
	{"``id`", "`id"}, // Invalid behaviour
}

func TestWriteKey(t *testing.T) {
	for _, d := range keysDataProvider {
		t.Run(d.Provided, func(t *testing.T) {
			assert.Equal(t, d.Expected, NewStatementBuilder().WriteKey(d.Provided).Build().GetSQL())
		})
	}
}

var namedsDataProvider = []struct {
	ExpectedColumn  string
	ExpectedSchema  string
	ExpectedSorting string
	Source          query.Named
}{
	{"`foo`", "`foo`", "`foo`", query.String("foo")},
	{"`foo`", "`foo`", "`foo`", query.String("`foo`")},
	{"`bar` as `zzz`", "`bar`", "`bar`", query.AliasedName{Name: "bar", Alias: "zzz"}},
	{"`bar` as `zzz`", "`bar`", "`bar`", query.AliasedName{Name: "`bar`", Alias: "`zzz`"}},
}

func TestWriteNameds(t *testing.T) {
	var b *StatementBuilder
	for _, d := range namedsDataProvider {
		t.Run(d.ExpectedColumn, func(t *testing.T) {
			b = NewStatementBuilder()
			if assert.NoError(t, b.WriteColumn(d.Source)) {
				assert.Equal(t, d.ExpectedColumn, b.Build().GetSQL())
			}

			b = NewStatementBuilder()
			if assert.NoError(t, b.WriteSchema(d.Source)) {
				assert.Equal(t, d.ExpectedSchema, b.Build().GetSQL())
			}

			b = NewStatementBuilder()
			b.WriteNamed(d.Source)
			assert.Equal(t, d.ExpectedSorting, b.Build().GetSQL())
		})
	}
}
