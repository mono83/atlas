package exec

import (
	"database/sql"
	"errors"
	"time"
)

// Process contains information about running queries
type Process struct {
	ID       int
	User     string
	Host     string
	Database string
	Command  string
	Elapsed  time.Duration
	State    string
	Info     string
}

// ShowFullProcesslist performs SHOW FULL PROCESSLIST command and returns slice of running queries
func ShowFullProcesslist(db *sql.DB) ([]Process, error) {
	if db == nil {
		return nil, errors.New("no database provided")
	}

	rows, err := db.Query("SHOW FULL PROCESSLIST")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var optDb, optInfo *string
	var t int
	var list []Process
	for rows.Next() {
		p := Process{}

		optDb = nil
		if err := rows.Scan(&p.ID, &p.User, &p.Host, &optDb, &p.Command, &t, &p.State, &optInfo); err != nil {
			return nil, err
		}

		p.Elapsed = time.Second * time.Duration(t)
		if optDb != nil {
			p.Database = *optDb
		}
		if optInfo != nil {
			p.Info = *optInfo
		}

		list = append(list, p)
	}

	return list, nil
}
