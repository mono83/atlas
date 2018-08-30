package match

// Type describes matcher type and uses primarily in rules
type Type byte

// Invert returns inverted match
func (o Type) Invert() Type {
	if o == Unknown || o.IsCustom() || !o.IsStandard() {
		return Unknown
	}

	if o%2 == 0 {
		return Type(byte(o) - 1)
	}

	return Type(byte(o) + 1)
}

// IsCustom returns true if type is custom type
func (o Type) IsCustom() bool {
	return o > 127
}

// IsStandard returns true if operation is in standard
// operations pool
func (o Type) IsStandard() bool {
	i := byte(o)
	return i >= lower && i <= upper
}

// Not reverses match condition
func Not(operation Type) Type {
	return operation.Invert()
}
