package match

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var typesDataProvider = []struct {
	IsStandard bool
	IsCustom   bool
	String     string
	Inverted   Type
	Type       Type
}{
	{false, false, "UNKNOWN", Unknown, Unknown},
	{false, false, "UNSUPPORTED #127", Unknown, Type(127)},

	{false, true, "CUSTOM #128", Unknown, Type(128)},
	{false, true, "CUSTOM #255", Unknown, Type(255)},

	{true, false, "IS NULL", NotIsNull, IsNull},
	{true, false, "NOT IS NULL", IsNull, NotIsNull},
	{true, false, "EQUALS", NotEquals, Equals},
	{true, false, "IN", NotIn, In},
	{true, false, "NOT IN", In, NotIn},
	{true, false, "GREATER THAN", LesserThanEquals, GreaterThan},
	{true, false, "GREATER THAN EQUALS", LesserThan, GreaterThanEquals},
	{true, false, "LOWER THAN", GreaterThanEquals, LowerThan},
	{true, false, "LOWER THAN EQUALS", GreaterThan, LowerThanEquals},
}

func TestType(t *testing.T) {
	for _, d := range typesDataProvider {
		t.Run(d.Type.String(), func(t *testing.T) {
			assert.Equal(t, d.IsStandard, d.Type.IsStandard())
			assert.Equal(t, d.IsCustom, d.Type.IsCustom())
			assert.Equal(t, d.Inverted, d.Type.Invert())
			assert.Equal(t, d.String, d.Type.String())
		})
	}
}
