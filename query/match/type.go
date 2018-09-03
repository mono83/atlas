package match

// Type describes matcher type and uses primarily in rules
type Type byte

// Invert returns inverted match
func (t Type) Invert() Type {
	if t == Unknown || t.IsCustom() || !t.IsStandard() {
		return Unknown
	}

	if t%2 == 0 {
		return Type(byte(t) - 1)
	}

	return Type(byte(t) + 1)
}

// IsCustom returns true if type is custom type
func (t Type) IsCustom() bool {
	return t > 127
}

// IsStandard returns true if operation is in standard
// operations pool
func (t Type) IsStandard() bool {
	i := byte(t)
	return i >= lower && i <= upper
}

// Not reverses match condition
func Not(operation Type) Type {
	return operation.Invert()
}
