package mycnf

import "github.com/go-sql-driver/mysql"

// Config contains connection settings with connection name, because .my.cnf
// allows multiple connections configuration
type Config struct {
	mysql.Config

	ConnectionName string
}
