package mysql

// PartialSQL contains SQL chunk with placeholders, used within
type PartialSQL struct {
	SQL          string
	Placeholders []interface{}
}
