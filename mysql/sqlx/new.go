package sqlx

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mono83/xray"
	"github.com/mono83/xray/args"
)

// ConnectMySQL establishes connection to MySQL server
func ConnectMySQL(c mysql.Config) (*Instance, error) {
	log := xray.BOOT.Fork().With(args.Host(c.User + "@" + c.Addr))
	log.Debug("Connecting to MySQL at :host using SQLx adapter")
	db, err := sqlx.Connect("mysql", c.FormatDSN())
	if err == nil {
		err = db.Ping()
	}
	if err != nil {
		return nil, err
	}

	log.Info("Established connection to MySQL at :host using SQLx adapter")

	return &Instance{DB: db}, nil
}
