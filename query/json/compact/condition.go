package compact

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/mono83/atlas/query"
)

func mapCondition(src query.Condition) condition {
	if x, ok := src.(condition); ok {
		return x
	}

	c := condition{t: src.GetType()}
	for _, x := range src.GetRules() {
		c.r = append(c.r, mapRule(x))
	}
	for _, x := range src.GetConditions() {
		c.c = append(c.c, mapCondition(x))
	}
	return c
}

type condition struct {
	t query.Logic
	r []rule
	c []condition
}

func (c condition) GetType() query.Logic { return c.t }
func (c condition) GetRules() []query.Rule {
	l := len(c.r)
	if l == 0 {
		return nil
	}

	res := make([]query.Rule, l)
	for i, j := range c.r {
		res[i] = j
	}
	return res
}
func (c condition) GetConditions() []query.Condition {
	l := len(c.c)
	if l == 0 {
		return nil
	}

	res := make([]query.Condition, l)
	for i, j := range c.c {
		res[i] = j
	}
	return res
}

func (c condition) MarshalJSON() ([]byte, error) {
	var t string
	if c.t == query.And {
		t = "AND"
	} else if c.t == query.Or {
		t = "OR"
	} else {
		return nil, errors.New("unsupported condition type")
	}

	rules := c.r
	conditions := c.c
	if rules == nil {
		rules = []rule{}
	}
	if conditions == nil {
		conditions = []condition{}
	}

	return json.Marshal([]interface{}{t, rules, conditions})
}

func (c *condition) UnmarshalJSON(src []byte) error {
	var arr []json.RawMessage
	if err := json.Unmarshal(src, &arr); err != nil {
		return err
	}
	if arr == nil {
		return errors.New("nil condition value")
	} else if len(arr) != 3 {
		return errors.New("condition array must contain exactly 3 elements")
	}

	// First element is type
	var t string
	if err := json.Unmarshal(arr[0], &t); err != nil {
		return err
	}
	var l query.Logic
	switch strings.ToUpper(t) {
	case "AND":
		l = query.And
	case "OR":
		l = query.Or
	default:
		return errors.New("unsupported condition type")
	}
	// Second element is rules array
	var rules []rule
	if err := json.Unmarshal(arr[1], &rules); err != nil {
		return err
	}
	// Third element is conditions array
	var conditions []condition
	if err := json.Unmarshal(arr[2], &conditions); err != nil {
		return err
	}

	*c = condition{t: l, r: rules, c: conditions}
	return nil
}
