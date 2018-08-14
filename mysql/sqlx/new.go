package sqlx

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// ConnectMySQL establishes connection to MySQL server
func ConnectMySQL(c mysql.Config) (*Instance, error) {
	db, err := sqlx.Connect("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	return &Instance{DB: db}, nil
}
