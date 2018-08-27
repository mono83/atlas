package match

// Type describes matcher type and uses primarily in rules
type Type int

// Invert returns inverted match
func (o Type) Invert() Type {
	if o == Unknown || !o.Valid() {
		return Unknown
	}

	if o%2 == 0 {
		return Type(int(o) - 1)
	}

	return Type(int(o) + 1)
}

// Valid returns true if operation is valid
func (o Type) Valid() bool {
	i := int(o)
	return i >= lower && i <= upper
}

// Not reverses match condition
func Not(operation Type) Type {
	return operation.Invert()
}
