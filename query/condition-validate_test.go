package query

import (
	"fmt"
	"testing"
)

var conditionValidateDataProvider = []struct {
	Error     string
	Condition Condition
}{
	{"", CommonCondition{Type: And}},
	{"error on top condition level: nil condition", nil},
	{"error on top condition level: unsupported logic type", CommonCondition{}},
	{"error on second condition level: nil condition", CommonCondition{Type: And, Conditions: []Condition{nil}}},
	{"error on third condition level: nil condition", CommonCondition{Type: And, Conditions: []Condition{CommonCondition{Type: And, Conditions: []Condition{nil}}}}},
	{"error on top condition level: rule #0(starting from 0): nil rule", CommonCondition{Type: And, Rules: []Rule{nil}}},
}

func TestConditionValidate(t *testing.T) {
	for _, d := range conditionValidateDataProvider {
		t.Run(fmt.Sprintf("%v", d), func(t *testing.T) {
			err := ValidateCondition(d.Condition)
			if err == nil {
				if d.Error != "" {
					t.Errorf(`expected error "%s" but got nothing`, d.Error)
				}
			} else {
				if d.Error != err.Error() {
					t.Errorf(`expected "%s" but got "%s"`, d.Error, err.Error())
				}
			}
		})
	}
}
