package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/mono83/atlas/query"
)

// NewDAOX return data access object, backed by sqlx
func NewDAOX(db *sqlx.DB) query.ReadOnlyDAO {
	return invoker{db: db}
}

type invoker struct {
	db *sqlx.DB
}

func (i invoker) Select(def query.SelectDef, target interface{}) error {
	sql, placeholders, err := QueryToSQL(def)
	if err != nil {
		return err
	}

	return i.db.Select(target, sql, placeholders...)
}
