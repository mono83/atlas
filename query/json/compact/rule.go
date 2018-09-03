package compact

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mono83/atlas/query"
	"github.com/mono83/atlas/query/match"
	"strings"
)

func mapRule(src query.Rule) (response rule) {
	if x, ok := src.(rule); ok {
		return x
	}
	if src != nil {
		response.l = src.GetLeft()
		response.r = src.GetRight()
		response.t = src.GetType()
	}
	return
}

type rule struct {
	l, r interface{}
	t    match.Type
}

func (r rule) GetLeft() interface{}  { return r.l }
func (r rule) GetRight() interface{} { return r.r }
func (r rule) GetType() match.Type   { return r.t }

func (r rule) MarshalJSON() ([]byte, error) {
	if r.t.IsCustom() {
		return nil, errors.New("custom rule types are not supported")
	} else if !r.t.IsStandard() {
		return nil, errors.New("not standard rule types are not supported")
	}

	def, ok := match.Definitions[r.t]
	if !ok {
		return nil, errors.New("rule type definition not found")
	}

	left, right := r.l, r.r
	if n, ok := r.l.(query.Named); ok {
		left = n.GetName()
	}
	if n, ok := r.r.(query.Named); ok {
		right = n.GetName()
	}

	return json.Marshal([]interface{}{
		strings.ToLower(def.Name),
		left,
		right,
	}[0 : def.Args+1])
}

func (r *rule) UnmarshalJSON(src []byte) error {
	var arr []interface{}
	if err := json.Unmarshal(src, &arr); err != nil {
		return err
	}

	if arr == nil {
		return errors.New("nil rule value")
	} else if len(arr) == 0 {
		return errors.New("empty rule array")
	}

	op, ok := arr[0].(string)
	if !ok {
		return errors.New("first element of rule array not string")
	}

	// Resolving type
	res, ok := ruleTypeFromString[op]
	if !ok {
		return fmt.Errorf(`unable to parse rule type "%s"`, op)
	}

	if len(arr) < res.d.Args+1 {
		return fmt.Errorf(
			`invalid rule array elements counts, must be at lease %d but got %d`,
			res.d.Args+1, len(arr),
		)
	}

	var right interface{}
	if len(arr) > 2 {
		right = arr[2]
	}

	*r = rule{
		t: res.t,
		l: arr[1],
		r: right,
	}

	return nil
}

var ruleTypeFromString map[string]struct {
	t match.Type
	d match.Def
}

func init() {
	ruleTypeFromString = map[string]struct {
		t match.Type
		d match.Def
	}{}
	for t, d := range match.Definitions {
		data := struct {
			t match.Type
			d match.Def
		}{t: t, d: d}

		for _, name := range d.Names() {
			name = strings.ToLower(name)
			ruleTypeFromString[name] = data
		}
	}
}
