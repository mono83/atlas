package sqlx

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/mono83/atlas/sql"
)

// Instance is wrapper over SQLx
type Instance struct {
	*sqlx.DB
}

// Invoke is sql.Invoker interface implementation
func (i *Instance) Invoke(stmt sql.Statement, target interface{}) error {
	if i.DB == nil {
		return errors.New("DB not configured")
	}
	return i.DB.Select(target, stmt.GetSQL(), stmt.GetPlaceholders()...)
}

// Modify is sql.Modifier interface implementation
func (i *Instance) Modify(stmt sql.Statement) (int64, error) {
	if i.DB == nil {
		return 0, errors.New("DB not configured")
	}
	if stmt == nil {
		return 0, errors.New("empty statement")
	}

	res, err := i.DB.Exec(stmt.GetSQL(), stmt.GetPlaceholders()...)
	if err != nil {
		return 0, err
	}

	if id, e := res.LastInsertId(); e == nil && id != 0 {
		return id, nil
	}

	rc, _ := res.RowsAffected()
	return rc, nil
}
