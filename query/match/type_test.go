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
	{false, false, "Unknown", Unknown, Unknown},
	{false, false, "Unsupported #127", Unknown, Type(127)},

	{false, true, "Custom #128", Unknown, Type(128)},
	{false, true, "Custom #255", Unknown, Type(255)},

	{true, false, "Is_Null", NotIsNull, IsNull},
	{true, false, "Not_Is_Null", IsNull, NotIsNull},
	{true, false, "Equal", NotEquals, Equals},
	{true, false, "In", NotIn, In},
	{true, false, "Not_In", In, NotIn},
	{true, false, "Greater_Than", LesserThanEquals, GreaterThan},
	{true, false, "Greater_Than_Equals", LesserThan, GreaterThanEquals},
	{true, false, "Lesser_Than", GreaterThanEquals, LowerThan},
	{true, false, "Lesser_Than_Equals", GreaterThan, LowerThanEquals},
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
