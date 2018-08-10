package mysql

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

// MySQLCmd is root command for mysql operations
var MySQLCmd = &cobra.Command{
	Use:   "mysql",
	Short: "MySQL specific operations",
}

func init() {
	MySQLCmd.AddCommand(
		statusCmd,
	)

	MySQLCmd.PersistentFlags().StringP("addr", "a", "localhost:3306", "Database address")
	MySQLCmd.PersistentFlags().StringP("user", "u", "", "Database username")
	MySQLCmd.PersistentFlags().StringP("password", "p", "", "Database password")
}

func getConnection(cmd *cobra.Command) (*sql.DB, error) {
	config := mysql.Config{
		Net: "tcp",
	}

	config.Addr, _ = cmd.Flags().GetString("addr")
	config.User, _ = cmd.Flags().GetString("user")
	config.Passwd, _ = cmd.Flags().GetString("password")

	// Reading .my.cnf
	myCnf := userHomeDir() + string(os.PathSeparator) + ".my.cnf"
	if info, err := os.Stat(myCnf); err == nil && !info.IsDir() {
		// Reading file
		bts, _ := ioutil.ReadFile(myCnf)
		for _, line := range strings.Split(string(bts), "\n") {
			line = strings.TrimSpace(line)
			if len(line) == 0 || line[0] == '#' || line[0] == '[' {
				continue
			}

			if chunks := strings.Split(line, "="); len(chunks) == 2 {
				chunks[0] = strings.TrimSpace(chunks[0])
				chunks[1] = strings.TrimSpace(chunks[1])
				switch chunks[0] {
				case "user":
					if len(config.User) == 0 {
						config.User = chunks[1]
					}
				case "password":
					if len(config.Passwd) == 0 {
						config.Passwd = chunks[1]
					}
				}
			}
		}
	}

	// Establishing connection
	return sql.Open("mysql", config.FormatDSN())
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
