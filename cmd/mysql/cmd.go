package mysql

import (
	"github.com/go-sql-driver/mysql"
	"github.com/mono83/atlas/mysql/sqlx"
	"github.com/mono83/atlas/sql"
	"github.com/spf13/cobra"
)

var mysqlAddr, mysqlUser, mysqlPassword, mysqlDatabase string

// Cmd is main CLI command for MySQL
var Cmd = &cobra.Command{
	Use:   "mysql",
	Short: "CLI MySQL toolset",
}

func init() {
	Cmd.PersistentFlags().StringVarP(&mysqlAddr, "addr", "a", "", "MySQL address, like 127.0.0.1:3306")
	Cmd.PersistentFlags().StringVarP(&mysqlUser, "user", "u", "", "MySQL username")
	Cmd.PersistentFlags().StringVarP(&mysqlPassword, "password", "p", "", "MySQL password")
	Cmd.PersistentFlags().StringVarP(&mysqlDatabase, "database", "d", "", "MySQL database name")

	Cmd.AddCommand(
		statusCmd,
	)
}

func connect(c *cobra.Command) (sql.Modifier, error) {
	// Building DSN
	config := mysql.Config{
		User:                 mysqlUser,
		Passwd:               mysqlPassword,
		DBName:               mysqlDatabase,
		Addr:                 mysqlAddr,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	// Connecting
	return sqlx.ConnectMySQL(config)
}
