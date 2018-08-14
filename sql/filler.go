package sql

import "errors"

// Filler describes entity models, that can produce statements to
// read themselves from database
type Filler interface {
	GetStatement(args ... interface{}) (Statement, error)
}

// SelectFill perform database select using Statement, obtained
// from Filler ant then writes response to Filler
func SelectFill(i Invoker, f Filler, args ... interface{}) error {
	if i == nil {
		return errors.New("invoker not provided")
	}
	if f == nil {
		return errors.New("filler not provided")
	}

	// Building statement
	stmt, err := f.GetStatement(args...)
	if err != nil {
		return err
	}

	// Invoking
	return i.Invoke(stmt, f)
}
