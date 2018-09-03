package match

// Def contains addition definitions for rule match type
type Def struct {
	Name    string
	Aliases []string
	Math    []string
	Args    int
}

// Definitions contains data for known definitions
var Definitions map[Type]Def

func init() {
	Definitions = map[Type]Def{}
	Definitions[IsNull] = Def{Args: 1, Name: "Is_Null"}
	Definitions[NotIsNull] = Def{Args: 1, Name: "Not_Is_Null", Aliases: []string{"Is_Not_Null", "IsNotNull"}}
	Definitions[Equals] = Def{Args: 2, Name: "Equal", Aliases: []string{"Equals", "Eq"}, Math: []string{"="}}
	Definitions[NotEquals] = Def{Args: 2, Name: "Not_Equal", Aliases: []string{"Not_Equals", "NotEqual", "NotEquals", "Neq"}, Math: []string{"!=", "≠"}}
	Definitions[In] = Def{Args: 2, Name: "In"}
	Definitions[NotIn] = Def{Args: 2, Name: "Not_In"}
	Definitions[GreaterThan] = Def{Args: 2, Name: "Greater_Than", Aliases: []string{"GreaterThan", "Gt"}, Math: []string{">"}}
	Definitions[GreaterThanEquals] = Def{Args: 2, Name: "Greater_Than_Equals", Aliases: []string{"GreaterThanEquals", "Greater_Than_Or_Equals", "GreaterThanOrEquals", "Gte"}, Math: []string{">=", "≥"}}
	Definitions[LesserThan] = Def{Args: 2, Name: "Lesser_Than", Aliases: []string{"LesserThan", "Lt"}, Math: []string{"<"}}
	Definitions[LesserThanEquals] = Def{Args: 2, Name: "Lesser_Than_Equals", Aliases: []string{"LesserThanEquals", "Lesser_Than_Or_Equals", "LesserThanOrEquals", "Lte"}, Math: []string{"<=", "≤"}}
}
