package preset

import (
	"github.com/mono83/atlas/sql"
	"github.com/mono83/atlas/types"
	"sort"
	"strings"
	"time"
)

// Process contains information about running queries
type Process struct {
	ID           int64                 `db:"Id"`
	User         string                `db:"User"`
	Host         string                `db:"Host"`
	Database     types.OptString       `db:"db"`
	Command      string                `db:"Command"`
	Elapsed      types.DurationSeconds `db:"Time"`
	State        types.OptString       `db:"State"`
	Info         types.OptString       `db:"Info"`
	RowsSent     int                   `db:"Rows_sent"`
	RowsExamined int                   `db:"Rows_examined"`
}

// IsIDLE return true if process in IDLE state
func (p Process) IsIDLE() bool {
	return p.Command == "Sleep"
}

// ProcessList contains list of processes
type ProcessList []Process

// Len return amount of items
func (p ProcessList) Len() int {
	return len(p)
}

// All returns full list of process information
func (p ProcessList) All() []Process {
	return p
}

// WithElapsed returns only that processes, that lasts more that provided duration (inclusive)
func (p ProcessList) WithElapsed(d time.Duration) ProcessList {
	var list []Process
	for _, process := range p {
		if process.Elapsed.Duration >= d {
			list = append(list, process)
		}
	}
	return ProcessList(list)
}

// NotIDLE returns only not IDLE processes
func (p ProcessList) NotIDLE() ProcessList {
	var list []Process
	for _, process := range p {
		if !process.IsIDLE() {
			list = append(list, process)
		}
	}
	return ProcessList(list)
}

// SortByTimeAsc sorts processlist by elapsed time in ASC order
func (p ProcessList) SortByTimeAsc() {
	sort.Slice(p, func(i, j int) bool {
		return p[i].Elapsed.Duration < p[j].Elapsed.Duration
	})
}

// GetStatement returns statement, that could be used to obtain data
// Also it is sql.Filler interface implementation
func (p *ProcessList) GetStatement(args ...interface{}) (sql.Statement, error) {
	return sql.StringStatement("SHOW FULL PROCESSLIST"), nil
}

// Utilization contains information about database utilization per user
// obtained from process list
type Utilization struct {
	User, Host string
	Processes  int
	Total      time.Duration
}

// GetUtilizationByUser returns utilization data per user
func (p ProcessList) GetUtilizationByUser() []Utilization {
	m := map[string]*Utilization{}

	var u *Utilization
	var ok bool
	for _, process := range p {
		host := process.Host
		if last := strings.LastIndex(host, ":"); last > -1 {
			host = host[0:last]
		}
		key := process.User + "@" + host

		if u, ok = m[key]; !ok {
			u = &Utilization{User: process.User, Host: host}
			m[key] = u
		}

		u.Processes++
		u.Total += process.Elapsed.Duration
	}

	var response []Utilization
	for _, v := range m {
		response = append(response, *v)
	}

	return response
}
